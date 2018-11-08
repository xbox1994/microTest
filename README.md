# microTest

#### 安装本应用依赖
`dep ensure` 

#### gRPC文件编译
* 安装protoc(https://github.com/protocolbuffers/protobuf/releases)
* protoc-gen-micro(`go get -u github.com/micro/protoc-gen-micro`)

编译proto文件夹中的hello.proto文件命令：

`protoc --proto_path=proto --proto_path=vendor --micro_out=proto --go_out=proto hello.proto`

#### 服务注册中心
* consul(https://learn.hashicorp.com/consul/getting-started/install.html)

使用方式：`consul agent -dev`

#### api网关使用
`micro api`
