package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	serverAddr string
	version   = "0.0.0"
)

var rootCmd = &cobra.Command{
	Use:     "wework-cli",
	Short:   "企微会话存档 CLI 客户端",
	Long:    "wework-cli 是一个命令行工具，用于连接 WeworkMsg 服务端拉取会话消息和下载媒体文件。",
	Version: version,
}

func init() {
	rootCmd.PersistentFlags().StringVar(&serverAddr, "server", "", "服务端地址 (默认 http://localhost:8888，可设置环境变量 WEWORK_SERVER)")
}

func GetServerAddr() string {
	if serverAddr != "" {
		return serverAddr
	}
	if env := os.Getenv("WEWORK_SERVER"); env != "" {
		return env
	}
	return "http://localhost:8888"
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
