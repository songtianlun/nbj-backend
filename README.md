# MinePin-backend

Dev deploy with [render](https://render.com).

## Usage

```bash
cp ./conf/config.sample.yaml ./conf/config.yaml
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

程序配置优先级 `环境变量` > `配置文件`

## Limit

- 登陆可以使用 `邮箱/手机号` + `密码`


## ToDo

 - [x] 启动默认 `release` ，开发环境自动 `debug`；
 - [x] 端口配置 `PORT` 变量最高，其次 `MINEPIN_PORT`，再其次配置文件；
 - [x] 日志默认输出到终端，配置后输出到文件；
 - [x] 简单接口鉴权；
 - [ ] 注册、登陆、修改、删除功能。