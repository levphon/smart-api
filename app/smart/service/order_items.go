// @Author sunwenbo
// 2024/7/13 21:15
package service

import (
	"errors"
	"fmt"
	"go-admin/app/smart/models"

	"github.com/go-admin-team/go-admin-core/sdk/pkg"

	"gorm.io/gorm"

	"go-admin/app/smart/service/dto"
	cDto "go-admin/common/dto"

	"github.com/go-admin-team/go-admin-core/sdk/service"
)

type OrderItems struct {
	service.Service
}

// Get 获取SysDept对象
func (e *OrderItems) Get(d *dto.OrderitemsGetReq, model *models.OrderItems) error {
	var err error
	var data models.OrderItems

	db := e.Orm.Model(&data).
		First(model, d.GetId())
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("db error:%s", err)
		return err
	}
	if err = db.Error; err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SysDept对象
func (e *OrderItems) Insert(c *dto.OrderItemsInsertReq) error {
	var err error
	var data models.OrderItems
	c.Generate(&data)
	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	err = tx.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	deptPath := pkg.UIntToString(data.ID) + "/"

	var mp = map[string]string{}
	mp["dept_path"] = deptPath
	if err := tx.Model(&data).Update("dept_path", deptPath).Error; err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Update 修改SysDept对象
func (e *OrderItems) Update(c *dto.OrderItemsUpdateReq) error {
	var err error
	var model = models.OrderItems{}
	tx := e.Orm.Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	tx.First(&model, c.GetId())
	c.Generate(&model)

	db := tx.Save(&model)
	if err = db.Error; err != nil {
		e.Log.Errorf("UpdateSysDept error:%s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除SysDept
func (e *OrderItems) Remove(d *dto.OrderItemsDeleteReq) error {
	var err error
	var data models.OrderItems

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

// GetSysDeptList 获取组织数据
func (e *OrderItems) getList(c *dto.OrderItemsGetPageReq, list *[]models.OrderItems) error {
	var err error
	var data models.OrderItems

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).
		Find(list).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// SetOrderItems 设置Orderitems页面数据
func (e *OrderItems) SetOrderItems(c *dto.OrderItemsGetPageReq) (m []models.OrderItems, err error) {
	var list []models.OrderItems
	err = e.getList(c, &list)
	for i := 0; i < len(list); i++ {
		if list[i].ID != 0 {
			continue
		}
		info := e.deptPageCall(&list, list[i])
		m = append(m, info)
	}
	return
}

func (e *OrderItems) deptPageCall(deptlist *[]models.OrderItems, menu models.OrderItems) models.OrderItems {
	list := *deptlist
	min := make([]models.OrderItems, 0)
	for j := 0; j < len(list); j++ {
		if menu.ID != list[j].ID {
			continue
		}
		mi := models.OrderItems{}
		mi.ID = list[j].ID
		mi.Title = list[j].Title
		mi.BindTempLate = list[j].BindTempLate
		mi.CategoryID = list[j].CategoryID
		mi.Link = list[j].Link
		mi.Favorite = list[j].Favorite
		mi.Icon = list[j].Icon
		mi.UpdatedAt = list[j].UpdatedAt
		mi.CreatedAt = list[j].CreatedAt
		ms := e.deptPageCall(deptlist, mi)
		min = append(min, ms)
	}
	fmt.Println("menu:", menu)
	return menu
}
