# OpenBMC

**OpenBMC** 是一个基于 **Go**、**Vue.js** 和 **SQLite** 的轻量级 **BMC（Baseboard Management Controller）** 管理系统，旨在提供高效的硬件监控和远程管理。它结合了前端现代化技术和后端强大的功能，使用 **JSON** 格式存储数据库，支持基本的硬件操作、资源管理、设备控制等功能。

---

## 项目结构

* **前端**：基于 **Vue.js** 和 **TypeScript** 构建，使用 **Element Plus** 和 **TailwindCSS** 进行界面设计，集成 **ECharts** 用于数据可视化。
* **后端**：使用 **Go** 和 **Echo** 框架构建，利用 **GORM** 操作 **SQLite** 数据库，提供 RESTful API 用于数据交互和设备管理。
* **数据库**：使用 **SQLite** 存储设备信息和管理数据，以 **JSON** 格式存储应用配置。

---

## 主要功能

### 前端（Vue.js）

* 使用 **Vue.js** + **TypeScript** 构建，组件化管理，便于扩展和维护。
* **Element Plus** 用于快速构建 UI，提供丰富的组件。
* **TailwindCSS** 提供高效的样式设计，快速开发响应式界面。
* **ECharts** 用于图表和数据可视化，提供实时监控和图形展示。

### 后端（Go）

* **Echo** 用于创建高效的 HTTP Web 服务，处理前端请求。
* **GORM** 用于与 **SQLite** 数据库交互，提供 ORM 支持。
* 提供 RESTful API 支持，进行 **设备管理** 和 **资源操作**。

---

## 安装与运行

### 1. 克隆项目

```bash
git clone https://github.com/hexdeep/openbmc.git
cd openbmc
```

### 2. 前端

#### 安装依赖

```bash
cd frontend
npm install
```

#### 启动前端

```bash
npm run serve
```

前端会在 **localhost:8080** 启动，提供完整的管理界面。

### 3. 后端

#### 安装依赖

```bash
cd backend
go mod tidy
```

#### 启动后端

```bash
go run main.go
```

后端会在 **localhost:8081** 启动，处理来自前端的 API 请求。

### 4. 数据库

数据库使用 **SQLite**，无需额外配置，首次启动时自动生成数据库文件。

---

## 未来计划

* **更多硬件支持**：目前支持基本的硬件监控，未来计划增加更多的硬件接口支持。
* **权限系统**：未来可能增加用户权限和管理系统，当前为单用户管理模式。
* **功能扩展**：根据需求，未来可能增加更多的设备控制、硬件诊断和日志管理功能。

---

## 贡献

我们欢迎任何贡献！请参考以下步骤来贡献代码：

1. Fork 本仓库
2. 创建一个新的分支：`git checkout -b feature-name`
3. 提交你的更改：`git commit -am 'Add new feature'`
4. 推送到分支：`git push origin feature-name`
5. 提交 Pull Request

---

## 许可证

此项目遵循 **MIT 许可证**，具体内容请参见 [LICENSE](LICENSE) 文件。

---

### 注意事项

由于项目细节尚未完全确定，当前版本仅为初步框架，未来将根据需求和开发进度进行不断迭代和优化。

