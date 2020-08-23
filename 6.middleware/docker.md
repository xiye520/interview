## docker
* [Docker 有什么优势](https://www.zhihu.com/question/22871084/answer/88293837)
* 华为云技术宅基地的回答 [Docker 有什么优势？](https://www.zhihu.com/question/22871084/answer/635699901)
* [Docker到底是什么？为什么它这么火！](https://blog.csdn.net/truelove12358/article/details/54923729)
* [Docker 背后的内核知识——Namespace 资源隔离](https://www.infoq.cn/article/docker-kernel-knowledge-namespace-resource-isolation)
* [Docker容器实现原理及容器隔离性踩坑介绍](http://dockone.io/article/8148)
* [Docker资源隔离和限制实现原理](http://lionheartwang.github.io/blog/2018/03/18/dockerzi-yuan-ge-chi-he-xian-zhi-shi-xian-yuan-li/)
* [由浅入深docker系列： (5)资源隔离](https://zhuanlan.zhihu.com/p/67021108)
* 阿里巴巴高级工程师 [从零开始入门 K8s | 详解 K8s 容器基本概念](https://mp.weixin.qq.com/s?__biz=MzA4ODg0NDkzOA==&mid=2247487889&idx=1&sn=bc920579ae724205a9f522b752e93800&chksm=9022ae74a7552762fe25c5f93e43992767b0a13c99b950af97c5dbaef72355332891b36ccc55&xtrack=1&scene=90&subscene=93&sessionid=1589096137&clicktime=1589096682&enterid=1589096682&ascene=56&devicetype=android-29&version=27000e37&nettype=WIFI&abtest_cookie=AAACAA%3D%3D&lang=zh_CN&exportkey=Awxa6if4eJF1k8LT5tLPPKI%3D&pass_ticket=ufAQ%2FTAY5SXGIpk4jbfHacgC6g893tSWAQe%2FEY0trCvWBS0u3KWLWRlxV8HTJZ2K&wx_header=1)
* []()
* []()
* []()

## 简单总结

关于Docker实现原理，简单总结如下：

- 使用Namespaces实现了系统环境的隔离，Namespaces允许一个进程以及它的子进程从共享的宿主机内核资源（网络栈、进程列表、挂载点等）里获得一个仅自己可见的隔离区域，让同一个Namespace下的所有进程感知彼此变化，对外界进程一无所知，仿佛运行在一个独占的操作系统中；一般，Docker容器需要并且Linux内核也提供了这6种资源的namespace的隔离：

	>  UTS : 主机与域名
  >
  >  IPC : 信号量、消息队列和共享内存
  >
  >  PID : 进程编号
  >
  >  NETWORK : 网络设备、网络栈、端口等
  >
  >  mount : 挂载点(文件系统)
  >
  >  User : 用户和用户组
- 使用CGroups限制这个环境的资源使用情况，比如一台16核32GB的机器上只让容器使用2核4GB。使用CGroups还可以为资源设置权重，计算使用量，操控任务（进程或线程）启停等；
- 使用镜像管理功能，利用Docker的镜像分层、写时复制、内容寻址、联合挂载技术实现了一套完整的容器文件系统及运行环境，再结合镜像仓库，镜像可以快速下载和共享，方便在多环境部署。

正因为Docker不像虚机虚拟化一个Guest OS，而是利用宿主机的资源，和宿主机共用一个内核，所以会存在下面问题：

> 注意：存在问题并不一定说就是安全隐患，Docker作为最重视安全的容器技术之一，在很多方面都提供了强安全性的默认配置，其中包括：容器root用户的 Capability 能力限制，Seccomp系统调用过滤，Apparmor的 MAC 访问控制，ulimit限制，pid-limits的支持，镜像签名机制等。

1、Docker是利用CGroups实现资源限制的，只能限制资源消耗的最大值，而不能隔绝其他程序占用自己的资源;

2、Namespace的6项隔离看似完整，实际上依旧没有完全隔离Linux资源，比如/proc 、/sys 、/dev/sd*等目录未完全隔离，SELinux、time、syslog等所有现有Namespace之外的信息都未隔离。



Linux 内核中实现了六种 namespace，按照引入的先后顺序，列表如下：

![img](https://pic4.zhimg.com/80/v2-e7ade4269dd31d43713fe4ccf02c6bef_720w.jpg)

## 如何为进程提供一个独立的运行环境？

- 针对不同进程使用同一个文件系统所造成的问题而言，Linux 和 Unix 操作系统可以通过 chroot 系统调用将子目录变成根目录，达到视图级别的隔离；进程在 chroot 的帮助下可以具有独立的文件系统，对于这样的文件系统进行增删改查不会影响到其他进程；

- 因为进程之间相互可见并且可以相互通信，使用 Namespace 技术来实现进程在资源的视图上进行隔离。在 chroot 和 Namespace 的帮助下，进程就能够运行在一个独立的环境下了；

- 但在独立的环境下，进程所使用的还是同一个操作系统的资源，一些进程可能会侵蚀掉整个系统的资源。为了减少进程彼此之间的影响，可以通过 Cgroup 来限制其资源使用率，设置其能够使用的 CPU 以及内存量。

  

  其实，容器就是一个视图隔离、资源可限制、独立文件系统的进程集合。所谓“视图隔离”就是能够看到部分进程以及具有独立的主机名等；控制资源使用率则是可以对于内存大小以及 CPU 使用个数等进行限制。容器就是一个进程集合，它将系统的其他资源隔离开来，具有自己独立的资源视图。

  容器具有一个独立的文件系统，因为使用的是系统的资源，所以在独立的文件系统内不需要具备内核相关的代码或者工具，我们只需要提供容器所需的二进制文件、配置文件以及依赖即可。只要容器运行时所需的文件集合都能够具备，那么这个容器就能够运行起来。


## Docker的本质
Docker容器本质上是宿主机上的进程。Docker 通过namespace实现了资源隔离，通过cgroups实现资源限制，通过写时复制机制(Copy-on-write)实现了高效的文件操作。

2 其实docker是一个内核的搬运工
所以虽然docker帮助我们准备好了rootfs地址，镜像里面的文件，以及各种资源隔离的配置，但是在启动一个容器的时候，它只是调用系统中早已内置的可以隔离资源的方法，而kernel支持这些方法，也是在创建进程的方法上做了一层资源隔离的扩展而已。

这就解释了docker两个特性：

启动速度快，因为本质来说容器和进程差别没有想象中的大，共享了很多代码，流程也差的不多
linux内核版本有最低的要求，因为linux是在某个版本后开始支持隔离特性

## Docker的优势是什么
是软件交付领域的一种「标准化」，这种标准化的具体产物，简单来说就是「镜像（image）」


* 1.Docker是一个对Linux cgroup, namespace....包装并提供便利借口的的一个开源项目，使其看起来可以更像“虚拟机”
* 2.实现更轻量级的虚拟化，方便快速部署
* 3.感觉 对开发的影响不大，但对于部署来说，是场革命。可以极大的减少部署的时间成本和人力成本
  
## Docker背后的内核知识

## cgroups资源限制机制

cgroups是Linux内核提供的一种机制。

这种机制可以根据需求把一系列系统任务及其子任务整合(或分隔)到按资源划分等级的不同组内，限制了被namespace隔离起来的资源，并为资源设置权重、计算使用量(CPU, Memory, IO等)、操控任务启停等。

从而为系统资源管理提供了一个统一的框架。

### cgroup功能

cgroup提供的主要功能如下：

- 资源限制：限制任务使用的资源总额，并在超过这个配额时发出提示
- 优先级分配：分配CPU时间片数量及磁盘IO带宽大小、控制任务运行的优先级
- 资源统计：统计系统资源使用量，如CPU使用时长、内存用量等
- 任务控制：对任务执行挂起、恢复等操作

### cgroup子系统

cgroups的资源控制系统，每种子系统独立地控制一种资源。

- cpu：使用调度程序控制任务对CPU的使用。
- cpuacct(CPU Accounting)：自动生成cgroup中任务对CPU资源使用情况的报告。
- cpuset：为cgroup中的任务分配独立的CPU(多处理器系统时)和内存。
- devices：开启或关闭cgroup中任务对设备的访问
- freezer：挂起或恢复cgroup中的任务
- memory：设定cgroup中任务对内存使用量的限定，并生成这些任务对内存资源使用情况的报告
- perf(Linux CPU性能探测器)_event：使cgroup中的任务可以进行统一的性能测试
- net_cls(Docker未使用)：通过等级识别符标记网络数据包，从而允许Linux流量监控程序(Traffic Controller)识别从具体cgroup中生成的数据包