// @Author sunwenbo
// 2024/9/29
package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/smart/models"
	"go-admin/app/smart/service"
	"go-admin/app/smart/service/dto"
)

type UserFavorites struct {
	api.Api
}

// AddFavorite 添加收藏
// @Summary 添加收藏
// @Description 添加收藏
// @Tags 收藏
// @Accept application/json
// @Product application/json
// @Param data body dto.UserFavoriteInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/favorites [post]
// @Security Bearer
func (e UserFavorites) AddFavorite(c *gin.Context) {
	s := service.UserFavorites{}

	req := dto.UserFavoriteInsertReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetCreateBy(user.GetUserId(c))
	userID := user.GetUserId(c)
	err = s.AddFavorite(userID, &req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("添加失败 err:%v", err))
		return
	}
	e.OK(req.OrderItemID, "添加成功")
}

// RemoveFavorite 取消收藏
// @Summary 取消收藏
// @Description 取消收藏
// @Tags 收藏
// @Accept application/json
// @Product application/json
// @Param data body dto.UserFavoriteDeleteReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "取消成功"}"
// @Router /api/v1/favorites [delete]
// @Security Bearer
func (e UserFavorites) RemoveFavorite(c *gin.Context) {

	s := service.UserFavorites{}
	req := dto.UserFavoriteDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, "绑定请求失败")
		return
	}
	userId := user.GetUserId(c)

	err = s.RemoveFavorite(userId, &req)

	if err != nil {
		e.Error(500, err, "删除失败")
		return
	}
	e.OK(req.GetId(), "删除成功")
}

// GetUserFavorites 获取用户收藏的工单项
// @Summary 获取用户收藏的工单项
// @Description 获取用户收藏的工单项
// @Tags 收藏
// @Param userId query int true "用户ID"
// @Success 200 {object} response.Response{data=[]models.UserFavorites} "{"code": 200, "data": [...]}"
// @Router /api/v1/favorites [get]
// @Security Bearer
func (e UserFavorites) GetUserFavorites(c *gin.Context) {
	s := service.UserFavorites{}
	// 请求结构体 id
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	var favorites []models.UserFavorites
	userID := user.GetUserId(c)
	// 根据用户 ID 查询收藏的工单
	err = s.GetUserFavoritesByUserID(userID, &favorites)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(favorites, "查询成功")
}
