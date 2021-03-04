package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	f, err := os.Open("index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	z := html.NewTokenizer(f)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			fmt.Printf("%+v\n", tt)
			fmt.Printf("%+v\n", z.Err())
			fmt.Println("エラートークン↑")
			return
		}
		// Process the current token.
		token := z.Token()
		fmt.Printf("%+v  |%+v|\n", tt, token)

		attr := token.Attr
		fmt.Printf("%+v\n", attr)
		for _, a := range attr {
			fmt.Printf("%+v  %+v  %+v\n", a.Namespace, a.Key, a.Val)
		}
	}
}
