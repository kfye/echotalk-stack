# echotalk-stack

EchoTalk 基建/部署仓：单台 2核2G 服务器上的容器栈，由 Portainer 管理。
栈含 **Nginx + MySQL 8 + Redis + Go 占位服务**（Day 1 基建）。

## 快速开始
1. 准备环境变量：`cp .env.example .env` 并填密码。
2. 部署：Portainer → Stacks → Add stack（推荐 Git 仓库方式），或服务器上 `docker compose up -d --build`。
3. 验收：`curl http://<服务器IP>/health` 返回 200 且 `deps.mysql/redis` 均为 `true`。

详细说明与 Portainer 操作步骤见 **[docs/容器栈部署指引.md](docs/容器栈部署指引.md)**。

## 结构
- `docker-compose.yml` — 栈定义（每容器限内存、DB/Redis 仅内网）
- `nginx/` — 反代 + 管理端静态托管
- `mysql/conf.d/` — MySQL 调瘦配置
- `redis/` — Redis 限内存配置
- `placeholder/` — Go 占位服务（Day 2 起由 backend 仓替换）
- `admin/dist/` — 管理端构建产物（Day 6 替换为 Vue 产物）
