package main
import(
	"github.com/panjf2000/ants"
	"fmt"
	"net/http"
	"time"
	"bytes"
	"io"
	"flag"
)

// 定义几个变量，用于接收命令行的参数值
var url        string
var number     int
var sleep      int

func main() {
    // &user 就是接收命令行中输入 -u 后面的参数值，其他同理
    flag.StringVar(&url, "url", "http://127.0.0.1:8888/?hello=123", "请求url")
    flag.IntVar(&number, "n", 20, "请求协程数")
    flag.IntVar(&sleep, "s", 0, "每个协程睡眠时间")
	// 解析命令行参数写入注册的flag里
    flag.Parse()
	p, _ := ants.NewPoolWithFunc(number, func(payload interface{}) {
		url, ok := payload.(string)
		if !ok {
			fmt.Println(payload)
			return
		}
		if sleep > 0 {
			time.Sleep(time.Duration(int64(sleep)) * time.Second)
		}
		fmt.Println(get(url))
	})

	for {
		p.Invoke(url)
	}
}

func get(url string) string {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for{
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	return result.String()
}