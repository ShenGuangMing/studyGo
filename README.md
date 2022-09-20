# 环境开发
## 安装Go
- [下载链接](https://studygolang.com/dl)
- windows留意安装路径
## 配置环境变量
- 添加GOROOT路径，value是go所在的bin上层目录
- 在Path中引用，%GOROOT%\bin
> **注意**
> 
> Go安装的时候默认配置了GOROOT、GOPATH环境变量，GOPATH可以根据自己情况进行修改
> 
## 环境变量的含义
- GOROOT是go的安装目录，go原生工具就在该目录下
- GOPATH通常存放自己开发的代码或第三方依赖
- GO111MODULE=on go会忽略GOPATH和vendor文件夹，只根据go.mod下载依赖。从go
1.16开始默认值就是on
- GOPROXY：下载依赖库时走那个镜像代理，可以公司内部自建镜像
- PATH下的二进制文件可以在任意目录下运行
- 在$go_path目录创建三个子目录：src、bin、pkg
## Go Modules依赖包查找机制
- 下载第三方依赖存储在$go_path/pkg/mod
- go install生成的可执行文件存储在$go_path/bin
- 依赖包的查找顺序
  - 工作目录
  - $go_path/pkg/mod
  - $go_path/src

# go程序
## Hello world
```go
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}

```
## 引入外部包
```go
package main

import (
	"basic/src/function"
	"fmt"
	"gonum.org/v1/gonum/stat"
)

func main() {
	a := 10
	b := 20
	c := function.Add(a, b)
	fmt.Println("hello world", c)
	arr := []float64{1, 2, 3, 4, 5, 6}
	v := stat.Variance(arr, nil)
	fmt.Printf("方差=%f\n", v)
}
```

# go命令介绍
## go常用命令
- go help：查看帮助文档
  - go help build
- go build：对源码和依赖的文件进行打包，生成可执行文件
  - go build -o my_first_go_exe entrance_class/demo.go
- go install：编译并安装包或依赖，安装到$go_path/bin下
  - go install entrance_class/demo.go
- go get：把依赖库添加到当前module中，如果本机之前从未下载过则先下载并安装（install）
  - go get github.com/tinylib/msgp 会在go_path/pkg/mod目录下
生成gitgub.com/tinylib/msgp目录，同时在go_path/bin下生成msgp可执行文件
- go mod: module相关命令
  - go module init module name
  - go mod tipy通过扫描当前项目中的所有代码来添加未被记录的依赖至go.mod文件文件中删除不再被使用的依赖
- go run：编译并运行程序
- go test：执行测试代码
- go tool：执行go自带的工具
  - go tool pprof对cpu、内存和协程进行监控
  - go tool trace跟踪协程的执行过程
- go vet：检查代码中的静态错误
- go fmt：对代码文件进行格式化，如果用了IDE这个命令就不需要了
  - go fmt entrance class/demo.go
- go doc：查看go标准库或第三方库的帮助文档
  - go doc fmt
  - go doc gonum.org/v1/gonum/stat
- go version：

# 标识符和关键字
## 命名方式
go变量、常量、自定义类型、包、函数的命名方式必须遵循一下规则
- 首字符可以是任意Unicode字符火下划线
- 首字符之外部分可以式Unicode字符、下划线和数字
- 名字的长度无限制
> 理论上名字里可以有汉字，甚至可以全是汉字，但实际不要这样做
> 
## 关键字

|       |         |        |             |          |
|-------|---------|--------|-------------|----------|
| break | default | func   | interface   | select   |
| case  | defer   | go     | map         | struct   |
| chan  | else    | goto   | package     | switch   |
| const | if      | range  | type        | continue |
| for   | import  | return | fallthrough | var      |

## 保留字
### 常量
true false iota nil
### 数据类型
int int8 int16 int32  int64
uint uint8 uint16 uint32 uint65
uintptr float32 float64 complex128 complex64
bool byte rune string error
### 函数
make len cap new append copy close delete complex real image panic recover

## 操作符与表达式
### 算术运算符

| 运算符 | 描述  |
|-----|-----|
| +   | 相加  |
| -   | 相见  |
| *   | 相乘  |
| /   | 相除  |
| %   | 求余  |
### 位运算

| 运算符  | 描述                      |
|------|-------------------------|
| （&）  | 参与运算的两个数对应的二进制相与（同1才为1） |
| （或）  | 有一个1为1                  |
| （^）  | 异或                      |
| （<<） | 左移位                     |
| （>>） | 右移位                     |

