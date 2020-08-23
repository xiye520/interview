golang后台开发面试题汇总

# 一、golang
## channal有哪几种，应用场景
## go启动过程
## go的反射
## go如何跨平台
## goroutine的调度原理
## golang的gc原理
## golang的内存分配原理
## cgo是什么，为什么要有cgo
## Golang 的 GC 触发时机是什么
答：阈值触发；主动触发；两分钟定时触发；
## channel的注意点
答：1.channel一定记得close。goroutine记得return或者中断，不然容易造成goroutine占用大量CPU。
2.channel是通过注册相关goroutine id实现消息通知的。
## 切片注意点
答：从slice创建slice的时候，注意原slice的操作可能导致底层数组变化。如果要创建一个很长的slice，尽量创建成一个slice里存引用，这样可以分批释放，避免gc在低配机器上stop the world
## 关于go内存分配
```
>32Kb,直接从mheap上分配，这就是所谓的大对象
<=16B使用mcache的tiny分配器分配，这就是所谓的小对象，一定要控制小对象的数量，多了会增加gc压力；
(16,32Kb]的内存对象，首先计算对象的规格大小（剩下66+67-1=133种类型），然后从mcache中拿出相应规格大小的mspan来分配；

内存分配过程中，
先从当前p的mcache种分配相应大小的内存，如没有向下一级申请    --无锁
如果mcache中没有相应规格大小的mspan，则向mcentral申请        --该class大小的mcentral的局部锁
如果mcentral中没有相应规格大小的mspan，则向mheap申请         --该class大小的mheap的全局锁
如果mheap中也没有合适大小大小的mspan，则向操作系统申请
```

----------------------

# 二、linux、操作系统知识
## 进程和线程的概念，进程、线程、协程的区别
## 进程间通信的方式有哪些？哪种效率高
## linux操作系统cgroup、namespace隔离
## 用户态、内核态的区别
## docker的隔离机制，
## 什么是缓冲区溢出？有什么危害？其原因是什么？
```
缓冲区溢出是指当计算机向缓冲区填充数据时超出了缓冲区本身的容量，溢出的数据覆盖在合法数据上。

危害有以下两点：
1.程序崩溃，导致拒绝额服务
2.跳转并且执行一段恶意代码

造成缓冲区溢出的主要原因是程序中没有仔细检查用户输入。
```
## fork
## poll、epoll、select网络模型
## 堆和栈的区别
## 浅谈乐观锁与悲观锁
## 线程间的通信方式
答：如golang可使用channel，或者全局变量，也可以通过socket，如走mq、grpc，如订阅redis键值修改等



----------------------

# 三、中间件
## rabbitMQ的原理，为什么要用mq队列



----------------------

# 四、数据库、缓存
## innodb和myisam的区别，为什么要加索引
## 数据库索引中B.tree、B+tree数据结构，mysql用的b+tree是如何实现查找
## 为什么用cassandra，优势
## redis中用了哪些数据类型
string(字符串),hash(哈希),list(列表),set(集合)及zset(sorted set:有序集合)
## redis中hashmap，跳表
## innodb引擎的4大特性
## redis有哪些架构模式？各自的特点是什么






# 五、算法与数据结构

## hashmap，跳表，几种树优缺点
## redis一些数据结构（比如zset），深度，广度遍历
## 链表，复杂度和消耗时间的分析和评估

## 堆排、快排、选择、冒泡，常见的八种排序及其复杂度
## 一百亿条数据如何排序
## 线性表
## 二叉树
## 红黑树
## 平衡树
## Radix树
## 八叉树
## 梅克尔树
## 一致性hash算法
## raft算法

----------------------

# 六、网络
## tcp四次挥手、三次握手
## tcp、udp的区别 
## 为啥要用proterbuf、msgpack
## jwt协议
## gin、beego网络框架
## http 301、302的区别？504和500有什么区别？
```
200 OK 客户端请求成功
301 Moved Permanently 请求永久重定向
302 Moved Temporarily 请求临时重定向
304 Not Modified 文件未修改，可以直接使用缓存的文件。
400 Bad Request 由于客户端请求有语法错误，不能被服务器所理解。
401 Unauthorized 请求未经授权。这个状态代码必须和WWW-Authenticate报头域一起使用
403 Forbidden 服务器收到请求，但是拒绝提供服务。服务器通常会在响应正文中给出不提供服务的原因
404 Not Found 请求的资源不存在，例如，输入了错误的URL
500 Internal Server Error 服务器发生不可预期的错误，导致无法完成客户端的请求。
503 Service Unavailable 服务器当前不能够处理客户端的请求，在一段时间之后，服务器可能会恢复正常。
```

## rpc底层协议是什么  socket？
## grpc和http协议的区别，grpc对比http有哪些优缺点
## 长连接短连接的区别？各自的优缺点？使用场景
## http和https的区别
## 三次握手的具体实现，time_wait原理？
## 计算机网络中不同层有用到哪些协议
## https建立连接过程
## 大小端模式
## 浏览器输入一个url发生了什么





# 七、系统、场景设计

## 高并发系统的限流如何实现？

## 高并发秒杀系统的设计

## 负载均衡如何设计？