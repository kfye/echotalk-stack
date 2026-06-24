# 把 nginx 配置与管理端静态页直接打进镜像，避免 Portainer git stack 的单文件 bind 挂载坑。
# 构建上下文为仓库根目录（compose 里 context: .），故路径含 nginx/ 与 admin/ 前缀。
FROM nginx:1.27-alpine
COPY nginx/nginx.conf       /etc/nginx/nginx.conf
COPY nginx/conf.d/          /etc/nginx/conf.d/
COPY admin/dist/            /usr/share/nginx/html/
EXPOSE 80
