[简体中文](README.md) | [English](README.en.md)

# go-web-admin

![版本](https://img.shields.io/badge/release-0.1.1-blue.svg)
![语言](https://img.shields.io/badge/language-golang1.12-blue.svg)
![base](https://img.shields.io/badge/env-goframe1.11.1-red.svg)

> 基于 GoFrame 框架的 Go Web 后台管理基础框架，实现登录认证 + RBAC 权限管理等后端基本组件，开箱即可作为后台类项目的骨架。

## 📖 项目介绍

go-web-admin 基于 [GoFrame](https://github.com/gogf/gf) 框架，完成 Go Web 后端的基本组件开发。项目按「api / service / model」分层组织业务代码，内置 JWT 登录认证，并利用 casbin 将 user、role、menu 自动关联实现权限验证。

适合作为 Go 后台管理类项目的起点，也适合学习 GoFrame 分层结构、JWT 与 casbin RBAC 的整合方式。

## ✨ 功能特性

- 登录认证，JWT 校验（全局中间件）
- 权限验证（casbin）：用户关联角色、角色关联菜单（菜单含 path + method），启动时自动加载权限，变更后自动重建
- 用户 user、权限组 role、菜单 menu 的增删改查
- admin 用户拥有全部权限，不做权限匹配；`/token`、`/userInfo` 接口免验证
- 全局 CORS 中间件、统一响应封装、分页工具

权限模型（`config/rbac_model.conf`）：`角色(role.name, menu.path, menu.method)`，`用户(user.username, role.name)`。例如用户 hequan 属于 test 组，test 组拥有 `GET /api/v1/users` 权限，则 hequan 请求该地址时校验通过。

## 🛠 技术栈

| 组件 | 版本 | 用途 |
| --- | --- | --- |
| GoFrame | v1.11.2 | Web 框架 |
| gorm | v1.9.10 | ORM |
| casbin | v2.1.2 | 权限验证 |
| jwt-go | v3.2.0 | 登录令牌 |
| MySQL | - | 数据库 |
| sha1 | - | 密码散列 |

## 🚀 快速开始

1. 部署 MySQL，创建库 `go_web_admin`
2. 导入 `docfile/sql/go_web_admin.sql`
3. 修改配置文件 `config/config.toml`：

```toml
[setting]
    logpath = "log/"        # 日志目录
    JwtSecret = "123456789" # JWT 签名密钥
    PageSize = 10           # 默认分页大小

[database]
    host = "192.168.100.50:3306"
    user = "root"
    pass = "123456"
    name = "go_web_admin"
    type = "mysql"
    TablePrefix = "go_"
```

```bash
go run main.go   # 监听 :8000，默认账户密码 admin / 123456
```

## 📡 API 说明

请求和响应均使用 JSON 格式。登录获取 token 后，在请求头携带 `Authorization: Token xxxxxxxx` 访问业务接口。

| 接口 | 方法 | 说明 |
| --- | --- | --- |
| `/token` | POST | 登录获取 token（免验证） |
| `/userInfo` | GET | 当前用户信息（免验证） |
| `/api/v1/users/*id` | GET / POST / PUT / DELETE | 用户增删改查 |
| `/api/v1/roles/*id` | GET / POST / PUT / DELETE | 角色增删改查 |
| `/api/v1/menus/*id` | GET / POST / PUT / DELETE | 菜单增删改查 |

统一响应状态码：`200` 请求成功、`201` 创建/修改成功、`204` 删除成功、`400` 参数错误、`401` 未登录、`403` 禁止访问、`404` 未找到、`500` 系统错误。

## 📁 目录结构

```
- app     业务逻辑层
    - api       业务接口（接收/解析用户输入参数）
    - model     数据模型（数据库操作）
    - service   业务逻辑封装层
- boot        项目初始化参数设置
- config      配置文件（config.toml、rbac_model.conf）
- docfile     项目文档（SQL 脚本）
- library     公共库（jwt、permission、response、e、util、inject）
- router      路由统一注册管理
- test        单元测试
```

## 📄 License

[MIT](LICENSE)

## 👤 作者

- 何全
