# go map解析
* [Golang map 的底层实现](https://www.jianshu.com/p/aa0d4808cbb8)
* [深入理解 Go map：初始化和访问元素](https://eddycjy.com/posts/go/map/2019-03-05-map-access/)
* [深入理解 Go map：赋值和扩容迁移](https://eddycjy.com/posts/go/map/2019-03-24-map-assign/)
* [深度解密Go语言之sync.map](https://mp.weixin.qq.com/s?__biz=MjM5MDUwNTQwMQ==&mid=2257484131&idx=1&sn=a241eb4b5d869aae91c6c54ec7e89c44&chksm=a53919b5924e90a3800afefe8c8ef7cb4cf0fd8e4d793b6282fffe93ee733ba3ce0d3999e714&scene=126&sessionid=1589300701&key=5fa94eed5565a8ca5039ce5305624ecfdda2f8fa9639961ab266c17f1074ba6d29ea9ba37091ef511fbdec05b9627a29d5fd4e7fd7eb7ab9291ac6270ba1aae2fc1d338927b20019f4f6cd9122838c63&ascene=1&uin=NzA3NzMxNjgx&devicetype=Windows+10+x64&version=62090070&lang=zh_CN&exportkey=A2YgmnOOk%2B3khVEs8pXZCZ4%3D&pass_ticket=bkqB6VCumkFhyQKCK4SHdAAH0HEdiBIJmzogo9YxEbT%2FW6ohTqn%2B4jrs%2B3EgoZYg)
* [深入理解Go-sync.Map原理剖析](https://juejin.im/post/5d74d562f265da03ab4273e1)
* []()
* []()
* []()
* []()





## 一、map底层数据结构

### hmap

```cgo
type hmap struct {
	count     int
	flags     uint8
	B         uint8
	noverflow uint16
	hash0     uint32
	buckets    unsafe.Pointer
	oldbuckets unsafe.Pointer
	nevacuate  uintptr
	extra *mapextra
}

type mapextra struct {
	overflow    *[]*bmap
	oldoverflow *[]*bmap
	nextOverflow *bmap
}
```
* count：map 的大小，也就是 len() 的值。代指 map 中的键值对个数
* flags：状态标识，主要是 goroutine 写入和扩容机制的相关状态控制。并发读写的判断条件之一就是该值
* B：桶，最大可容纳的元素数量，值为 负载因子（默认 6.5） * 2 ^ B，是 2 的指数
* noverflow：溢出桶的数量
* hash0：哈希因子
* buckets：保存当前桶数据的指针地址（指向一段连续的内存地址，主要存储键值对数据）
* oldbuckets，保存旧桶的指针地址
* nevacuate：迁移进度
* extra：原有 buckets 满载后，会发生扩容动作，在 Go 的机制中使用了增量扩容，如下为细项：
    * overflow 为 hmap.buckets （当前）溢出桶的指针地址
    * oldoverflow 为 hmap.oldbuckets （旧）溢出桶的指针地址
    * nextOverflow 为空闲溢出桶的指针地址
在这里我们要注意几点，如下：

如果 keys 和 values 都不包含指针并且允许内联的情况下。会将 bucket 标识为不包含指针，使用 extra 存储溢出桶就可以避免 GC 扫描整个 map，节省不必要的开销
在前面有提到，Go 用了增量扩容。而 buckets 和 oldbuckets 也是与扩容相关的载体，一般情况下只使用 buckets，oldbuckets 是为空的。但如果正在扩容的话，oldbuckets 便不为空，buckets 的大小也会改变
当 hint 大于 8 时，就会使用 *mapextra 做溢出桶。若小于 8，则存储在 buckets 桶中

### bmap
```
bucketCntBits = 3
bucketCnt     = 1 << bucketCntBits
...
type bmap struct {
	tophash [bucketCnt]uint8
}
```

* tophash：key 的 hash 值高 8 位
* keys：8 个 key
* values：8 个 value
* overflow：下一个溢出桶的指针地址（当 hash 冲突发生时）
实际 bmap 就是 buckets 中的 bucket，一个 bucket 最多存储 8 个键值对

#### tophash
tophash 是个长度为 8 的数组，代指桶最大可容纳的键值对为 8。

存储每个元素 hash 值的高 8 位，如果 tophash [0] <minTopHash，则 tophash [0] 表示为迁移进度

#### keys 和 values
在这里我们留意到，存储 k 和 v 的载体并不是用 k/v/k/v/k/v/k/v 的模式，而是 k/k/k/k/v/v/v/v 的形式去存储。这是为什么呢？

```
map[int64]int8
```
在这个例子中，如果按照 k/v/k/v/k/v/k/v 的形式存放的话，虽然每个键值对的值都只占用 1 个字节。但是却需要 7 个填充字节来补齐内存空间。最终就会造成大量的内存 “浪费”

但是如果以 k/k/k/k/v/v/v/v 的形式存放的话，就能够解决因对齐所 “浪费” 的内存空间

因此这部分的拆分主要是考虑到内存对齐的问题，虽然相对会复杂一点，但依然值得如此设计

## 二、golang map遍历顺序为何随机？

map是有删除的，那个空的bucket会被新的数据填充，这样会导致遍历结果不一致，再一个就是map扩容时位置也会bucket内的数据也会有迁移。所以官方直接在map遍历里面引入随机开头点的遍历机制，防止别人在依赖于map的遍历顺序来写业务，会出乱子的！