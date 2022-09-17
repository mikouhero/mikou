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
