package repository

type Users struct {
	id       int32  `json:"id"`
	login    string `json:"login"`
	fullName string `json:"full_name,omitempty"`
}
