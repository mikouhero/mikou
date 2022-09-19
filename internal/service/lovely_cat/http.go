package lovely_cat

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var host string = "http://127.0.0.1:8073/send"

// sendHttp 调用可爱猫接口
func (w *WechatService) sendHttp(data map[string]interface{}) []byte {
	//JSON序列化
	configData, _ := json.Marshal(data)
	param := bytes.NewBuffer([]byte(configData))

	//构建http请求
	client := &http.Client{}
	req, err := http.NewRequest("POST", host, param)

	if err != nil {
		return nil
	}

	//发送请求
	res, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	//返回结果
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil
	}

	return body
}
