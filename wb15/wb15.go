package main

import (
	"fmt"
	"strings"
)

var justString string

func someFunc() {
	// Создадим функцию, которая возвращает срез из 1024 строк
	v := bigString(1 << 10)
	/* Воспользуемся функцией Clone, которая копирует определенные элементы(в нашем случае - первые 100),
	и при этом не ссылается на общий срез, по этому garbage collector может удалить неиспользуемые элементы среза*/
	justString = strings.Clone(v[:100])
	fmt.Println(justString)
}

func main() {
	someFunc()
}

func bigString(a int) string {
	var b string
	for a > 0 {
		a--
		b += "qwerty"
	}
	return b
}