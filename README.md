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

# 人生的前三个go程序
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




