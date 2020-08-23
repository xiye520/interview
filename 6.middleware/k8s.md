# K8S--kubernetes

* [只花10分钟就能了解Kubernetes！]( http://blog.itpub.net/31547898/viewspace-2214070/)
* [K8S——基础概念](https://baijiahao.baidu.com/s?id=1651889690904189224&wfr=spider&for=pc)
* [K8s Service原理介绍](https://www.cnblogs.com/wn1m/p/11288131.html)
* [Kubernetes笔记（二）：了解k8s的基本组件与概念](http://www.bubuko.com/infodetail-3539517.html?__cf_chl_jschl_tk__=e381e103ac715e35c7ada08f7a54dc952370d59c-1589194438-0-AShPat97oORyEIfmzi4f5Hn1Y-QCEQHIvvWxffLEjOaxE3qP4eLy_0OPuq5hmca6ftVv8LTZmFMNtrMdbH42dmztJN2-SBK2zqCB2NfukV8EaXIqEd931d-IUrVr_0A86NQCRAD1JY16NifWCCAHstXDdiieU4mTWXYuzNyg0AF1sHrQLgyN3uQWHT2bQdekbZ63g8Se2BrX_07dYqQbkYjQENWCiasXMiu64-YJ11c-L6vhYBS6TF3zHLhDFKj_6PqGT5sXxDRXf4waxc6WaziQ77R9Xy0ttjBql1FHdvAhlnOSvOiLvQ6XGsC3DIANWA)
* [Kubernetes 本地快速启动（基于 Docker）](https://eddycjy.com/posts/kubernetes/2020-05-01-install/)
* [Containers vs Serverless：本质区别是什么？](http://blog.itpub.net/69923336/viewspace-2664936/)
* []()
* []()
* []()

# 

当Docker 成为流行趋势之后，Kubernetes获得快速发展，成为最常用的容器编排工具。那么问题来了，Kubernetes为什么重要？什么是Kubernetes？什么时候使用？如何使用？本文将按照各个要点，逐一总结！

## 一、前言

### 1、Kubernetes为什么重要？
在Kubernetes没出现以前，我们都认为容器是云管理平台应用部署的最佳工具。容器为软件开发和运维提供了新视角，通过使用容器，软件开发人员可以很容易地将应用打包，它既可以把应用程序拆分为分布式组件，也可以整体移植一个应用，而不像传统的虚拟机那样需要安装复杂的插件。
当整个世界开始走向分布式架构，当企业变得更加依赖网络、依赖计算，当单块应用开始迁移到微服务，我们的用户更希望通过微服务单独扩展关键功能，并能够处理数百万客户业务。因此，Docker容器、Mesos和AWS ECS等工具出现在企业应用名录中，这些工具为用户创建和部署微服务提供了更好的一致性、可移植性和更简单的操作方式。
但是，一旦企业应用变得更成熟和复杂，就需要在多台机器上运行多个容器。用户需要确定匹配多少个容器，容器的存储方式、存储数据量、性能需求等，如果这些工作全部用人工来统计，那简直是噩梦。
为解决容器的编排需求，Kubernetes应运而生！
### 2、什么是Kubernetes？
当容器管理成为Docker时代最重要的需求，谷歌做出了大胆决定，开放内部项目Borg。为了进一步增强容器管理功能，谷歌又开发了Kubernetes。这是一个开源项目，可自动化、大规模部署和管理容器应用过程。
于是，2014年年中，Kubernetes正式诞生，并在很短的时间内成长为开源社区，受到来自谷歌、Red Hat和许多其他公司工程师的热捧。
kubernetes，简称K8s，是用8代替8个字符“ubernete”而成的缩写。简单理解，kubernetes是一个开源的容器管理系统，用于管理云平台中多个主机上的容器化的应用。Kubernetes的目标是，让部署容器化的应用变得简单和高效,Kubernetes提供了应用部署，规划，更新和维护的一种新机制。
kubernetes的一些功能包括:管理容器集群、提供部署应用的工具、根据需要扩展应用、对现有容器应用的更改进行管理、能优化容器下的底层硬件的使用。另外，kubernetes还能管理跨机器的容器，解决Docker跨机器容器之间的通讯问题。
实际上，Kubernetes提供了比基础框架更多的内容，用户可以选择不同的应用框架、语言、监视工具和日志管理的类型等。虽然kubernetes不能完全当作服务平台，但绝对是一个很完整的PaaS。

## 二、K8S——基础概念

**Container**

Container（容器）是一种便携式、轻量级的操作系统级虚拟化技术。它使用 namespace 隔离不同的软件运行环境，并通过镜像自包含软件的运行环境，从而使得容器可以很方便的在任何地方运行。

由于容器体积小且启动快，因此可以在每个容器镜像中打包一个应用程序。

一对一的关系

**Pod**

Kubernetes 使用 Pod 来管理容器，每个 Pod 可以包含一个或多个紧密关联的容器。一对多的关系

Pod 是一组紧密关联的容器集合，它们共享 PID、IPC、Network 和 UTS namespace，是 Kubernetes 调度的基本单位。Pod 内的多个容器共享网络和文件系统，可以通过进程间通信和文件共享这种简单高效的方式组合完成服务。

![img](https://pics1.baidu.com/feed/6a63f6246b600c33e8c1aa0bd8e7860adbf9a1ee.png?token=ddecd8a0bf2364ad53ef6420dccb1533&s=18215C329B244903524DF0C60300C0B2)

apiVersion: v1 kind: Pod metadata: name: nginx labels: app: nginx spec: containers: - name: nginx image: nginx ports: - containerPort: 80

**Node**

Node 是 Pod 真正运行的主机，可以是物理机，也可以是虚拟机。为了管理 Pod，每个 Node 节点上至少要运行 container runtime（比如 docker 或者 rkt）、kubelet 和 kube-proxy 服务。

![img](https://pics4.baidu.com/feed/7e3e6709c93d70cfaa3b0ac13a770605b8a12b6e.png?token=5ce44bd66c15a2d9dff41f24e367570b&s=78A1187269DBD4CE08F9B1C7030010B0)

**Namespace**

Namespace 是对一组资源和对象的抽象集合，比如可以用来将系统内部的对象划分为不同的项目组或用户组。常见的 pods, services, replication controllers 和 deployments 等都是属于某一个 namespace 的（默认是 default），而 node, persistentVolumes 等则不属于任何 namespace。

**Service**

Service 是应用服务的抽象，通过 labels 为应用提供负载均衡和服务发现。匹配 labels 的 Pod IP 和端口列表组成 endpoints，由 kube-proxy 负责将服务 IP 负载均衡到这些 endpoints 上。

每个 Service 都会自动分配一个 cluster IP（仅在集群内部可访问的虚拟地址）和 DNS 名，其他容器可以通过该地址或 DNS 来访问服务，而不需要了解后端容器的运行。

![img](https://pics3.baidu.com/feed/54fbb2fb43166d22294419798488def29152d259.png?token=36127cee8b40aaa0fef5811589518caa&s=1A207C2249DFC1EB5C68746F0300E0F1)

apiVersion: v1 kind: Service metadata: name: nginx spec: ports: - port: 8078 # the port that this service should serve on name: http # the container on each pod to connect to, can be a name # (e.g. 'www') or a number (e.g. 80) targetPort: 80 protocol: TCP selector: app: nginx

**Label**

Label 是识别 Kubernetes 对象的标签，以 key/value 的方式附加到对象上（key 最长不能超过 63 字节，value 可以为空，也可以是不超过 253 字节的字符串）。

Label 不提供唯一性，并且实际上经常是很多对象（如 Pods）都使用相同的 label 来标志具体的应用。

Label 定义好后其他对象可以使用 Label Selector 来选择一组相同 label 的对象（比如 ReplicaSet 和 Service 用 label 来选择一组 Pod）。Label Selector 支持以下几种方式：

等式，如 app=nginx 和 env!=production集合，如 env in (production, qa)多个 label（它们之间是 AND 关系），如 app=nginx,env=test

**Annotations**

Annotations 是 key/value 形式附加于对象的注解。不同于 Labels 用于标志和选择对象，Annotations 则是用来记录一些附加信息，用来辅助应用部署、安全策略以及调度策略等。比如 deployment 使用 annotations 来记录 rolling update 的状态。

### configMap

configMap里放的是配置信息
configMap的主要作用是：让配置信息与镜像文件结藕，配置信息可以通过configMap注入。

##  三、Kubernetes的工作原理

![img](http://img.blog.itpub.net/blog/2018/09/12/c5fc9386f292bd52.png?x-oss-process=style/bb)

KubernetesMaster:提供集群管理控制中心，是最主要的控制单元，管理各系统之间的工作负载和通信。Master组件可以在集群中任何节点上运行。但是为了简单起见，通常在一台VM/机器上启动所有Master组件，并且不会在此VM/机器上运行用户容器。

Etcd:一个开源的用于配置共享和服务发现的高性能的键值存储系统，由CoreOS团队开发，也是CoreOS的核心组件。Kubernetes使用“Etcd”存储集群的配置数据，对集群内部组建进行协调。

API- server: 是接收和修改REST请求的中央控制系统，用作控制集群的前端。此外，这是唯一与Etcd集群通信的东西，确保数据存储在Etcd中。

scheduler：是关键角色，它决定了任务何时被调度运行，也决定一次任务运行中，哪些节点可以被执行。被判定执行的节点会被scheduler通过MQ或FaaS发送给worker执行。不同业务的任务有独立的scheduler负责调度，发送任务到指定的Worker上。

Controller：是一个控制器，在后台运行许多不同的控制器进程，用以控制集群的共享状态，并执行例行任务。当服务发生任何更改时，控制器会发现更改，并开始以新的状态工作。

Worker Node:也被称为Kubernetes或Minion节点，它包含管理容器之间的网络(如Docker)和主节点之间的通信信息，并按照计划将资源分配给容器

Kubelet: Kubelet确保节点中的所有容器都在运行，并处于健康状态。Kubelet负责pod的创建，以及是不是想要的状态。如果Node失败，Controller会观察到这个变化，并在另一个健康的pod上启动pod。
Container:是微服务的最低级别，放置在pod中，需要外部IP地址才能查看外部进程。

Kube Proxy:充当网络代理和负载均衡器。此外，它将请求转发到集群中的隔离网络中，主要负责为Pod对象提供代理。

cAdvisor:充当助理，负责监视和收集关于每个节点上的资源使用和性能指标的数据。

Kubernetes的优点：

1）部署简单，更具开放性

Kubernetes可以在一个或多个云环境、虚拟机或裸机上运行容器，这意味着它可以部署在任何基础设施上。此外，它兼容多个平台，使得多云策略高度灵活和可用。

2）更强大的工作负载能力和可伸缩性

Kubernetes为应用扩展提供了几个大的特性，比如：水平扩展、自动缩放、手动缩放、可复制控制器创建的Pods等。另外，Kubernetes还提供了更强大的高可用性、健康检查、流量控制和负载均衡器、自动转出和回滚等。

总结来看，Kubernetes为开发云应用奠定了更坚实的基础，Kubernetes和其他管理编配工具一样，如Marathon的Apache Mesos、Docker Swarm和AWS EC2等，提供了很棒的特性，但其他工具的份量都比Kubernetes小。

