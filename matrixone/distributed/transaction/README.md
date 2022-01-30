## 单机事务编程题
使用`Golang`、`C++`、`Java`任意语言实现一个单机的支持事务的KV内存存储服务。

### 1. 要求
* 采用悲观锁实现
* 实现`SI`级别的事务隔离级别
* 支持死锁检测
* 支持如下KV操作
  * PUT 保存一个KV
  * GET 查询一个KV
  * DELETE 删除一个KV
  * SCAN 从某一个KEY开始，顺序的查询指定个数的记录
* 采用C/S架构，自定义基于TCP的二进制私有协议对外提供服务（不能使用现有的协议来实现，比如HTTP）
* 实现访问KV服务的客户端

### 2. 如何提交

调试完毕后，发送代码到邮件 zx AT matrixorigin DOT cn 

**WARNING: DO NOT DIRECTLY FORK THIS REPO. DO NOT PUSH PROJECT SOLUTIONS PUBLICLY.**
