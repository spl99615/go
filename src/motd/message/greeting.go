package message

import "fmt"

func Greeting(name, message string) (greeting string) {
	greeting = fmt.Sprintf("%s,%s", message, name)
	return
}
