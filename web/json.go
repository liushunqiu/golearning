package web

import "encoding/json"

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"llll"`
	Age       int    `json:"age"`
}

func CustomJson() {
	buf, err := json.Marshal(User{
		FirstName: "aaa",
		LastName:  "sss",
		Age:       22,
	})
	if err != nil {
		println("序列化错误")
	}
	println(string(buf))
}
