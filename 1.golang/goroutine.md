## goroutine
* 每个 goroutine (协程) 默认占用内存远比 Java 、C 的线程少（goroutine：2KB ，线程：8MB）
* Rob Pike：“一个Goroutine是一个与其他goroutines 并发运行在同一地址空间的Go函数或方法。一个运行的程序由一个或更多个goroutine组成。它与线程、协程、进程等不同。它是一个goroutine。”


* Goroutine和其他语言的协程（coroutine）在使用方式上类似，但从字面意义上来看不同（一个是Goroutine，一个是coroutine），再就是协程是一种协作任务控制机制，在最简单的意义上，协程不是并发的，而Goroutine支持并发的。因此Goroutine可以理解为一种Go语言的协程。同时它可以运行在一个或多个线程上。

* goroutine不同于thread，threads是操作系统中的对于一个独立运行实例的描述，不同操作系统，对于thread的实现也不尽相同；但是，操作系统并不知道goroutine的存在，goroutine的调度是有Golang运行时进行管理的。启动thread虽然比process所需的资源要少，但是多个thread之间的上下文切换仍然是需要大量的工作的（寄存器/Program Count/Stack Pointer/...），Golang有自己的调度器，许多goroutine的数据都是共享的，因此goroutine之间的切换会快很多，启动goroutine所耗费的资源也很少，一个Golang程序同时存在几百个goroutine是很正常的。goroutine是Go语言中的轻量级线程实现，由Go运行时（runtime）管理.goroutine 比thread 更易用、更高效、更轻便。


#### Go线程实现模型MPG
M指的是Machine，一个M直接关联了一个内核线程。由操作系统管理。
P指的是”processor”，代表了M所需的上下文环境，也是处理用户级代码逻辑的处理器。它负责衔接M和G的调度上下文，将等待执行的G与M对接。
G指的是Goroutine，其实本质上也是一种轻量级的线程。包括了调用栈，重要的调度信息，例如channel等。

Go的调度器内部有三个重要的结构：M，G，P
M:代表真正的内核OS线程，和POSIX里的thread差不多，真正干活的人
G:代表一个goroutine，它有自己的栈，instruction pointer和其他信息（正在等待的channel等等），用于调度。
P:代表调度的上下文，可以把它看做一个局部的调度器，使go代码在一个线程上跑，它是实现从N:1到N:M映射的关键。

#### 参考资料
1. [goroutine理解，并发的实现原理、MPG调度模型](https://zhuanlan.zhihu.com/p/60613088)
