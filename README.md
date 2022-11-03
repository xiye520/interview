golang后端知识点总结，作为一名后端开发人员需要了解掌握以下知识点及其背后原理：

* 1.首先要选准一门好的生产语言，比如golang，要了解MPG调度、基于tcmalloc的内存分配模型、gc，channel、context、反射，切片、map等基础知识的底层原理，掌握pprof、火焰图分析调优；

* 2.linux操作系统相关的知识要了解，进程、线程、僵尸进程、堆、栈，IPC（进程间通信），锁，用户态内核态，select、poll、epoll等io多路复用模型，常用的shell脚本，linux下开发调试、vim；

* 3.网络，OSI七层网络模型，http、tcp、udp，https的tls传输加密，tcp握手挥手、滑动窗口、拥塞控制，更高级的http2.0、quic等的原理；

* 4.mysql相关知识，索引、锁机制，事务特性、四大隔离机制，存储引擎，b+tree索引，索引下推、最左匹配原则，分库分表，多机房数据同步，sql调优；
  redis缓存相关知识：五大基础数据类型及其使用场景，源码层数据结构实现（如dict、intset、skiplist等），缓存雪崩、穿透、击穿，aof、rdb持久化及其实现原理，分布式锁、并发竞争、双写一致性，集群高可用、哨兵机制，主从同步机制；；

* 5.算法与数据结构，从链表到堆栈、队列、各种树（二叉树、二叉查找树、AVL树、B树、B+树、红黑树），八大排序算法，分治法、动态规划算法、贪心算法、回朔法、、kmp算法、快慢指针，使用map空间换时间等，

* 6.中间件，微服务，rabbitmq、kafka、k8s、docker、etcd、grpc、Prometheus、nginx，熔断、限流组件，网关，负载均衡；

* 7.设计模式，工厂模式、单例模式、观察者模式、责任链模式、模板模式、组合模式等；



## 引用说明

其中0.Alibaba目录中的均为从[HIT-Alibaba](https://github.com/HIT-Alibaba/interview)此github仓库中拷贝来的，主要是想做一个内部备份，感谢该作者的开源分享！
另外，这份总结中很多知识点都说从各个博客中归纳总结过来的，所以，感谢各位的分享，如果有引用忘记标明了，请提issue或者pr，本人看到后会第一时间纠正，谢谢。

[写在19年初的后端社招面试经历(两年经验): 蚂蚁 头条 PingCAP](https://github.com/aylei/interview)


### 推荐链接

* [SSH隧道技术----端口转发，socket代理](https://www.cnblogs.com/fbwfbi/p/3702896.html)

* [开源许可证教程](http://www.ruanyifeng.com/blog/2017/10/open-source-license-tutorial.html)

* [VIM中使用正则匹配中文](https://my.oschina.net/hotleave/blog/341500)

* [Linux sed 命令](http://man.linuxde.net/sed)

* [go generate 生成代码](https://www.cnblogs.com/majianguo/p/6653919.html)

* [简单围观一下有趣的 //go: 指令](https://eddycjy.com/posts/go/talk/2019-03-31-go-ins/)

* [Prometheus 查询语言](https://www.jianshu.com/p/3bdc4cfa08da)

* [工程师应该怎么学习--xargin](https://xargin.com/how-to-learn/)

