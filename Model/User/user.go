package User

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func GetUsers() (ret *User) {
	return &User{
		Id:   1,
		Name: "中文",
	}
}
