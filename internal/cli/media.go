package cli

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	mediaSdkFileID string
	mediaTimeout   int64
	mediaProxy     string
	mediaPasswd    string
	mediaOutput    string
)

var mediaCmd = &cobra.Command{
	Use:   "media",
	Short: "下载媒体文件",
	Long:  "从 WeworkMsg 服务端下载企微会话中的媒体文件（图片、语音、视频等）。",
	Example: `  wework-cli media --sdk-file-id xxx --output image.jpg
  wework-cli media --sdk-file-id xxx --server http://10.0.0.1:8888 --output file.dat`,
	RunE: runMedia,
}

func init() {
	mediaCmd.Flags().StringVar(&mediaSdkFileID, "sdk-file-id", "", "媒体文件 ID（必填）")
	mediaCmd.Flags().Int64Var(&mediaTimeout, "timeout", 3, "请求超时时间（秒）")
	mediaCmd.Flags().StringVar(&mediaProxy, "proxy", "", "代理地址")
	mediaCmd.Flags().StringVar(&mediaPasswd, "passwd", "", "代理密码")
	mediaCmd.Flags().StringVar(&mediaOutput, "output", "", "保存到文件路径")

	_ = mediaCmd.MarkFlagRequired("sdk-file-id")

	rootCmd.AddCommand(mediaCmd)
}

func runMedia(cmd *cobra.Command, args []string) error {
	payload := map[string]interface{}{
		"sdk_file_id": mediaSdkFileID,
		"timeout":     mediaTimeout,
	}
	if mediaProxy != "" {
		payload["proxy"] = mediaProxy
	}
	if mediaPasswd != "" {
		payload["passwd"] = mediaPasswd
	}

	server := GetServerAddr()
	PrintInfo("连接服务端: %s", server)
	PrintInfo("下载媒体文件: %s", mediaSdkFileID)

	client := NewClient(server)
	resp, err := client.Post("/get_media_data", payload)
	if err != nil {
		return fmt.Errorf("下载媒体文件失败: %w", err)
	}

	var base64Data string
	if err := json.Unmarshal(resp.Data, &base64Data); err != nil {
		return fmt.Errorf("解析媒体数据失败: %w", err)
	}

	if mediaOutput != "" {
		binaryData, err := base64.StdEncoding.DecodeString(base64Data)
		if err != nil {
			return fmt.Errorf("base64 解码失败: %w", err)
		}
		if err := os.WriteFile(mediaOutput, binaryData, 0644); err != nil {
			return fmt.Errorf("写入文件失败: %w", err)
		}
		PrintInfo("已保存文件: %s (%d bytes)", mediaOutput, len(binaryData))
	} else {
		fmt.Println(base64Data)
	}

	return nil
}
