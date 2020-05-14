# 日志服务

定位：只是收集服务业务的重要日志。用来报警提醒。

## 提供http和GRPC接口

- 记录所有其它服务器日志
- 对外提供接口查询记录日志

## 规范日志格式

级别|服务名|时间秒|其它

## 存储方式

- 文件存储，7天；

## 性能指标


## 参考资源

[业界日志系统架构](https://juejin.im/post/5d79cbc96fb9a06b0b1ca2df)
[有赞百亿级日志系统架构设计](https://www.infoq.cn/article/eVg_NLKEL6eD8WQwltrQ)
[审计日志格式](http://publib.boulder.ibm.com/tividd/td/TWS/SH19-4555-00/zh_CN/HTML/ig_mst114.htm)
[最佳日志实践](https://zhuanlan.zhihu.com/p/27363484)

## beat

日志收集

## search

日志查询

## 运行

[查看app/sv-log/readme.md](../../app/sv-log/readme.md)