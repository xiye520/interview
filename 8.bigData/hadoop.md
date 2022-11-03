# hadoop

* [Hadoop到底是干什么用的？](https://www.zhihu.com/question/333417513)
* [Hadoop、Hive、Spark 之间是什么关系？](https://cloud.tencent.com/developer/article/1042387)
* [大数据组件Presto，Spark SQL，Hive相互关系](https://blog.csdn.net/yilulvxing/article/details/86220888)
* []()
* []()
* []()
* 

> 一句话概括：Hadoop就是存储海量数据和分析海量数据的工具。
>
> HDFS(Hadoop Distributed FileSystem)：分布式文件系统

**稍专业点的解释**

Hadoop是由java语言编写的，在分布式服务器集群上存储海量数据并运行分布式分析应用的开源框架，其核心部件是HDFS与MapReduce。

​       HDFS是一个分布式文件系统：引入存放文件元数据信息的服务器Namenode和实际存放数据的服务器Datanode，对数据进行分布式储存和读取。

　　MapReduce是一个计算框架：MapReduce的核心思想是把计算任务分配给集群内的服务器里执行。通过对计算任务的拆分（Map计算/Reduce计算）再根据任务调度器（JobTracker）对任务进行分布式计算。

**1.3、记住下面的话：**

​       Hadoop的框架最核心的设计就是：HDFS和MapReduce。HDFS为海量的数据提供了存储，则MapReduce为海量的数据提供了计算。

​       把HDFS理解为一个分布式的，有冗余备份的，可以动态扩展的用来存储大规模数据的大硬盘。

​       把MapReduce理解成为一个计算引擎，按照MapReduce的规则编写Map计算/Reduce计算的程序，可以完成计算任务。

**2、Hadoop能干什么**

大数据存储：分布式存储

日志处理：擅长日志分析

ETL:数据抽取到oracle、mysql、DB2、mongdb及主流数据库

机器学习: 比如Apache Mahout项目

搜索引擎:Hadoop + lucene实现

数据挖掘：目前比较流行的广告推荐，个性化广告推荐

Hadoop是专为离线和大规模数据分析而设计的，并不适合那种对几个记录随机读写的在线事务处理模式。

实际应用：

（1）Flume+Logstash+Kafka+Spark Streaming进行实时日志处理分析

![img](https://pic3.zhimg.com/50/v2-fdaa7b4c0d7e2ae7d7c746d36d57c94c_720w.jpg?source=1940ef5c)![img](https://pic3.zhimg.com/80/v2-fdaa7b4c0d7e2ae7d7c746d36d57c94c_1440w.jpg?source=1940ef5c)

（2）酷狗音乐的大数据平台

![img](https://pic1.zhimg.com/50/v2-0e338a5f6ee11fd90e739a379c12982f_720w.jpg?source=1940ef5c)![img](https://pic1.zhimg.com/80/v2-0e338a5f6ee11fd90e739a379c12982f_1440w.jpg?source=1940ef5c)

**3、怎么使用Hadoop**

**3.1、Hadoop集群的搭建**

无论是在windows上装几台虚拟机玩Hadoop，还是真实的服务器来玩，说简单点就是把Hadoop的安装包放在每一台服务器上，改改配置，启动就完成了Hadoop集群的搭建。

**3.2、上传文件到Hadoop集群**

Hadoop集群搭建好以后，可以通过web页面查看集群的情况，还可以通过Hadoop命令来上传文件到hdfs集群，通过Hadoop命令在hdfs集群上建立目录，通过Hadoop命令删除集群上的文件等等。

**3.3、编写map/reduce程序**

通过集成开发工具（例如eclipse）导入Hadoop相关的jar包，编写map/reduce程序，将程序打成jar包扔在集群上执行，运行后出计算结果。