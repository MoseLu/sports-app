# 项目架构设计

## 系统架构

### 整体架构

```
sports-app/
├── backend/          # 后端服务
│   ├── config/       # 配置文件
│   ├── controllers/  # 控制器
│   ├── models/       # 数据模型
│   ├── routes/       # 路由配置
│   ├── utils/        # 工具函数
│   ├── database/     # 数据库相关
│   └── main.go       # 入口文件
├── src/              # 前端应用
│   ├── assets/       # 静态资源
│   ├── boot/         # 启动配置
│   ├── components/   # 组件
│   ├── css/          # 样式文件
│   ├── i18n/         # 国际化
│   ├── layouts/      # 布局组件
│   ├── pages/        # 页面组件
│   │   ├── auth/     # 认证相关页面
│   │   ├── community/# 社区相关页面
│   │   ├── error/    # 错误页面
│   │   ├── exercise/ # 运动相关页面
│   │   └── user/     # 用户相关页面
│   ├── router/       # 路由配置
│   ├── stores/       # 状态管理
│   └── App.vue       # 根组件
├── public/           # 公共资源
├── docs/             # 项目文档
└── logs/             # 项目日志
```

## 技术选型

### 后端技术栈

- 编程语言：Go 1.21+
- Web框架：Gin
- ORM框架：GORM
- 数据库：MySQL 5.7+
- 认证：JWT (RSA加密)
- 日志：标准库log

### 前端技术栈

- 框架：Vue 3 + TypeScript
- 状态管理：Pinia
- UI框架：Quasar
- 路由：Vue Router
- HTTP客户端：Axios
- 构建工具：Vite
- 图表库：ECharts

## 核心模块

### 1. 用户认证模块

- 功能：
  - 用户注册
  - 用户登录
  - JWT token管理
  - 密码加密存储
  - 密码重置
  - 邮箱验证码
- 技术实现：
  - RSA非对称加密
  - bcrypt密码哈希
  - 数据库索引优化
  - 验证码服务
  - 邮件服务（模拟）

### 2. 验证码服务模块

- 功能：
  - 验证码生成
  - 验证码存储
  - 验证码验证
  - 自动清理过期验证码
- 技术实现：
  - 数据库存储
  - 定时清理任务
  - 验证码加密
  - 邮件发送（模拟）

### 3. 运动记录模块

- 功能：
  - 运动记录创建
  - 运动记录查询
  - 运动记录统计
  - 运动数据可视化
- 技术实现：
  - 数据库事务处理
  - 数据缓存优化
  - 图表展示
  - 响应式布局
  - 表格固定列

### 4. 社区模块

- 功能：
  - 动态发布
  - 动态浏览
  - 点赞评论
  - 用户关注
- 技术实现：
  - 实时数据更新
  - 分页加载
  - 图片上传
  - 瀑布流布局

## 安全设计

### 1. 认证安全

- JWT使用RSA非对称加密
- 密码使用bcrypt加密存储
- Token过期机制
- 自动刷新Token
- 验证码有效期控制
- 验证码使用次数限制

### 2. 数据安全

- 数据库连接加密
- 敏感数据加密存储
- SQL注入防护
- XSS防护

### 3. 接口安全

- 请求参数验证
- 接口访问控制
- 错误信息处理
- 请求频率限制

## 性能优化

### 1. 前端优化

- 组件懒加载
- 路由懒加载
- 图片懒加载
- 数据缓存
- 表格性能优化
  - 固定列实现
  - 虚拟滚动
  - 数据分页
  - 响应式布局

### 2. 后端优化

- 数据库连接池
- 查询优化
- 缓存策略
- 并发控制
- 接口性能优化
  - 数据聚合
  - 批量处理
  - 异步处理
  - 结果缓存

## 部署架构

### 自动化部署流程

1. 代码提交触发 GitHub Actions 工作流
2. 工作流执行构建和测试
3. 构建成功后自动部署到阿里云服务器
4. 使用 systemd 管理应用进程
5. 自动重启服务以应用更新

### 部署架构图

```
[GitHub Repository] --> [GitHub Actions] --> [阿里云服务器]
                           |
                           v
                    [构建和测试] --> [部署] --> [systemd服务]
```

### 技术栈

- GitHub Actions: CI/CD 工具
- 阿里云服务器: 生产环境
- systemd: 进程管理
- 自动化部署脚本: 部署流程

### 部署流程

1. 代码提交到 master 分支
2. GitHub Actions 自动触发工作流
3. 执行构建和测试
4. 构建成功后自动部署
5. 使用 systemd 管理进程
6. 自动重启服务

### 监控和维护

- 部署状态监控
- 错误日志记录
- 性能监控
- 自动告警机制

### 开发环境

- 本地开发服务器
- 本地数据库
- 开发工具链
- 热重载支持

### 生产环境

- 应用服务器
- 数据库服务器
- 负载均衡
- 监控系统
- CDN加速

## 扩展性设计

### 1. 模块化设计

- 清晰的模块划分
- 低耦合高内聚
- 接口标准化
- 组件复用

### 2. 可扩展性

- 插件化架构
- 配置化设计
- 接口版本控制
- 主题定制

### 3. 性能扩展

- 水平扩展支持
- 缓存机制
- 异步处理
- 预加载策略

## 监控与运维

### 1. 日志系统

- 访问日志
- 错误日志
- 性能日志
- 操作日志
- 系统日志

### 2. 监控系统

- 性能监控
- 错误监控
- 用户行为监控
- 系统资源监控

### 3. 告警系统

- 错误告警
- 性能告警
- 安全告警
- 业务告警

# 项目架构文档

## 系统架构

### 前端架构

#### 1. 技术栈

- Vue 3 + TypeScript
- Quasar Framework
- Pinia 状态管理
- Vue Router
- ECharts 数据可视化

#### 2. 目录结构

```
src/
├── assets/          # 静态资源
├── boot/           # 启动文件
├── components/     # 通用组件
├── css/           # 全局样式
├── i18n/          # 国际化
├── layouts/       # 布局组件
├── pages/         # 页面组件
├── router/        # 路由配置
├── services/      # 服务层
├── stores/        # 状态管理
└── types/         # 类型定义
```

#### 3. 核心模块

##### 3.1 仪表盘模块

- ExerciseStats.vue：运动统计组件
  - 数据可视化
  - 时间范围筛选
  - 运动类型筛选
  - 响应式设计
- ExerciseList.vue：运动记录列表
  - 表格展示
  - 筛选功能
  - 暗色模式支持
  - 移动端适配

##### 3.2 状态管理

- exercise-store：运动记录状态
- user-store：用户状态
- app-store：应用状态

##### 3.3 路由系统

- 路由守卫
- 权限控制
- 页面过渡

### 后端架构

#### 1. 技术栈

- Go
- Gin Web框架
- GORM ORM框架
- MySQL数据库
- JWT认证

#### 2. 目录结构

```
backend/
├── config/        # 配置文件
├── controllers/   # 控制器
├── middleware/    # 中间件
├── models/        # 数据模型
├── routes/        # 路由配置
├── services/      # 业务逻辑
└── utils/         # 工具函数
```

#### 3. 核心模块

##### 3.1 运动记录模块

- 记录管理
  - 创建记录
  - 更新记录
  - 删除记录
  - 查询记录
- 统计分析
  - 时间范围统计
  - 运动类型统计
  - 每日统计
  - 趋势分析

##### 3.2 数据模型

```go
type Exercise struct {
    ID          int64     `json:"id"`
    UserID      int64     `json:"user_id"`
    SportTypeID int64     `json:"sport_type_id"`
    Duration    int       `json:"duration"`
    Distance    float64   `json:"distance"`
    Calories    int       `json:"calories"`
    Date        time.Time `json:"date"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type Stats struct {
    TotalDuration    int       `json:"total_duration"`
    ExerciseCount    int       `json:"exercise_count"`
    AverageDuration  float64   `json:"average_duration"`
    AverageCalories  float64   `json:"average_calories"`
    DailyDuration    []int     `json:"daily_duration"`
    DailyCount       []int     `json:"daily_count"`
}
```

## 安全设计

### 1. 认证与授权

- JWT认证
- 路由权限控制
- 数据访问控制

### 2. 数据安全

- 密码加密存储
- 敏感数据脱敏
- 数据验证

### 3. 接口安全

- 请求频率限制
- 参数验证
- 错误处理

## 性能优化

### 1. 前端优化

- 组件懒加载
- 图片优化
- 缓存策略

### 2. 后端优化

- 数据库索引
- 查询优化
- 连接池配置

## 部署架构

### 1. 开发环境

- 本地开发服务器
- 开发数据库
- 开发工具链

### 2. 生产环境

- Nginx反向代理
- MySQL主从复制
- 负载均衡

## 监控与运维

### 1. 日志系统

- 错误日志
- 访问日志
- 性能日志

### 2. 监控指标

- 接口响应时间
- 数据库性能
- 服务器资源

## 扩展性设计

### 1. 模块化

- 组件复用
- 服务解耦
- 接口标准化

### 2. 配置化

- 环境配置
- 功能开关
- 主题定制
