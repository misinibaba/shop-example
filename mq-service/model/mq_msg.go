package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type MqMsg struct {
	ID          int    `gorm:"primary_key" json:"id"`
	FromService string `gorm:"column:fromService" json:"fromService"`
	Message     string `gorm:"column:message" json:"message"`
	Topic       string `json:"topic"`
	Key         string `json:"key"`
	UniId       string `gorm:"column:uniId" json:"uniId"`
	Status      int    `json:"status"`
}

const (
	MSGDELETE   = -1
	MSGPRESEND  = 0
	MSGCOMMITED = 1 // 如果需要用到这个状态，那么就是commit后先改变状态为这个，然后一个后台任务轮询为commit的更改为sended；多一步我认为可以防止kafka宕机，然后切换kv库存储
	MSGSENDED   = 2
)
func (mqMsg MqMsg) TableName() string {
	return "mqMsg"
}

func GetMqMsg(mqMsgId int) (*MqMsg, error) {
	var mqMsg MqMsg
	err := db.Where("id = ? ", mqMsgId).First(&mqMsg).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &mqMsg, nil
}

func delayTwoSeconds(db *gorm.DB) *gorm.DB{
	m, _ := time.ParseDuration("-2s")
	return db.Where("createAt < ?", time.Now().Add(m).Format("2006-01-02 15:04:05"))
}

func GetMqMsgList(where map[string]interface{}) ([]MqMsg, error) {
	var mqMsgList []MqMsg
	err := db.Table("mqMsg").Scopes(delayTwoSeconds).Find(&mqMsgList, where).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return mqMsgList, nil
}

func UpdateMqMsgStatus(mqMsgId int, status int) error {
	err := db.Table("mqMsg").Where("id = ? ", mqMsgId).Update("status", status).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	return nil
}

func PreSendMqMsg(fromService, message, topic, key, uniId string) (int, error) {
	mqMsg := &MqMsg{FromService: fromService, Message: message, Topic: topic, Key: key, UniId: uniId, Status: MSGPRESEND}
	err := db.Table("mqMsg").Save(&mqMsg).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}
	return mqMsg.ID, nil
}