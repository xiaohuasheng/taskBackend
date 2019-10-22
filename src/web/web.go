package web
import (
	"io"
	"log"
	"net/http"
)
func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!")
}
func HttpDemo() {
	//http://192.168.3.97:8080/hello
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
