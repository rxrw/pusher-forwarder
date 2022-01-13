package parsers

// 根据请求体解析用
type DingtalkEntity struct {
	Msgtype string       `json:"msgtype"`
	Text    DingtalkText `json:"text"`
}

type DingtalkText struct {
	Content string         `json:"content"`
	At      DingtalkTextAt `json:"at"`
}

type DingtalkTextAt struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

// 自己定义一个标准格式，
func (d DingtalkEntity) ToStandard() *StandardEntity {
	return &StandardEntity{
		Title:   "钉钉消息通知",
		Content: d.Text.Content,
	}
}
