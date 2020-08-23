# 4.数据库、缓存

mysql相关知识：	索引、锁机制，事务特性、四大隔离机制，存储引擎，b+tree索引，索引下推、最左匹配原则，分库分表，多机房数据同步，left join、inner join等，sql调优；

redis缓存相关知识：	五大基础数据类型及其使用场景，源码层数据结构实现（如dict、intset、skiplist等），缓存雪崩、穿透、击穿，aof、rdb持久化及其实现原理，分布式锁、并发竞争、双写一致性，集群高可用、哨兵机制，主从同步机制；

* [拜托，别再问我什么是 B+ 树了](https://mp.weixin.qq.com/s/vXnG1SVYeybaLeGGVnh2vA)
* [以B tree和B+ tree的区别来分析mysql索引实现](https://www.jianshu.com/p/0371c9569736)
*  码猿技术专栏 [Redis高频面试题及答案](https://mp.weixin.qq.com/s/AgWmiuvT0v6_2qVIYwV8Bg)
*  码猿技术专栏[【吊打面试官】Mysql大厂高频面试题！！！](https://mp.weixin.qq.com/s/DdOeiuZZ5onuAJXPTOzTUw)
* [我以为我对Mysql索引很了解，直到我遇到了阿里的面试官](https://juejin.im/post/5d23ef4ce51d45572c0600bc)
* [mysql知识总结](https://www.jianshu.com/p/2d97e48513d3)
* [《MySQL技术内幕：InnoDB存储引擎(第2版)﻿》书摘](https://www.jianshu.com/p/3eca0b18cf51)
* [Mysql从入门到入神之（二）Select 和Update的执行过程](https://juejin.im/post/5e806d23518825739728583d)
* [可能是全网最好的MySQL重要知识点](https://www.jianshu.com/p/5dd5993f981b)
* [真正理解Mysql的四种隔离级别](https://www.jianshu.com/p/8d735db9c2c0/)
* [数据库连接池到底应该设多大？这篇文章可能会颠覆你的认知](https://mp.weixin.qq.com/s/dQFSrXEmgBMh1PW835rlwQ)
* [Redis和mysql数据怎么保持数据一致的？](https://juejin.im/post/5c96fb795188252d5f0fdff2)
* [如何保证缓存(redis)与数据库(MySQL)的一致性](https://developer.aliyun.com/article/712285)
* []()
* []()


## innodb和myisam的区别，为什么要加索引
## 数据库索引中B.tree、B+tree数据结构，mysql用的b+tree是如何实现查
## 为什么用cassandra，优势
## redis中用了哪些数据类型

```
Redis目前支持5种数据类型，分别是：
String（字符串）
List（列表）
Hash（字典）
Set（集合）
Sorted Set（有序集合）
```
## redis中hashmap





## 分布式的环境下， MySQL和Redis如何保持数据的一致性？

数据库和缓存之间一般不需要强一致性。

一般缓存是这样的：

- 读的顺序是先读缓存，后读数据库
- 写的顺序是先写数据库，然后写缓存
- 每次更新了相关的数据，都要把该缓存清理掉
- 为了避免极端条件下造成的缓存与数据库之间的数据不一致，缓存需要设置一个失效时间。时间到了，缓存自动被清理，达到缓存和数据库数据的“最终一致性”

