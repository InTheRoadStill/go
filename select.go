package main

import(
	"fmt"
	"time"
	"os"
	"strconv"
)

func init(){
	fmt.Println("init loading...")
}

func main(){
	var s chan string = make(chan string)
	go func(){
		i := 1
		 for{
		 	time.Sleep(time.Second * 5)
		 	s <- strconv.Itoa(i)
		 	i++
		 }
	}()
	for{
		select{
		    case value := <-s:
				txt, err := os.OpenFile(`test.txt`, os.O_APPEND|os.O_CREATE, 0666)
				if err != nil {
					panic(err)
				}
				value = value + "\r\n"
				// 写入文件
			    n, err := txt.Write([]byte(value))
			    // 当 n != len(b) 时，返回非零错误
			    if err == nil && n != len(value) {
			        println(`错误代码：`, n)
			        panic(err)
			    }
			    txt.Close()
		}
	}
}