package main

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"fmt"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		f := fib()

		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Hello World!!\n")

		for _, e := range os.Environ() {
			pair := strings.Split(e, "=")
			io.WriteString(w, pair[0]+"="+pair[1]+"\n")
		}

		for i := 1; i <= 90; i++ {
			io.WriteString(w, strconv.Itoa(f())+"\n")
		}

		fmt.Println("Hello world called - this is the log message")

	})
	http.ListenAndServe(":80", nil)
}

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
