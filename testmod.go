package testmod

import "fmt"

func SayHello(name string) string {
	return fmt.Sprintf("你好: %s", name)
}