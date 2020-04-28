gin使用**Model binding**绑定请求的内容到一个类型上。
#### 使用
使用方法：
```golang
package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	route := gin.Default()
	route.Any("/testing", startPage)
	route.Run(":8085")
}

func startPage(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "Success")
}
```
#### 源码
1. c.ShouldBindQuery(&person)
```golang
// ShouldBindQuery is a shortcut for c.ShouldBindWith(obj, binding.Query).
func (c *Context) ShouldBindQuery(obj interface{}) error {
	return c.ShouldBindWith(obj, binding.Query)
}
```
ShouldBindQuery 传入一个接口类型的obj，执行了  c.ShouldBindWith(obj, binding.Query)

2.  c.ShouldBindWith(obj, binding.Query)
```golang
// ShouldBindWith binds the passed struct pointer using the specified binding engine.
// See the binding package.
func (c *Context) ShouldBindWith(obj interface{}, b binding.Binding) error {
	return b.Bind(c.Request, obj)
}
```
ShouldBindWith 方法接收两个参数，第一个参数传一个接口类型的obj，第二个参数传入一个 binding.Binding的实现。
```golang
// Binding describes the interface which needs to be implemented for binding the
// data present in the request such as JSON request body, query parameters or
// the form POST.
type Binding interface {
	Name() string
	Bind(*http.Request, interface{}) error
}
```
ShouldBindWith调用binding.Binding的Bind方法。

3. binding.Query
```golang
package binding
var (
	...
	Query         = queryBinding{}
	...
)
```
binding.Query 初始化了一个queryBinding的结构体
```golang
// Copyright 2017 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import "net/http"

type queryBinding struct{}

func (queryBinding) Name() string {
	return "query"
}

// Bind方法传入一个*http.Request和一个interface{}类型
func (queryBinding) Bind(req *http.Request, obj interface{}) error {
	// 获取到请求的url里的query的参数
	values := req.URL.Query()
	// 调用mapForm方法，传入obj 和url.Value
	if err := mapForm(obj, values); err != nil {
		return err
	}
	return validate(obj)
}
```

4. mapForm
mapForm接收一个interface{}的指针和一个map[string][]string类型的参数
```golang
func mapForm(ptr interface{}, form map[string][]string) error {
	return mapFormByTag(ptr, form, "form")
}
```
调用mapFormByTag方法

5. mapFormByTag
mapFormByTag 接收一个tag的字符串，第4步的时候第三个参数传入的是**form**
```golang
func mapFormByTag(ptr interface{}, form map[string][]string, tag string) error {
	return mappingByPtr(ptr, formSource(form), tag)
}
```
`formSource(form)` 把form转换成formSource结构
```golang
type formSource map[string][]string
```
formSource实现了setter接口
```golang
// setter tries to set value on a walking by fields of a struct
type setter interface {
	TrySet(value reflect.Value, field reflect.StructField, key string, opt setOptions) (isSetted bool, err error)
}

var _ setter = formSource(nil)
```
调用mappingByPtr方法

6. mapFormByTag
mappingByPtr 接收三个参数 第一个是interface{}的指针类型，第二个是setter的实现，第三个是tag，这里传入的是`form`
```golang
func mappingByPtr(ptr interface{}, setter setter, tag string) error {
	_, err := mapping(reflect.ValueOf(ptr), emptyField, setter, tag)
	return err
}
```
调用mapping方法，传入ptr反射的Value,空的reflect.SturctField结构emptyField,setter接口实现和tag`form`
```golang
var emptyField = reflect.StructField{}
```

7. mapping
```golang
// mapping 
func mapping(value reflect.Value, field reflect.StructField, setter setter, tag string) (bool, error) {
	// 这里会判断传入的结构体字段类型field,如果标签form为-，就跳过。
	if field.Tag.Get(tag) == "-" { // just ignoring this field
		return false, nil
	}
	
	// 获取到value的类别
	var vKind = value.Kind()
	
		// 这里的意思是value如果是一个指针类型，并且它是nil，那么会初始化一个该类型的零值。
	if vKind == reflect.Ptr {
		var isNew bool
		vPtr := value
		
		if value.IsNil() {
			isNew = true
				// New返回一个Value类型值，该值持有一个**指向类型为typ的新申请的零值的指针**，返回值的Type为PtrTo(typ)。
				// 这里的value的Kind必须是是Array、Chan、Map、Ptr或Slice
			vPtr = reflect.New(value.Type().Elem())
		}
		// 取这个指针所对应的值vPtr.Elem()，再对它的值调用mapping
		isSetted, err := mapping(vPtr.Elem(), field, setter, tag)
		if err != nil {
			return false, err
		}
		// 如果指针为nil，给该指针新申请了一个零值空间，并且指针对应的值已经被设置好了。
		if isNew && isSetted {
			// 那么把value的值设置成vPtr
			// 到这里可以理解为什么要用vPtr=value，因为 vPtr.Elem()的值可能会设置失败，那么不应该更改原值。（类似于事务）
			value.Set(vPtr)
		}
		// 返回是否设置成功
		return isSetted, nil
	}
	
	// 如果vKind不是结构体，或者field不是匿名字段执行下面的语句
	if vKind != reflect.Struct || !field.Anonymous {
		// tryToSetValue
		ok, err := tryToSetValue(value, field, setter, tag)
		if err != nil {
			return false, err
		}
		if ok {
			return true, nil
		}
	}
	
	// 匿名结构体
	if vKind == reflect.Struct {
		tValue := value.Type()

		var isSetted bool
		// 拿到结构体的每个字段进行遍历
		for i := 0; i < value.NumField(); i++ {
			sf := tValue.Field(i)
			// 这个字段的包路径如果不为空，并且不是匿名字段就跳过 比如types.Person,匿名结构体内嵌套了有名结构体，会跳过 [？？？]
			if sf.PkgPath != "" && !sf.Anonymous { // unexported
				continue
			}
			// 用结构体的值调用mapping
			ok, err := mapping(value.Field(i), tValue.Field(i), setter, tag)
			if err != nil {
				return false, err
			}
			isSetted = isSetted || ok
		}
		return isSetted, nil
	}
	return false, nil
}
```
tryToSetValue:
```golang
func tryToSetValue(value reflect.Value, field reflect.StructField, setter setter, tag string) (bool, error) {
	var tagValue string
	var setOpt setOptions
	// 获取到tag
	tagValue = field.Tag.Get(tag)
	// head方法用来将tag拿到的值用 ","分割，tagValue为第一个，opts为剩余部分
	tagValue, opts := head(tagValue, ",")
	
	// 如果没有设置tag，默认tag的值等于字段名
	if tagValue == "" { // default value is FieldName
		tagValue = field.Name
	}
	// field是emptyField的时候返回false
	if tagValue == "" { // when field is "emptyField" variable
		return false, nil
	}

	var opt string
	// 循环拆分tag
	for len(opts) > 0 {
		opt, opts = head(opts, ",")
		// 如果拆分出的opt是default= xxx这种形式的，那么就是有默认值，默认值=v
		if k, v := head(opt, "="); k == "default" {
			setOpt.isDefaultExists = true
			setOpt.defaultValue = v
		}
	}
	
	// 调用TrySet方法设置值。
	return setter.TrySet(value, field, tagValue, setOpt)
}
```

8. TrySet
TrySet 试图把request的值设置到接收对象中
```golang
// TrySet tries to set a value by request's form source (like map[string][]string)
func (form formSource) TrySet(value reflect.Value, field reflect.StructField, tagValue string, opt setOptions) (isSetted bool, err error) {
	return setByForm(value, field, form, tagValue, opt)
}
```
setByForm:
```golang
func setByForm(value reflect.Value, field reflect.StructField, form map[string][]string, tagValue string, opt setOptions) (isSetted bool, err error) {
	vs, ok := form[tagValue]
	if !ok && !opt.isDefaultExists {
		return false, nil
	}

	switch value.Kind() {
	case reflect.Slice:
		if !ok {
			vs = []string{opt.defaultValue}
		}
		return true, setSlice(vs, value, field)
	case reflect.Array:
		if !ok {
			vs = []string{opt.defaultValue}
		}
		if len(vs) != value.Len() {
			return false, fmt.Errorf("%q is not valid value for %s", vs, value.Type().String())
		}
		return true, setArray(vs, value, field)
	default:
		var val string
		if !ok {
			val = opt.defaultValue
		}

		if len(vs) > 0 {
			val = vs[0]
		}
		return true, setWithProperType(val, value, field)
	}
}
```


#### 项目中遇到的一个坑
curl -X GET 'localhost:8080'
```golang
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	s := gin.Default()
	s.GET("", testB)
	s.Run(":8080")
}

func testB(ctx *gin.Context) {
	in := struct {
		Name *uint
		Age  *uint
	}{}
	fmt.Printf("%+v", in)
	ctx.ShouldBindQuery(&in)
	fmt.Printf("%+v", in)
}
```

同样的一段代码,用gin的不同版本执行出来的结果是不一样的。
之前的版本打印结果是：
`{Name:<nil> Age:<nil>}{Name:0xc0000b65d8 Age:0xc0000b65e8}`
之后的版本打印结果是：
`{Name:<nil> Age:<nil>}{Name:<nil> Age:<nil>}`
是因为gin在19年3月份的时候对这里的实现进行了修改。
https://github.com/gin-gonic/gin/commit/0d50ce859745354fa83dcf2bf4c972abed25e53b#diff-e1ee2d6085c74d622d08bb3927e7036c

