# quic协议
* 腾讯技术工程 [科普：QUIC协议原理分析](https://zhuanlan.zhihu.com/p/32553477)

#### 改进的拥塞控制
   
   TCP 的拥塞控制实际上包含了四个算法：慢启动，拥塞避免，快速重传，快速恢复 [22]。
   
   QUIC 协议当前默认使用了 TCP 协议的 Cubic 拥塞控制算法 [6]，同时也支持 CubicBytes, Reno, RenoBytes, BBR, PCC 等拥塞控制算法。