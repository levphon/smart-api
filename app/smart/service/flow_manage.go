// @Author sunwenbo
// 2024/7/13 21:15
package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/smart/models"
	"gorm.io/gorm"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/smart/service/dto"
)

type FlowManage struct {
	service.Service
}

// 分页获取FlowManage所有的数据
func (e *FlowManage) GetFlowManagePage(pageNum, limit int, objects *[]models.FlowManage, total *int64) error {

	// 计算偏移量
	offset := (pageNum - 1) * limit

	// 查询总数，不应用分页限制
	err := e.Orm.Model(&models.FlowManage{}).Count(total).Error
	if err != nil {
		e.Log.Errorf("查询总数失败: %s", err)
		return fmt.Errorf("查询总数失败: %s", err)
	}
	// 查询并分页获取订单项数据
	db := e.Orm.Limit(limit).Offset(offset).Find(objects)

	if err := db.Error; err != nil {
		e.Log.Errorf("分页查询流程失败: %s", err)
		return fmt.Errorf("分页查询流程失败: %s", err)
	}
	return nil
}

// 根据ID获取流程
func (e *FlowManage) Get(d *dto.FlowManageGetReq, model *models.FlowManage) error {
	var err error
	var data models.FlowManage
	var total int64
	err = e.Orm.Model(&models.FlowManage{}).Count(&total).Error
	if err != nil {
		e.Log.Errorf("查询总数失败: %s", err)
		return fmt.Errorf("查询总数失败: %s", err)
	}

	db := e.Orm.Model(&data).First(model, d.GetId())
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		e.Log.Errorf("db error:%s", err)
		return fmt.Errorf("查看对象不存在或无权查看")
	}
	if err = db.Error; err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Flow对象
func (e *FlowManage) Insert(c *dto.FlowManageInsertReq) error {
	var err error
	var data models.FlowManage
	c.Generate(&data)
	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var extFlowManage models.FlowManage
	if err = tx.Where("name = ?", data.Name).First(&extFlowManage).Error; err == nil {
		// 如果存在相同标题的订单项，返回相应的错误消息
		return fmt.Errorf(fmt.Sprintf("flow name '%v' already exists", extFlowManage.Name))
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 如果查询发生其他错误，返回错误
		return err
	}

	err = tx.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Clone Flow对象
func (e *FlowManage) Clone(id string, c *gin.Context) error {
	var err error
	var data models.FlowManage
	if err = e.Orm.First(&data, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			e.Log.Errorf("Flow with ID '%v' not found", id)
			return fmt.Errorf(fmt.Sprintf("Flow with ID '%v' not found", id))
		}
		e.Log.Errorf("Error querying flow with ID '%v': %s", id, err)
		return fmt.Errorf(fmt.Sprintf("Error querying flow with ID '%v': %s"))
	}

	// 复制流程数据
	newFlow := data
	newFlow.ID = 0                                   // 清除ID，以便创建新记录
	newFlow.Name = fmt.Sprintf("%s-copy", data.Name) // 添加 "-copy" 后缀
	newFlow.SetCreateBy(user.GetUserId(c))
	newFlow.Creator = user.GetUserName(c)
	// 创建新的流程记录
	if err = e.Orm.Create(&newFlow).Error; err != nil {
		e.Log.Errorf("Error creating cloned flow: %s", err)
		return err
	}

	return nil
}

// Update Flow
func (e *FlowManage) Update(c *dto.FlowManageUpdateReq) error {
	var err error
	var model = models.FlowManage{}

	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// 根据 ID 查找要更新的记录
	if err = tx.First(&model, c.GetId()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			e.Log.Errorf("FlowManage record with ID '%v' does not exist", c.GetId())
			return fmt.Errorf("record with ID '%v' does not exist", c.GetId())
		}
		e.Log.Errorf("Error querying FlowManage with ID '%v': %s", c.GetId(), err)
		return fmt.Errorf("error querying FlowManage with ID '%v': %s", c.GetId(), err)
	}

	// 使用 Generate 方法生成数据
	c.Generate(&model)

	// 使用 Updates 方法更新字段，确保嵌套结构体 StrucTure 能被正确更新
	if err = tx.Model(&model).Updates(model).Error; err != nil {
		e.Log.Errorf("Failed to update FlowManage with ID '%v': %s", c.GetId(), err)
		return fmt.Errorf("failed to update FlowManage with ID '%v': %s", c.GetId(), err)
	}

	return nil
}

// Remove 删除Flow
func (e *FlowManage) Remove(d *dto.FlowManageDeleteReq) error {
	var err error
	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	var data models.FlowManage
	// 查询要删除的流程
	if err = tx.First(&data, d.GetId()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("flow with ID '%v' not found", d.GetId())
		}
		e.Log.Errorf("Error querying flow with ID '%v': %s", d.GetId(), err)
		return fmt.Errorf("error querying flow with ID '%v': %s", d.GetId(), err)
	}

	var tempData models.FlowTemplates

	// 查询 flowtemplates 表，检查是否存在 bindflow 等于当前流程的 ID
	if err = tx.Where("bindFlow = ?", data.ID).First(&tempData).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			// 查询出错
			return fmt.Errorf("未能找到绑定的流程记录，err: '%v'", err)
		}
	} else {
		// 找到绑定记录，返回错误提示
		return fmt.Errorf("当前流程 '%v' 已经绑定了模板: '%v', 请先解绑然后再删除流程", data.Name, tempData.Name)
	}

	// 执行删除操作
	db := e.Orm.Model(&data).Delete(&data, d.GetId())
	if err = db.Error; err != nil {
		err = db.Error
		e.Log.Errorf("Delete error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		return err
	}
	return nil
}
