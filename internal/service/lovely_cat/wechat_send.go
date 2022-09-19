package lovely_cat

import (
	"encoding/json"
	"net/url"

	"fmt"
)

// 发送私/群聊普通文本消息
func (w *WechatService) SendTextMsg(robwxid, to_wxid, msg string) {
	param := map[string]interface{}{
		"type":       SEND_TEXT_MESSAGE,
		"robot_wxid": robwxid,
		"to_wxid":    to_wxid,
		"msg":        url.QueryEscape(msg),
	}
	fmt.Println(param)
	b := w.sendHttp(param)
	data := response{}
	_ = json.Unmarshal(b, &data)
	fmt.Println(data)
}

// 发送群聊at消息
func (w *WechatService) SendGroupAtMsg(robwxid, to_wxid, at_wxid, at_name, msg string) {
	param := map[string]interface{}{
		"type":       SEND_GROUP_AT_MSG,
		"robot_wxid": robwxid,
		"to_wxid":    to_wxid,
		"at_wxid":    at_wxid,
		"at_name":    at_name,
		"msg":        url.QueryEscape(msg),
	}

	fmt.Println(param)
	b := w.sendHttp(param)
	data := response{}
	_ = json.Unmarshal(b, &data)
	fmt.Println(data)
}

func (w *WechatService) SendImageMsg() {

}
func (w *WechatService) SendVideoMsg() {

}
func (w *WechatService) SendFileMsg() {

}
func (w *WechatService) SendEmojiMsg() {

}
func (w *WechatService) SendLinkMsg() {

}
func (w *WechatService) SendMusicMsg() {

}
