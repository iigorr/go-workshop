package main

import (
	"fmt"
	"github.com/getwe/figlet4go"
)

func main() {
	ascii := figlet4go.NewAsciiRender()
	renderStr, _ := ascii.Render("Hello World!")
	fmt.Println(renderStr)
}
