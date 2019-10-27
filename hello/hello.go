package hello

import (
	"fmt"
)

// SayHello return "Hello: %s"
func SayHello(name string) string {
	return fmt.Sprintf("Hello: %s", name)
}