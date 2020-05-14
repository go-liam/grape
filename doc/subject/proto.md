# info

## run更新

```shell 
#安装protobuf
cd proto
./protoc.sh

# 安装golang for protobuf插件
go get github.com/golang/protobuf/protoc-gen-go
go get -u -v github.com/golang/protobuf/protoc-gen-go
```


## 生成文件

```shell
source ~/.bash_profile

./proto/**/protoc.sh
```