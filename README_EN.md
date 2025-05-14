# Sports App

A modern sports tracking application built with Quasar Framework and Go backend.

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
