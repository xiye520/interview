# 1.golang

​		要了解MPG调度、基于tcmalloc的内存分配方式、gc，channel、context、反射，切片、map底层原理，pprof性能分析调优；



* 1.向已经关闭的channel写数据**[**http://play.golang.org/p/vl5d5tkfl7**](http://play.golang.org/p/vl5d5tkfl7)
* 2.http://www.ituring.com.cn/article/37642  延伸阅读**
* 3.实验楼**[**https://www.shiyanlou.com/search?search=go**](https://www.shiyanlou.com/search?search=go)
* 4.golang的channel使用： http://my.oschina.net/goskyblue/blog/191149
* 5.https://www.zybuluo.com/Gestapo/note/32082#golang编程百例-bygestapo
* [【汇总】Go 面试每天一篇](https://mp.weixin.qq.com/s/rEXhrAqEOg9Ja4wYomOsGw)



* [Golang 协程Goroutine到底是怎么回事？（一）](https://studygolang.com/articles/27864)
* 知乎--谢大回答 [为什么要使用 Go 语言？Go 语言的优势在哪里？](https://www.zhihu.com/question/21409296/answer/18184584)
* 知乎--腾讯云技术社区的回答 [为什么要使用 Go 语言？Go 语言的优势在哪里？](https://www.zhihu.com/question/21409296/answer/421089971)
* [为什么很多公司转型 Go 语言开发？Go 语言能做什么](https://blog.csdn.net/qfliweimin/article/details/89027754)
* golang大端小端转换 [Golang binary包——byte数组如何转int？](https://blog.cyeam.com/hash/2014/07/29/go_bytearraytoint)
* [判断相等的deepequal](https://studygolang.com/articles/12944)



* [探索golang程序启动过程](https://cbsheng.github.io/posts/%E6%8E%A2%E7%B4%A2golang%E7%A8%8B%E5%BA%8F%E5%90%AF%E5%8A%A8%E8%BF%87%E7%A8%8B/)
* 码农桃花源 [Go 程序是怎样跑起来的](https://mp.weixin.qq.com/s?__biz=MjM5MDUwNTQwMQ==&mid=2257483812&idx=1&sn=3bc022cc699e24c0639e9ca6b321d552&chksm=a53918f2924e91e488c786c308353ee963df3e1bccb577bc9b03dd94f9551e4172401133becd&mpshare=1&scene=1&srcid=&from=groupmessage&ascene=1&devicetype=android-28&version=2700043c&nettype=ctnet&abtest_cookie=BQABAAoACwASABMAFQAFACOXHgBWmR4AyJkeAPiZHgAKmh4AAAA%3D&lang=zh_CN&pass_ticket=GgtiA7tB83Ck%2FG6zvWmWdU9u%2BWifTx49tZ9%2Bov8VqDBB%2BtFPLIPLNZ5rExmEkRER&wx_header=1)
* 滴滴技术 [详尽干货！从源码角度看 Golang 的调度](https://studygolang.com/articles/20651)
* [万字长文深入浅出 Golang Runtime](https://zhuanlan.zhihu.com/p/95056679)
* [Golang-gopark函数和goready函数原理分析](https://blog.csdn.net/u010853261/article/details/85887948)
* []()
* []()
* []()
* []()

## Go程序启动

```
// The bootstrap sequence is:
//
// call osinit
// call schedinit
// make & queue new G
// call runtime·mstart
//
// The new G calls runtime·main.
```

1. 做一些初始化的操作
2. 创建出一个goroutine结构 runtime.main 函数
3. 执行runtime.mstart 函数
4. 汇编引导结束，之后就由golang的函数main入口运行

初始化的时候，会创建几个线程（M）

1. sysmon特殊线程
2. 垃圾回收的线程

goroutine的主动切出

1. Gosched : 把当前G放入到队列中，然后切出
2. gopark/goparkunlock ： 保存上下文，直接切出
3. goready ： 唤醒G（把G重新入队）

gopark函数做的主要事情分为两点：

1. 解除当前goroutine的m的绑定关系，将当前goroutine状态机切换为等待状态；
2. 调用一次schedule()函数，在局部调度器P发起一轮新的调度。



## channal有哪几种，应用场景
## go启动过程
## go的反射
## go如何跨平台
go直接编译成本地码，不同的平台要重新编译，编译出来的程序能直接运行(C和C++也这样，但C/C++要自己管理内存)，不像java要有jvm才能运行，所以发布更简单，直接编译好的程序拷过去就能运行了

## golang的gc原理
## golang的内存分配原理
内存分配
Go提供了两种分配原语new与make：

func new(Type) *Type
func make(t Type, size ...IntegerType) Type
new(T)用于分配内存，它返回一个指针，指向新分配的，类型为T的零值，通过new来申请的内存都会被置零。这意味着如果设计了某种数据结构，那么每种类型的零值就不必进一步初始化了。

make(T,args)的目的不同于new(T)，它只用于创建切片（slice）、映射（map）、信道（channel），这三种类型本质上与引用数据类型，它们在使用前必须初始化。make返回类型为一个类型为T的已初始化的值，而非*T。

下面是new与make的对比：
```
var p *[]int = new([]int) // 分配切片结构；*p == nil；基本没用
var v []int = make([]int, 100) // 切片 v 现在引用了一个具有 100 个 int 元素的新数组
// 没必要的复杂：
var p *[]int = new([]int)
*p = make([]int, 100, 100)
// 习惯用法：
v := make([]int, 100)
```
## cgo是什么，为什么要有cgo



## 为何要用go语言，它有什么优势？
#### 1、Go有什么优势
* 可直接编译成机器码，不依赖其他库，glibc的版本有一定要求，部署就是扔一个文件上去就完成了。
* 静态类型语言，但是有动态语言的感觉，静态类型的语言就是可以在编译的时候检查出来隐藏的大多数问题，动态语言的感觉就是有很多的包可以使用，写起来的效率很高。
* 语言层面支持并发，这个就是Go最大的特色，天生的支持并发，我曾经说过一句话，天生的基因和整容是有区别的，大家一样美丽，但是你喜欢整容的还是天生基因的美丽呢？Go就是基因里面支持的并发，可以充分的利用多核，很容易的使用并发。
* 内置runtime，支持垃圾回收，这属于动态语言的特性之一吧，虽然目前来说GC不算完美，但是足以应付我们所能遇到的大多数情况，特别是Go1.1之后的GC。
* 简单易学，Go语言的作者都有C的基因，那么Go自然而然就有了C的基因，那么Go关键字是25个，但是表达能力很强大，几乎支持大多数你在其他语言见过的特性：继承、重载、对象等。
* 丰富的标准库，Go目前已经内置了大量的库，特别是网络库非常强大，我最爱的也是这部分。
* 内置强大的工具，Go语言里面内置了很多工具链，最好的应该是gofmt工具，自动化格式化代码，能够让团队review变得如此的简单，代码格式一模一样，想不一样都很困难。
* 跨平台编译，如果你写的Go代码不包含cgo，那么就可以做到window系统编译linux的应用，如何做到的呢？Go引用了plan9的代码，这就是不依赖系统的信息。
* 内嵌C支持，前面说了作者是C的作者，所以Go里面也可以直接包含c代码，利用现有的丰富的C库。

#### 2、Go适合用来做什么
* 服务器编程，以前你如果使用C或者C++做的那些事情，用Go来做很合适，例如处理日志、数据打包、虚拟机处理、文件系统等。
* 分布式系统，数据库代理器等
* 网络编程，这一块目前应用最广，包括Web应用、API应用、下载应用、
* 内存数据库，前一段时间google开发的groupcache，couchbase的部分组建
* 云平台，目前国外很多云平台在采用Go开发，CloudFoundy的部分组建，前VMare的技术总监自己出来搞的apcera云平台。

#### 3、Go成功的项目
* nsq：bitly开源的消息队列系统，性能非常高，目前他们每天处理数十亿条的消息
* docker:基于lxc的一个虚拟打包工具，能够实现PAAS平台的组建。
* packer:用来生成不同平台的镜像文件，例如VM、vbox、AWS等，作者是vagrant的作者
* skynet：分布式调度框架Doozer：分布式同步工具，类似ZooKeeper
* Heka：mazila开源的日志处理系统
* cbfs：couchbase开源的分布式文件系统
* tsuru：开源的PAAS平台，和SAE实现的功能一模一样
* groupcache：memcahe作者写的用于Google下载系统的缓存系统
* god：类似redis的缓存系统，但是支持分布式和扩展性
* gor：网络流量抓包和重放工具

#### 4、Go还存在的缺点以下缺点是我自己在项目开发中遇到的一些问题：
* Go的import包不支持版本，有时候升级容易导致项目不可运行，所以需要自己控制相应的版本信息
* Go的goroutine一旦启动之后，不同的goroutine之间切换不是受程序控制，runtime调度的时候，需要严谨的逻辑，不然goroutine休眠，过一段时间逻辑结束了，突然冒出来又执行了，会导致逻辑出错等情况。
* GC延迟有点大，我开发的日志系统伤过一次，同时并发很大的情况下，处理很大的日志，GC没有那么快，内存回收不给力，后来经过profile程序改进之后得到了改善。
* pkg下面的图片处理库很多bug，还是使用成熟产品好，调用这些成熟库imagemagick的接口比较靠谱

##

