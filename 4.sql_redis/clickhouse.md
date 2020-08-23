# Clickhouse

* [clickhouse官方中文文档](https://clickhouse.tech/docs/zh/)
* [ClickHouse概述](https://www.jianshu.com/p/350b59e8ea68)
* [最快开源 OLAP 引擎！ClickHouse 在头条的技术演进](https://www.infoq.cn/article/NTwo*yR2ujwLMP8WCXOE)

ClickHouse 是由号称“俄罗斯 Google”的 Yandex 公司开源的面向 OLAP 的分布式列式数据库，能够使用 SQL 查询生成实时数据报告。

## ClickHouse 简介

ClickHouse 是由号称“俄罗斯 Google”的 Yandex 开发而来，在 2016 年开源，在计算引擎里算是一个后起之秀，在内存数据库领域号称是最快的。大家从网上也能够看到，它有几倍于 GreenPlum 等引擎的性能优势。

如果大家研究过它的源码，会发现其实它采用的技术并不新。ClickHouse 是一个列导向数据库，是原生的向量化执行引擎。它在大数据领域没有走 Hadoop 生态，而是采用 Local attached storage 作为存储，这样整个 IO 可能就没有 Hadoop 那一套的局限。它的系统在生产环境中可以应用到比较大的规模，因为它的线性扩展能力和可靠性保障能够原生支持 shard + replication 这种解决方案。它还提供了一些 SQL 直接接口，有比较丰富的原生 client。另外就是它比较快。

大家选择 ClickHouse 的首要原因是它比较快，但其实它的技术没有什么新的地方，为什么会快？我认为主要有三个方面的因素：

1. 它的数据剪枝能力比较强，分区剪枝在执行层，而存储格式用局部数据表示，就可以更细粒度地做一些数据的剪枝。它的引擎在实际使用中应用了一种现在比较流行的 LSM 方式。
2. 它对整个资源的垂直整合能力做得比较好，并发 MPP+ SMP 这种执行方式可以很充分地利用机器的集成资源。它的实现又做了很多性能相关的优化，它的一个简单的汇聚操作有很多不同的版本，会根据不同 Key 的组合方式有不同的实现。对于高级的计算指令，数据解压时，它也有少量使用。
3. 我当时选择它的一个原因，ClickHouse 是一套完全由 C++ 模板 Code 写出来的实现，代码还是比较优雅的。

#### **可以应用以下场景：**

1.电信行业用于存储数据和统计数据使用。

2.新浪微博用于用户行为数据记录和分析工作。

3.用于广告网络和RTB,电子商务的用户行为分析。

4.信息安全里面的日志分析。

5.检测和遥感信息的挖掘。

6.商业智能。

7.网络游戏以及物联网的数据处理和价值数据分析。

8.最大的应用来自于Yandex的统计分析服务Yandex.Metrica，类似于谷歌Analytics(GA)，或友盟统计，小米统计，帮助网站或移动应用进行数据分析和精细化运营工具，据称Yandex.Metrica为世界上第二大的网站分析平台。ClickHouse在这个应用中，部署了近四百台机器，每天支持200亿的事件和历史总记录超过13万亿条记录，这些记录都存有原始数据（非聚合数据），随时可以使用SQL查询和分析，生成用户报告。

> 五．ClickHouse 和一些技术的比较

1.商业OLAP数据库

例如：HP Vertica, Actian the Vector,

区别：ClickHouse是开源而且免费的

2.云解决方案

例如：亚马逊RedShift和谷歌的BigQuery

区别：ClickHouse可以使用自己机器部署，无需为云付费

3.Hadoop生态软件

例如：Cloudera Impala, Spark SQL, Facebook Presto , Apache Drill

区别：

ClickHouse支持实时的高并发系统

ClickHouse不依赖于Hadoop生态软件和基础

ClickHouse支持分布式机房的部署

4.开源OLAP数据库

例如：InfiniDB, MonetDB, LucidDB

区别：这些项目的应用的规模较小，并没有应用在大型的互联网服务当中，相比之下，ClickHouse的成熟度和稳定性远远超过这些软件。

5.开源分析，非关系型数据库

例如：Druid , Apache Kylin

区别：ClickHouse可以支持从原始数据的直接查询，ClickHouse支持类SQL语言，提供了传统关系型数据的便利。

> 六．总结

在大数据分析领域中，传统的大数据分析需要不同框架和技术组合才能达到最终的效果，在人力成本，技术能力和硬件成本上以及维护成本让大数据分析变得成为昂贵的事情。让很多中小型企业非常苦恼，不得不被迫租赁第三方大型公司的数据分析服务。

ClickHouse开源的出现让许多想做大数据并且想做大数据分析的很多公司和企业耳目一新。ClickHouse 正是以不依赖Hadoop 生态、安装和维护简单、查询速度快、可以支持SQL等特点在大数据分析领域越走越远。

