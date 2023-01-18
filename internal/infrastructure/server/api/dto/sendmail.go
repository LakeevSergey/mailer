package dto

type SendMail struct {
	Code        string            `json:"code"`
	SendTo      []string          `json:"send_to"`
	SendFrom    *SendFrom         `json:"send_from"`
	Params      map[string]string `json:"params"`
	Title       string            `json:"title"`
	Body        string            `json:"body"`
	Attachments []int64           `json:"attachments"`
}
