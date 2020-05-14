# Bazel 构建


## 检查工具

bazel --version
source ~/.bash_profile

## 项目构建
```shell script
bazel build //app/api-cms:api-cms
bazel run //app/api-cms:api-cms

bazel run //app/api-www:api-www

bazel run //app/gateway:gateway

bazel run //app/auth-rbac:auth-rbac

```

## 资源

[资源1](https://juejin.im/post/5d2f31b9f265da1b6a34cae0)

[go构建](https://zhuanlan.zhihu.com/p/95998597)

- 构建

```shell script
使用gazelle生成build文件
使用生成的build文件构建项目代码
# 使用gazelle自动生成build文件
bazel run //:gazelle

# 使用生成的build文件构建
# 构建根目录下所有target
bazel build //...
# 构建特定taget
bazel build //path/of/target:target_name

```

- 清理（无法加载包时，就要清理下）

```shell script
bazel clean --expunge
```

- 自动添加外部依赖

用 bazel run //:gazelle update-repos repo-uri 可以从 go.mod 导入对应依赖包。

比如想要往项目中增加 gin 的 segmentio 的 go client 包，只需要在项目根目录下执行命令：

```shell script
 bazel run //:gazelle update-repos github.com/gin-gonic/gin
```

- Gazelle 便会自动增加一条依赖到 WORKSPACE 文件：

```shell script
go_repository(
    name = "com_github_gin_gonic_gin",
    importpath = "github.com/gin-gonic/gin",
    sum = "h1:88crIK23zO6TqlQBt+f9FrPJNKm9ZEr7qjp9vl/d5TM=",
    version = "v1.6.2",
)
```


## 常见bug,没找到的包，在下面加上了

```shell script
 bazel run //:gazelle update-repos github.com/mattn/go-isatty
 bazel run //:gazelle update-repos github.com/gin-contrib/sse
# 找不到包 @com_github_go_playground_validator_v10
bazel run //:gazelle update-repos github.com/go-playground/validator/v10
bazel run //:gazelle update-repos gopkg.in/yaml.v2
bazel run //:gazelle update-repos github.com/ugorji/go
bazel run //:gazelle update-repos github.com/ugorji/go/codec
bazel run //:gazelle update-repos github.com/go-playground/universal-translator
bazel run //:gazelle update-repos github.com/leodido/go-urn
bazel run //:gazelle update-repos github.com/go-playground/locales

```
