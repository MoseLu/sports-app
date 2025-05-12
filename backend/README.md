# 运动打卡应用后端

## 项目概述

这是一个基于 Go 语言开发的运动打卡应用后端服务，使用 Gin 框架和 GORM ORM。

## 技术栈

- Go 1.21+
- Gin Web 框架
- GORM ORM
- MySQL 数据库
- JWT 认证

## 项目结构

```
backend/
├── config/         # 配置文件
├── controllers/    # 控制器
├── models/         # 数据模型
├── utils/          # 工具函数
├── database/       # 数据库相关
├── main.go         # 入口文件
└── README.md       # 项目文档
```

## 功能特性

- 用户认证（注册、登录、登出）
- JWT token 认证
- 密码加密存储
- 数据库连接池优化
- 高性能查询优化

## 环境要求

- Go 1.21 或更高版本
- MySQL 5.7 或更高版本
- 至少 1GB 可用内存

## 快速开始

### 1. 安装依赖

```bash
go mod download
```

### 2. 配置数据库

创建 MySQL 数据库：

```sql
CREATE DATABASE sports_app CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 3. 配置环境变量（可选）

```bash
export DB_HOST=localhost
export DB_PORT=3306
export DB_USER=root
export DB_PASSWORD=123456
export DB_NAME=sports_app
```

### 4. 运行服务

```bash
go run main.go
```

服务将在 http://localhost:8080 启动

## API 文档

### 认证接口

#### 用户注册

```http
POST /api/auth/register
Content-Type: application/json

{
    "username": "testuser",
    "password": "123456",
    "email": "test@example.com"
}
```

#### 用户登录

```http
POST /api/auth/login
Content-Type: application/json

{
    "username": "testuser",
    "password": "123456"
}
```

#### 用户登出

```http
POST /api/auth/logout
Authorization: Bearer <token>
```

## 性能优化

- 使用原生 SQL 查询优化性能
- 配置数据库连接池
- 优化密码加密过程
- 减少不必要的日志记录

## 开发规范

- 遵循 Go 代码规范
- 使用有意义的变量名
- 添加必要的注释
- 编写单元测试

## 部署说明

1. 编译项目：

```bash
go build -o sports-app
```

2. 运行服务：

```bash
./sports-app
```

## 维护者

- 项目负责人：[你的名字]
- 联系方式：[你的邮箱]

## 许可证

MIT License

# 运动应用后端 API 文档

## 认证相关 API

### 用户注册

- **URL**: `/register`
- **Method**: `POST`
- **描述**: 注册新用户
- **请求体**:

```json
{
  "username": "string", // 用户名
  "password": "string", // 密码
  "email": "string" // 邮箱
}
```

- **响应**:

```json
{
  "message": "用户创建成功"
}
```

### 用户登录

- **URL**: `/login`
- **Method**: `POST`
- **描述**: 用户登录
- **请求体**:

```json
{
  "username": "string", // 用户名
  "password": "string" // 密码
}
```

- **响应**:

```json
{
  "message": "登录成功",
  "token": "string", // JWT token
  "user": {
    "id": "number", // 用户ID
    "username": "string", // 用户名
    "email": "string" // 邮箱
  }
}
```

### 用户登出

- **URL**: `/api/logout`
- **Method**: `POST`
- **描述**: 用户登出
- **认证**: 需要 Bearer Token
- **响应**:

```json
{
  "message": "登出成功"
}
```

## 用户相关 API

### 获取用户信息

- **URL**: `/api/users/:username`
- **Method**: `GET`
- **描述**: 获取指定用户的基本信息
- **认证**: 需要 Bearer Token
- **响应**:

```json
{
  "id": "number", // 用户ID
  "username": "string" // 用户名
}
```

### 获取个人资料

- **URL**: `/api/users/profile`
- **Method**: `GET`
- **描述**: 获取当前登录用户的个人资料
- **认证**: 需要 Bearer Token
- **响应**:

```json
{
  "id": "number", // 用户ID
  "username": "string", // 用户名
  "email": "string" // 邮箱
}
```

### 更新个人资料

- **URL**: `/api/users/profile`
- **Method**: `PUT`
- **描述**: 更新当前登录用户的个人资料
- **认证**: 需要 Bearer Token
- **请求体**:

```json
{
  "email": "string" // 新邮箱
}
```

- **响应**:

```json
{
  "id": "number", // 用户ID
  "username": "string", // 用户名
  "email": "string" // 更新后的邮箱
}
```

## 运动记录相关 API

### 获取运动记录列表

- **URL**: `/api/records`
- **Method**: `GET`
- **描述**: 获取当前用户的所有运动记录
- **认证**: 需要 Bearer Token
- **响应**:

```json
[
  {
    "id": "number", // 记录ID
    "user_id": "number", // 用户ID
    "sport_type": "string", // 运动类型
    "duration": "number", // 运动时长(分钟)
    "calories": "number", // 消耗卡路里
    "created_at": "string" // 创建时间
  }
]
```

### 创建运动记录

- **URL**: `/api/records`
- **Method**: `POST`
- **描述**: 创建新的运动记录
- **认证**: 需要 Bearer Token
- **请求体**:

```json
{
  "sport_type": "string", // 运动类型
  "duration": "number", // 运动时长(分钟)
  "calories": "number" // 消耗卡路里
}
```

- **响应**:

```json
{
  "id": "number", // 记录ID
  "user_id": "number", // 用户ID
  "sport_type": "string", // 运动类型
  "duration": "number", // 运动时长(分钟)
  "calories": "number", // 消耗卡路里
  "created_at": "string" // 创建时间
}
```

### 更新运动记录

- **URL**: `/api/records/:id`
- **Method**: `PUT`
- **描述**: 更新指定ID的运动记录
- **认证**: 需要 Bearer Token
- **请求体**:

```json
{
  "sport_type": "string", // 运动类型
  "duration": "number", // 运动时长(分钟)
  "calories": "number" // 消耗卡路里
}
```

- **响应**:

```json
{
  "id": "number", // 记录ID
  "user_id": "number", // 用户ID
  "sport_type": "string", // 运动类型
  "duration": "number", // 运动时长(分钟)
  "calories": "number", // 消耗卡路里
  "created_at": "string" // 创建时间
}
```

### 删除运动记录

- **URL**: `/api/records/:id`
- **Method**: `DELETE`
- **描述**: 删除指定ID的运动记录
- **认证**: 需要 Bearer Token
- **响应**:

```json
{
  "message": "记录已删除"
}
```

### 获取运动统计

- **URL**: `/api/records/stats`
- **Method**: `GET`
- **描述**: 获取当前用户的运动统计信息
- **认证**: 需要 Bearer Token
- **响应**:

```json
{
  "total_duration": "number", // 总运动时长(分钟)
  "exercise_count": "number", // 运动次数
  "average_duration": "number", // 平均运动时长(分钟)
  "average_calories": "number", // 平均消耗卡路里
  "daily_duration": [
    // 每日运动时长(分钟)
    "number",
    "number",
    "number",
    "number",
    "number",
    "number",
    "number"
  ],
  "daily_count": [
    // 每日运动次数
    "number",
    "number",
    "number",
    "number",
    "number",
    "number",
    "number"
  ]
}
```

## 错误响应格式

所有 API 在发生错误时都会返回以下格式的响应：

```json
{
  "error": "string", // 错误描述
  "code": "string", // 错误代码
  "details": "string" // 详细错误信息（可选）
}
```

### 常见错误代码

- `AUTH_REQUIRED`: 需要认证
- `INVALID_TOKEN_FORMAT`: 无效的token格式
- `INVALID_TOKEN`: 无效的token
- `UNAUTHORIZED`: 未授权
- `NOT_FOUND`: 资源不存在
- `BAD_REQUEST`: 请求参数错误
- `INTERNAL_ERROR`: 服务器内部错误
