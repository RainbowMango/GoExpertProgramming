## 前言
可变参函数是指函数的某个参数可有可无，即这个参数个数可以是0个或多个。
声明可变参数函数的方式是在参数类型前加上`...`前缀。

比如`fmt`包中的`Println`:
```go
func Println(a ...interface{})
```

本节我们会总结一下其使用方法，顺便了解一下其原理，以避免在使用过程中进入误区。

## 函数特征
我们先写一个可变参函数：
```go
func Greeting(prefix string, who ...string) {
    if who == nil {
        fmt.Printf("Nobody to say hi.")
        return
    }

    for _, people := range who{
        fmt.Printf("%s %s\n", prefix, people)
    }
}
```
`Greeting`函数负责给指定的人打招呼，其参数`who`为可变参数。

这个函数几乎把可变参函数的特征全部表现出来了：
- 可变参数必须在函数参数列表的尾部，即最后一个（如放前面会引起编译时歧义）；
- 可变参数在函数内部是作为切片来解析的；
- 可变参数可以不填，不填时函数内部当成`nil`切片处理；
- 可变参数必须是相同类型的（如果需要是不同类型的可以定义为interface{}类型）；

## 使用举例
我们使用`testing`包中的Example函数来说明上面`Greeting`函数（函数位于sugar包中）用法。

### 不传值
调用可变参函数时，可变参部分是可以不传值的，例如：
```go
func ExampleGreetingWithoutParameter() {
    sugar.Greeting("nobody")
    // OutPut:
    // Nobody to say hi.
}
```
这里没有传递第二个参数。可变参数不传递的话，默认为nil。

### 传递多个参数
调用可变参函数时，可变参数部分可以传递多个值，例如：
```go
func ExampleGreetingWithParameter() {
    sugar.Greeting("hello:", "Joe", "Anna", "Eileen")
    // OutPut:
    // hello: Joe
    // hello: Anna
    // hello: Eileen
}
```
可变参数可以有多个。多个参数将会生成一个切片传入，函数内部按照切片来处理。

### 传递切片
调用可变参函数时，可变参数部分可以直接传递一个切片。参数部分需要使用`slice...`来表示切片。例如：
```go
func ExampleGreetingWithSlice() {
    guest := []string{"Joe", "Anna", "Eileen"}
    sugar.Greeting("hello:", guest...)
    // OutPut:
    // hello: Joe
    // hello: Anna
    // hello: Eileen
}
```
此时需要注意的一点是，切片传入时不会生成新的切片，也就是说函数内部使用的切片与传入的切片共享相同的存储空间。说得再直白一点就是，如果函数内部修改了切片，可能会影响外部调用的函数。

## 总结
- 可变参数必须要位于函数列表尾部；
- 可变参数是被当作切片来处理的；
- 函数调用时，可变参数可以不填；
- 函数调用时，可变参数可以填入切片；

> 赠人玫瑰手留余香，如果觉得不错请给个赞~
> 
> 本篇文章已归档到GitHub项目，求星~ [点我即达](https://github.com/RainbowMango/GoExpertProgramming)