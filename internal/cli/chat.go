package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	chatSeq     uint64
	chatLimit   uint64
	chatTimeout int64
	chatProxy   string
	chatPasswd  string
	chatOutput  string
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "拉取会话消息",
	Long:  "从 WeworkMsg 服务端拉取企微会话消息。",
	Example: `  wework-cli chat
  wework-cli chat --seq 1000 --limit 50
  wework-cli chat --server http://10.0.0.1:8888 --output result.json`,
	RunE: runChat,
}

func init() {
	chatCmd.Flags().Uint64Var(&chatSeq, "seq", 0, "消息 seq 起始值")
	chatCmd.Flags().Uint64Var(&chatLimit, "limit", 100, "拉取消息条数上限")
	chatCmd.Flags().Int64Var(&chatTimeout, "timeout", 3, "请求超时时间（秒）")
	chatCmd.Flags().StringVar(&chatProxy, "proxy", "", "代理地址")
	chatCmd.Flags().StringVar(&chatPasswd, "passwd", "", "代理密码")
	chatCmd.Flags().StringVar(&chatOutput, "output", "", "输出到文件（默认输出到 stdout）")

	rootCmd.AddCommand(chatCmd)
}

func runChat(cmd *cobra.Command, args []string) error {
	payload := map[string]interface{}{
		"seq":     chatSeq,
		"limit":   chatLimit,
		"timeout": chatTimeout,
	}
	if chatProxy != "" {
		payload["proxy"] = chatProxy
	}
	if chatPasswd != "" {
		payload["passwd"] = chatPasswd
	}

	server := GetServerAddr()
	PrintInfo("连接服务端: %s", server)
	PrintInfo("参数: seq=%d, limit=%d, timeout=%d", chatSeq, chatLimit, chatTimeout)

	client := NewClient(server)
	resp, err := client.Post("/get_chat_data", payload)
	if err != nil {
		return fmt.Errorf("拉取会话消息失败: %w", err)
	}

	var dataList []interface{}
	if err := json.Unmarshal(resp.Data, &dataList); err != nil {
		return fmt.Errorf("解析会话数据失败: %w", err)
	}

	PrintInfo("成功拉取 %d 条消息", len(dataList))

	if chatOutput != "" {
		if err := WriteJSON(dataList, chatOutput); err != nil {
			return err
		}
		PrintInfo("已写入文件: %s", chatOutput)
	} else {
		if err := PrintJSON(dataList); err != nil {
			return err
		}
	}

	return nil
}
