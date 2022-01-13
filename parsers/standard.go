package parsers

type StandardEntity struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	URL     string
}

func (s StandardEntity) ToPushover() PushoverEntity {
	return PushoverEntity{
		Title:   s.Title,
		Message: s.Content,
		URL:     s.URL,
	}
}
