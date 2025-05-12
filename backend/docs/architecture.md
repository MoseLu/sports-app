# 体育应用后端架构文档

## 项目结构

```
sports-app/backend/
├── config/             # 配置相关
│   └── config.go      # 配置结构和加载逻辑
├── controllers/        # 控制器层
│   └── user.go        # 用户相关控制器
├── database/          # 数据库相关
│   └── db.go         # 数据库连接和初始化
├── docs/              # 文档
│   └── architecture.md # 本文档
├── models/            # 数据模型
│   └── user.go       # 用户模型
├── routes/            # 路由定义
│   └── routes.go     # 路由配置
├── go.mod            # Go模块定义
└── main.go           # 应用入口
```

## 数据库配置

项目使用MySQL作为主数据库，配置如下：

### 默认配置
- 主机：localhost
- 端口：3306
- 用户名：root
- 密码：（默认为空）
- 数据库名：sports_app
- 字符集：utf8mb4
- 时区：Local

### 环境变量配置
可以通过以下环境变量自定义数据库配置：
- `DB_HOST`：数据库主机地址
- `DB_PORT`：数据库端口
- `DB_USER`：数据库用户名
- `DB_PASSWORD`：数据库密码
- `DB_NAME`：数据库名称

### 连接字符串格式
```
{user}:{password}@tcp({host}:{port})/{dbname}?charset=utf8mb4&parseTime=True&loc=Local
```

## 用户模型

用户模型包含以下字段：
- `ID`：主键，自增
- `CreatedAt`：创建时间
- `UpdatedAt`：更新时间
- `DeletedAt`：删除时间（软删除）
- `Username`：用户名（唯一）
- `Password`：密码（加密存储）
- `Email`：邮箱（唯一）
- `Role`：用户角色
- `Timezone`：时区

### 密码处理
- 使用bcrypt进行密码加密
- 加密强度（cost）设置为14
- 在创建和更新用户时自动进行密码加密

## 服务器配置

- 默认端口：8080
- 可通过`SERVER_PORT`环境变量自定义

## 依赖管理

主要依赖：
- gin-gonic/gin：Web框架
- gorm.io/gorm：ORM框架
- gorm.io/driver/mysql：MySQL驱动
- golang-jwt/jwt：JWT认证
- golang.org/x/crypto：密码加密 

## 时间处理

### 时间同步机制
- 后端统一使用UTC时间存储和计算
- 用户模型增加时区字段，记录用户所在时区
- API响应中包含时间戳和时区信息
- 前端负责根据用户时区进行时间转换和显示

### 打卡提醒功能
- 后端职责：
  - 存储打卡时间规则
  - 提供打卡时间配置API
  - 记录用户打卡历史
  - 提供打卡时间查询接口

- 前端职责：
  - 实现本地提醒功能
  - 处理时区转换
  - 管理提醒设置
  - 离线提醒支持

### 时间相关API
- `GET /api/timezones`：获取支持的时区列表
- `GET /api/user/timezone`：获取用户时区设置
- `PUT /api/user/timezone`：更新用户时区设置
- `GET /api/checkin/rules`：获取打卡规则
- `POST /api/checkin/records`：提交打卡记录 