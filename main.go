package main

import(
	"fmt"
)

func main() {
	name := GetUserInfo("Введите имя")
	surname := GetUserInfo("Введите фамилию")
	fmt.Println(name, surname)
}

func GetUserInfo(str string) string {
var info string = "" 

fmt.Println(str) 

if _, err := fmt.Scanf("%s\n", &info); err != nil{
fmt.Println(err)
return info
}
return info
}