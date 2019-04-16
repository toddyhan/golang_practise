# go modules管理依赖

## golang版本要求
go version在1.11以上

## 激活go modules
在命令行下，执行：
```bash
export GO111MODULE=on
```
由于很多包被墙（ golang.org/x/net等x下的包，cloud.google.com/下的包等 )，需要设置代理，比较便捷的方法是设置GOPROXY：
```bash
export GOPROXY=https://goproxy.io
```
当然如果自己有vpn或ss，也可以自行设置，总之能在go mod时下载到所有依赖包即可。
## go mod使用方法

目前网上搜到的方法基本都是demo级别的例子，不过基本够用。开发过程中用到的包基本都能直接mod到。自己开发的包也不一定必须加上github.com前缀，毕竟并不是所有公司的产品代码都会开源。

可能大部分人的情况是，早就用vendor或godep等其他第三方依赖管理工具进行管理了。现在1.11之后的版本发布了go mod功能，想从这些第三方工具切过来（不是必须的，完全取决于自己的需要）。这是后比较简单的方式是。

在代码目录下执行init命令，生成mod和sum文件：
```bash
go mod init test
```
生成的mod文件中，第一行指定了你自己包的名称。但此时go.mod文件是空的，并没有真正下载依赖，可以执行tidy命令，出发下载和更新依赖：
```bash
go mod tidy
```
执行成功无报错的话，再次打开文件，可以看到依赖及版本/tag等信息了。

此时可以go build，编译自己的代码啦。
## go mod优势

* 代码里面需要哪个包，尽管用。go mod给你解决依赖问题，全自动。
* 不用像vendor一样将依赖放到指定目录。go mod默认将依赖下载到cache目录下（$GOPATH/pkg/mod/cache），且一旦下载后，全居都可以用，不管你是否在$GOPATH下。之前团队内部用vendor时，由于vendor比较大，全量上库会导致gerrit操作非常慢。习惯做法是把vendor依赖放到一个远程路径下，保证大家都能获取到，每次换了编译环境，都要重新手动把vendor放到指定目录下才能用。现在用go mod了，不再存在这个问题了。
* 可以在团队内部建立golang私服了。利用replace功能，把依赖全部替换为私服地址获取，可大大提升获取速度，提高写作效率。

## TODO

* 第一次执行时，由于需要下载的依赖太多，go mod卡住了。ctrl-c取消掉，再次执行后成功。看来还不是很稳定（go1.12.4)
* 使用GOPROXY的方式，速度还是太慢了。
* 