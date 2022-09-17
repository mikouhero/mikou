package lovely_cat

// response 可爱猫返回的数据结构
type response struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

// FriendList 好友列表内部json结构
type FriendList struct {
	Wxid      string `json:"wxid"`
	Nickname  string `json:"nickname"`
	RobotWxid string `json:"robot_wxid"`
	Note      string `json:"note"`
	WxNum     string `json:"wx_num"`
}

// LoginUser  登录用户列表结构
type LoginUser struct {
	Wxid             string `json:"wxid"`
	WxNum            string `json:"wx_num"`
	Nickname         string `json:"nickname"`
	Headimgurl       string `json:"headimgurl"`
	Headimgbase64    string `json:"headimgbase64"`
	Status           int    `json:"status"`
	LoginTime        int    `json:"loginTime"`
	WxWindHandle     int    `json:"wx_wind_handle"`
	WxPid            int    `json:"wx_pid"`
	Signature        string `json:"signature"`
	Backgroundimgurl string `json:"backgroundimgurl"`
}

// GroupList  群聊列表
type GroupList struct {
	Wxid        string `json:"wxid"`
	WxNum       string `json:"wx_num"`
	Nickname    string `json:"nickname"`
	RobotWxid   string `json:"robot_wxid"`
	MemberCount int    `json:"memberCount"`
	IsManager   int    `json:"isManager"`
}
