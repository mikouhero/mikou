package lovely_cat

// 可爱猫事件推送标识 集合
const (
	_                      = iota
	PRIVATE_MESSAGE        = 100 * iota // 私聊
	GROUP_MESSAGE                       // 群聊
	_                                   // 忽略
	GROUP_USER_INCR                     // 群成员增加
	RECEIVE_FRIEND_REQUEST              // 好友请求
	QR_CODE_COLLECTION                  // 二维码收款
	TRANSFER_ACCOUNTS                   // 转账
	SOFT_INIT                           // 软件启动
	NEW_ACCOUNT_LOGIN                   // 新账号登录

)

const (
	GROUP_USER_DESC = 410 // 群成员减少
	ACCOUNT_LOGOUT  = 910 // 账号下线
)

const (
	SEND_TEXT_MESSAGE = 100 + iota //发送文本消息
	_
	SEND_GROUP_AT_MSG //发送群消息并艾特某人
	SEND_IMAGE_MSG    //发送图片消息
	SEND_VIDEO_MSG    //发送视频消息
	SEND_FILE_MSG     //发送文件消息
	SEND_EMOJI_MSG    //发送动态表情
	SEND_LINK_MSG     //发送分享链接
	SEND_MUSIC_MSG    //发送音乐消息
)

const (
	_                       = iota
	GET_ROBOT_NAME          = 200 + iota //取指定登录账号的昵称
	GET_ROBOT_HEADIMGURL                 //取指定登录账号的头像
	GET_LOGGED_ACCOUNT_LIST              //取登录账号列表
	GET_FRIEND_LIST                      //取好友列表
	GET_GROUP_LIST                       //取群聊列表
	GET_GROUP_MEMBER_LIST                //取群成员列表
	GET_GROUP_MEMBER                     //取群成员资料
)

const (
	_                   = iota
	ACCEPT_TRANSFER     = 300 + iota //接收好友转账
	AGREE_GROUP_INVITE               //同意群聊邀请
	AGREE_FRIEND_VERIFY              //同意好友请求
	MODIFY_FRIEND_NOTE               //修改好友备注
	DELETE_FRIEND                    //删除好友
	REMOVE_GROUP_MEMBER              //踢出群成员
	MODIFY_GROUP_NAME                //修改群名称
	MODIFY_GROUP_NOTICE              //修改群公告
	BUILDING_GROUP                   //建立新群
	QUIT_GROUP                       //退出群建立新群聊
	INVITE_IN_GROUP                  //邀请加入群聊
)
