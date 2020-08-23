## redis持久化
redis的持久化有两种方式，RDB和AOF 

#### RDB: 

------------------------------
在指定的时间间隔内，执行指定次数的写操作，则会将内存中的数据写入到磁盘中。即在指定目录下生成一个dump.rdb文件。Redis 重启会通过加载dump.rdb文件恢复数据。

配置方法：

* save 900 1
* save 300 10
* save 60 10000

意思是， 900秒内有1个更改，300秒内有10个更改以及60秒内有10000个更改，则将内存中的数据快照写入磁盘。

* 恢复方法
------------------------------
将dump.rdb 文件拷贝到redis的安装目录的bin目录下，重启redis服务即可

#### AOF

------------------------------
采用日志的形式来记录每个写操作，并追加到文件中。Redis 重启的会根据日志文件的内容将写指令从前到后执行一次以完成数据的恢复工作。

* 配置方法
```
appendonly yes
appendfilename "appendonly.aof"
#指定更新条件
# appendfsync always
appendfsync everysec
# appendfsync no
#配置重写触发机制
auto-aof-rewrite-percentage 100
auto-aof-rewrite-min-size 64mb

```
#### 恢复方法

将appendonly.aof 文件拷贝到redis的安装目录的bin目录下，重启redis服务即可。如果因为某些原因导致appendonly.aof 文件格式异常，从而导致数据还原失败，可以通过命令redis-check-aof –fix appendonly.aof 进行修复

#### 区别
RDB通过fork的方式进行处理，性能更好 AOF备份所有的操作，数据更完整，但是效率略差，文件相对较大 实际环境中，可以两者同时使用