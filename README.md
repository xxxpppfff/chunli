# X-bulletproof-backend

一个基于 Go + Gin + GORM 的现代化后端 API 服务。

## 项目结构

```
.
├── config/          # 配置管理
│   └── config.go
├── controllers/     # 控制器层（处理 HTTP 请求）
│   └── user_controller.go
├── data/            # 数据操作示例
├── global/          # 全局变量（数据库连接等）
│   └── enter.go
├── middleware/      # 中间件（CORS、日志、错误处理）
│   ├── cors.go
│   ├── logger.go
│   └── recovery.go
├── models/          # 数据模型
│   └── user_model.go
├── routes/          # 路由配置
│   └── routes.go
├── services/        # 服务层（业务逻辑）
│   └── user_service.go
├── utils/           # 工具函数
│   └── response/    # 统一响应格式
│       └── response.go
├── go.mod
├── go.sum
└── main.go          # 程序入口
```

## 功能特性

- ✅ RESTful API 设计
- ✅ 统一的响应格式
- ✅ CORS 跨域支持
- ✅ 请求日志记录
- ✅ 错误恢复处理
- ✅ 数据库自动迁移
- ✅ 配置管理（支持环境变量）
- ✅ 分页查询支持

## 快速开始

### 1. 安装依赖

```bash
go mod download
```

### 2. 配置数据库

修改 `config/config.go` 中的默认数据库配置，或通过环境变量设置：

```bash
export DB_HOST=127.0.0.1
export DB_PORT=3306
export DB_USER=root
export DB_PASSWORD=123456
export DB_NAME=go_test
export SERVER_PORT=8080
```

### 3. 运行程序

```bash
go run main.go
```

服务器将在 `http://localhost:8080` 启动。

## API 文档

### 健康检查

```
GET /health
```

### 用户管理 API

#### 创建用户
```
POST /api/users
Content-Type: application/json

{
  "name": "张三",
  "age": 25,
  "e_mail": "zhangsan@example.com"
}
```

#### 获取用户列表（支持分页）
```
GET /api/users?page=1&pageSize=10
```

#### 获取用户详情
```
GET /api/users/{id}
```

#### 更新用户
```
PUT /api/users/{id}
Content-Type: application/json

{
  "name": "李四",
  "age": 30,
  "e_mail": "lisi@example.com"
}
```

#### 删除用户
```
DELETE /api/users/{id}
```

## 响应格式

### 成功响应
```json
{
  "code": 200,
  "message": "success",
  "data": {}
}
```

### 错误响应
```json
{
  "code": 400,
  "message": "错误信息"
}
```

## 技术栈

- **Web 框架**: Gin
- **ORM**: GORM
- **数据库**: MySQL
- **语言**: Go 1.23+

## 开发建议

1. 添加新的功能模块时，按照 MVC 架构：
   - `models/` - 定义数据模型
   - `services/` - 实现业务逻辑
   - `controllers/` - 处理 HTTP 请求
   - `routes/` - 注册路由

2. 使用统一的响应格式 `utils/response`

3. 通过环境变量管理配置，便于不同环境部署

4. 数据库操作统一使用 `global.DB`

## License

MIT