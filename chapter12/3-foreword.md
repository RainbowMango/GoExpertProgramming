在Go v1.11版本中，Module特性被首次引入，这标志着Go的依赖管理开始进入第三个阶段。

Go Module 相比GOPATH和vendor而言功能强大得多，它基本上完全解决了GOPATH和vendor时代遗留的问题。
我们知道，GOPATH时代最大的困扰是无法让多个项目共享同一个pakage的不同版本，在vendor时代，通过把每个项目依赖的package放到vendor
中可以解决这个困扰，但是使用vendor的问题是无法很好的管理依赖的package，比如升级package。

虽然Go Module能够解决GOPATH和vendor时代遗留的问题，但需要注意的是Go Module不是GOPATH 和 vendor的演进，理解这个对于接下来
正确理解Go Module非常重要。

Go Module更像是一种全新的依赖管理方案，它涉及一系列的特性，但究其核心，它主要解决两个重要的问题：
- 准确的记录项目依赖；
- 可重复的构建；

准确的记录项目依赖，是指你的项目依赖哪些package、以及package的版本可以非常精确。比如你的项目依赖`github.com/prometheus/client_golang`,
且必须是`v1.0.0`版本，那么你可以通过Go Module指定（具体指定方法后面会介绍），任何人在任何环境下编译你的项目，
都必须要使用`github.com/prometheus/client_golang`的`v1.0.0`版本。

可重复的构建是指，项目无论在谁的环境中（同平台）构建，其产物都是相同的。回想一下GOPATH时代，虽然大家拥有同一个项目的代码，但由于各自
的GOPATH中`github.com/prometheus/client_golang`版本不一样，虽然项目可以构建，但构建出的可执行文件很可能是不同的。
可重复构建至关重要，避免出现“我这运行没问题，肯定是你环境问题”等类似问题出现。

一旦项目的依赖被准确记录了，就很容易做到重复构建。

事实上，Go Module是一个非常复杂的特性，一下子全盘托出其特性，往往会让人产生疑惑，所以接下来的章节，我们希望逐个介绍其特性，
并且，尽可以附以实例，希望大家也跟我一样手动实践一下，以加深认识。
