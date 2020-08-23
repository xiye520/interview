# golang进阶

## 并发

并发的意义就是让一个程序同时做多件事情，其目的只是为了能让程序同时做另一件事情而已，而不是为了让程序运行的更快(如果是多核处理器，而且任务可以分成相互独立的部分，那么并发确实可以让事情解决的更快)。

golang从语言级别上对并发提供了支持，而且在启动并发的方式上直接添加了语言级的关键字，不必非要按照固定的格式来定义线程函数，也不必因为启动线程的时候只能给线程函数传递一个参数而烦恼。

goroutine 协程类似于线程，但是是更轻量的线程

## 协程间的同步与通信(sync.WaitGroup以及channel)

- 1、sync.WaitGroup
- 2、channel channel是一种golang内置的类型，英语的直译为"通道"，其实，它真的就是一根管道，而且是一个先进先出的数据结构。

我们能对channel进行的操作只有4种：

- (1) 创建chennel (通过make()函数)
- (2) 放入数据 (通过 channel <- data 操作)
- (3) 取出数据 (通过 <-channel 操作)
- (4) 关闭channel (通过close()函数)

------

channel的3种性质入如下：

- (1) channel是一种自动阻塞的管道。如果管道满了，一个对channel放入数据的操作就会阻塞，直到有某个routine从channel中取出数据，这个放入数据的操作才会执行。相反同理，如果管道是空的，一个从channel取出数据的操作就会阻塞，直到某个routine向这个channel中放入数据，这个取出数据的操作才会执行。这是channel最重要的一个性质！！！

------

- (2)channel分为有缓冲的channel和无缓冲的channel。两种channel的创建方法如下：

```
ch := make(chan int) 	//无缓冲的channel，同等于make(chan int, 0)
ch := make(chan int, 5) //一个缓冲区大小为5的channel
```

无缓冲通道与有缓冲通道的主要区别为：无缓冲通道存取数据是同步的，即如果通道中无数据，则通道一直处于阻塞状态；有缓冲通道存取数据是异步的，即存取数据互不干扰，只有当通道中已满时，存数据操作，通道阻塞；当通道中为空时，取数据操作，通道阻塞。

因此，使用无缓冲的channel时，放入操作和取出操作不能在同一个routine中，而且应该是先确保有某个routine对它执行取出操作，然后才能在另一个routine中执行放入操作，否则会发生死锁现象

使用带缓冲的channel时，因为有缓冲空间，所以只要缓冲区不满，放入操作就不会阻塞，同样，只要缓冲区不空，取出操作就不会阻塞。而且，带有缓冲的channel的放入和取出操作可以用在同一个routine中。但是，一定要注意放入和取出的速率问题，否则也会发生死锁现象

------

- (3)关闭后的channel可以取数据，但是不能放数据。而且，channel在执行了close()后并没有真的关闭，channel中的数据全部取走之后才会真正关闭。

## go调度 goroutine

[浅析goroutine调度器](https://tonybai.com/2017/06/23/an-intro-about-goroutine-scheduler/)

## golang ast词法分析

[golang 和 ast](http://xargin.com/ast/)

我们可以基于 ast 做很多静态分析、自动化和代码生成的事情

## channel实现

[Golang， 以17个简短代码片段，切底弄懂 channel 基础](https://www.cnblogs.com/linguanh/p/6248301.html#4192477)

## 三色标记

## 指针

Go语言提供了指针。指针是一种直接存储了变量的内存地址的数据类型。在其它语言中，比如C语言，指针操作是完全不受约束的。在另外一些语言中，指针一般被处理为“引用”，除了到处传递这些指针之外，并不能对这些指针做太多事情。Go语言在这两种范围中取了一种平衡。指针是可见的内存地址，&操作符可以返回一个变量的内存地址，并且*操作符可以获取指针指向的变量内容，但是在Go语言里没有指针运算，也就是不能像c语言里可以对指针进行加或减操作。

## 方法和接口

方法是和命名类型关联的一类函数。Go语言里比较特殊的是方法可以被关联到任意一种命名类型。接口是一种抽象类型，这种类型可以让我们以同样的方式来处理不同的固有类型，不用关心它们的具体实现，而只需要关注它们提供的方法。

## Map、slice、interface

参考博客： [golang: 常用数据类型底层结构分析](https://www.cnblogs.com/moodlxs/p/4133121.html)

基础类型

- 源码在：$GOROOT/src/pkg/runtime/runtime.h
- int8、uint8、int16、uint16、int32、uint32、int64、uint64、float32、float64分别对应于C的类型，这个只要有C基础就很容易看得出来。uintptr和intptr是无符号和有符号的指针类型，并且确保在64位平台上是8个字节，在32位平台上是4个字节，uintptr主要用于golang中的指针运算。而intgo和uintgo之所以不命名为int和uint，是因为int在C中是类型名，想必uintgo是为了跟intgo的命名对应吧。intgo和uintgo对应golang中的int和uint。从定义可以看出int和uint是可变大小类型的，在64位平台上占8个字节，在32位平台上占4个字节。所以如果有明确的要求，应该选择int32、int64或uint32、uint64。byte类型的底层类型是uint8
- 数据类型分为静态类型和底层类型，相对于以上代码中的变量b来说，byte是它的静态类型，uint8是它的底层类型。

------

rune类型

rune是int32的别名，用于表示unicode字符。通常在处理中文的时候需要用到它，当然也可以用range关键字。

------

string类型

string类型的底层是一个C struct。

```
struct String
{
        byte*   str;
        intgo   len;
};
```

成员str为字符数组，len为字符数组长度。golang的字符串是不可变类型，对string类型的变量初始化意味着会对底层结构的初始化。至于为什么str用byte类型而不用rune类型，这是因为golang的for循环对字符串的遍历是基于字节的，如果有必要，可以转成rune切片或使用range来迭代

内建函数len对string类型的操作是直接从底层结构中取出len值，而不需要额外的操作，当然在初始化时必需同时初始化len的值。

------

slice类型

slice类型的底层同样是一个C struct。

```
struct  Slice
{               // must not move anything
    byte*   array;      // actual data
    uintgo  len;        // number of elements
    uintgo  cap;        // allocated number of elements
}
```

包括三个成员。array为底层数组，len为实际存放的个数，cap为总容量。使用内建函数make对slice进行初始化，也可以类似于数组的方式进行初始化。当使用make函数来对slice进行初始化时，第一个参数为切片类型，第二个参数为len，第三个参数可选，如果不传入，则cap等于len。通常传入cap参数来预先分配大小的slice，避免频繁重新分配内存。

由于切片指向一个底层数组，并且可以通过切片语法直接从数组生成切片，所以需要了解切片和数组的关系，否则可能就会不知不觉的写出有bug的代码。

------

golang的map实现是hashtable，源码在：$GOROOT/src/pkg/runtime/hashmap.c

------

interface实际上是一个结构体，包括两个成员，一个是指向数据的指针，一个包含了成员的类型信息。Eface是interface{}底层使用的数据结构。因为interface中保存了类型信息，所以可以实现反射。反射其实就是查找底层数据结构的元数据。完整的实现在：$GOROOT/src/pkg/runtime/iface.c 。

## context

[Golang 并发 与 context标准库](https://mp.weixin.qq.com/s/FJLH4o7Y1TG9I0seiNwR_w)

context是一个很好的解决多goroutine下通知传递和元数据的Go标准库。由于Go中的goroutine之间没有父子关系，因此也不存在子进程退出后的通知机制。多个goroutine协调工作涉及 通信，同步，通知，退出 四个方面：

- 通信：chan通道是各goroutine之间通信的基础。注意这里的通信主要指程序的数据通道。
- 同步：可以使用不带缓冲的chan；sync.WaitGroup为多个gorouting提供同步等待机制；mutex锁与读写锁机制。
- 通知：通知与上文通信的区别是，通知的作用为管理，控制流数据。一般的解决方法是在输入端绑定两个chan，通过select收敛处理。这个方案可以解决简单的问题，但不是一个通用的解决方案。
- 退出：简单的解决方案与通知类似，即增加一个单独的通道，借助chan和select的广播机制(close chan to broadcast)实现退出。

但由于Go之间的goroutine都是平等的，因此当遇到复杂的并发结构时处理退出机制则会显得力不从心。因此Go1.7版本开始提供了context标准库来解决这个问题。他提供两个功能：退出通知和元数据传递。他们都可以传递给整个goroutine调用树的每一个goroutine。同时这也是一个不太复杂的，适合初学Gopher学习的一段源码。

## hashmap

[全网把Map中的hash()分析的最透彻的文章，别无二家](https://juejin.im/post/5ab99afff265da23a2291dee)

Hash，一般翻译做“散列”，也有直接音译为“哈希”的，就是把任意长度的输入，通过散列算法，变换成固定长度的输出，该输出就是散列值。这种转换是一种压缩映射，也就是，散列值的空间通常远小于输入的空间，不同的输入可能会散列成相同的输出，所以不可能从散列值来唯一的确定输入值。简单的说就是一种将任意长度的消息压缩到某一固定长度的消息摘要的函数。

所有散列函数都有如下一个基本特性：根据同一散列函数计算出的散列值如果不同，那么输入值肯定也不同。但是，根据同一散列函数计算出的散列值如果相同，输入值不一定相同。

两个不同的输入值，根据同一散列函数计算出的散列值相同的现象叫做碰撞。

## go逃逸

## malloc、锁、原子操作

## etcd

[Etcd超全解：原理阐释及部署设置的最佳实践](https://mp.weixin.qq.com/s?__biz=MzIyMTUwMDMyOQ==&mid=2247490456&idx=1&sn=6805596d29e299fae5c670efc126beb0&chksm=e83a9d5edf4d14480a1532ac5d89bbef8eaa76ba597563f09b4c998bc1e15b46ee53b171b08e&mpshare=1&scene=23&srcid=0222Hjcnsz3yR4EykfVKwSXL#rd)

## grpc、负载均衡LB

[gRPC服务发现&负载均衡](https://segmentfault.com/a/1190000008672912#articleHeader1)

## $GOOS and 和 $GOARCH

这两个环境变量表示目标代码的操作系统和CPU 类型。$GOOS 选项有linux 、 freebsd 、darwin (Mac OS X 10.5 or 10.6) 和 nacl (Chrome 的Native Client 接口，还未完成) 。$GOARCH 的 选项有amd64 (64-bit x86 ，目前最成熟) 、386 (32-bit x86) 、 和arm (32-bit ARM ，还未完成) 。下面是$GOOS 和 $GOARCH 的可能组合：

$GOARCH 和$GOOS 环境变量表示的是目标代码运行环境，和当前使用的平台是无关的。这个对于交叉编译是很方便的。

| $GOOS   | $GOARCH |
| ------- | ------- |
| darwin  | 386     |
| darwin  | amd64   |
| freebsd | 386     |
| freebsd | amd64   |
| linux   | 386     |
| linux   | amd64   |
| linux   | arm     |
| nacl    | 386     |

## golang plan9 汇编入门，带你打通应用和底层

[golang汇编入门知识分析博客](https://github.com/cch123/asmshare/blob/master/layout.md)

## 反射

###### 什么是反射？

反射就是程序能够在运行时检查变量和值，求出它们的类型。

- 清晰优于聪明。而反射并不是一目了然的。 --Rob Pike 关于使用反射的格言

###### 参考文章：

[Go 系列教程 —— 34. 反射](https://studygolang.com/articles/13178)

## Go 语言机制

本系列文章总共四篇，主要帮助理解 Go 语言中一些语法结构和其背后的设计原则，包括指针、栈、堆、逃逸分析和值/指针传递：

- [Go 语言机制（四篇）](https://studygolang.com/subject/74)
- [Go 语言机制之栈和指针](https://studygolang.com/articles/12443)
- [Go 语言机制之逃逸分析（Language Mechanics On Escape Analysis）](https://studygolang.com/subject/74)
- [Go 语言机制之内存剖析（Language Mechanics On Memory Profiling）](https://studygolang.com/articles/12445)
- [Go 语言机制之数据和语法的设计哲学（Design Philosophy On Data And Semantics）](https://studygolang.com/articles/12487)

## golang中字符串拼接方法

- +=
- fmt.sprintf
- append
- buffer.WriteString
- copy

效率：copy > append > buf.WriteString > += > fmt.Sprintf

[golang各种字符串拼接性能对比](https://www.jianshu.com/p/101755e7777c)

## Printf标准输出

Printf有一大堆这种转换，Go程序员称之为动词（verb）。下面的表格虽然远不是完整的规范，但展示了可用的很多特性：

- %d 十进制整数
- %x, %o, %b 十六进制，八进制，二进制整数。
- %f, %g, %e 浮点数： 3.141593 3.141592653589793 3.141593e+00
- %t 布尔：true或false
- %c 字符（rune） (Unicode码点)
- %s 字符串
- %q 带双引号的字符串"abc"或带单引号的字符'c'
- %v 变量的自然形式（natural format）
- %T 变量的类型
- %% 字面上的百分号标志（无操作数）