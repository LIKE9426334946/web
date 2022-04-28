package main

/*
GOROOT和GOPATH
GOROOT就是Go安装的根目录，GOPATH就是Go主要的工作目录，go install和go get等会用到GOPATH环境变量
GOPATH下有三个重要目录，bin pkg src
bin:Go编译可执行文件的存放路径
pkg:Go编译包时生成的.a中间文件存放路径
src:Go标准库源码路径

1.17.1版本
修改go env的GOPROXY
go env -w GOPROXY="https://goproxy.cn,direct"
go mod init demo
先初始化一个mod文件，然后编写main.go文件，之后go mod tidy生成依赖关系文件

windows显示环境变量
echo %GOROOT%

go env 显示Go环境变量的值

*/
