package logic

import "fmt"

var Clock = new(clockLogic)

type clockLogic struct{}

func (c *clockLogic) ZeroTime() {
	fmt.Println("0点到了.")
}
