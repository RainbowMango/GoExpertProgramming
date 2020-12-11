Golang 内置方法`append`用于向切片中追加一个或多个元素，实际项目中比较常见。

其原型如下：
```golang
func append(slice []Type, elems ...Type) []Type
```

本节不会对`append`的使用方式详细展开，而是重点介绍几个使用中常见的误区或者陷阱。

## 热身

按照惯例，我们先拿几个小题目来检测一下对`append`的理解是否足够深刻。

### 题目一
函数`Validation()`用于一些合法性检查，每遇到一个错误，就生成一个新的`error`并追加到切片`errs`中，
最后返回包含所有错误信息的切片。
为了简单起见，假定函数发现了三个错误，如下所示：
```golang
func Validation() []error {
	var errs []error

	append(errs, errors.New("error 1"))
	append(errs, errors.New("error 2"))
	append(errs, errors.New("error 3"))

	return errs
}
```

请问函数`Validation()`有什么问题？

### 题目二
函数`ValidateName()`用于检查某个名字是否合法，如果不为空则认为合法，否则返回一个`error`。
类似的，还可以有很多检查项，比如检查性别、年龄等，我们统称为子检查项。
函数`Validations()`用于收集所有子检查项的错误信息，将错误信息汇总到一个切片中返回。

请问函数`Validations()`有什么问题？

```golang
func ValidateName(name string) error {
	if name != "" {
		return nil
	}

	return errors.New("empty name")
}

func Validations(name string) []error {
	var errs []error

	errs = append(errs, ValidateName(name))

	return errs
}
```

## 陷阱

前面的热身题目均来源于实际项目（已经做了最大程度的精简），分别代表一个本节将要介绍的陷阱。

### 陷阱一： append 会改变切片的地址

`append`的本质是向切片中追加数据，而随着切片中元素逐渐增加，当切片底层的数组将满时，切片会发生扩容，
扩容会导致产生一个新的切片（拥有容量更大的底层数组），更多关于切片的信息，请查阅切片相关章节。

`append`每个追加元素，都有可能触发切片扩容，也即有可能返回一个新的切片，这也是`append`函数声明中返回值为切片的原因。实际使用中应该总是接收该返回值。

上述题目一中，由于初始切片长度为0，所以实际上每次`append`都会产生一个新的切片并迅速抛弃（被gc回收）。
原始切片并没有任何改变。需要特别说明的是，不管初始切片长度为多少，不接收`append`返回都是有极大风险的。

另外，目前有很多的工具可以自动检查出类似的问题，比如`Goland`IDE就会给出很明显的提示。

### 陷阱二： append 可以追加nil值

向切片中追加一个`nil`值是完全不会报错的，如下代码所示：
```
slice := append(slice, nil)
```
经过追加后，slice的长度递增1。

实际上`nil`是一个预定义的值，即空值，所以完全有理由向切片中追加。

上述题目二中，就是典型的向切片中追加`nil`（当名字为空时）的问题。单纯从技术上讲是没有问题，但在题目二场景中就有很大的问题。

题目中函数用于收集所有错误信息，没有错误就不应该追加到切片中。因后，后续极有可能会根据切片的长度来判断是否有错误发生，比如：

```golang
func foo() {
	errs := Validations("")
	
	if len(errs) > 0 {
		println(errs)
		os.Exit(1)
	}
}
```
如果向切片中追加一个`nil`元素，那么切片长度则不再为0，程序很可能因此而退出，更糟糕的是，这样的切片是没有内容会打印出来的，这无疑又增加了定位难度。

> 赠人玫瑰手留余香，如果觉得不错请给个赞~
> 你的鼓励将成为我继续写作的动力！ 
>
> 本篇文章已归档到GitHub项目，求星~ [点我即达](https://github.com/RainbowMango/GoExpertProgramming)