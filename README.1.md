# platformDrawConsole
开奖分发控制器

#### build
```
./controll.sh build
```

### dev 环境
#### 镜像打包
```
docker build --build-arg config=./conf/app.dev.ini -t platdrawconsole .
```

#### 停止进程
```
docker stop platdrawconsole && docker rm -f platdrawconsole
```

#### docker 启动
```
docker run --restart=always --name platdrawconsole  -idt -p 15850:15850 -v /var/log/ruok/platdrawconsole:/home/platformdrawconsole/logs  platdrawconsole:latest
```

#### 查看日志
```
tailf /var/log/ruok/platdrawconsole/drawconsole.log
```
