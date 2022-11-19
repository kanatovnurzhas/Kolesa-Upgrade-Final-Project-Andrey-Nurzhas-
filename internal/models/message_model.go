package models

type Message struct {
	Message string `json:"message"`
}

//type MessageModel struct {
//	Db *gorm.DB
//}
//
//func (m *MessageModel) Create(message Message) error {
//
//	result := m.Db.Create(&message)
//
//	return result.Error
//}
