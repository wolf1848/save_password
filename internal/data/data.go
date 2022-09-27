package data

type User struct {
	Id       int    `json:"id"`
	Login    string `json:"login"`
	FullName string `json:"full_name"`
	Pwd      string `json:"pwd"`
}
