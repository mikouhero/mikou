package lovely_cat

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mikou/global"
	dao "mikou/internal/dao/v1"
	"net/http"
	"net/url"
)

var host string = "http://127.0.0.1:8073/send"

type WechatService struct {
	ctx context.Context // 上下文
	dao *dao.Dao        // dao 数据层
}

//  NewWechatService 可爱猫微信接口服务层
func NewWechatService(ctx context.Context) WechatService {
	svc := WechatService{ctx: ctx}
	svc.dao = dao.New(global.DBEngineV2)
	return svc
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

// GetLoggedAccountList 获取当前登录用户的列表
func (w *WechatService) GetLoggedAccountList() []LoginUser {
	param := map[string]interface{}{
		"type": GET_LOGGED_ACCOUNT_LIST,
	}
	b := w.sendHttp(param)
	data := response{}
	_ = json.Unmarshal(b, &data)
	lists, _ := url.QueryUnescape(data.Data)
	var user []LoginUser
	_ = json.Unmarshal([]byte(lists), &user)
	return user
}

// GetRobotHeadimgurl  获取头像
func (w *WechatService) GetRobotHeadimgurl(robwxid string) {
	param := map[string]interface{}{
		"type":       GET_ROBOT_HEADIMGURL,
		"robot_wxid": robwxid,
	}
	b := w.sendHttp(param)
	data := response{}
	_ = json.Unmarshal(b, &data)
	fmt.Println(string(b))
	// todo  未返回结构

}

// GetFriendList  取好友列表
func (w *WechatService) GetFriendList(robwxid string, is_refresh int) []FriendList {
	param := map[string]interface{}{
		"type":       GET_FRIEND_LIST,
		"robot_wxid": robwxid,
		"is_refresh": is_refresh,
	}
	b := w.sendHttp(param)
	data := response{}
	_ = json.Unmarshal(b, &data)
	lists, _ := url.QueryUnescape(data.Data)
	var friendList []FriendList
	_ = json.Unmarshal([]byte(lists), &friendList)
	return friendList
}

// GetGroupList 获取群聊
func (w *WechatService) GetGroupList(robwxid string, is_refresh int) []GroupList {
	param := map[string]interface{}{
		"type":       GET_GROUP_LIST,
		"robot_wxid": robwxid,
		"is_refresh": is_refresh,
	}
	b := w.sendHttp(param)
	data := response{}
	_ = json.Unmarshal(b, &data)
	lists, _ := url.QueryUnescape(data.Data)
	var list []GroupList
	_ = json.Unmarshal([]byte(lists), &list)
	return list
}

func (w *WechatService) GetGroupMemberList(robwxid, group_wxid string, is_refresh int) {
	param := map[string]interface{}{
		"type":       GET_GROUP_MEMBER_LIST,
		"robot_wxid": robwxid,
		"group_wxid": group_wxid,
		"is_refresh": 1,
	}
	b := w.sendHttp(param)
	fmt.Println((string(b)))

	data := response{}
	_ = json.Unmarshal(b, &data)
	fmt.Println((data))
	lists, _ := url.QueryUnescape(data.Data)
	fmt.Println(string(lists))
	//var list []GroupList
	//_ = json.Unmarshal([]byte(lists), &list)
}

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
