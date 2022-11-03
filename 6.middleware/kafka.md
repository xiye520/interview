* [Kafka的分区数和消费者个数](https://www.jianshu.com/p/dbbca800f607)

* [Kafka 工作原理](https://www.jianshu.com/p/6cbe28a44543)

* [Kafka 入门介绍](https://lotabout.me/2018/kafka-introduction/)

* [再过半小时，你就能明白kafka的工作原理了](https://zhuanlan.zhihu.com/p/68052232)

* [Kafka学习之路 （一）Kafka的简介](https://www.cnblogs.com/qingyunzong/p/9004509.html)

* []()

### 术语

- Broker
   Kafka集群包含一个或多个服务器，这种服务器被称为broker，可以水平扩展，一般broker数量越多，集群吞吐率越高，而且kafka 每个节点可以有多个 broker
- Producer
   负责发布消息到Kafka broker，可以是web前端产生的page view，或者是服务器日志，系统CPU、memory等
- Consumer
   消费消息。每个consumer属于一个特定的consumer group（可为每个consumer指定group name，若不指定group name则属于默认的group）。使用consumer high level API时，同一topic的一条消息只能被同一个consumer group内的一个consumer消费，但多个consumer group可同时消费这一消息。
- Zookeeper
   通过Zookeeper管理集群配置，选举leader，以及在consumer group发生变化时进行rebalance
- Topic
   每条发布到Kafka集群的消息都有一个类别，这个类别被称为topic。（物理上不同topic的消息分开存储，逻辑上一个topic的消息虽然保存于一个或多个broker上但用户只需指定消息的topic即可生产或消费数据而不必关心数据存于何处）
- Partition
   parition是物理上的概念，每个topic包含一个或多个partition，创建topic时可指定parition数量。每个partition对应于一个文件夹，该文件夹下存储该partition的数据和索引文件
- Segment
   partition物理上由多个segment组成，每一个segment 数据文件都有一个索引文件对应
- Offset
   每个partition都由一系列有序的、不可变的消息组成，这些消息被连续的追加到partition中。partition中的每个消息都有一个连续的序列号叫做offset,用于partition唯一标识一条消息.

### Push vs. Pull

push模式很难适应消费速率不同的消费者，因为消息发送速率是由broker决定的。push模式的目标是尽可能以最快速度传递消息，但是这样很容易造成consumer来不及处理消息，典型的表现就是拒绝服务以及网络拥塞。而pull模式则可以根据consumer的消费能力以适当的速率消费消息。
 所以我们一般在 Kafka 前面再加一个 Log Server，可以用 LevelDB 缓存，作为一个缓冲，提高峰值处理能力

### Topic & Partition

每条消费都必须指定它的topic，为了使得Kafka的吞吐率可以水平扩展，物理上把topic分成一个或多个partition，每个partition在物理上对应一个文件夹，该文件夹下存储这个partition的所有消息和索引文件。