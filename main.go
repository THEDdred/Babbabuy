package main

import(
	"fmt"
)

func main() {
	age := GetAge()
	CheckAge(age, 18)
}

func GetAge() int {
	var age int
	fmt.Println("Сколько тебе лет?")

	if _, err := fmt.Scanf("%d\n", &age); err !=nil {
	fmt.Println(err)
	return age
}
	return age
}

func CheckAge(age, limitAge int) {
	if age >= limitAge {
	printAge(age)
	return
} else if age == 18 {
		fmt.Println("Проходи, бродяга")
		return
}
	printError()
}

func printError() {

	fmt.Println("вы не можете пройти!")
}

func printAge(age int) {
	fmt.Printf("твой возраст - %d.\nМожешь пройти", age)
}
//.dlpsa.p