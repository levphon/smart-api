# Smart-API Work Order System

The Smart-API Work Order System is a backend management system developed in Go, featuring a decoupled architecture with Gin as the backend framework, and Vue.js and Element UI for the frontend. The system aims to efficiently handle work order management, supporting task allocation, execution monitoring, real-time updates, and more.

<img align="right" width="320" src="https://github.com/sunwenbo/golang/raw/master/logo.png">

[![Build Status](https://github.com/wenjianzhang/go-admin/workflows/build/badge.svg)](https://github.com/go-admin-team/go-admin)
[![Release](https://img.shields.io/github/release/go-admin-team/go-admin.svg?style=flat-square)](https://github.com/go-admin-team/go-admin/releases)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/go-admin-team/go-admin)

[English](https://github.com/go-admin-team/go-admin/blob/master/README.md) | 简体中文

Based on Gin + Vue + Element UI OR Arco Design OR Ant Design, the front-end and back-end separated permission management system is extremely simple to initialize, requiring only configuration file modification for database connection. The system supports multi-command operations, and migration commands can simplify the initialization of database information, while service commands can easily start the API service.

[Frontend Project](https://github.com/sunwenbo/smart-ui-new)

[Backend Project](https://github.com/sunwenbo/smart-api)

## 🎬 Online Experience

[Click to access the online demo](https://smart-api.example.com)

> Default login account: `admin` / `123456`

## ✨ Architectural Features

- Follows RESTful API design specifications

- Based on the GIN WEB API framework, providing rich middleware support (user authentication, CORS, access logs, trace IDs, etc.)

- RBAC access control model based on Casbin

- JWT authentication

- Supports Swagger documentation (based on swaggo)

- Database storage based on GORM, supporting various types of databases

- Simple model mapping in configuration files for quick access to desired configurations

- Code generation tools

- Form building tools

- Multi-command mode

- Multi-tenant support

- TODO: Unit tests

## 🤩 Functional Features

- **Custom Volunteer Order form** : Customize various types of work order form structures according to actual needs
- **Customized approval process** : You can flexibly set the processing personnel of the approval node and multiple people to assist in approval
- **Task Work Order Management**: Comprehensive management of task creation, allocation, tracking, and execution status.
- **Real-Time Task Monitoring**: Real-time updates and feedback on task execution status via WebSocket.
- **Role Permission Control**: Fine-grained role and permission management based on the Casbin RBAC model.
- **Rating and Comment Functionality**: Users can rate work orders upon completion and support multiple comment records.
- **Log Recording**: Detailed operational and task execution logs for subsequent auditing and analysis.
- **Chart Statistics**: Intuitive presentation of statistical data on work orders, including weekly and monthly statistics, and personal submission rankings.

## 🎁 Built-in Features

1. Multi-Tenant: The system supports multi-tenancy by default, with separation by database—one database per tenant.
2. User Management: Users are system operators; this function primarily handles user configuration in the system.
3. Department Management: Configure the organization's structure (company, department, group) with tree structure support for data permissions.
4. Position Management: Configure the positions held by system users.
5. Menu Management: Configure system menus, operation permissions, button permission identifiers, and interface permissions.
6. Role Management: Role menu permission distribution, setting role data scope permissions according to organizations.
7. Dictionary Management: Maintenance of fixed data commonly used in the system.
8. Parameter Management: Dynamic configuration of commonly used parameters in the system.
9. Operation Logs: Normal operation log recording and querying; abnormal information log recording and querying.
10. Login Logs: System login log recording and querying, including login exceptions.
11. API Documentation: Automatically generate relevant API documentation based on business code.
12. Code Generation: Generate corresponding CRUD business operations based on table structures, fully visualized, allowing basic business implementation without coding.
13. Form Building: Custom page styling, drag-and-drop page layout.
14. Service Monitoring: View basic information about servers.
15. Content Management: Demo feature with category management and content management. Refer to it for quick entry.
16. Scheduled Tasks: Automated tasks, currently supporting interface calls and function calls.

## 🛰️ System Modules

1. **User Management**: Supports management of user information and permission allocation.
2. **Task Management**: Provides allocation, tracking, and status monitoring of work orders.
3. **Role Management**: Implements fine-grained permission control through roles.
4. **Log Management**: Includes system operation logs and task execution logs.
5. **Service Monitoring**: Real-time viewing of server performance and operational status.
6. **Rating and Comment**: Supports user ratings and comments on work orders upon task completion.

## 🔧 Technology Stack

- **Backend**: Go, Gin, GORM, JWT, Casbin
- **Frontend**: Vue.js, Element UI, Axios, WebSocket
- **Database**: MySQL / PostgreSQL / SQLite
- **Other Tools**: Docker, Swagger, GIT

## 🚀 Quick Start

### Preparation

You need to install [go] [gin] [node](http://nodejs.org/) and [git](https://git-scm.com/) locally.

## 📦 Environment Requirements

- Go 1.18 or higher
- Node.js v14.16.0 or higher
- npm version: 6.14.11
- MySQL or other compatible databases
- Docker (optional)

### Backend Installation Steps

1. Clone the project:

    ```bash
    git clone https://github.com/sunwenbo/smart-api.git
    cd smart-api
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Configure database connection information:

   Modify the database connection information in `config/settings.yml`, ensuring the database configuration is correct.

4. Initialize the database:

    ```bash
    # Initial configuration requires database resource information initialization
    # Use on macOS or Linux
    $ ./smart-api migrate -c config/settings.dev.yml
   
    # ⚠️ Note: Use on Windows
    $ smart-api.exe migrate -c config/settings.dev.yml

    # Start the project, can also use IDE for debugging
    # Use on macOS or Linux
    $ ./smart-api server -c config/settings.yml
    
    # ⚠️ Note: Use on Windows
    $ smart-api.exe server -c config/settings.yml
    ```

5. How to add data to the sys_api table

   When starting the project, using `-a true` will automatically add missing interface data.
   ```bash
   ./smart-api server -c config/settings.yml -a true
   ```
6. Document generation:
   ```bash
   go generate
   ```
7. Start the service:
   ```
   ./smart-api server -c config/settings.yml
   ```

#### Using Docker for Compilation and Start

```azure

```

#### Cross Compilation
   ```bash
   # windows
   env GOOS=windows GOARCH=amd64 go build main.go
   
   # or
   # linux
   env GOOS=linux GOARCH=amd64 go build main.go
   ```

### Frontend Installation Steps

1. Clone the frontend project:

    ```bash
    git clone https://github.com/sunwenbo/smart-ui-new.git
    cd smart-ui-new
    ```

2. Install dependencies:

    ```bash
    npm config set registry https://registry.npmmirror.com/  
    npm config get registry
    npm install --legacy-peer-deps 
    ```

3. Start the development server:

    ```bash
    npm run dev
    ```

4. Access the address:

   Open your browser and go to `http://localhost:9527` to view the frontend interface.


## Usage Instructions

1. **Login to the system**: Use the default administrator account `admin` to log in.
2. **Create Work Orders**: Go to the task management page, click "Create Work Order", fill in the relevant information, and assign a handler.
3. **Task Execution Monitoring**: You can view the execution status, results, and logs of tasks in real-time through the task execution interface in the system.
4. **Work Order Rating and Comments**: After a work order is completed, users can rate the work order and use the comment feature to provide feedback on the handling process.

## System Screenshots


### Login Interface
![Login Interface](https://example.com/screenshot/login.png)

### Work Order Management
![Work Order Management](https://example.com/screenshot/order-management.png)

### Real-time Task Monitoring
![Real-time Task Monitoring](https://example.com/screenshot/task-monitor.png)

## Contribution Guidelines

Contributions from the community developers are welcome. If you want to participate in this project, please follow these steps:

1. **Fork this repository**: Click the `Fork` button at the top right to copy the project to your GitHub account.
2. **Clone your Fork**: Run the following command in your terminal to clone the project to your local machine:

    ```bash
    git clone https://github.com/your-username/smart-api-backend.git
    ```

3. **Create a branch**: In the cloned project directory, create a new branch for development:

    ```bash
    git checkout -b feature-branch-name
    ```
4. **Commit your changes**: After development, use `git add` and `git commit` to submit your code changes.

    ```bash
    git add .
    git commit -m "Describe your changes"
    ```

5. **Submit a Pull Request**: On GitHub, submit a Pull Request, ensuring your code passes all tests and adheres to coding standards.

## Frequently Asked Questions

### How to modify the database configuration?

The database configuration is stored in `config/settings.yml`, where you can modify the `database` configuration item according to your needs.

### How to add a new API interface?

Create new handling logic in the `api` directory and register the route in `router`. The service layer logic should be placed in the `service` directory.

### How to implement WebSocket real-time task monitoring?

During task execution, the system pushes task status update information to the frontend via WebSocket. The frontend uses the `WebSocket` interface to receive and dynamically display task execution logs.

## Open Source License

The Smart-API Work Order System is open-sourced under the [MIT License](LICENSE), allowing individuals and organizations to use and modify it for free.


## 📨 Interaction

<table>
   <tr>
    <td><img src="https://github.com/sunwenbo/golang/raw/master/wx.jpeg" width="180px"></td>
  </tr>
  <tr>
    <td>WeChat</td>
  </tr>
</table>

## 💎 Contributors

Thanks to the following developers for their contributions to this project:

- [Developer 1](https://github.com/developer1)
- [Developer 2](https://github.com/developer2)
- [Developer 3](https://github.com/developer3)


> If you have any questions or suggestions, please submit an [issue](https://github.com/your-repo/smart-api-backend/issues).

## JetBrains Open Source License Support

The `smart-api` project has always been developed using JetBrains' GoLand IDE, based on **free JetBrains Open Source license(s)**, expressing my gratitude for their support.

<a href="https://www.jetbrains.com/?from=kubeadm-ha" target="_blank"><img src="https://raw.githubusercontent.com/panjf2000/illustrations/master/jetbrains/jetbrains-variant-4.png" width="250" align="middle"/></a>

## 🤝 Special Thanks

1. [ant-design](https://github.com/ant-design/ant-design)
2. [ant-design-pro](https://github.com/ant-design/ant-design-pro)
3. [arco-design](https://github.com/arco-design/arco-design)
4. [arco-design-pro](https://github.com/arco-design/arco-design-pro)
5. [gin](https://github.com/gin-gonic/gin)
6. [casbin](https://github.com/casbin/casbin)
7. [spf13/viper](https://github.com/spf13/viper)
8. [gorm](https://github.com/jinzhu/gorm)
9. [gin-swagger](https://github.com/swaggo/gin-swagger)
10. [jwt-go](https://github.com/dgrijalva/jwt-go)
11. [vue-element-admin](https://github.com/PanJiaChen/vue-element-admin)
12. [ruoyi-vue](https://gitee.com/y_project/RuoYi-Vue)
13. [form-generator](https://github.com/JakHuang/form-generator)


## 🤟 Donation

> If you find this project helpful, you can buy the author a drink to show your encouragement :tropical_drink:

<img class="no-margin" src="https://github.com/sunwenbo/golang/raw/master/wxPay.jpeg"  height="200px" >

## 🤝 Links

[Go Developer Growth Roadmap](http://www.golangroadmap.com/)

## 🔑 License

[MIT](https://github.com/sunwenbo/smart-api/blob/main/LICENSE.md)

Copyright (c) 2022 sunwenbo
