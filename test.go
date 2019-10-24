package main

import(
	"fmt"
)
// range只有第一次初始化地址
func main() {
	arr := [...]int{1,2,3}
	var m map[int]*int = make(map[int]*int)
	for k,v := range arr {
		fmt.Println(v)
		m[k] = &v
	}
	fmt.Println(m)
}