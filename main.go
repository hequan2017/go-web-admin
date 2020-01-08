package main

import (
	"github.com/gogf/gf/frame/g"
	_ "go-web-admin/boot"
)

func main() {
	s := g.Server()
	s.Run()
}
