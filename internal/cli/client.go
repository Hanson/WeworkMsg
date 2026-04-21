package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type APIResponse struct {
	ErrCode int             `json:"err_code"`
	Data    json.RawMessage `json:"data"`
}

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(serverAddr string) *Client {
	return &Client{
		baseURL: serverAddr,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) Post(endpoint string, payload interface{}) (*APIResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("请求参数序列化失败: %w", err)
	}

	url := c.baseURL + endpoint
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("连接服务端失败 (%s): %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("服务端返回 HTTP %d", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	var apiResp APIResponse
	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	if apiResp.ErrCode != 0 {
		return nil, fmt.Errorf("服务端错误 (err_code=%d): %s", apiResp.ErrCode, string(apiResp.Data))
	}

	return &apiResp, nil
}
