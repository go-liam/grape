# info

参考资源

[资源1](https://github.com/jaegertracing/jaeger/blob/master/examples/hotrod/docker-compose.yml)

[资源2](https://www.jaegertracing.io/download/#docker-images)

[资源3](https://github.com/xinliangnote/go-jaeger-demo)

## run

docker-compose up -d

## url

启动成功后，访问 http://localhost:16686 就可以看到 Jaeger UI。

```shell
curl http://localhost:16686
curl http://localhost:7041
curl http://localhost:6831
```

# 监控

http://localhost:14269/metrics
