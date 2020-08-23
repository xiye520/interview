# gRPC

* [从源码透析gRPC调用原理](https://cloud.tencent.com/developer/article/1189548)
* 煎鱼grpc系列 [4.1 gRPC及相关介绍](https://eddycjy.gitbook.io/golang/di-4-ke-grpc/install)

## 客户端

首先，我们以Github官网上的example为示例来一览gRPC client端的使用，从而跟踪其调用的逻辑个原理。总的来看，调用的过程基本就是分为三步：

- 创建connection
- 创建业务客户端实例
- 调用RPC接口

## 服务端

对于Server端，我们同样地根据Github上的官网示例来展开说明。总的来看，grpc在server端的调用逻辑如下，基本就是分为四步：

- 创建端口监听listener
- 创建server实例
- 注册服务（并未真正开始服务）
- 启动服务端

最后，简单以一个图示来展示grpc服务端的调用流程：

![img](https://ask.qcloudimg.com/draft/2276093/goteqrvm97.jpg?imageView2/2/w/1620)

gRPC Server简化调用流程

