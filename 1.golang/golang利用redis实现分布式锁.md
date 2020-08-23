# Golang利用redis实现分布式锁
#### 原理
使用SETNX命令(SET if Not eXists)

* SETNX key value

将 key 的值设为 value，当且仅当 key 不存在。若给定的 key 已经存在，则 SETNX 不做任何动作。

设置成功，返回 1 。

设置失败，返回 0 。

为防止获取锁之后，忘记删除，成功后再设置一个过期时间

以上就是利用redis实现分布式锁的原理

#### 代码
[代码](./redis_lock.go)