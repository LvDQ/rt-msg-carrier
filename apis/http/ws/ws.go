package ws

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"sync"
	"time"

	"rt-msg-carrier/log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type User struct {
	Conn       *websocket.Conn
	Type       string
	Name       string
	Id         string
	mu         sync.Mutex
	UpdateTime time.Time
}

type Message struct {
	Type    string      `json:"type" binding:"required"`
	TopicId string      `json:"topic_id" binding:"required"`
	Data    interface{} `json:"data" binding:"required"`
}

type UserSlices []*User

var UserList = make(map[string]UserSlices)
var upgrader = websocket.Upgrader{}

func init() {
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// 解决跨域问题
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}

func (u *User) SendMsg(ctx context.Context, logger *logrus.Logger, msg *Message) error {
	switch {
	case msg.Type == "json":
		json_string, _ := json.Marshal(msg.Data)
		logger.Debug("msg.Type is 'json': " + string(json_string))
		err := u.Conn.WriteJSON(msg.Data)
		if err != nil {
			return err
		}
		// u.Conn.WriteString(websocket.TextMessage, msg.Data)
	case msg.Type == "text":
		data := msg.Data.(string)
		logger.Debug("msg.Type is 'text': " + string(data))
		err := u.Conn.WriteMessage(websocket.TextMessage, []byte(data))
		if err != nil {
			return err
		}
	default:
		logger.Debug("Default msg Type")
	}
	return nil
}

func deleteTopicUser(ctx context.Context, logger *logrus.Logger, topicId string, user *User) {
	for idx, duser := range UserList[topicId] {
		if user == duser {
			logger.Debug("deleted idx:", idx)
			user.Conn.Close()
			UserList[topicId] = append(UserList[topicId][:idx], UserList[topicId][idx+1:]...)
		}
	}
	if len(UserList[topicId]) <= 0 {
		delete(UserList, topicId)
	}
}

func PostWsMessage(ctx context.Context, logger *logrus.Logger, msg *Message) (bool, error) {
	if userlice, ok := UserList[msg.TopicId]; ok {
		for _, user := range userlice {
			user.mu.Lock()
			defer user.mu.Unlock()
			err := user.SendMsg(ctx, logger, msg)
			if err != nil {
				deleteTopicUser(ctx, logger, msg.TopicId, user)
				// return false, err
				logger.Debug("drop the conn with one user")
			}
		}
	} else {
		err := errors.New("Empty users are listing to this topic.")
		return false, err
	}
	return true, nil
}

// @BasePath /rt-msg-carrier/v1/
// Hello world godoc
// @Summary ws send message example
// @Schemes
// @Description do send ws message by topic
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} is_success:true
// @Router /ws [post]
func PostWsMessageHandler(c *gin.Context) {
	var msg Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"is_success": false, "error": err.Error()})
		return
	}
	ctx := context.Background()

	if ok, err := PostWsMessage(ctx, log.Logger, &msg); !ok {
		log.Logger.Error("postWsMessage error: %v", err)
		c.JSON(http.StatusNoContent, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"is_success": true,
	})
}

// @BasePath /rt-msg-carrier/v1/
// Websocket connection
// @Summary ws connection
// @Schemes
// @Description do ws connection by topic
// @Tags example
// @Accept application/json
// @Produce application/json
// @Success 200
// @Router /ws [get]
func NewWsConn(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Logger.Error(err)
		return
	}
	user := &User{
		Conn: ws,
	}
	topic, ok := c.GetQuery("ws_topic")

	if ok {
		UserList[topic] = append(UserList[topic], user)
	} else {
		log.Logger.Error("GetConn error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"is_success": false, "error": err.Error()})
		return
	}
	// c.JSON(http.StatusOK, gin.H{"is_success": "true"})
	// return
}
