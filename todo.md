# 测试工具
生成覆盖率
```
go test -coverprofile=coverage.out ./...

coverage: 92.2% of statements
ok      github.com/melodywen/go-ioc        0.219s
```
测试结果：
```
go tool cover -html=coverage.out 
```

常用工具：
```
https://github.com/gojp/goreportcard
```
```
goreportcard-cli -v
```