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

复制 libWeWorkFinanceSdk_C.so 动态库文件到系统动态链接库默认文件夹下，或者复制到任意文件夹并在当前文件夹下执行 export LD_LIBRARY_PATH=$(pwd)命令设置动态链接库检索地址

执行 `./WeworkMsg`

## 运行

* 下载最新版本 https://github.com/Hanson/WeworkMsg/releases/
* 在 WeworkMsg 目录下新建文件 `.env` 和 `private_key.pem`，并进行配置
* 运行 `./WeworkMsg`

## 接口调用

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

![image](https://github.com/Hanson/WeworkMsg/assets/10583423/abff87b3-c6b6-4246-902e-a34929dc9373)

## 相关项目
* 开源scrm https://github.com/juhe-scrm/juhe-scrm
* 高级企微接口 https://github.com/hanson/vbot
* 聚合聊天 https://juhebot.com


![cd788338d401375c814f0fd66f4fbb81](https://github.com/Hanson/vbot/assets/10583423/034ce0fb-12c2-4ce0-8335-cf5132b17bca)
