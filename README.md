# kotkit
玩转tiktok

# 前置条件
1. 安装go
2. 安装mysql，本地3306端口
3. 安装minio，本地9000端口（需要本地设置secretAccessKey）
4. 安装ffmpeg
5. 安装etcd，本地2379端口

# 如何运行
1. 在cmd/api目录下启动http服务
```shell
go run main.go
```

2. 在cmd/user目录下执行（Linux/MacOS）
```shell
sh build.sh
sh output/bootstrap.sh
```

# 如何测试
本地下载Postman进行接口测试