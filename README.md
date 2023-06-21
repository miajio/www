# www
www.miajio.com 平台源码

miajio(米亚西奥)网是一个开源的网站,将采用bbs论坛模式及社区模式搭建平台业务

其商业模式待定

## 技术栈
平台使用gin作为web服务开发,前端使用原生Bootstrap5与JQuery结合

数据库使用mysql

缓存层将采用redis

后续将结合实际业务判定是否搭载minio进行文件管理

## 如何编译与启动
需要在你拥有mysql服务,请自行查找安装方式并将其安装

安装myusql数据库后你需要将 sql-resources 文件夹下的sql文件全部执行一遍

需要有go语言环境,请自行查找安装方式并将其安装,同时配置好go语言环境变量等

然后将该源码库下载到你本地,并进入该项目目录中

使用 go mod tidy 指令将该项目使用的第三方库安装下来 (如果你无法安装,请将你的 GO111MODULE 环境变量设置为 on 同时将你的 GOPROXY 环境变量设置为 https://goproxy.cn,direct)
```
go mod tidy
```

然后使用 go build -o main-linux .\main.go 指令将其编译为可执行文件(如果你当前电脑与你的服务器不是统一操作系统,那么请自行查找 go 语言编译方式将其进行编译)

然后基于你的服务信息修改 config.toml 文件,如果你不是使用https,那么你需要将 [server] 中的 useHttps 设定为false,否则你将需要设定下列三个参数
```
httpsKey="你的https秘钥"
httpsPem="你的https pem文件"
httpsHost="当用户访问的是你80端口时,你需要进行重定向的地址"
```

最后将编译出来的文件,main-linux, config.toml, static文件夹及内部文件, 均放到你服务器的同一文件夹下,然后使用启动命令启动即可

如果是windows可以直接双击启动

如果是linux则可以使用 ./main-linux 或者使用 nohup 进行后台启动

```
nohup ./main-linux &
```

## config.toml
```
[email]
name="miajio" # 发送者邮件名称
mailAddr="miajio@163.com" # 发送者邮件地址
smtpAddr="smtp.163.com:25"
hostAddr="smtp.163.com"
password="XXXXX" # 邮箱授权码

[mysql]
host="127.0.0.1:3306" # 数据库地址
user="root" # 数据库用户名
password="123456" # 数据库密码
database="miajio" # 数据库名
charset="utf8mb4" # 字符集
parseTime="True" # 是否格式化时间
loc="Local" # 时区

[redis]
host="127.0.0.1:6379" # redis 服务器地址
password="" # redis 密码
db=0 # 默认连接库

[server]
port=":8080" # 服务端口
useHttps=false # 是否使用https
httpsKey=""
httpsPem=""
httpsHost=""

```

