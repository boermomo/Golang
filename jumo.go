package main

import (
	"fmt"
)

type Snowpea interface {
	shoot(a string)
	retard()
}

type Repeatshoot interface {
	repeat()
}
type Repeater struct {
	Snowpea
}

func (op1 *Repeater) shoot(a string) {
	fmt.Printf("不论遇到什么%s,都不要害怕！\n", a)
	fmt.Println("微笑着面对他！")
}

func (op2 *Repeater) retard() {
	fmt.Println("巨魔时代前来考古")
}
func (op3 *Repeater) repeat() {
	for i := 0; i < 3; i++ {
		fmt.Println("奥里给！！")
	}
}
func main() {
	var ans Snowpea
	var ans2 Repeatshoot
	jumo := new(Repeater)
	ans = jumo
	ans2 = jumo
	ans.shoot("憨憨")
	ans2.repeat()
}
