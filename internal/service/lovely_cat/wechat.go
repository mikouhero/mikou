package lovely_cat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var host string = "http://127.0.0.1:8073/send"

type response struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

type friendList struct {
	Wxid      string `json:"wxid"`
	Nickname  string `json:"nickname"`
	RobotWxid string `json:"robot_wxid"`
	Note      string `json:"note"`
	WxNum     string `json:"wx_num"`
}

type WechatService struct {
}

func NewWechatService() WechatService {
	return WechatService{}
}

/**
* @todo 接口列表
发送文本消息              send_text_msg()
发送群消息并艾特某人      send_group_at_msg()
发送图片消息              send_image_msg()
发送视频消息              send_video_msg()
发送文件消息              send_file_msg()
发送动态表情              send_emoji_msg()
发送分享链接              send_link_msg()
发送音乐消息              send_music_msg()
取指定登录账号的昵称      get_robot_name()
取指定登录账号的头像      get_robot_headimgurl()
取登录账号列表            get_logged_account_list()
取好友列表                get_friend_list()
取群聊列表                get_group_list()
取群成员资料              get_group_member()
取群成员列表              get_group_member_list()
接收好友转账              accept_transfer()
同意群聊邀请              agree_group_invite()
同意好友请求              agree_friend_verify()
修改好友备注              modify_friend_note()
删除好友                  delete_friend()
踢出群成员                remove_group_member()
修改群名称                modify_group_name()
修改群公告                modify_group_notice()
建立新群                  building_group()
退出群聊                  quit_group()
邀请加入群聊              invite_in_group()
*/
func (w *WechatService) SendTextMsg() {

}

/**
* 取好友列表
* @access public
* @param  string $robwxid    账户id
* @param  string $is_refresh 是否刷新
* @return string 当前框架已登录的账号信息列表
 */
func (w *WechatService) GetFriendList(robwxid string, is_refresh int) {

	m := make(map[string]interface{})
	m["type"] = 204
	m["robwxid"] = robwxid
	m["is_refresh"] = is_refresh

	b := w.sendHttp(m)
	data := response{}
	json.Unmarshal(b, &data)
	datas, _ := url.QueryUnescape(data.Data)
	var a []friendList
	json.Unmarshal([]byte(datas), &a)

}

func (w *WechatService) sendHttp(data map[string]interface{}) []byte {
	//JSON序列化
	configData, _ := json.Marshal(data)
	param := bytes.NewBuffer([]byte(configData))

	//构建http请求
	client := &http.Client{}
	req, err := http.NewRequest("POST", host, param)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	//发送请求
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	//返回结果
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return body
}
