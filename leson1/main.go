package main

// для инита необходимо выполнить  go mod init leson1
// go build main.go
import (
	"fmt"
	"time"
)

func main() {
	name := "Go Developers"
	fmt.Println("Azure for", name)
	fmt.Println("System time - ", time.Now())
}
