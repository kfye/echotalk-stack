# echotalk-stack

EchoTalk 基建/部署仓：单台 2核2G 服务器上的容器栈，由 Portainer 管理。
栈含 **Nginx + MySQL 8 + Redis + Go 占位服务**（Day 1 基建）。

## 快速开始
1. 准备环境变量：`cp .env.example .env` 并填密码。
2. 部署：Portainer → Stacks → Add stack（推荐 Git 仓库方式），或服务器上 `docker compose up -d --build`。
3. 验收：`curl http://<服务器IP>/health` 返回 200 且 `deps.mysql/redis` 均为 `true`。

详细说明与 Portainer 操作步骤见 **[docs/容器栈部署指引.md](docs/容器栈部署指引.md)**。

## 结构
- `docker-compose.yml` — **完整栈**（Day 1）：Nginx + Go 占位 + MySQL + Redis
- `docker-compose.db.yml` — 精简栈：只起 MySQL + Redis（纯 web editor 可粘贴部署）
- `.env.example` — 环境变量样例（复制为 `.env` / 填进 Portainer）
- `nginx/` — 反代 + 管理端静态托管
- `mysql/conf.d/` — MySQL 调瘦配置
- `redis/` — Redis 限内存配置
- `placeholder/` — Go 占位服务（Day 2 起由 backend 仓替换）
- `admin/dist/` — 管理端构建产物（Day 6 替换为 Vue 产物）

## 部署须知
完整栈含 Go 构建上下文与多处配置挂载，**需仓库文件在场**——用 Portainer 的 **Repository** 方式，
或先把仓库 `git clone` 到服务器再 `docker compose up -d --build`。纯 web-editor 粘贴只适用于
`docker-compose.db.yml`（仅 DB/Redis，无构建与挂载）。详见 [docs/容器栈部署指引.md](docs/容器栈部署指引.md)。
