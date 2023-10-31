package db

type User struct {
	ID         int64  `json:"id,omitempty"`
	Phone      string `json:"phone,omitempty"`
	Password   string `json:"password,omitempty"`
	Sex        int32  `json:"sex,omitempty"`
	Nickname   string `json:"nickname,omitempty"`
	CreateTime int64  `json:"create_time,omitempty"`
	UpdateTime int64  `json:"update_time,omitempty"`
}

func RegisterUser(user *User) {
	dbCon.Create(user)
}
