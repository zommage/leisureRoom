# leisureRoom
leisureRoom

#### build
```
./controll.sh build
```

### dev 环境
#### 镜像打包
```
docker build --build-arg config=./conf/app.dev.ini -t leisureroom .
```

#### 停止进程
```
docker stop leisureroom && docker rm -f leisureroom
```

#### docker 启动
```
docker run --restart=always --name leisureroom  -idt -p 17007:17007 -v /var/log/ruok/leisureroom:/home/leisureroom/logs  leisureroom:latest
```

#### 查看日志
```
tailf /var/log/ruok/leisureroom/leisureroom.log
```
