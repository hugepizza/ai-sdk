package aiqqsdk

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Httpserver 放在有公网IP的机器上测试callback
func Httpserver() {
	http.Handle("/", http.FileServer(http.Dir("/usr/local/share")))
	http.HandleFunc("/audio", func(resp http.ResponseWriter, req *http.Request) {
		bs, _ := ioutil.ReadAll(req.Body)
		fmt.Println(string(bs))
	})
	err := http.ListenAndServe(":20019", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
