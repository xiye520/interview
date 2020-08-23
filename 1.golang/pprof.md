# golang pprof调优
* [Go性能调优](https://www.liwenzhou.com/posts/Go/performance_optimisation/)
* 这一篇超经典--[[教你如何找到Go内存泄露【精编实战】](https://www.cnblogs.com/sunsky303/p/11077030.html)](https://www.cnblogs.com/sunsky303/p/11077030.html)
* 煎鱼 [Go 大杀器之性能剖析 PProf](https://eddycjy.com/posts/go/tools/2018-09-15-go-tool-pprof/)
* 煎鱼 [Go 大杀器之跟踪剖析 trace](https://eddycjy.com/posts/go/tools/2019-07-12-go-tool-trace/)
* [如何获取运行的go程序的profile信息](https://studygolang.com/articles/12314)
* Dave.cheney [High Performance Go Workshop](https://dave.cheney.net/high-performance-go-workshop/dotgo-paris.html)
* 看云 [go tool pprof](https://www.kancloud.cn/cattong/go_command_tutorial/261357)
* [通过 profiling 定位 golang 性能问题 - 内存篇](https://www.infoq.cn/article/f69uvzJUOmq276HBp1Qb)
* []()
* []()
* []()
* []()
* []()

## 一、pprof功能

* CPU Profiling：CPU 分析，按照一定的频率采集所监听的应用程序 CPU（含寄存器）的使用情况，可确定应用程序在主动消耗 CPU 周期时花费时间的位置
* Memory Profiling：内存分析，在应用程序进行堆分配时记录堆栈跟踪，用于监视当前和历史内存使用情况，以及检查内存泄漏
* Block Profiling：阻塞分析，记录 goroutine 阻塞等待同步（包括定时器通道）的位置
* Mutex Profiling：互斥锁分析，报告互斥锁的竞争情况

#### pprof三大命令

> top：显示正运行到某个函数goroutine的数量
>
> traces：显示所有goroutine的调用栈
>
> list：列出代码详细的信息


#### 通过pprof可以得到什么

* cpu（CPU Profiling）: $HOST/debug/pprof/profile，默认进行 30s 的 CPU Profiling，得到一个分析用的 profile 文件
* block（Block Profiling）：$HOST/debug/pprof/block，查看导致阻塞同步的堆栈跟踪
* goroutine：$HOST/debug/pprof/goroutine，查看当前所有运行的 goroutines 堆栈跟踪
* heap（Memory Profiling）: $HOST/debug/pprof/heap，查看活动对象的内存分配情况
* mutex（Mutex Profiling）：$HOST/debug/pprof/mutex，查看导致互斥锁的竞争持有者的堆栈跟踪
* threadcreate：$HOST/debug/pprof/threadcreate，查看创建新 OS 线程的堆栈跟踪

#### 通过交互式终端使用

###### 1.查看cpu

go tool pprof http://localhost:6060/debug/pprof/profile?seconds=60

执行该命令后，需等待 60 秒（可调整 seconds 的值），pprof 会进行 CPU Profiling。结束后将默认进入 pprof 的交互式命令模式，可以对分析的结果进行查看或导出。具体可执行 pprof help 查看命令说明

```
(pprof) top10
Showing nodes accounting for 25.92s, 97.63% of 26.55s total
Dropped 85 nodes (cum <= 0.13s)
Showing top 10 nodes out of 21
      flat  flat%   sum%        cum   cum%
    23.28s 87.68% 87.68%     23.29s 87.72%  syscall.Syscall
     0.77s  2.90% 90.58%      0.77s  2.90%  runtime.memmove
     0.58s  2.18% 92.77%      0.58s  2.18%  runtime.freedefer
     0.53s  2.00% 94.76%      1.42s  5.35%  runtime.scanobject
     0.36s  1.36% 96.12%      0.39s  1.47%  runtime.heapBitsForObject
     0.35s  1.32% 97.44%      0.45s  1.69%  runtime.greyobject
     0.02s 0.075% 97.51%     24.96s 94.01%  main.main.func1
     0.01s 0.038% 97.55%     23.91s 90.06%  os.(*File).Write
     0.01s 0.038% 97.59%      0.19s  0.72%  runtime.mallocgc
     0.01s 0.038% 97.63%     23.30s 87.76%  syscall.Write

如果我们在top命令后加入标签--cum，那么输出的列表就是以累积取样计数为顺序的。
(pprof) top --cum
Showing nodes accounting for 2.95s, 64.41% of 4.58s total
Dropped 34 nodes (cum <= 0.02s)
Showing top 10 nodes out of 42
      flat  flat%   sum%        cum   cum%
     0.01s  0.22%  0.22%      4.43s 96.72%  runtime.systemstack
         0     0%  0.22%      4.04s 88.21%  runtime.gcBgMarkWorker
         0     0%  0.22%      4.01s 87.55%  runtime.gcBgMarkWorker.func2
     0.11s  2.40%  2.62%      4.01s 87.55%  runtime.gcDrain
     1.92s 41.92% 44.54%      3.57s 77.95%  runtime.scanobject
     0.61s 13.32% 57.86%      0.71s 15.50%  runtime.heapBitsForObject
     0.29s  6.33% 64.19%      0.59s 12.88%  runtime.greyobject
         0     0% 64.19%      0.42s  9.17%  runtime.bgsweep
     0.01s  0.22% 64.41%      0.42s  9.17%  runtime.gosweepone
         0     0% 64.41%      0.41s  8.95%  runtime.gosweepone.func1

(pprof) top20 --cum -fmt\..* -os\..*

```

* flat：给定函数上运行耗时
* flat%：同上的 CPU 运行耗时总比例
* sum%：给定函数累积使用 CPU 总比例
* cum：当前函数加上它之上的调用运行总耗时
* cum%：同上的 CPU 运行耗时总比例
  最后一列为函数名称，在大多数的情况下，我们可以通过这五列得出一个应用程序的运行情况，加以优化

###### 2.查看内存

 go tool pprof http://localhost:6060/debug/pprof/heap

* -inuse_space：分析应用程序的常驻内存占用情况
* -alloc_objects：分析应用程序的内存临时分配情况

3.对比文件，可对比前后两次的损耗：

```
$ go tool pprof -base xauthority_20200516-023226.pprof xauthority_20200516-235856.pprof
Local symbolization failed for xauthority: open /opt/data/xrocket/package/xauthority/xauthority-2.3.77-release/bin/xauthority: no such file or directory
Some binary filenames not available. Symbolization may be incomplete.
Try setting PPROF_BINARY_PATH to the search path for local binaries.
File: xauthority
Type: cpu
Time: May 16, 2020 at 2:32am (CST)
Duration: 10.14s, Total samples = 19.96s (196.85%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 14780ms, 74.05% of 19960ms total
Dropped 26 nodes (cum <= 99.80ms)
Showing top 10 nodes out of 60
      flat  flat%   sum%        cum   cum%
    3130ms 15.68% 15.68%     6690ms 33.52%  runtime.pcvalue
    2570ms 12.88% 28.56%     3070ms 15.38%  runtime.step
    2360ms 11.82% 40.38%     5320ms 26.65%  runtime.scanblock
    1880ms  9.42% 49.80%    17430ms 87.32%  runtime.gentraceback
    1370ms  6.86% 56.66%     1510ms  7.57%  runtime.heapBitsForObject
     990ms  4.96% 61.62%    11070ms 55.46%  runtime.scanframeworker
     740ms  3.71% 65.33%      880ms  4.41%  runtime.findfunc
     720ms  3.61% 68.94%     1360ms  6.81%  runtime.greyobject
     520ms  2.61% 71.54%      520ms  2.61%  runtime.epollwait
     500ms  2.51% 74.05%      500ms  2.51%  runtime.readvarint
(pprof)
```



## 二、go tool trace

单单使用 PProf 有时候不一定足够完整，因为在真实的程序中还包含许多的隐藏动作，例如 Goroutine 在执行时会做哪些操作？执行/阻塞了多长时间？在什么时候阻止？在哪里被阻止的？谁又锁/解锁了它们？GC 是怎么影响到 Goroutine 的执行的？这些东西用 PProf 是很难分析出来的，但如果你又想知道上述的答案的话，你可以用` go tool trace` 来打开新世界的大门。

```go
import (
	"os"
	"runtime/trace"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	ch := make(chan string)
	go func() {
		ch <- "EDDYCJY"
	}()

	<-ch
}
```

生成跟踪文件：

```
$ go run main.go 2> trace.out
```

启动可视化界面：

```
$ go tool trace trace.out
```

* View trace：查看跟踪
* Goroutine analysis：Goroutine 分析
* Network blocking profile：网络阻塞概况
* Synchronization blocking profile：同步阻塞概况
* Syscall blocking profile：系统调用阻塞概况
* Scheduler latency profile：调度延迟概况
* User defined tasks：用户自定义任务
* User defined regions：用户自定义区域
* Minimum mutator utilization：最低 Mutator 利用率

在刚开始查看问题时，除非是很明显的现象，否则不应该一开始就陷入细节，因此我们一般先查看 “Scheduler latency profile”，我们能通过 Graph 看到整体的调用开销情况


## 三、获取内存分配解读

```
curl http://localhost:9999/debug/pprof/heap > mem.prof
```

上一句是将内存的分配信息保存到文件中

```
go tool pprof mem.prof
```

这一句是查看分配之后还没有释放的内存（in-use memory）：要么正在使用，要么确实没有及时释放；这对于检查内存泄漏很有帮助。

```
go tool pprof -alloc_space mem.prof
```

通过这个命令我们可以查看分配历史的统计，知道什么地方分配内存过于频繁，是否可以复用。

```
top N
```

该子命令会列出N处分配内存最多的代码所在的函数，在输入top命令前输入cum或flat可以使得top列出的列表按cum或flat列排序，flat是单个函数自身（不计函数里面调用的其它函数）的内存占用，cum是函数总的占用。

如果要看某个函数具体是在什么地方分配的内存，可以使用list子命令查看：

```
list func_name_in_top_list
```

该命令会显示分配内存的代码和行号。

## Dave讲了以下几点：

内存profiling记录的是堆内存分配的情况，以及调用栈信息，并不是进程完整的内存情况，猜测这也是在go pprof中称为heap而不是memory的原因。
栈内存的分配是在调用栈结束后会被释放的内存，所以并不在内存profile中。
内存profiling是基于抽样的，默认是每1000次堆内存分配，执行1次profile记录。
因为内存profiling是基于抽样和它跟踪的是已分配的内存，而不是使用中的内存，（比如有些内存已经分配，看似使用，但实际以及不使用的内存，比如内存泄露的那部分），所以不能使用内存profiling衡量程序总体的内存使用情况。
Dave个人观点：使用内存profiling不能够发现内存泄露。

## 总结

#### goroutine泄露的本质

goroutine泄露的本质是channel阻塞，无法继续向下执行，导致此goroutine关联的内存都无法释放，进一步造成内存泄露。

#### goroutine泄露的发现和定位

利用好go pprof获取goroutine profile文件，然后利用3个命令top、traces、list定位内存泄露的原因。

#### goroutine泄露的场景

泄露的场景不仅限于以下两类，但因channel相关的泄露是最多的。

-  1.channel的读或者写：
   - 1.无缓冲channel的阻塞通常是写操作因为没有读而阻塞
   - 2.有缓冲的channel因为缓冲区满了，写操作阻塞
   - 3.期待从channel读数据，结果没有goroutine写
-  2.select操作，select里也是channel操作，如果所有case上的操作阻塞，goroutine也无法继续执行。

#### 编码goroutine泄露的建议

为避免goroutine泄露造成内存泄露，启动goroutine前要思考清楚：

    1.goroutine如何退出？
    2.是否会有阻塞造成无法退出？如果有，那么这个路径是否会创建大量的goroutine？



## 四、Go语言项目中的性能优化主要有以下几个方面：
```
CPU profile：报告程序的 CPU 使用情况，按照一定频率去采集应用程序在 CPU 和寄存器上面的数据
Memory Profile（Heap Profile）：报告程序的内存使用情况
Block Profiling：报告 goroutines 不在运行状态的情况，可以用来分析和查找死锁等性能瓶颈
Goroutine Profiling：报告 goroutines 的使用情况，有哪些 goroutine，它们的调用关系是怎样的
```
#### 1. 采集性能数据
Go语言内置了获取程序的运行数据的工具，包括以下两个标准库：

* runtime/pprof：采集工具型应用运行数据进行分析
* net/http/pprof：采集服务型应用运行时数据进行分析

pprof开启后，每隔一段时间（10ms）就会收集下当前的堆栈信息，获取格格函数占用的CPU以及内存资源；最后通过对这些采样数据进行分析，形成一个性能分析报告。

注意，我们只应该在性能测试的时候才在代码中引入pprof。

## 2. 工具型应用
如果你的应用程序是运行一段时间就结束退出类型。那么最好的办法是在应用退出的时候把 profiling 的报告保存到文件中，进行分析。对于这种情况，可以使用runtime/pprof库。 首先在代码中导入runtime/pprof工具：
```
import "runtime/pprof"
```
#### CPU性能分析
开启CPU性能分析：
```
pprof.StartCPUProfile(w io.Writer)
```

停止CPU性能分析：
```
pprof.StopCPUProfile()
```
应用执行结束后，就会生成一个文件，保存了我们的 CPU profiling 数据。得到采样数据之后，使用go tool pprof工具进行CPU性能分析。

#### 内存性能优化
记录程序的堆栈信息
```
pprof.WriteHeapProfile(w io.Writer)
```
得到采样数据之后，使用go tool pprof工具进行内存性能分析。

go tool pprof默认是使用-inuse_space进行统计，还可以使用-inuse-objects查看分配对象的数量。

## 3. 服务型应用
如果你的应用程序是一直运行的，比如 web 应用，那么可以使用net/http/pprof库，它能够在提供 HTTP 服务进行分析。
* /debug/pprof/profile：访问这个链接会自动进行 CPU profiling，持续 30s，并生成一个文件供下载
* /debug/pprof/heap： Memory Profiling 的路径，访问这个链接会得到一个内存 Profiling 结果的文件
* /debug/pprof/block：block Profiling 的路径
* /debug/pprof/goroutines：运行的 goroutines 列表，以及调用关系

## 4. go tool pprof命令
不管是工具型应用还是服务型应用，我们使用相应的pprof库获取数据之后，下一步的都要对这些数据进行分析，我们可以使用go tool pprof命令行工具。

go tool pprof最简单的使用方式为:
```
go tool pprof [binary] [source]
```
其中:
* binary 是应用的二进制文件，用来解析各种符号；
* source 表示 profile 数据的来源，可以是本地的文件，也可以是 http 地址。

注意事项： 获取的 Profiling 数据是动态的，要想获得有效的数据，请保证应用处于较大的负载（比如正在生成中运行的服务，或者通过其他工具模拟访问压力）。否则如果应用处于空闲状态，得到的结果可能没有任何意义。


## 5. pprof与性能测试结合
go test命令有两个参数和 pprof 相关，它们分别指定生成的 CPU 和 Memory profiling 保存的文件：
```
-cpuprofile：cpu profiling 数据要保存的文件地址
-memprofile：memory profiling 数据要报文的文件地址
```
我们还可以选择将pprof与性能测试相结合，比如：

```
比如下面执行测试的同时，也会执行 CPU profiling，并把结果保存在 cpu.prof 文件中：
go test -bench . -cpuprofile=cpu.prof
```

```
比如下面执行测试的同时，也会执行 Mem profiling，并把结果保存在 cpu.prof 文件中：
go test -bench . -memprofile=./mem.prof
```
需要注意的是，Profiling 一般和性能测试一起使用，这个原因在前文也提到过，只有应用在负载高的情况下 Profiling 才有意义。


## 单独测内存
你可以通过下面的方式产生两个时间点的堆的profile,之后使用pprof工具进行分析。
* 1.首先确保你已经配置了pprof的http路径， 可以访问http://ip:port/debug/pprof/查看(如果你没有修改默认的pprof路径)
* 2.导出时间点1的堆的profile: curl -s http://127.0.0.1:8080/debug/pprof/heap > base.heap, 我们把它作为基准点
* 3.喝杯茶，等待一段时间后导出时间点2的堆的profile: curl -s http://127.0.0.1:8080/debug/pprof/heap > current.heap
* 4.现在你就可以比较这两个时间点的堆的差异了: go tool pprof --base base.heap current.heap

操作和正常的go tool pprof操作一样， 比如使用top查看使用堆内存最多的几处地方的内存增删情况：



## 五、实践 Tips

以下是一些从其它项目借鉴或者自己总结的实践经验，它们只是建议，而不是准则，实际项目中应该以性能分析数据来作为优化的参考，避免过早优化。

* 1.对频繁分配的对象，使用 sync.Pool 对象池减少分配时GC压力
* 2.自动化的 DeepCopy 是非常耗时的，其中涉及到反射，内存分配，容器(如 map)扩展等，大概比手动拷贝慢一个数量级
* 3.用 atomic.Load/StoreXXX，atomic.Value, sync.Map 等代替 Mutex。(优先级递减)
* 4.使用高效的第三方库，如用fasthttp替代 net/http
* 5.在开发环境加上-race编译选项进行竞态检查
* 6.在开发或线上环境开启 net/http/pprof，方便实时pprof
* 7.将所有外部IO(网络IO，磁盘IO)做成异步
