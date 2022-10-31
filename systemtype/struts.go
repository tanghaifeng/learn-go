package systemtype

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u *User) GetUserInfo() {
	fmt.Println(u.Name)
	fmt.Println(u.Age)
}

func (u *User) SetInfo(user User) {
	if user.Name != "" {
		u.Name = user.Name
	}
	if user.Age != 0 {
		u.Age = user.Age
	}
}

func GetUser(info UserInfo) {
	info.GetUserInfo()
}
