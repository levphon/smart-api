// @Author sunwenbo
// 2024/8/5 19:45
package service

import (
	"errors"
	"fmt"
	"go-admin/app/smart/models"
	"go-admin/app/smart/service/dto"
	models2 "go-admin/common/models"
	"gorm.io/gorm"
	"regexp"
	"time"
)

// 定义结构体，用于返回节点名称和 assignValue

/*
    -- 节点 --
	start: 开始节点
	userTask: 审批节点
	receiveTask: 处理节点
	scriptTask: 任务节点
	end: 结束节点

    -- 网关 --
    exclusiveGateway: 排他网关
    parallelGateway: 并行网关
    inclusiveGateway: 包容网关

*/

type NodeInfo struct {
	Name        string `json:"name"`
	AssignValue int    `json:"assignValue"`
}

// Handle OrderWorks
func (e *OrderWorks) Handle(c *dto.OrderWorksHandleReq, handle int) error {
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

	// 获取当前节点
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
	// 根据当前节点查找目标节点
	switch c.ActionType {
	// 1 为同意 0为拒绝
	case "1":
		targetNodeId = findNextNode(edges, cutNodeId)
	case "0":
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

	// 判断如果 targetNodeInfo.AssignValue 等于 0，说明当前节点是start或者是end
	if targetNodeInfo.AssignValue == 0 {
		model.CurrentNode = targetNodeInfo.Name
		// 将工单创建人设置为当前处理人
		model.CurrentHandlerID = model.CreateBy // 假设 workOrderCreator 是工单创建人的变量
		model.CurrentHandler = model.Creator
	} else {
		model.CurrentNode = targetNodeInfo.Name
		model.CurrentHandlerID = targetNodeInfo.AssignValue
		model.CurrentHandler, err = GetHandlerNameByID(tx, targetNodeInfo.AssignValue)
	}

	if err != nil {
		e.Log.Errorf("Error getting handler name for ID '%v': %s", targetNodeInfo.AssignValue, err)
		return fmt.Errorf("error getting handler name for ID '%v': %s", targetNodeInfo.AssignValue, err)
	}
	beforeUpdate := model.UpdatedAt

	// 保存更新
	if err = tx.Save(&model).Error; err != nil {
		e.Log.Errorf("Failed to update order works: %s", err)
		return fmt.Errorf("failed to update order works: %s", err)
	}

	// 记录操作历史
	historyReq := dto.OrderWorksHistReq{
		Title:          model.Title,
		Transfer:       "工单流转", // 假设流转操作类型保存在 ActionType 中
		Remark:         "工单处理",
		CurrentNode:    model.CurrentNode,
		Status:         c.ActionType,
		HandlerId:      handle,
		HandleTime:     models2.JSONTime(time.Now()),
		HandleDuration: calculateHandleDuration(beforeUpdate),
	}
	if err = RecordOperationHistory(tx, &historyReq); err != nil {
		e.Log.Errorf("Failed to record operation history: %s", err)
		return fmt.Errorf("failed to record operation history: %s", err)
	}

	return nil
}

// 辅助函数：根据节点名称查找节点 ID
func findNodeByName(nodes models.StrucTure, currentNode string) string {
	// 获取 nodes 列表
	nodeList, ok := nodes["nodes"].([]interface{})
	if !ok {
		_ = fmt.Errorf("nodes is not a slice of interface{}")
		return ""
	}

	for _, node := range nodeList {
		// 确保 node 是一个 map[string]interface{}
		n, ok := node.(map[string]interface{})
		if !ok {
			_ = fmt.Errorf(`node is not a map[string]interface{}`)
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
	// 编译正则表达式
	re := regexp.MustCompile(`start|end`)

	// 如果nodeId等于 start，则将工单创建人赋值给当前处理人，同时将id赋值给assignValue
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
			// 检查 nodeId 是否包含 "start"
			if re.MatchString(nodeId) {
				return &NodeInfo{
					Name:        n["label"].(string),
					AssignValue: 0,
				}
			}

			assignValues := n["assignValue"].([]interface{})
			if !ok {
				continue
			}
			if len(assignValues) > 0 {
				assignValue, ok := assignValues[0].(float64) // 假设 assignValue 是 float64 类型
				if !ok {
					fmt.Println("assignValue[0] is not a float64")
					continue
				}
				return &NodeInfo{
					Name:        n["label"].(string),
					AssignValue: int(assignValue),
				}
			}
		}
	}
	return nil
}
