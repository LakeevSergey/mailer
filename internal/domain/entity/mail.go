package entity

type Mail struct {
	SendTo      []string
	SendFrom    SendFrom
	Title       string
	Body        string
	Attachments []File
}
