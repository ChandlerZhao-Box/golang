package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/once", readBodyOnce)
	http.HandleFunc("/url/query", queryParams)
	http.HandleFunc("/url/wholeUrl", wholeUrl)
	http.HandleFunc("/form", form)
	http.HandleFunc("/header", header)
	http.ListenAndServe(":8080", nil)
}

func header(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "request header: %v\n", r.Header)
}

func form(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "before parse form %v\n", r.Form)
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "parse form err: %v", r.Form)
	}

	fmt.Fprintf(w, "after parse form %v\n", r.Form)
}

func wholeUrl(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(r.URL)
	fmt.Fprintf(w, string(data))
}

func queryParams(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	name := values["name"][0]
	fmt.Println(name)
	fmt.Fprintf(w, "query is %v\n", values)
}

func readBodyOnce(w http.ResponseWriter, r *http.Request) {


	body, err:= io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed!")
		//这里要返回，不然就会执行后续的代码
		return
	}

	//类型转换， 将[]byte类型转化为string
	fmt.Fprintf(w, "read the data: %s \n", string(body))

	//尝试再次读取，什么也读不到，也是也不会报错
	body, err = io.ReadAll(r.Body)
	if err != nil {
		//不会进入到这里
		fmt.Fprintf(w, "read the data one more time got err: %v", err)
		return
	}

	fmt.Fprintf(w, "read the data one more time: %s and the data length: %d ", string(body), len(body))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is home")
}
