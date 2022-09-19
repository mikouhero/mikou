package lovely_cat

import (
	"encoding/json"
	"fmt"
	"net/url"
)

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

//GetGroupMemberList 获取群成员列表
func (w *WechatService) GetGroupMemberList(robwxid, group_wxid string, is_refresh int) []GroupMemberList {
	param := map[string]interface{}{
		"type":       GET_GROUP_MEMBER_LIST,
		"robot_wxid": robwxid,
		"group_wxid": group_wxid,
		"is_refresh": 1,
	}
	b := w.sendHttp(param)

	data := response{}
	_ = json.Unmarshal(b, &data)
	lists, _ := url.QueryUnescape(data.Data)
	var list []GroupMemberList
	_ = json.Unmarshal([]byte(lists), &list)
	return list
}

//GetGroupMemberList 获取成员信息
func (w *WechatService) GetGroupMember(robwxid, group_wxid, member_wxid string) {
	param := map[string]interface{}{
		"type":        GET_GROUP_MEMBER,
		"robot_wxid":  robwxid,
		"group_wxid":  group_wxid,
		"member_wxid": member_wxid,
	}
	b := w.sendHttp(param)

	data := response{}
	_ = json.Unmarshal(b, &data)

	fmt.Println(string(b))
	lists, _ := url.QueryUnescape(data.Data)
	var user FriendList
	_ = json.Unmarshal([]byte(lists), &user)
	fmt.Println(user)
}
