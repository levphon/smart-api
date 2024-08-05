// @Author sunwenbo
// 2024/8/5 19:45
package service

import (
	"errors"
	"fmt"
	"go-admin/app/smart/models"
	"go-admin/app/smart/service/dto"
	"gorm.io/gorm"
)

// 定义结构体，用于返回节点名称和 assignValue
type NodeInfo struct {
	Name        string `json:"name"`
	AssignValue int    `json:"assignValue"`
}

// Handle OrderWorksry
func (e *OrderWorks) Handle(c *dto.OrderWorksHandleReq) error {
	var err error
	var model = models.OrderWorks{}

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
			e.Log.Errorf("order works with ID '%v' not exists", c.GetId())
			return fmt.Errorf("order works with ID '%v' not exists", c.GetId())
		}
		e.Log.Errorf("Error querying order works with ID '%v': %s", c.GetId(), err)
		return fmt.Errorf("error querying order works with ID '%v': %s", c.GetId(), err)
	}
	// 获取绑定的流程数据
	flowData := model.BindFlowData

	// 获取当前节点和目标节点
	cutNodeId := findNodeByName(flowData.StrucTure, model.CurrentNode)
	if cutNodeId == "" {
		e.Log.Errorf("Current node '%v' not found in flow structure", model.CurrentNode)
		return fmt.Errorf("current node '%v' not found in flow structure", model.CurrentNode)
	}

	// 获取边
	edges, ok := flowData.StrucTure["edges"].([]interface{})
	if !ok {
		e.Log.Errorf("Edges not found or not a slice")
		return fmt.Errorf("edges not found or not a slice")
	}
	targetNodeId := ""
	// 查找目标节点
	switch c.ActionType {
	case "approve":
		targetNodeId = findNextNode(edges, cutNodeId)
	case "reject":
		targetNodeId = findPreviousNode(edges, cutNodeId)
	default:
		e.Log.Errorf("Invalid action type '%v'", c.ActionType)
		return fmt.Errorf("invalid action type '%v'", c.ActionType)
	}

	if targetNodeId == "" {
		targetNodeId = cutNodeId
	}

	// 获取目标节点的名称和处理人
	targetNodeInfo := findNodeInfoById(flowData.StrucTure, targetNodeId)

	if targetNodeInfo == nil {
		e.Log.Errorf("Target node information not found for ID '%v'", targetNodeId)
		return fmt.Errorf("target node information not found for ID '%v'", targetNodeId)
	}
	model.CurrentNode = targetNodeInfo.Name

	if targetNodeId == "end" {
		model.CurrentHandlerID = model.CreateBy // Assuming CreatorID is the ID of the work order creator
	} else {
		model.CurrentHandlerID = targetNodeInfo.AssignValue
	}

	model.CurrentHandler, err = e.getHandlerNameByID(model.CurrentHandlerID)
	if err != nil {
		e.Log.Errorf("Error getting handler name for ID '%v': %s", model.CurrentHandlerID, err)
		return fmt.Errorf("error getting handler name for ID '%v': %s", model.CurrentHandlerID, err)
	}

	// 保存更新
	if err = tx.Save(&model).Error; err != nil {
		e.Log.Errorf("Failed to update order works: %s", err)
		return fmt.Errorf("failed to update order works: %s", err)
	}

	return nil

	//// 记录操作历史
	//if err = recordOperationHistory(tx, c, "Updated order status"); err != nil {
	//	e.Log.Errorf("Failed to record operation history: %s", err)
	//	return fmt.Errorf("failed to record operation history: %s", err)
	//}
	//
	//return nil
}

// 辅助函数：根据当前节点 查找下一个节点（目标节点)
// 辅助函数：根据节点名称查找节点 ID
func findNodeByName(nodes models.StrucTure, currentNode string) string {
	// 获取 nodes 列表
	nodeList, ok := nodes["nodes"].([]interface{})
	if !ok {
		fmt.Println("nodes is not a slice of interface{}")
		return ""
	}

	for _, node := range nodeList {
		// 确保 node 是一个 map[string]interface{}
		n, ok := node.(map[string]interface{})
		if !ok {
			fmt.Println("node is not a map[string]interface{}")
			continue
		}

		if n["label"] == currentNode {
			return n["id"].(string)
		}
	}
	return ""
}

// 辅助函数：根据当前节点查找下一个节点
func findNextNode(edges []interface{}, cutNodeId string) string {
	for _, edge := range edges {
		e, ok := edge.(map[string]interface{})
		if !ok {
			continue
		}
		if e["source"] == cutNodeId {
			return e["target"].(string)
		}
	}
	return ""
}

// 辅助函数：根据当前节点查找上一个节点
func findPreviousNode(edges []interface{}, cutNodeId string) string {
	for _, edge := range edges {
		e, ok := edge.(map[string]interface{})
		if !ok {
			continue
		}
		if e["target"] == cutNodeId {
			return e["source"].(string)
		}
	}
	return cutNodeId // Return the current node ID if no previous node is found
}

// 新增函数：根据节点ID查找节点名称和 assignValue
func findNodeInfoById(nodes models.StrucTure, nodeId string) *NodeInfo {
	nodeList, ok := nodes["nodes"].([]interface{})
	if !ok {
		fmt.Println("nodes is not a slice of interface{}")
		return nil
	}

	for _, node := range nodeList {
		n, ok := node.(map[string]interface{})
		if !ok {
			fmt.Println("node is not a map[string]interface{}")
			continue
		}

		if n["id"] == nodeId {
			assignValues := n["assignValue"].([]interface{})
			if len(assignValues) > 0 {
				assignValue, _ := assignValues[0].(float64) // 假设 assignValue 是 float64 类型
				return &NodeInfo{
					Name:        n["label"].(string),
					AssignValue: int(assignValue),
				}
			}
		}
	}
	return nil
}
