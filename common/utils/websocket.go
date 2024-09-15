package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WebSocketManager 管理WebSocket连接和消息广播
type WebSocketManager struct {
	TaskChannels map[int]*TaskChannel
	mu           sync.Mutex
}

// TaskChannel 管理每个任务的连接和消息
type TaskChannel struct {
	Clients    map[string]*websocket.Conn
	Broadcast  chan Message
	Register   chan *websocket.Conn
	Unregister chan *websocket.Conn
}

// 单例实例
var Manager = &WebSocketManager{
	TaskChannels: make(map[int]*TaskChannel),
}

// Message 用于传递任务信息的结构
type Message struct {
	Type        string `json:"type"`         // 消息类型: start, complete, error
	TaskID      int    `json:"task_id"`      // 任务ID
	TaskName    string `json:"task_name"`    // 任务名称
	Username    string `json:"username"`     // 主机用户名
	Host        string `json:"host"`         // 主机地址
	Port        int    `json:"port"`         // 主机端口
	Command     string `json:"command"`      // 执行的命令
	Output      string `json:"output"`       // 标准输出
	ErrorOutput string `json:"error_output"` // 标准错误输出
	StartTime   string `json:"start_time"`   // 执行开始时间
	EndTime     string `json:"end_time"`     // 执行结束时间
	Duration    string `json:"duration"`     // 执行时长
}

// Start 方法保留作为WebSocket管理器的初始化方法，但无需参数
func (wm *WebSocketManager) Start() {
	log.Println("WebSocket Manager initialized...")
}

// 确保任务通道存在的方法
func (wm *WebSocketManager) ensureTaskChannel(taskID int) *TaskChannel {
	wm.mu.Lock()
	defer wm.mu.Unlock()

	if _, exists := wm.TaskChannels[taskID]; !exists {
		taskChannel := &TaskChannel{
			Clients:    make(map[string]*websocket.Conn),
			Broadcast:  make(chan Message),
			Register:   make(chan *websocket.Conn),
			Unregister: make(chan *websocket.Conn),
		}
		wm.TaskChannels[taskID] = taskChannel
		go taskChannel.start(taskID)
		log.Printf("Task channel started for task %d...", taskID)
	}
	return wm.TaskChannels[taskID]
}

// start 方法启动单个任务通道
func (tc *TaskChannel) start(taskID int) {
	for {
		select {
		case conn := <-tc.Register:
			id := conn.RemoteAddr().String()
			tc.Clients[id] = conn
			fmt.Printf("Client %s connected to task %d\n", id, taskID)

		case conn := <-tc.Unregister:
			id := conn.RemoteAddr().String()
			delete(tc.Clients, id)
			fmt.Printf("Client %s disconnected from task %d\n", id, taskID)

		case message := <-tc.Broadcast:
			for _, conn := range tc.Clients {
				if err := conn.WriteJSON(message); err != nil {
					log.Printf("WebSocket write error: %v", err)
				}
			}
		}
	}
}

func (wm *WebSocketManager) getTaskChannel(taskID int) *TaskChannel {
	wm.mu.Lock()
	defer wm.mu.Unlock()
	return wm.TaskChannels[taskID]
}

func (wm *WebSocketManager) BroadcastMessage(taskID int, message Message) {
	taskChannel := wm.ensureTaskChannel(taskID)
	taskChannel.Broadcast <- message
}

// WsHandler 处理WebSocket连接
func WsHandler(c *gin.Context) {
	taskIDStr := c.Param("id")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		http.Error(c.Writer, "invalid task ID", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}
	defer conn.Close()

	taskChannel := Manager.ensureTaskChannel(taskID)
	taskChannel.Register <- conn

	// 发送欢迎消息
	welcomeMsg := []byte(fmt.Sprintf("Welcome to the WebSocket server for task ID %d!", taskID))
	if err := conn.WriteMessage(websocket.TextMessage, welcomeMsg); err != nil {
		log.Printf("WebSocket write error: %v", err)
		return
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("WebSocket read error: %v", err)
			taskChannel.Unregister <- conn
			break
		}
		fmt.Printf("Received message: %s\n", message)
	}
}
