前面我们介绍了使用GOPATH时的痛点：多项目无法共享同一个GOPATH。
其实本节介绍的vendor机制也没有彻底解决这个痛点，但是它提供了一个机制让项目的依赖隔离而不互相干扰。

自Go 1.6版本起，vendor机制正式启用，它允许把项目的依赖全部放到一个位于本项目的vendor目录中，这个vendor目录可以简单理解成私有的GOPATH目录。即编译时，优先从vendor中寻找依赖包，如果vendor中找不到再到GOPATH中寻找。

## vendor目录位置
一个项目可以有多个vendor目录，分别位于不同的目录级别，但建议每个项目只在根目录放置一个vendor目录。

假如你有一个`github.com/constabulary/example-gsftp`项目，项目目录结构如下：
```
$GOPATH
|	src/
|	|	github.com/constabulary/example-gsftp/
|	|	|	cmd/
|	|	|	|	gsftp/
|	|	|	|	|	main.go
```
其中 `main.go`中依赖如下几个包：
```
import (
	"golang.org/x/crypto/ssh"
	"github.com/pkg/sftp"
)
```

在没有使用vendor目录时，若想编译这个项目，那么GOPATH目录结构应该是如下所示：
```
$GOPATH
|	src/
|	|	github.com/constabulary/example-gsftp/
|	|	golang.org/x/crypto/ssh
|	|	github.com/pkg/sftp
```
即，所有依赖的包，都位于`$GOPATH/src`下。

为了把所使用到的`golang.org/x/crypto/ssh` 和 `github.com/pkg/sftp`版本固化下来，那么可以使用vendor机制。

在项目`github.com/constabulary/example-gsftp`根目录下，创建一个vendor目录，并把`golang.org/x/crypto/ssh` 和 `github.com/pkg/sftp`存放到此处，让其成为项目的一部分。如下所示：
```
$GOPATH
|	src/
|	|	github.com/constabulary/example-gsftp/
|	|	|	cmd/
|	|	|	|	gsftp/
|	|	|	|	|	main.go
|	|	|	vendor/
|	|	|	|	github.com/pkg/sftp/
|	|	|	|	golang.org/x/crypto/ssh/
```
使用vendor的好处是在项目`github.com/constabulary/example-gsftp`发布时，把其所依赖的软件一并发布，编译时不会受到GOPATH目录的影响，即便GOPATH下也有一个同名但不同版本的依赖包。

## 搜索顺序
上面的例子中，在编译main.go时，编译器搜索依赖包顺序为：
1. 从`github.com/constabulary/example-gsftp/cmd/gsftp/`下寻找vendor目录，没有找到，继续从上层查找；
2. 从`github.com/constabulary/example-gsftp/cmd/`下寻找vendor目录，没有找到，继续从上层查找；
3. 从`github.com/constabulary/example-gsftp/`下寻找vendor目录，从vendor目录中查找依赖包，结束；

如果`github.com/constabulary/example-gsftp/`下的vendor目录中没有依赖包，则返回到GOPATH目录继续查找，这就是前面介绍的GOPATH机制了。

从上面的搜索顺序可以看出，实际上vendor目录可以存在于项目的任意目录的。但非常不推荐这么做，因为如果vendor目录过于分散，很可能会出现同一个依赖包，在项目的多个vendor中出现多次，这样依赖包会多次编译进二进制文件，从而造成二进制大小急剧变大。同时，也很可能出现一个项目中使用同一个依赖包的多个版本的情况，这种情况往往应该避免。

## vendor存在的问题
vendor很好的解决了多项目间的隔离问题，但是位于vendor中的依赖包无法指定版本，某个依赖包，在把它放入vendor的那刻起，它就固定在当时版本，项目的使用者很难识别出你所使用的依赖版本。

比起这个，更严重的问题是上面提到的二进制急剧扩大问题，比如你依赖某个开源包A和B，但A中也有一个vendor目录，其中也放了B，那么你的项目中将会出现两个开源包B。再进一步，如果这两个开源包B版本不一致呢？如果二者不兼容，那后果将是灾难性的。

但是，不得不说，vendor能够解决绝大部分项目中的问题，如果你项目在使用vendor，也绝对没有问题。一直到Go 1.11版本，官方社区推出了Modules机制，从此Go的版本管理走进第三个时代。

> 赠人玫瑰手留余香，如果觉得不错请给个赞~
> 
> 本篇文章已归档到GitHub项目，求星~ [点我即达](https://github.com/RainbowMango/GoExpertProgramming)