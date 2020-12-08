## 题目：
基于 [errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup) 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

