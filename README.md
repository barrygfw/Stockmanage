### 基于GoLang 的gin框架实现的库存管理 - RESTfulAPI api系统

orm 采用 gorm  
使用 jwt 鉴权

#### 目录结构
```sh
.
├── README.md
├── conf                        配置文件
│   └── app.ini
├── db                          mysql
│   └── db.go
├── middleware                  中间件相关
│   └── jwt                     jwt中间件
├── models                      orm
│   └── 
├── keycenter                   keycenter工具
├── log                         日志，对logz的进一步封装，增加logid
│   └── log.go
├── pkg                         基础包
│   └── e                       错误信息包
│       ├── code.go             错误码
│       ├── msg.go              相关错误信息
│   └── logging                 日志包
│   └── setting                  
│       ├── setting.go          配置处理
│   └── util                    工具包
├── routers        
│   └──                         路由管理
├── runtime        
│   └──                         运行缓存
├── service        
│   └──                         业务逻辑包
```