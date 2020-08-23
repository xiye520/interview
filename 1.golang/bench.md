# golang单元测试
```
1，测试单个文件，一定要带上被测试的原文件
    go test -v  wechat_test.go wechat.go

2，测试单个方法
    go test -v -test.run TestRefreshAccessToken

go test -test.bench=".*"

```