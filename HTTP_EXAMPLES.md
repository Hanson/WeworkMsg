# HTTP 接口调用示例

服务端提供两个 HTTP 接口，均以 POST JSON 方式调用。

## 接口说明

### 获取会话消息

```
POST /get_chat_data
```

请求参数：

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| seq | uint64 | 是 | 消息 seq 起始值，首次拉取传 0 |
| limit | uint64 | 是 | 拉取条数上限 |
| timeout | int | 是 | 超时时间（秒） |
| proxy | string | 否 | 代理地址 |
| passwd | string | 否 | 代理密码 |

### 获取媒体文件

```
POST /get_media_data
```

请求参数：

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| sdk_file_id | string | 是 | 媒体文件 ID |
| timeout | int | 是 | 超时时间（秒） |
| proxy | string | 否 | 代理地址 |
| passwd | string | 否 | 代理密码 |

## PHP

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

## Python

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

## Go

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

如需要其他语言例子，可以提 [issue](https://github.com/Hanson/WeworkMsg/issues)。
