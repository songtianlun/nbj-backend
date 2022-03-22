# MineGin Server

a go web server template with gin.

## Usage

```bash
$ cp ./conf/config.sample.yaml ./conf/config.yaml
$  openssl rand 16 -hex
#  openssl rand 32 -hex
# 配置必选项，或根据提示配置必要的环境变量
$ go mod tidy
# use [air](https://github.com/cosmtrek/air)
$ curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
$ air
```

## Libs

 - [`gin`](https://gin-gonic.com) to web server.
 - [`gopsutil`](https://github.com/shirou/gopsutil) to system check.
 - [`viper`](https://github.com/spf13/viper) to read config.
 - [`zap`](https://github.com/uber-go/zap) to log.
 - [`lumberjack`](https://github.com/natefinch/lumberjack) to log rolling.
 - [`GORM`](https://gorm.io/zh_CN/) for ORM.
 - [`go-jwt`](https://github.com/dgrijalva/jwt-go) for JWT.


## Config

程序配置优先级 `环境变量` > `配置文件`，二选一即可。

## Limit

- 登陆使用 `邮箱` + `密码`

## ToDo

 - [x] 启动默认 `release` ，开发环境自动 `debug`；
 - [x] 端口配置 `PORT` 变量最高，其次 `MINEGIN_PORT`，再其次配置文件；
 - [x] 日志默认输出到终端，配置后输出到文件；
 - [x] 简单接口鉴权；
 - [x] 注册、登陆、修改、删除功能；
 - [x] 用户配置 CRUD.