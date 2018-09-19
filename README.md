# backend

包含后台api和后台ui：

- 后台api使用go进行开发
- ui使用vue进行开发

## 后端技术栈

- 框架路由使用 [gin](https://github.com/gin-gonic/gin) 路由
- 中间件使用 [gin](https://github.com/gin-gonic/gin) 框架的中间件
- 数据库组件 [gorm](https://github.com/jinzhu/gorm)
- 文档使用 [swagger](https://swagger.io/) 生成
- 配置文件解析库 [viper](https://github.com/spf13/viper)
- 使用 [JWT](https://jwt.io/) 进行身份鉴权认证
- 校验器 [validator](https://gopkg.in/go-playground/validator.v9)  也是 gin 框架默认的校验器，当前最新是v9版本
- 第三方包管理工具 [govendor](https://github.com/kardianos/govendor)

## 前端技术栈

> 针对后端页面的技术栈

- [vue](https://cn.vuejs.org/)
- [element](http://element-cn.eleme.io/#/zh-CN)
