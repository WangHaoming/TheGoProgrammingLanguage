package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"shoes": 50, "socks": 5}

	// 创建mux，然后把mux作为参数传递给ListenAndServe
	// mux := http.NewServeMux()
	// mux.Handle("/list", http.HandlerFunc(db.list))
	// mux.Handle("/price", http.HandlerFunc(db.price))
	// log.Fatal(http.ListenAndServe("localhost:8000", mux))

	// 直接用http默认的mux net/http包提供了一个全局的ServeMux实例DefaultServerMux和包级别的http.Handle和http.HandleFunc函数。
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type database map[string]dollars
type dollars float32

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %.2f\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%.2f\n", price)
}
