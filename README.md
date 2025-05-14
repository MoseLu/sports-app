# Sports App / 运动打卡应用

<div align="center">
  <img src="https://img.shields.io/badge/Vue-3.x-brightgreen" alt="Vue 3.x">
  <img src="https://img.shields.io/badge/Quasar-2.x-blue" alt="Quasar 2.x">
  <img src="https://img.shields.io/badge/Go-1.21+-00ADD8" alt="Go 1.21+">
  <img src="https://img.shields.io/badge/License-MIT-yellow" alt="License">
</div>

<div align="center">
  <a href="#english">English</a> |
  <a href="#chinese">中文</a>
</div>

<div id="english">

A modern sports tracking application built with Quasar Framework and Go backend.
一个使用 Quasar 框架和 Go 后端构建的现代运动打卡应用。

## Features

### Frontend

- Modern and responsive UI
- Cross-platform compatibility
- Hot-code reloading for development
- Built-in linting and formatting
- Production-ready build system
- Dark mode support
- Internationalization
- Data visualization

### Backend

- RESTful API
- JWT Authentication
- Database integration
- Email verification
- Data statistics
- Performance optimization

## Tech Stack

### Frontend

- Vue 3 + TypeScript
- Quasar Framework
- Pinia (State Management)
- Vue Router
- ECharts (Data Visualization)
- Axios (HTTP Client)

### Backend

- Go 1.21+
- Gin Web Framework
- GORM (ORM)
- MySQL 5.7+
- JWT (Authentication)

## Prerequisites

- Node.js (v14 or higher)
- Go 1.21 or higher
- MySQL 5.7 or higher
- npm or yarn package manager

## Installation

### Install the dependencies

```bash
# Frontend
cd sports-app
yarn
# or
npm install

# Backend
cd backend
go mod download
```

## Development

### Start the frontend in development mode

```bash
quasar dev
```

### Start the backend server

```bash
cd backend
go run main.go
```

### Lint the files

```bash
yarn lint
# or
npm run lint
```

### Format the files

```bash
yarn format
# or
npm run format
```

## Production

### Build the app for production

```bash
quasar build
```

### Deploy the backend

```bash
cd backend
go build
./sports-app
```

### Automated Deployment

This project uses GitHub Actions for automated deployment:

1. Deployment workflow is triggered when code is pushed to the master branch
2. The workflow automatically builds both frontend and backend code
3. After successful build, it automatically deploys to Aliyun server
4. Uses systemd service to manage the application process
5. Automatically restarts the service to apply updates

The deployment process is fully automated, requiring no manual intervention.

## Project Structure

```
sports-app/
├── backend/          # Backend service
│   ├── config/       # Configuration
│   ├── controllers/  # Controllers
│   ├── models/       # Data models
│   ├── routes/       # Routes
│   ├── utils/        # Utilities
│   └── main.go       # Entry point
├── src/              # Frontend application
│   ├── assets/       # Static assets
│   ├── components/   # Components
│   ├── pages/        # Pages
│   ├── stores/       # State management
│   └── App.vue       # Root component
├── docs/             # Documentation
└── logs/             # Logs
```

## Features in Development

- Community features
- Personal profile system
- Social features
- Data analysis

## License

This project is licensed under the MIT License.

</div>

<div id="chinese">

一个使用 Quasar 框架和 Go 后端构建的现代运动打卡应用。

## 功能特点

### 前端

- 现代化响应式界面
- 跨平台兼容性
- 开发热重载
- 内置代码检查和格式化
- 生产环境构建系统
- 暗色模式支持
- 国际化支持
- 数据可视化

### 后端

- RESTful API接口
- JWT认证
- 数据库集成
- 邮箱验证
- 数据统计
- 性能优化

## 技术栈

### 前端

- Vue 3 + TypeScript
- Quasar Framework
- Pinia (状态管理)
- Vue Router
- ECharts (数据可视化)
- Axios (HTTP客户端)

### 后端

- Go 1.21+
- Gin Web框架
- GORM (ORM)
- MySQL 5.7+
- JWT (认证)

## 环境要求

- Node.js (v14或更高)
- Go 1.21或更高
- MySQL 5.7或更高
- npm或yarn包管理器

## 安装

### 安装依赖

```bash
# 前端
cd sports-app
yarn
# 或
npm install

# 后端
cd backend
go mod download
```

## 开发

### 启动前端开发服务器

```bash
quasar dev
```

### 启动后端服务器

```bash
cd backend
go run main.go
```

### 代码检查

```bash
yarn lint
# 或
npm run lint
```

### 代码格式化

```bash
yarn format
# 或
npm run format
```

## 生产环境

### 构建生产版本

```bash
quasar build
```

### 部署后端

```bash
cd backend
go build
./sports-app
```

### 自动化部署

本项目使用 GitHub Actions 实现自动化部署流程：

1. 当代码推送到 master 分支时，自动触发部署工作流
2. 工作流会自动构建前端和后端代码
3. 构建完成后，自动部署到阿里云服务器
4. 使用 systemd 服务管理应用进程
5. 自动重启服务以应用更新

部署流程完全自动化，无需手动操作。

## 项目结构

```
sports-app/
├── backend/          # 后端服务
│   ├── config/       # 配置
│   ├── controllers/  # 控制器
│   ├── models/       # 数据模型
│   ├── routes/       # 路由
│   ├── utils/        # 工具函数
│   └── main.go       # 入口文件
├── src/              # 前端应用
│   ├── assets/       # 静态资源
│   ├── components/   # 组件
│   ├── pages/        # 页面
│   ├── stores/       # 状态管理
│   └── App.vue       # 根组件
├── docs/             # 文档
└── logs/             # 日志
```

## 开发中的功能

- 社区功能
- 个人资料系统
- 社交功能
- 数据分析

## 许可证

本项目采用 MIT 许可证。

</div>
