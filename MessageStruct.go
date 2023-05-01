type messageType string

const (
	Message              messageType = "MESSAGE"
	Like                 messageType = "LIKE"
)

type Message struct {
	Id        string            `json:"id"`
	UserId    string            `json:"userId"`
	Message   string            `json:"message"`
	Type      messageType       `json:"messageType"`
	HasRead   map[string]string `json:"hasRead"`
	TimeStamp int64             `json:"timeStamp"`
	IsDeleted bool              `json:"isDeleted"`
}
