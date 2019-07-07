package testmod

import (
	"fmt"
	"time"
)

func SayHello(name string) string {
	return fmt.Sprintf("你好: %s --%s", name, time.Now())
}