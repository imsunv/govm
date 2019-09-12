export GO111MODULE=on 
go mod init

从模块的根目录执行时，./... 模式匹配当前模块中的所有包。 go build 将根据需要自动添加缺失或未转换的依赖项，以满足此特定构建调用的导入：
$ go build ./...

测试模块
$ go test ./...