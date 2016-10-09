package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

var url = flag.String("url", "", "-url ws://127.0.0.1")
var headers = flag.String("headers", "", "-headers name1=value1,name2=value2")

func main() {
	flag.Parse()
	if *url == "" {
		log.Fatal("URL can't be empty.")
	}
	headersMap := make(map[string]string)
	if *headers != "" {
		headersSeq := strings.Split(*headers, ",")
		for _, headerString := range headersSeq {
			splittedHeader := strings.Split(headerString, "=")
			headersMap[splittedHeader[0]] = splittedHeader[1]
		}
	}
	done := make(chan bool)
	c := NewClient(*url, headersMap)

	go func() {
		for {
			_, r, err := c.ReadMessage()
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Println("\n", color(string(r)))
		}
	}()

	r := bufio.NewReader(os.Stdin)
	go func() {
		for {
			line, err := r.ReadString('\n')
			if err != nil && err.Error() != "unexpected newline" {
				fmt.Println(err)
				return
			}
			line = strings.TrimSpace(line)
			if err = c.WriteMessage(1, []byte(line)); err != nil {
				fmt.Println(err.Error())
			}
		}
	}()

	<-done

}

func NewClient(url string, headers map[string]string) *websocket.Conn {
	r, _ := http.NewRequest("GET", url, nil)
	r.Header.Add("Content-Type", "application/json")
	for key, value := range headers {
		r.Header.Add(key, value)
	}
	c, _, err := websocket.DefaultDialer.Dial(url, r.Header)
	if err != nil {
		log.Fatal("errrr ", err)
	} else {
		fmt.Println("Connected!")
	}

	return c
}

func color(msg string) string {
	return ("\033[36m" + msg + "\033[0m")
}
