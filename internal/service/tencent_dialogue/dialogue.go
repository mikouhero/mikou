package tencent_dialogue

import (
	"context"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tbp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tbp/v20190627"
	"mikou/global"
	dao "mikou/internal/dao/v1"
)

type DialogueService struct {
	ctx context.Context // 上下文
	dao *dao.Dao        // dao 数据层
}

// NewDialogueService
func NewDialogueService(ctx context.Context) DialogueService {
	svc := DialogueService{ctx: ctx}
	svc.dao = dao.New(global.DBEngineV2)
	return svc
}

// Reset  重置
func (ds DialogueService) Reset(TerminalId string) {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	// 密钥可前往https://console.cloud.tencent.com/cam/capi网站进行获取
	credential := common.NewCredential(
		global.TencentSetting.DialogueSecretId,
		global.TencentSetting.DialogueSecretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "tbp.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := tbp.NewClient(credential, "", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := tbp.NewTextResetRequest()

	request.BotId = common.StringPtr(global.TencentSetting.DialogueBotId)
	request.BotEnv = common.StringPtr(global.TencentSetting.DialogueBotEnv)
	request.TerminalId = common.StringPtr(TerminalId)

	// 返回的resp是一个TextResetResponse的实例，与请求对象对应
	response, err := client.TextReset(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
}

// Text 发送文本
func (ds DialogueService) Text(TerminalId, InputText string) ([]*tbp.Group, error) {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	// 密钥可前往https://console.cloud.tencent.com/cam/capi网站进行获取
	credential := common.NewCredential(
		global.TencentSetting.DialogueSecretId,
		global.TencentSetting.DialogueSecretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "tbp.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := tbp.NewClient(credential, "", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := tbp.NewTextProcessRequest()

	request.BotId = common.StringPtr(global.TencentSetting.DialogueBotId)
	request.BotEnv = common.StringPtr(global.TencentSetting.DialogueBotEnv)
	request.TerminalId = common.StringPtr(TerminalId)
	request.InputText = common.StringPtr(InputText)

	// 返回的resp是一个TextProcessResponse的实例，与请求对象对应
	response, err := client.TextProcess(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return nil, err
	}
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return response.Response.ResponseMessage.GroupList, err

}
