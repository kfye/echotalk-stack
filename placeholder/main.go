// EchoTalk Go 占位服务
// 目的：Day 1 验证容器栈链路 —— 暴露 /health，并探测 MySQL/Redis 连通性。
// Day 2 起由真实 backend 仓替换本服务。
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

// tcpReachable 在超时内 dial 一个 host:port，返回是否可达。
func tcpReachable(host, port string) bool {
	addr := net.JoinHostPort(host, port)
	conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		return false
	}
	_ = conn.Close()
	return true
}

type health struct {
	Status string          `json:"status"`
	Time   string          `json:"time"`
	Deps   map[string]bool `json:"deps"`
}

func checkDeps() health {
	deps := map[string]bool{
		"mysql": tcpReachable(env("MYSQL_HOST", "mysql"), env("MYSQL_PORT", "3306")),
		"redis": tcpReachable(env("REDIS_HOST", "redis"), env("REDIS_PORT", "6379")),
	}
	status := "ok"
	for _, ok := range deps {
		if !ok {
			status = "degraded"
		}
	}
	return health{Status: status, Time: time.Now().Format(time.RFC3339), Deps: deps}
}

func main() {
	// 容器 healthcheck 入口：本进程自检 /health，200 即退出 0。
	hc := flag.Bool("healthcheck", false, "probe local /health and exit")
	flag.Parse()
	if *hc {
		resp, err := http.Get("http://127.0.0.1:8080/health")
		if err != nil || resp.StatusCode != http.StatusOK {
			os.Exit(1)
		}
		os.Exit(0)
	}

	mux := http.NewServeMux()

	// 永远返回 200，仅表示「服务存活」；依赖状态放在 body 里供观察。
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(checkDeps())
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "EchoTalk placeholder service. See /health")
	})

	// 启动时打印一次依赖探测结果，方便在 Portainer 日志里确认链路。
	h := checkDeps()
	b, _ := json.Marshal(h)
	log.Printf("startup deps: %s", b)

	log.Println("listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
