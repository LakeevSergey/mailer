package entity

type SendMail struct {
	Code     string
	SendTo   []string
	SendFrom *SendFrom
	Params   map[string]string
	Title    string
	Body     string
}
