package main

import (
	"github.com/panjf2000/ants"
	"fmt"
	"time"
)

func init(){
	fmt.Println("init ...")
}

func main(){
	runTimes := 100
	defer ants.Release()
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		time.Sleep(time.Second * 2)
		fmt.Println(i)
	})
	defer p.Release()
	for i := 0; i < runTimes; i++ {
		p.Invoke(int32(i))
	}
	time.Sleep(time.Second * 2)
}