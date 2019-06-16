### 安装`Gin`
```bash
go get -u -v github.com/gin-gonic/gin
go get -u -v github.com/satori/go.uuid
go get -u -v github.com/jinzhu/gorm
go get -u -v github.com/dgrijalva/jwt-go
go get -u -v github.com/uber-go/zap
go get -u -v github.com/pkg/errors
go get -u -v github.com/stretchr/testify
go get -u -v gopkg.in/natefinch/lumberjack.v2
#https://github.com/natefinch/lumberjack
mv $GOPATH/src/github.com/uber-go/zap $GOPATH/src/go.uber.org/zap
```
#### 初始化项目目录
```
gin-blog/
├── conf
├── middleware
├── models
├── pkg
├── routers
```
- conf：用于存储配置文件
- middleware：应用中间件
- models：应用数据库模型
- pkg：第三方包
- routers 路由逻辑处理
