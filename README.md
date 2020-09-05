## Gin+mongodb的后端接口服务程序

> 这是一个简单的Gin+Mongodb后端服务程序。
>首先，自己需要搭建一个mongodb数据库。
>修改配置config.default.yaml中的内容，并修改文件为config.yaml,
>最后运行go程序。


## 插件列表
```js
    github.com/aws/aws-sdk-go 
    github.com/dgrijalva/jwt-go
    github.com/gin-gonic/gin 
    github.com/go-playground/validator/v10 
    github.com/golang/protobuf
    github.com/json-iterator/go 
    github.com/klauspost/compress 
    github.com/kr/text 
    github.com/modern-go/concurrent 
    github.com/modern-go/reflect2 
    github.com/niemeyer/pretty 
    github.com/rs/zerolog 
    github.com/stretchr/testify 
    go.mongodb.org/mongo-driver 
    golang.org/x/crypto 
    golang.org/x/sync 
    golang.org/x/sys 
    google.golang.org/protobuf 
    gopkg.in/check.v1 
    gopkg.in/yaml.v2 
```

## 目录结构
```js
.
├── app
│   ├── api
│   ├── controllers
│   └── handler
├── middleware
├── models
└── utils
```

## 接口测试
### 1. 
### 2. getHTML
```js
url: http://localhost:5922/v1/getHTML
method: POST
type: form-data
params: 
{ url: http://www.baidu.com }
```
### 3. 