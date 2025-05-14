# Sports App / 运动打卡应用

A modern sports tracking application built with Quasar Framework and Go backend.
一个使用 Quasar 框架和 Go 后端构建的现代运动打卡应用。

## Features / 功能特点

### Frontend / 前端

- Modern and responsive UI / 现代化响应式界面
- Cross-platform compatibility / 跨平台兼容性
- Hot-code reloading for development / 开发热重载
- Built-in linting and formatting / 内置代码检查和格式化
- Production-ready build system / 生产环境构建系统
- Dark mode support / 暗色模式支持
- Internationalization / 国际化支持
- Data visualization / 数据可视化

### Backend / 后端

- RESTful API / RESTful API接口
- JWT Authentication / JWT认证
- Database integration / 数据库集成
- Email verification / 邮箱验证
- Data statistics / 数据统计
- Performance optimization / 性能优化

## Tech Stack / 技术栈

### Frontend / 前端

- Vue 3 + TypeScript
- Quasar Framework
- Pinia (State Management)
- Vue Router
- ECharts (Data Visualization)
- Axios (HTTP Client)

### Backend / 后端

- Go 1.21+
- Gin Web Framework
- GORM (ORM)
- MySQL 5.7+
- JWT (Authentication)

## Prerequisites / 环境要求

- Node.js (v14 or higher)
- Go 1.21 or higher
- MySQL 5.7 or higher
- npm or yarn package manager

## Installation / 安装

### Install the dependencies / 安装依赖

```bash
# Frontend / 前端
cd sports-app
yarn
# or
npm install

# Backend / 后端
cd backend
go mod download
```

## Development / 开发

### Start the frontend in development mode / 启动前端开发服务器

```bash
quasar dev
```

### Start the backend server / 启动后端服务器

```bash
cd backend
go run main.go
```

### Lint the files / 代码检查

```bash
yarn lint
# or
npm run lint
```

### Format the files / 代码格式化

```bash
yarn format
# or
npm run format
```

## Production / 生产环境

### Build the app for production / 构建生产版本

```bash
quasar build
```

### Deploy the backend / 部署后端

```bash
cd backend
go build
./sports-app
```

## Project Structure / 项目结构

```
sports-app/
├── backend/          # Backend service / 后端服务
│   ├── config/       # Configuration / 配置
│   ├── controllers/  # Controllers / 控制器
│   ├── models/       # Data models / 数据模型
│   ├── routes/       # Routes / 路由
│   ├── utils/        # Utilities / 工具函数
│   └── main.go       # Entry point / 入口文件
├── src/              # Frontend application / 前端应用
│   ├── assets/       # Static assets / 静态资源
│   ├── components/   # Components / 组件
│   ├── pages/        # Pages / 页面
│   ├── stores/       # State management / 状态管理
│   └── App.vue       # Root component / 根组件
├── docs/             # Documentation / 文档
└── logs/             # Logs / 日志
```

## Features in Development / 开发中的功能

- Community features / 社区功能
- Personal profile system / 个人资料系统
- Social features / 社交功能
- Data analysis / 数据分析

## License / 许可证

This project is licensed under the MIT License.
本项目采用 MIT 许可证。

# Triggering CI/CD workflow for master branch

# Another trigger for CI/CD workflow

# Final trigger for CI/CD workflow

测试提交
