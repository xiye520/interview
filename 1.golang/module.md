# Go Module

* [跳出Go module的泥潭](https://mp.weixin.qq.com/s?__biz=MzA4ODg0NDkzOA==&mid=2247487336&idx=1&sn=34ee7dbfa8b6352e87c4bb277240a5a2&source=41#wechat_redirect)

* [干货满满的 Go Modules 和 goproxy.cn](https://eddycjy.com/posts/go/go-moduels/2019-09-29-goproxy-cn/)

* [Go Modules 终极入门](https://eddycjy.com/posts/go/go-moduels/2020-02-28-go-modules/)

* [Golang modules 初探](https://my.oschina.net/qiangmzsx/blog/1934149)

* []()

* []()

* []()

* []()

* []()

* []()

* []()

* []()

## 什么是modules

现在都在说modules，那么它是什么？
到文档看看 [Modules, module versions, and more](https://tip.golang.org/cmd/go/#hdr-Modules__module_versions__and_more)：

```
A module is a collection of related Go packages. Modules are the unit of source code interchange and versioning. The go command has direct support for working with modules, including recording and resolving dependencies on other modules. Modules replace the old GOPATH-based approach to specifying which source files are used in a given build.
```

翻译一下：

```
模块是相关Go包的集合。modules是源代码交换和版本控制的单元。 go命令直接支持使用modules，包括记录和解析对其他模块的依赖性。modules替换旧的基于GOPATH的方法来指定在给定构建中使用哪些源文件。
```

可以得到两个重要信息：

- Go命令行支持modules操作
- modules用来替换GOPATH的

## 如何使用modules

modules是一个新的特性，那么就需要新的Golang版本进行支持了，可以到[官网](https://golang.org/dl/)下载，一定要是go 1.11及以上的版本（写博文的时候go 1.11刚刚出来）。 怎么部署就不在这里说了，相信初学者也是知道怎么做的。

还有人记得vendor刚刚出来时候golang提供的环境变量`GO15VENDOREXPERIMENT`吗？现在modules出来，按照惯例也提供了一个环境变量`GO111MODULE`，这个变量的三个1太有魔性了。

### GO111MODULE

`GO111MODULE`可以设置为三个字符串值之一：off，on或auto（默认值）。

- off，则go命令从不使用新模块支持。它查找vendor 目录和GOPATH以查找依赖关系;也就是继续使用“GOPATH模式”。
- on，则go命令需要使用模块，go 会忽略 GOPATH 和 vendor 文件夹，只根据 go.mod下载依赖。
- auto或未设置，则go命令根据当前目录启用或禁用模块支持。仅当当前目录位于GOPATH/src之外并且其本身包含go.mod文件或位于包含go.mod文件的目录下时，才启用模块支持。

### go mod 命令

`go mod`命令之前可以使用过了`go mod init`，下面我们把常用的`go mod`命令罗列一下：

- go mod init:初始化modules
- go mod download:下载modules到本地cache
- go mod edit:编辑go.mod文件，选项有-json、-require和-exclude，可以使用帮助go help mod edit
- go mod graph:以文本模式打印模块需求图
- go mod tidy:删除错误或者不使用的modules
- go mod vendor:生成vendor目录
- go mod verify:验证依赖是否正确
- go mod why：查找依赖

### go的 mod与get

go get这个命令大家应该不会陌生，这是下载go依赖包的根据，下载Go 1.11出来了，go get命令也与时俱进，支持了modules。 go get 来更新 module:

- 运行 go get -u 将会升级到最新的次要版本或者修订版本
- 运行 go get -u=patch 将会升级到最新的修订版本（比如说，将会升级到 1.0.1 版本，但不会升级到 1.1.0 版本）
- 运行 go get package@version将会升级到指定的版本号

运行go get如果有版本的更改，那么go.mod文件也会更改。



### 推 Go Modules 的人是谁

那么在上文中提到的 Russ Cox 何许人也呢，很多人应该都知道他，他是 Go 这个项目目前代码提交量最多的人，甚至是第二名的两倍还要多。

Russ Cox 还是 Go 现在的掌舵人（大家应该知道之前 Go 的掌舵人是 Rob Pike，但是听说由于他本人不喜欢特朗普执政所以离开了美国，然后他岁数也挺大的了，所以也正在逐渐交权，不过现在还是在参与 Go 的发展）。

Russ Cox 的个人能力相当强，看问题的角度也很独特，这也就是为什么他刚一提出 Go modules 的概念就能引起那么大范围的响应。虽然是被强推的，但事实也证明当下的 Go modules 表现得确实很优秀，所以这表明一定程度上的 “独裁” 还是可以接受的，至少可以保证一个项目能更加专一地朝着一个方向发展。

总之，无论如何 Go modules 现在都成了 Go 语言的一个密不可分的组件。

### go.mod

```
module example.com/foobar

go 1.13

require (
    example.com/apple v0.1.2
    example.com/banana v1.2.3
    example.com/banana/v2 v2.3.4
    example.com/pineapple v0.0.0-20190924185754-1b0db40df49a
)

exclude example.com/banana v1.2.4
replace example.com/apple v0.1.2 => example.com/rda v0.1.0 
replace example.com/banana => example.com/hugebanana
```

go.mod 是启用了 Go moduels 的项目所必须的最重要的文件，它描述了当前项目（也就是当前模块）的元信息，每一行都以一个动词开头，目前有以下 5 个动词:

- module：用于定义当前项目的模块路径。
- go：用于设置预期的 Go 版本。
- require：用于设置一个特定的模块版本。
- exclude：用于从使用中排除一个特定的模块版本。
- replace：用于将一个模块版本替换为另外一个模块版本。

这里的填写格式基本为包引用路径+版本号，另外比较特殊的是 `go $version`，目前从 Go1.13 的代码里来看，还只是个标识作用，暂时未知未来是否有更大的作用。

### go.sum

go.sum 是类似于比如 dep 的 Gopkg.lock 的一类文件，它详细罗列了当前项目直接或间接依赖的所有模块版本，并写明了那些模块版本的 SHA-256 哈希值以备 Go 在今后的操作中保证项目所依赖的那些模块版本不会被篡改。

### Global Caching

这个主要是针对 Go modules 的全局缓存数据说明，如下：

- 同一个模块版本的数据只缓存一份，所有其他模块共享使用。
- 目前所有模块版本数据均缓存在 `$GOPATH/pkg/mod`和 `$GOPATH/pkg/sum` 下，未来或将移至 `$GOCACHE/mod `和`$GOCACHE/sum` 下( 可能会在当 `$GOPATH` 被淘汰后)。
- 可以使用 `go clean -modcache` 清理所有已缓存的模块版本数据。

另外在 Go1.11 之后 GOCACHE 已经不允许设置为 off 了，我想着这也是为了模块数据缓存移动位置做准备，因此大家应该尽快做好适配。

## 快速迁移项目至 Go Modules

- 第一步: 升级到 Go 1.13。
- 第二步: 让 GOPATH 从你的脑海中完全消失，早一步踏入未来。
  - 修改 GOBIN 路径（可选）：`go env -w GOBIN=$HOME/bin`。
  - 打开 Go modules：`go env -w GO111MODULE=on`。
  - 设置 GOPROXY：`go env -w GOPROXY=https://goproxy.cn,direct` # 在中国是必须的，因为它的默认值被墙了。
- 第三步(可选): 按照你喜欢的目录结构重新组织你的所有项目。
- 第四步: 在你项目的根目录下执行 `go mod init ` 以生成 go.mod 文件。
- 第五步: 想办法说服你身边所有的人都去走一下前四步。

## 迁移后 go get 行为的改变

- 用 `go help module-get` 和 `go help gopath-get`分别去了解 Go modules 启用和未启用两种状态下的 go get 的行为

- 用

   

  ```
  go get
  ```

   

  拉取新的依赖

  - 拉取最新的版本(优先择取 tag)：`go get golang.org/x/text@latest`
  - 拉取 `master` 分支的最新 commit：`go get golang.org/x/text@master`
  - 拉取 tag 为 v0.3.2 的 commit：`go get golang.org/x/text@v0.3.2`
  - 拉取 hash 为 342b231 的 commit，最终会被转换为 v0.3.2：`go get golang.org/x/text@342b2e`
  - 用 `go get -u` 更新现有的依赖
  - 用 `go mod download` 下载 go.mod 文件中指明的所有依赖
  - 用 `go mod tidy` 整理现有的依赖
  - 用 `go mod graph` 查看现有的依赖结构
  - 用 `go mod init` 生成 go.mod 文件 (Go 1.13 中唯一一个可以生成 go.mod 文件的子命令)

- 用 `go mod edit` 编辑 go.mod 文件

- 用 `go mod vendor` 导出现有的所有依赖 (事实上 Go modules 正在淡化 Vendor 的概念)

- 用 `go mod verify` 校验一个模块是否被篡改过

这里我们注意到有两点比较特别，分别是：

- 第一点：为什么 “拉取 hash 为 342b231 的 commit，最终会被转换为 v0.3.2” 呢。这是因为虽然我们设置了拉取 @342b2e commit，但是因为 Go modules 会与 tag 进行对比，若发现对应的 commit 与 tag 有关联，则进行转换。
- 第二点：为什么不建议使用 `go mod vendor`，因为 Go modules 正在淡化 Vendor 的概念，很有可能 Go2 就去掉了。