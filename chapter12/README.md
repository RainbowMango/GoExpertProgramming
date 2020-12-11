Go语言依赖管理经历了三个重要的阶段：
- GOPATH；
- vendor；
- Go Module；

早期Go语言单纯使用GOPATH管理依赖，但GOPATH不方便管理依赖的多个版本，后来增加了vendor，允许把项目依赖连同项目源码一同管理。
自从Go 1.11版本引入了全新的依赖管理工具Go module，直到Go 1.14版本 Go module才走向成熟。

Go官方依赖管理演进过程中还有为数众多的第三方管理工具，比如`Glide`、`Govendor`等等，但随着Go module的推出，
这些工具终将逐步退出历史舞台，本章节不再涵盖此部分内容。

从GOPATH到vendor，再到Go module是个不断演进的过程，了解每种依赖管理的痛点可以更好的理解下一代依赖管理的设计初衷。
本章先从基础的GOPATH讲起，接着介绍vendor，最后再介绍Go module。
