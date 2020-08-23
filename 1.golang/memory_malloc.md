# 此篇讲的go内存分配更通俗易懂

* [简单易懂的 Go 内存分配原理解读](https://yq.aliyun.com/articles/652551)
* 《Go 语言设计与实现》 [7.1 内存分配器](https://draveness.me/golang/docs/part3-runtime/ch07-memory/golang-memory-allocator/)
* [图解Go语言内存分配](https://juejin.im/post/5c888a79e51d456ed11955a8)
* []()

##  前言

编写过C语言程序的肯定知道通过malloc()方法动态申请内存，其中内存分配器使用的是glibc提供的ptmalloc2。
除了glibc，业界比较出名的内存分配器有Google的tcmalloc和Facebook的jemalloc。二者在避免内存碎片和性能上均比glic有比较大的优势，在多线程环境中效果更明显。

Golang中也实现了内存分配器，原理与tcmalloc类似，简单的说就是维护一块大的全局内存，每个线程(Golang中为P)维护一块小的私有内存，私有内存不足再从全局申请。

另外，内存分配与GC（垃圾回收）关系密切，所以了解GC前有必要了解内存分配的原理。

## 1. 基础概念
为了方便自主管理内存，做法便是先向系统申请一块内存，然后将内存切割成小块，通过一定的内存分配算法管理内存。
以64位系统为例，Golang程序启动时会向系统申请的内存如下图所示：


预申请的内存划分为spans、bitmap、arena三部分。其中arena即为所谓的堆区，应用中需要的内存从这里分配。其中spans和bitmap是为了管理arena区而存在的。

arena的大小为512G，为了方便管理把arena区域划分成一个个的page，每个page为8KB,一共有512GB/8KB个页；

spans区域存放span的指针，每个指针对应一个page，所以span区域的大小为(512GB/8KB)*指针大小8byte = 512M

bitmap区域大小也是通过arena计算出来，不过主要用于GC。

* arena 是 Golang 中用于分配内存的连续虚拟地址区域。由 mheap 管理，堆上申请的所有内存都来自 arena。

* bitmap 用两个 bit 表示一个字的可用状态，那么算下来 bitmap 的大小为 16 G。读过 Golang 源码的同学会发现其实这段代码的注释里写的 bitmap 的大小为 32 G。其实是这段注释很久没有更新了，之前是用 4 个 bit 来表示一个字的可用状态，这真是一个悲伤的故事啊。

* spans 记录的 arena 区的块页号和对应的 mspan 指针的对应关系。比如 arena 区内存地址 x，对应的页号就是 page_num = (x - arena_start) / page_size，那么 spans 就会记录 spans[page_num] = x。如果 arena 为 512 G的话，spans 区的大小为 512 G / 8K * 8 = 512 M。这里值得注意的是 Golang 的内存管理虚拟地址页大小为 8k。

#### 2. 内存数据结构

为了方便自主管理内存，做法便是先向系统申请一块内存，然后将内存切割成小块，通过一定的内存分配算法管理内存。
以64位系统为例，Golang程序启动时会向系统申请的内存如下图所示：

![img](https://oscimg.oschina.net/oscnet/8c11a0048d1b9bff76199de7f14ebf8a9c0.jpg)

预申请的内存划分为spans、bitmap、arena三部分。其中arena即为所谓的堆区，应用中需要的内存从这里分配。其中spans和bitmap是为了管理arena区而存在的。

arena的大小为512G，为了方便管理把arena区域划分成一个个的page，每个page为8KB,一共有512GB/8KB个页；

spans区域存放span的指针，每个指针对应一个page，所以span区域的大小为(512GB/8KB)*指针大小8byte = 512M

bitmap区域大小也是通过arena计算出来，不过主要用于GC。

### 2.1 span
span是用于管理arena页的关键数据结构，每个span中包含1个或多个连续页，为了满足小对象分配，span中的一页会划分更小的粒度，而对于大对象比如超过页大小，则通过多页实现。

span划分了67种小对象类型，最大的对象是32K大小，超过32K大小的由特殊的class表示，该class ID为0，每个class只包含一个对象。

#### 2.1.1 class
跟据对象大小，划分了一系列class，每个class都代表一个固定大小的对象，以及每个span的大小。

mspan的Size Class共有67种，每种mspan分割的object大小是8*2n的倍数，这个是写死在代码里的
```
// path: /usr/local/go/src/runtime/sizeclasses.go

const (
	_MaxSmallSize   = 32768 //最大32KB
	smallSizeDiv    = 8
	smallSizeMax    = 1024
	largeSizeDiv    = 128
	_NumSizeClasses = 67
	_PageShift      = 13
)

var class_to_size = [_NumSizeClasses]uint16{0, 8, 16, 32, 48, 64, 80, 96, 112, 128, 144, 160, 176, 192, 208, 224, 240, 256, 288, 320, 352, 384, 416, 448, 480, 512, 576, 640, 704, 768, 896, 1024, 1152, 1280, 1408, 1536, 1792, 2048, 2304, 2688, 3072, 3200, 3456, 4096, 4864, 5376, 6144, 6528, 6784, 6912, 8192, 9472, 9728, 10240, 10880, 12288, 13568, 14336, 16384, 18432, 19072, 20480, 21760, 24576, 27264, 28672, 32768}

```


#### 2.1.2 span数据结构
span是内存管理的基本单位,每个span用于管理特定的class对象, 跟据对象大小，span将一个或多个页拆分成多个块进行管理。

src/runtime/mheap.go:mspan定义了其数据结构：
```
// path: /usr/local/go/src/runtime/mheap.go

type mspan struct {
	next *mspan			//链表前向指针，用于将span链接起来
	prev *mspan			//链表前向指针，用于将span链接起来
	startAddr uintptr // 起始地址，也即所管理页的地址
	npages    uintptr // 管理的页数
	
	nelems uintptr // 块个数，也即有多少个块可供分配

	allocBits  *gcBits //分配位图，每一位代表一个块是否已分配

	allocCount  uint16     // 已分配块的个数
	spanclass   spanClass  // class表中的class ID

	elemsize    uintptr    // class表中的对象大小，也即块大小
}
```

## 3. 内存分配过程
针对待分配对象的大小不同有不同的分配逻辑：

```
(0, 16B) 且不包含指针的对象： Tiny分配
(0, 16B) 包含指针的对象：正常分配
[16B, 32KB] : 正常分配
(32KB, -) : 大对象分配 其中Tiny分配和大对象分配都属于内存管理的优化范畴，这里暂时仅关注一般的分配方法。
```
以申请size为n的内存为例，分配步骤如下：
```
获取当前线程的私有缓存mcache
跟据size计算出适合的class的ID
从mcache的alloc[class]链表中查询可用的span
如果mcache没有可用的span则从mcentral申请一个新的span加入mcache中
如果mcentral中也没有可用的span则从mheap中申请一个新的span加入mcentral
从该span中获取到空闲对象地址并返回
```

#### 3.1内存管理3大组件：mcache、mcentral、mheap
###### 3.1.1 mcache
每个工作线程都会绑定一个mcache，本地缓存可用的mspan资源，这样就可以直接给Goroutine分配，因为不存在多个Goroutine竞争的情况，所以不会消耗锁资源。

mcache的结构体定义：
```
//path: /usr/local/go/src/runtime/mcache.go

type mcache struct {
    alloc [numSpanClasses]*mspan
}

numSpanClasses = _NumSizeClasses << 1
```
复制代码mcache用Span Classes作为索引管理多个用于分配的mspan，它包含所有规格的mspan。它是_NumSizeClasses的2倍，也就是67*2=134，为什么有一个两倍的关系，前面我们提到过：为了加速之后内存回收的速度，数组里一半的mspan中分配的对象不包含指针，另一半则包含指针。

对于无指针对象的mspan在进行垃圾回收的时候无需进一步扫描它是否引用了其他活跃的对象。

mcache在初始化的时候是没有任何mspan资源的，在使用过程中会动态地从mcentral申请，之后会缓存下来。当对象小于等于32KB大小时，使用mcache的相应规格的mspan进行分配。

###### 3.1.2 mcentral
mcentral：为所有mcache提供切分好的mspan资源。每个central保存一种特定大小的全局mspan列表，包括已分配出去的和未分配出去的。 每个mcentral对应一种mspan，而mspan的种类导致它分割的object大小不同。当工作线程的mcache中没有合适（也就是特定大小的）的mspan时就会从mcentral获取。

mcentral被所有的工作线程共同享有，存在多个Goroutine竞争的情况，因此会消耗锁资源。结构体定义：
```
//path: /usr/local/go/src/runtime/mcentral.go

type mcentral struct {
	lock      mutex     //互斥锁
	spanclass spanClass // span class ID
	nonempty  mSpanList // non-empty 指还有空闲块的span列表
	empty     mSpanList // 指没有空闲块的span列表
	nmalloc uint64      // 已累计分配的对象个数
}
```
empty表示这条链表里的mspan都被分配了object，或者是已经被cache取走了的mspan，这个mspan就被那个工作线程独占了。而nonempty则表示有空闲对象的mspan列表。每个central结构体都在mheap中维护。

简单说下mcache从mcentral获取和归还mspan的流程：

* 获取 加锁；从nonempty链表找到一个可用的mspan；并将其从nonempty链表删除；将取出的mspan加入到empty链表；将mspan返回给工作线程；解锁。
* 归还 加锁；将mspan从empty链表删除；将mspan加入到nonempty链表；解锁。

###### 3.1.3 mheap：
mheap：代表Go程序持有的所有堆空间，Go程序使用一个mheap的全局对象_mheap来管理堆内存。

当mcentral没有空闲的mspan时，会向mheap申请。而mheap没有资源时，会向操作系统申请新内存。mheap主要用于大对象的内存分配，以及管理未切割的mspan，用于给mcentral切割成小对象。

同时我们也看到，mheap中含有所有规格的mcentral，所以，当一个mcache从mcentral申请mspan时，只需要在独立的mcentral中使用锁，并不会影响申请其他规格的mspan。

mheap结构体定义：
```
//path: /usr/local/go/src/runtime/mheap.go

type mheap struct {
	lock      mutex

	spans []*mspan

	bitmap        uintptr 	//指向bitmap首地址，bitmap是从高地址向低地址增长的

	arena_start uintptr		//指示arena区首地址
	arena_used  uintptr		//指示arena区已使用地址位置

	central [67*2]struct {
		mcentral mcentral
		pad      [sys.CacheLineSize - unsafe.Sizeof(mcentral{})%sys.CacheLineSize]byte
	}
}
```
上图我们看到，bitmap和arena_start指向了同一个地址，这是因为bitmap的地址是从高到低增长的，所以他们指向的内存位置相同。

#### 分配流程
变量是在栈上分配还是在堆上分配，是由逃逸分析的结果决定的。通常情况下，编译器是倾向于将变量分配到栈上的，因为它的开销小，最极端的就是"zero garbage"，所有的变量都会在栈上分配，这样就不会存在内存碎片，垃圾回收之类的东西。

Go的内存分配器在分配对象时，根据对象的大小，分成三类：小对象（小于等于16B）、一般对象（大于16B，小于等于32KB）、大对象（大于32KB）。

大体上的分配流程
```
>32KB 的对象，直接从mheap上分配；

<=16B 的对象使用mcache的tiny分配器分配；

(16B,32KB] 的对象，首先计算对象的规格大小，然后使用mcache中相应规格大小的mspan分配；

如果mcache没有相应规格大小的mspan，则向mcentral申请
如果mcentral没有相应规格大小的mspan，则向mheap申请
如果mheap中也没有合适大小的mspan，则向操作系统申请
```


## 4. 总结
Golang内存分配是个相当复杂的过程，其中还掺杂了GC的处理，它的一个原则就是能复用的一定要复用。
    
> * 1.Go在程序启动时，会向操作系统申请一大块内存，并划分成spans、bitmap、arena区域，之后自行分配；
> * 2.Go内存管理的基本单元是mspan，它由若干个页组成，每种mspan可以分配特定大小的object。即arena区域按页划分成一个个小块
> * 3.mcache, mcentral, mheap是Go内存管理的三大组件，层层递进。mcache管理线程（所谓的`P`）在本地缓存的mspan；mcentral管理全局的mspan供所有线程使用。
> * 4.极小对象会分配在一个object中，以节省资源，使用tiny分配器分配内存；一般小对象通过mspan分配内存；大对象则直接由mheap分配内存。

* mheap管理Go的所有动态分配内存
* mcentral管理多个span供线程申请使用;
* mcache作为线程私有资源，资源来源于mcentral

