package v1

type MessageRequest struct {
	FinalFromName string `form:"final_from_name" json:"final_from_name"`
	FinalFromWxid string `form:"final_from_wxid" json:"final_from_wxid"`
	FromName      string `form:"from_name" json:"from_name"`
	FromWxid      string `form:"from_wxid" json:"from_wxid"`
	Msg           string `form:"msg" json:"msg"`
	Rid           int    `form:"rid" json:"rid"`
	RobotWxid     string `form:"robot_wxid" json:"robot_wxid"`
	Time          int    `form:"time" json:"time"`
	Type          int    `form:"type" json:"type"`
}

type GetFriendListRequest struct {
	RobotWxid string `form:"robot_wxid" json:"robot_wxid"`
	IsRefresh int    `form:"is_refresh" json:"is_refresh"`
}

type GetGroupMemberList struct {
	RobotWxid string `form:"robot_wxid" json:"robot_wxid"`
	GroupWxid string `form:"group_wxid" json:"group_wxid"`
	IsRefresh int    `form:"is_refresh" json:"is_refresh"`
}

type GetGroupMember struct {
	RobotWxid  string `form:"robot_wxid" json:"robot_wxid"`
	GroupWxid  string `form:"group_wxid" json:"group_wxid"`
	MemberWxid string `form:"member_wxid" json:"member_wxid"`
}

type SendMsgRequest struct {
	RobotWxid string `form:"robot_wxid" json:"robot_wxid"`
	Msg       string `form:"msg" json:"msg"`
	ToWxid    string `form:"to_wxid" json:"to_wxid"`
	AtWxid    string ` form:"at_wxid" json:"at_wxid"`
	AtName    string `form:"at_name" json:"at_name"`
	GroupWxid string `form:"group_wxid" json:"group_wxid"`
}
