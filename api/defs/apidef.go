package defs

// reqeusts
//data model
type UserCredentials struct {
	Username string `json:"user_name"` //打tag的方式, 自动序列化 user_name: xxx, pwd: xxx
	Pwd      string `json:"password"`
}
