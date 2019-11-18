package main

import (
	"io/ioutil"
	"net/http"
	"fmt"
	"github.com/panjf2000/ants"
)

type Request struct {
	Param  []byte
	Result chan []byte
}

func main() {
	pool, _ := ants.NewPoolWithFunc(100, func(payload interface{}) {
		request, ok := payload.(*Request)
		fmt.Println(request)
		if !ok {
			return
		}
		reverseParam := func(s []byte) []byte {
			for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
				s[i], s[j] = s[j], s[i]
			}
			return s
		}(request.Param)
		fmt.Println(string(reverseParam[:]))
		reverseParam = []byte("hello world")
		request.Result <- reverseParam
	})
	defer pool.Release()

	http.HandleFunc("/reverse", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("get it!")
		param, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "request error", http.StatusInternalServerError)
		}
		defer r.Body.Close()

		request := &Request{Param: param, Result: make(chan []byte)}

		// Throttle the requests traffic with ants pool. This process is asynchronous and
		// you can receive a result from the channel defined outside.
		if err := pool.Invoke(request); err != nil {
			http.Error(w, "throttle limit error", http.StatusInternalServerError)
		}
		fmt.Println("return")
		w.Write(<-request.Result)
	})

	http.ListenAndServe(":9003", nil)
}