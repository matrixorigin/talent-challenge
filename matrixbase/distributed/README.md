## 分布式编程题
实现一个高可用，多副本，强一致，持久化的KV存储系统。实现语言为`Golang`。

本项目提供了整体的框架，答题者按照这个框架完善`pkg/store`，实现自己的store。

答题者可以使用`make docker`来编译打包镜像，使用`docker-compose up`来启动3副本测试。

### 1. 要求
* 对外客户端接口使用HTTP接口，实现`GET`, `SET`, `DELETE`语义
* 强一致协议使用raft，借助etcd的raft来构建
* 持久化使用[Pebble](https://github.com/cockroachdb/pebble)来构建
* 数据采用3副本存储
* Raft-Log需要删除，不能无限制存储，删除的时机和逻辑需要答题者自己决定

### 2. 挑战
有些挑战的需求，上面的要求只是实现了多副本的kv存储，这个时候有很多的限制，所有的流量都经过一个leader节点，系统的数据存储规模受制于单机的磁盘大小。这个挑战需要解决以下几个问题：
* 系统的处理能力和机器数量成正比，即机器越多，性能越好
* 系统存储的数据规模和机器数量成正比，即机器越多，存储的数据规模越大
* 系统能够解决数据存储和流量的均衡问题

一些提示和思路：使用Multi-raft来解决，即考虑整个系统中数据如何分片，每个分片的数据使用一个raft-group来管理。

### 3. 如何测试
收到答题者的作品时，我们会按照以下的流程进行测试

* 启动2个节点
* 使用一个客户端以 1000/s 的速率持续的写入KV数据5分钟，并且记录服务端返回成功写入的数据
* 数据写入2分钟后，启动第三个节点
* 滚动重启3个节点
* 停止写入客户端
* 检查数据是否正确

### 3. 如何提交

调试完毕后，发送代码到邮件 zx AT matrixorigin DOT cn 

**WARNING: DO NOT DIRECTLY FORK THIS REPO. DO NOT PUSH PROJECT SOLUTIONS PUBLICLY.**