package main

import (
	"crypto/md5"
	"fmt"
	"time"
)

func renderAuth() string {
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%s", time.Now().Unix(), key))))
}

func main() {
	fmt.Println(renderAuth())
}
