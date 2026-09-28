[简体中文](README.md) | [English](README.en.md)

# go-web-admin

![Version](https://img.shields.io/badge/release-0.1.1-blue.svg)
![Language](https://img.shields.io/badge/language-golang1.12-blue.svg)
![base](https://img.shields.io/badge/env-goframe1.11.1-red.svg)

> A Go web admin boilerplate built on the GoFrame framework, implementing login authentication, RBAC permission control and other essential backend components — ready to use as a skeleton for admin projects.

## 📖 Introduction

go-web-admin is built on the [GoFrame](https://github.com/gogf/gf) framework and implements the basic components of a Go web backend. Business code is organized in api / service / model layers, with built-in JWT login authentication and casbin-based permission control that automatically associates users, roles and menus.

It works both as a starting point for Go admin projects and as a reference for learning GoFrame layering plus JWT/casbin RBAC integration.

## ✨ Features

- Login authentication with JWT verification (global middleware)
- Permission control via casbin: users are linked to roles, roles to menus (path + method); policies load automatically at startup and are rebuilt on change
- CRUD for users, roles and menus
- The `admin` user has full access without permission matching; `/token` and `/userInfo` endpoints are excluded from verification
- Global CORS middleware, unified response wrapper and pagination helper

Permission model (`config/rbac_model.conf`): `role(role.name, menu.path, menu.method)`, `user(user.username, role.name)`. For example, user hequan belongs to the test role which holds `GET /api/v1/users`, so his request to that path passes the check.

## 🛠 Tech Stack

| Component | Version | Purpose |
| --- | --- | --- |
| GoFrame | v1.11.2 | web framework |
| gorm | v1.9.10 | ORM |
| casbin | v2.1.2 | permission control |
| jwt-go | v3.2.0 | login token |
| MySQL | - | database |
| sha1 | - | password hashing |

## 🚀 Quick Start

1. Deploy MySQL and create the database `go_web_admin`
2. Import `docfile/sql/go_web_admin.sql`
3. Edit the config file `config/config.toml`:

```toml
[setting]
    logpath = "log/"        # log directory
    JwtSecret = "123456789" # JWT signing secret
    PageSize = 10           # default page size

[database]
    host = "192.168.100.50:3306"
    user = "root"
    pass = "123456"
    name = "go_web_admin"
    type = "mysql"
    TablePrefix = "go_"
```

```bash
go run main.go   # listens on :8000, default account admin / 123456
```

## 📡 API

Both requests and responses use JSON. After obtaining a token via login, send it in the `Authorization: Token xxxxxxxx` header.

| Endpoint | Method | Description |
| --- | --- | --- |
| `/token` | POST | login and get a token (no verification) |
| `/userInfo` | GET | current user info (no verification) |
| `/api/v1/users/*id` | GET / POST / PUT / DELETE | user CRUD |
| `/api/v1/roles/*id` | GET / POST / PUT / DELETE | role CRUD |
| `/api/v1/menus/*id` | GET / POST / PUT / DELETE | menu CRUD |

Response codes: `200` success, `201` created/updated, `204` deleted, `400` bad params, `401` not logged in, `403` forbidden, `404` not found, `500` server error.

## 📁 Directory Structure

```
- app     business layer
    - api       API handlers (parse request params)
    - model     data models (database access)
    - service   business logic layer
- boot        project initialization
- config      config files (config.toml, rbac_model.conf)
- docfile     docs (SQL scripts)
- library     shared libraries (jwt, permission, response, e, util, inject)
- router      unified route registration
- test        unit tests
```

## 📄 License

[MIT](LICENSE)

## 👤 Author

- 何全 (Hequan)
