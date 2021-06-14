package main

import "fmt"

func derefString(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}

func derefInt(val *int) int {
	if val == nil {
		return 0
	}
	return *val
}

func main() {
	fmt.Println(derefString(nil)) // ""
	fmt.Println(derefInt(nil))    // 0

	var str string = "Something"
	fmt.Println(derefString(&str)) // Something

	str = "another"
	fmt.Println(derefString(&str)) // another

	var anInt int = 10
	fmt.Println(derefInt(&anInt)) // 10

	anInt = 30
	fmt.Println(derefInt(&anInt)) // 30

}
