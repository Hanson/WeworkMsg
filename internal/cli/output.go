package cli

import (
	"encoding/json"
	"fmt"
	"os"
)

func PrintJSON(data interface{}) error {
	encoded, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("JSON 格式化失败: %w", err)
	}
	fmt.Println(string(encoded))
	return nil
}

func WriteJSON(data interface{}, filename string) error {
	encoded, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("JSON 格式化失败: %w", err)
	}
	if err := os.WriteFile(filename, encoded, 0644); err != nil {
		return fmt.Errorf("写入文件失败: %w", err)
	}
	return nil
}

func PrintInfo(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
}
