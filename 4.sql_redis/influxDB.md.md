# influxDB

* [InfluxDB](https://www.jianshu.com/p/68c471bf5533)

> InfluxDB（时序数据库），常用的一种使用场景：监控数据统计。每毫秒记录一下电脑内存的使用情况，然后就可以根据统计的数据，利用图形化界面（InfluxDB V1一般配合Grafana）制作内存使用情况的折线图；
>
> 可以理解为按时间记录一些数据（常用的监控数据、埋点统计数据等），然后制作图表做统计；
>
> **目前（2019-09-29）建议使用InfluxDB V1 版本**

## 1、什么是InfluxDB

从文章开票的介绍里能大概知道它的使用场景，下面介绍来自维基百科：

> InfluxDB是一个由InfluxData开发的开源时序型数据。它由Go写成，着力于高性能地查询与存储时序型数据。InfluxDB被广泛应用于存储系统的监控数据，IoT行业的实时数据等场景。

## 2、对常见关系型数据库（MySQL）的基础概念对比

| 概念         | MySQL    | InfluxDB                                                    |
| ------------ | -------- | ----------------------------------------------------------- |
| 数据库（同） | database | database                                                    |
| 表（不同）   | table    | measurement                                                 |
| 列（不同）   | column   | tag(带索引的，非必须)、field(不带索引)、timestemp(唯一主键) |

- tag set：**不同**的每组tag key和tag value的集合；
- field set：每组field key和field value的集合；
- retention policy：数据存储策略（默认策略为autogen）InfluxDB没有删除数据操作，规定数据的保留时间达到清除数据的目的；
- series：共同retention policy，measurement和tag set的集合；

- 示例数据如下： 其中census是measurement，butterflies和honeybees是field key，location和scientist是tag key

```sql
name: census
————————————
time                 butterflies     honeybees     location     scientist
2015-08-18T00:00:00Z      12             23           1         langstroth
2015-08-18T00:00:00Z      1              30           1         perpetua
2015-08-18T00:06:00Z      11             28           1         langstroth
2015-08-18T00:06:00Z      11             28           2         langstroth
```

示例中有三个tag set

## 3、注意点

- tag 只能为字符串类型
- field 类型无限制
- 不支持join
- 支持连续查询操作（汇总统计数据）：CONTINUOUS QUERY
- 配合Telegraf服务（Telegraf可以监控系统CPU、内存、网络等数据）
- 配合Grafana服务（数据展现的图像界面，将influxdb中的数据可视化）

## 4、常用InfluxQL

```sql
-- 查看所有的数据库
show databases;
-- 使用特定的数据库
use database_name;
-- 查看所有的measurement
show measurements;
-- 查询10条数据
select * from measurement_name limit 10;
-- 数据中的时间字段默认显示的是一个纳秒时间戳，改成可读格式
precision rfc3339; -- 之后再查询，时间就是rfc3339标准格式
-- 或可以在连接数据库的时候，直接带该参数
influx -precision rfc3339
-- 查看一个measurement中所有的tag key 
show tag keys
-- 查看一个measurement中所有的field key 
show field keys
-- 查看一个measurement中所有的保存策略(可以有多个，一个标识为default)
show retention policies;
```

