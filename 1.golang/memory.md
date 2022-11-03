# 内存
* [Golang 内存管理](http://legendtkl.com/2017/04/02/golang-alloc/)
* [tcmalloc 介绍](http://legendtkl.com/2015/12/11/go-memory/)
* [深入浅出Golang Runtime](https://mp.weixin.qq.com/s/TPIIjBycYuhXWOKTg1_2vA)
* [ptmalloc、tcmalloc与jemalloc对比分析](https://www.cyningsun.com/07-07-2018/memory-allocator-contrasts.html)

go的内存管理和tcmalloc（`thread-caching malloc`）很像，不妨先看看tcmalloc的实现。
## 1.tcmalloc
#### tcmalloc是什么
tcmalloc是google推出的一种内存分配器，常见的内存分配器还有c使用的glibc的ptmalloc2和google的jemalloc。相比于ptmalloc，tcmalloc性能更好，特别适用于高并发场景。

#### tcmalloc具体策略
tcmalloc分配的内存主要来自两个地方：全局缓存堆和进程的私有缓存。对于一些小容量的内存申请使用进程的私有缓存，私有缓存不足的时候可以再从全局缓存申请一部分作为私有缓存。对于大容量的内存申请则需要从全局缓存中进行申请。而大小容量的边界就是32k。缓存的组织方式是一个单链表数组，数组的每个元素是一个单链表，链表中的每个元素具有相同的大小。

#### go的内存分配
go语言的内存分配并不是和tcmalloc一模一样。
* 局部缓存并不是分配给进程或者线程，而是分配给P（这个还需要说一下go的goroutine实现）
* go的GC是stop the world，并不是每个进程单独进行GC。
* span的管理更有效率

## 2.golang内存
Golang 的内存管理基于 tcmalloc，可以说起点挺高的。但是 Golang 在实现的时候还做了很多优化，

连续的的虚拟内存布局（64 位）如下：
```
+-----------------------+---------------------+-----------------------+
|    spans 512M         |    bitmap 16G       |   arena 512           |
+-----------------------+---------------------+-----------------------+
```

4. 内存分配
先说一下给对象 object 分配内存的主要流程：

```
object size > 32K，则使用 mheap 直接分配。
object size < 16 byte，使用 mcache 的小对象分配器 tiny 直接分配。 （其实 tiny 就是一个指针，暂且这么说吧。）
object size > 16 byte && size <=32K byte 时，先使用 mcache 中对应的 size class 分配。
如果 mcache 对应的 size class 的 span 已经没有可用的块，则向 mcentral 请求。
如果 mcentral 也没有可用的块，则向 mheap 申请，并切分。
如果 mheap 也没有合适的 span，则想操作系统申请。

整个分配过程可以根据 object size 拆解成三部分：size < 16 byte, 16 byte <= size <= 32 K byte, size > 32 K byte。

4.1 size 小于 16 byte
对于小于 16 byte 的内存块，mcache 有个专门的内存区域 tiny 用来分配，tiny 是指针，指向开始地址。

4.2 size 大于 32 K byte
对于大于 32 Kb 的内存分配，直接跳过 mcache 和 mcentral，通过 mheap 分配。
对于大于 32 K 的内存分配都是分配整数页，先右移然后低位与计算需要的页数

4.3 size 介于 16 和 32K
对于 size 介于 16 ~ 32K byte 的内存分配先计算应该分配的 sizeclass，然后去 mcache 里面 alloc[sizeclass] 申请，
如果 mcache.alloc[sizeclass] 不足以申请，则 mcache 向 mcentral 申请，然后再分配。
mcentral 给 mcache 分配完之后会判断自己需不需要扩充，如果需要则想 mheap 申请。
```
## golang内存分配简介
* 类似于TCMalloc的结构
* 使用span机制来减少碎片. 每个span至少为一个页(go中的一个page为8KB). 每一种span用于一个范围的内存分配需求. 比
如16-32byte使用分配32byte的span, 112-128使用分配128byte的span.
* 一共有67个size范围, 8byte-32KB, 每个size有两种类型(scan和noscan, 表示分配的对象是否会包含指针)
* 多层次Cache来减少分配的冲突. per-P无锁的mcache, 全局67*2个对应不同size的span的后备mcentral, 全局1个的mheap.
* mheap中以treap的结构维护空闲连续page. 归还内存到heap时, 连续地址会进行合并.
* stack分配也是多层次和多class的.
* 对象由GC进行回收. sysmon会定时把空余的内存归还给操作系统

golang内存结构
```
1.10及以前
以下内存并不是初始化时就分配虚拟内存的:
arena的大小为512G, 为了方便管理把arena区域划分成一个个的page,
每个page 8KB, 一共有512GB/8KB个页
spans区域存放指向span的指针, 表示arean中对应的Page所属的span,
所以span区域的大小为(512GB/8KB)*指针大小8byte = 512M
bitmap主要用于GC, 用两个bit表示 arena中一个字的可用状态, 所以是
(512G/8个字节一个字)*2/8个bit每个字节=16G
1.11及以后:
改成了两阶稀疏索引的方式. 内存可以超过512G, 也可以允许不连续的内
存.
mheap中的areans字段是一个指针数组, 每个heapArena管理64M的内
存.
bitmap和spans和上面的功能一致.

```
mspan
使用span机制来减少碎片. 每个span至少分配1个page(8KB), 划分成固定大小的slot, 用于分配一定大小范围的内存需求

## Golang内存分配综合
   * 类似于TCMalloc的结构
   * 使用span机制来减少碎片. 每个span至少为一个页(go中的一个page为8KB). 每一种span用于一个范围的内存分配需求. 比如16-32byte使用分配32byte的span, 112-128使用分配128byte的span.
   * 一共有67个size范围, 8byte-32KB, 每个size有两种类型(scan和noscan, 表示分配的对象是否会包含指针)
   * 多阶Cache来减少分配的冲突. per-P无锁的mcache, 对应不同size(67*2)的全局mcentral, 全局的mheap.
   * go代码分配内存优先从当前p的mcache对应size的span中获取; 有的话, 再从对应size的mcentral中获取一个span; 还没有的话, 从mheap中sweep一个span; sweep不出来, 则从mheap中空闲块找到对应span大小的内存. mheap中如果还没有, 则从系统申请内存. 从无锁到全局1/(67*2)粒度的锁, 再到全局锁, 再到系统调用.
   * stack的分配也是多层次和多class的. 减少分配的锁争抢, 减少栈浪费.
   * mheap中以treap的结构维护空闲连续page. 归还内存到mheap时, 连续地址会进行合并. (1.11之前采用类似伙伴系统维护<1MB的连续page, treap维护>1MB的连续page)
   * 对象由GC进行释放. sysmon会定时把mheap空余的内存归还给操作系统

## 分配策略
多层次Cache来减少分配的冲突, 加快分配.
从无锁到粒度较低的锁, 再到全局一个锁, 或系统调用.

> 1. new, make最终调用mallocgc;
> 2. 大于32KB对象, 直接从mheap中分配, 构成一个span
> 3. 小于16byte且无指针(noscan), 使用tiny分配器, 合并分配.
> 4. 小于16byte有指针或16byte-32KB, 如果mcache中有对应class的空闲mspan, 则直接从该mspan中分配一个slot.
> 5. (mcentral.cacheSpan) mcache没有对应的空余span, 则从对应mcentral中申请一个有空余slot的span到mcache中. 再进行分配
> 6. ( mcentral.grow)对应mcentral没有空余span, 则向 mheap( mheap_.alloc)中申请一个span, 能sweep出span则返 回. 否则看mheap的free mTreap能否分配最大于该size的连续 页, 能则分配, 多的页放回 .
> 7. mheap的free mTreap无可用, 则调用sysAlloc(mmap)向系统申请.
> 8. 6, 7步中获得的内存构建成span, 返回给mcache, 分配对象

