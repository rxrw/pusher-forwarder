package parsers

// Generated by https://quicktype.io

type PushoverEntity struct {
	Token   string `json:"token"`
	User    string `json:"user"`
	Message string `json:"message"`
	Title   string `json:"title"`
	URL     string `json:"url"`
}

// 自己定义一个标准格式，
func (d PushoverEntity) ToStandard() *StandardEntity {
	return &StandardEntity{
		Title:   d.Title,
		Content: d.Message,
		URL:     d.URL,
	}
}

func (d *PushoverEntity) SetToken(user string, token string) {
	d.Token = token
	d.User = user
}

func (d *PushoverEntity) SetTitle(title string) {
	d.Title = title
}
