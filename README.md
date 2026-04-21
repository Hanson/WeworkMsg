# WeworkMsgSdk
全语言通用的企业微信会话存档SDK

## 介绍
此项目基于 https://github.com/NICEXAI/WeWorkFinanceSDK 实现

WeworkMsg 会以服务运行，对外暴露接口提供服务，开发者只需要进行 http 接口调用，无需关心各种复杂的兼容问题。

## 配置

* `.env` 里面配置好企业的相关信息
* 密钥信息粘贴在 private_key.pem 里面
* PORT 为服务的端口号

## 运行

* 下载最新版本 https://github.com/Hanson/WeworkMsg/releases/
* 复制 libWeWorkFinanceSdk_C.so 动态库文件到系统动态链接库默认文件夹下，或者复制到任意文件夹并在当前文件夹下执行 export LD_LIBRARY_PATH=$(pwd)命令设置动态链接库检索地址
* 在 WeworkMsg 目录下新建文件 `.env` 和 `private_key.pem`，并进行配置
* 运行 `./WeworkMsg`

## CLI 客户端

除了直接调用 HTTP 接口，项目还提供了命令行客户端 `wework-cli`，方便在本地快速拉取会话消息和下载媒体文件。

### 安装

下载最新版本 https://github.com/Hanson/WeworkMsg/releases/ 中对应平台的 `wework-cli` 二进制文件。

或自行编译：

```bash
make build-cli
```

### 使用

```bash
# 拉取会话消息
wework-cli chat --server http://your-server:8888 --seq 0 --limit 100

# 拉取消息并保存到文件
wework-cli chat --output result.json

# 下载媒体文件
wework-cli media --sdk-file-id xxx --output image.jpg

# 使用代理
wework-cli chat --proxy http://proxy:8080 --passwd yourpass

# 查看帮助
wework-cli --help
wework-cli chat --help
wework-cli media --help
```

### 配置

| 参数 | 环境变量 | 默认值 | 说明 |
|------|---------|--------|------|
| `--server` | `WEWORK_SERVER` | `http://localhost:8888` | 服务端地址，flag 优先于环境变量 |
| `--version` | - | - | 查看版本号 |

## HTTP 接口调用

服务提供了两个接口，均以 POST json 的方式进行调用

* `/get_chat_data` 获取会话列表，参数例子`{"seq":0,"limit":100,"timeout":3}`
* `/get_media_data` 获取媒体文件，参数例子`{"sdk_file_id":"xxx","timeout":3}`

### php

```php
<?php

$curl = curl_init();

curl_setopt_array($curl, array(
   CURLOPT_URL => 'http://localhost:8888/get_chat_data',
   CURLOPT_RETURNTRANSFER => true,
   CURLOPT_ENCODING => '',
   CURLOPT_MAXREDIRS => 10,
   CURLOPT_TIMEOUT => 0,
   CURLOPT_FOLLOWLOCATION => true,
   CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
   CURLOPT_CUSTOMREQUEST => 'POST',
   CURLOPT_POSTFIELDS =>'{
    "seq": 0,
    "limit": 100,
    "timeout": 3
}',
   CURLOPT_HTTPHEADER => array(
      'Content-Type: application/json'
   ),
));

$response = curl_exec($curl);

curl_close($curl);
echo $response;
```

### python

```python
import http.client
import json

conn = http.client.HTTPSConnection("localhost", 8888)
payload = json.dumps({
   "seq": 0,
   "limit": 100,
   "timeout": 3
})
headers = {
   'Content-Type': 'application/json'
}
conn.request("POST", "/get_chat_data", payload, headers)
res = conn.getresponse()
data = res.read()
print(data.decode("utf-8"))
```

### go

```go
package main

import (
   "fmt"
   "strings"
   "net/http"
   "io/ioutil"
)

func main() {

   url := "http://localhost:8888/get_chat_data"
   method := "POST"

   payload := strings.NewReader(`{"seq":0,"limit":100,"timeout":3}`)

   client := &http.Client {
   }
   req, err := http.NewRequest(method, url, payload)

   if err != nil {
      fmt.Println(err)
      return
   }
   req.Header.Add("Content-Type", "application/json")

   res, err := client.Do(req)
   if err != nil {
      fmt.Println(err)
      return
   }
   defer res.Body.Close()

   body, err := ioutil.ReadAll(res.Body)
   if err != nil {
      fmt.Println(err)
      return
   }
   fmt.Println(string(body))
}
```

如需要其他语言例子，可以提 issue

## 加交流群

![交流群](images/qrcode.jpg)

## 相关项目
* 开源scrm https://github.com/juhe-scrm/juhe-scrm
* 高级企微接口 https://github.com/hanson/vbot
* 聚合聊天 https://juhebot.com


![cd788338d401375c814f0fd66f4fbb81](https://github.com/Hanson/vbot/assets/10583423/034ce0fb-12c2-4ce0-8335-cf5132b17bca)
