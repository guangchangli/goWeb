# goWeb

### 1.go mod

 包管理，解决第三方依赖

```
add go mod
```

```
GO111MODULE off on auto
```

1. Off 

   ```
   寻找依赖包的方式将会沿用旧版本那种通过vendor目录或者GOPATH模式来查找
   ```

2. on

   ```
   使用go modules 不会去 gopath 寻找
   ```

3. auto

   ```
   go命令行将会根据当前目录来决定是否启用module功能。这种情况下可以分为两种
   当前目录在GOPATH/src之外且该目录包含go.mod文件
   当前文件在包含go.mod文件的目录下面。
   ```

```
当modules 功能启用时
依赖包的存放位置变更为$GOPATH/pkg
允许同一个package多个版本并存，且多个项目可以共享缓存的 module。
```

###  2. go mod 命令

- download  download modules to local cache(下载依赖包)
- edit  edit go.mod from tools or scripts（编辑go.mod）
- graph print module requirement graph (打印模块依赖图)
- Init initialize new module in current directory（在当前目录初始化mod）
- tidy add missing and remove unused modules(拉取缺少的模块，移除不用的模块)
- vendor make vendored copy of dependencies(将依赖复制到vendor下)
- verify verify dependencies have expected content (验证依赖是否正确）
- why explain why packages or modules are needed(解释为什么需要依赖)

### 3. go config in mac

```
export GOPROXY="https://mirrors.aliyun.com/goproxy/"
export GO111MODULE="auto"
export GOROOT=/usr/local/Cellar/go/1.13.7/libexec
export GOPATH=/Users/lgc/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
export GOBIN=$GOPATH/bin
```

