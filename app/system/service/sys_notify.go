package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	"github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"go-admin/app/system/models"
	"go-admin/app/system/service/dto"
	"go-admin/config"
	"gorm.io/gorm"
)

type SysNotify struct {
	service.Service
}

// SendNotify 发送飞书通知
func (e *SysNotify) SendMessage(c *dto.SysNotify) error {
	cfg := config.ExtConfig.Notify
	client := lark.NewClient(cfg.BotCredit.AppID, cfg.BotCredit.AppSecret)

	var users models.SysUser

	tx := e.Orm.Debug().Begin()

	if err := tx.Where("nick_name = ?", c.ReceiveName).First(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			e.Log.Errorf("db error: %s", err)
			return err
		}
	}
	// 获取用户ID，根据用户的手机号调用飞书的sdk进行查找
	receiveID, err := e.findUserIDByReceiveName(users.Phone, client)
	if err != nil {
		e.Log.Errorf("查找用户ID失败: %s", err)
		return err
	}

	content := fmt.Sprintf(`{"text":"你有一个工单需要处理: %s"}`, c.OrderName)

	// 创建请求对象
	req := larkim.NewCreateMessageReqBuilder().
		ReceiveIdType(`open_id`).
		Body(larkim.NewCreateMessageReqBodyBuilder().
			ReceiveId(receiveID).
			MsgType(`text`).
			Content(content).
			Build()).
		Build()

	// 发起请求
	resp, err := client.Im.Message.Create(context.Background(), req)

	// 处理错误
	if err != nil {
		e.Log.Errorf("发送通知失败: %s", err)
		return err
	}

	// 服务端错误处理
	if !resp.Success() {
		e.Log.Errorf("发送通知失败: %d - %s", resp.Code, resp.Msg)
		return fmt.Errorf("飞书通知发送失败: %d - %s", resp.Code, resp.Msg)
	}

	return nil
}

// 获取用户id
func (e *SysNotify) findUserIDByReceiveName(receivePhone string, client *lark.Client) (string, error) {
	req := larkcontact.NewBatchGetIdUserReqBuilder().
		UserIdType(`open_id`). // 获取用户的 OpenID
		Body(larkcontact.NewBatchGetIdUserReqBodyBuilder().
			Mobiles([]string{receivePhone}). // 使用传入的手机号
			IncludeResigned(true).           // 包括已离职的用户
			Build()).
		Build()

	// 发起请求
	resp, err := client.Contact.User.BatchGetId(context.Background(), req)

	// 处理错误
	if err != nil {
		return "", err
	}

	// 检查飞书 API 响应
	if !resp.Success() {
		e.Log.Errorf("飞书 API 错误: %d - %s", resp.Code, resp.Msg)
		return "", fmt.Errorf("飞书 API 错误: %d - %s", resp.Code, resp.Msg)
	}

	// 检查 UserList 和 UserId
	if resp.Data.UserList == nil || len(resp.Data.UserList) == 0 {
		return "", fmt.Errorf("未找到用户")
	}

	if resp.Data.UserList[0].UserId == nil {
		return "", fmt.Errorf("用户 ID 为空")
	}

	// 获取并返回 UserId
	userId := *resp.Data.UserList[0].UserId
	if userId == "" {
		return "", fmt.Errorf("用户 ID 为空")
	}

	return userId, nil
}
