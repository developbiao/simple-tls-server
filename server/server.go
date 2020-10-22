package server

import (
	"fmt"
	"go/build"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)
// Cached words ensure concurrency safety
var safeMap = struct{
	sync.RWMutex
	m map[string]bool
}{m: make(map[string]bool)}

// This handle function process user query string
func SayHelloServer(w http.ResponseWriter, req *http.Request) {
	 // Parsing client query string
	 req.URL.Query()
	 params := req.URL.Query()
	 receivedWords := strings.Split(params.Get("q"), ",")

	 // Check user specified words already exits in the map ?
	 result := make([]bool, len(receivedWords))
	 for k, v := range receivedWords {
	 	// read from the safeMap, take the read/write lock
	 	safeMap.Lock()
	 	_, ok := safeMap.m[v]
	 	if  !ok {
	 		// save new record
			safeMap.m[v] = true
		}
		// release read/write lock
		safeMap.Unlock()
		result[k] = ok
	 }

	 fmt.Println("Query:", receivedWords)
	 fmt.Println("Reply:", result)

	 w.Header().Set("Content-Type", "text/plain")
	 // We can format the []bool for client friendly easy read e.g: true,true,false
	 strResult := formatBooleansToStr(result)
	 w.Write([]byte(strResult))
}

// format boolean array to string
func formatBooleansToStr(booleans []bool) string {
	result := ""
	for _, val := range booleans {
		if val {
			result += "true,"
		}else {
			result += "false,"
		}
	}
	return strings.TrimSuffix(result, ",")
}


// Start server listen on 443 port
func StartHttpsServer() {
	// register route "/hello"
	http.HandleFunc("/hello", SayHelloServer)

	// Get certificate path
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		goPath = build.Default.GOPATH
	}
	certRoot := goPath + "/src/interview-go/cert/"

	// server listen on 443 port
	err := http.ListenAndServeTLS(":443", certRoot + "server.crt", certRoot + "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

