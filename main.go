package main

import (
	"encoding/base64"
	"fmt"
	"bufio"
	"os"
	"strings"
	"time"
    "log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
	WriteBufferSize: 1024,
  	CheckOrigin: func(r *http.Request) bool { return true },
}

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn) {
    for {
    // read in a message
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
    // print out that message for clarity
        fmt.Println(string(p))

        if err := conn.WriteMessage(messageType, p); err != nil {
            log.Println(err)
            return
        }

    }
}

// define our WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.Host)

  // upgrade this connection to a WebSocket
  // connection
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
  }
  // listen indefinitely for new messages coming
  // through on our WebSocket connection
    reader(ws)
}

func setupRoutes() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Simple Server")
  })
  // make our `/ws` endpoint to the `serveWs` function
	http.HandleFunc("/ws", serveWs)
	http.HandleFunc("/touch", touchREST)
}

func touchREST(w http.ResponseWriter, r *http.Request) {
	AddCors(&w)

	fmt.Print("I'm touched!")
}

func main() {

	setupRoutes()
    log.Fatal(http.ListenAndServe(":8080", nil))
	
}

func oldMainCode() {
	// Ingest data through console
	// -new- will have to get user value from UI
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text to Encode to Base64: ")
	text, _ := reader.ReadString('\n')

	// Remove Carriage Return from ingested data
	text = strings.TrimSuffix(text, "\n")

	// Begin clock for golang library
	startLib := time.Now()

	// Encode string in Base64 with Golang Library for comparison to our implementation
	libraryEncode := libraryEncode(text)

	// Capturing golang library runtime
	elapsedLib := time.Since(startLib)

	// Begin clock for golang library
	startImp := time.Now()

	// Comparing our implementation with golang's for validation
	valid := compareImplementations(implementedEncode(text), libraryEncode)

	if(!valid) {
		fmt.Println("Failed case in implementedEncode")
	}

	// Capturing golang library runtime
	elapsedImp := time.Since(startImp)

	// Printing run time for statistics
	fmt.Printf("Encoding to Base64 with Golang library took %s. While encoding to Base64 with Implemented Algorithm took %s .", elapsedLib, elapsedImp)
	fmt.Println()
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func libraryEncode(msg string) string{
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println("Golang Libray Base64 Encode: ")
	fmt.Println(encoded)
	fmt.Println()
	return encoded
}

func implementedEncode(msg string) string{
	overage := len(msg) % 3
	padding := 0
	
	if (overage > 0) {
		// Add empty characters to our string based on how much padding is needed
		for i := overage; i < 3; i++ {
			padding++
			msg = msg + "\000"
		}
	}

	// String codex to map 6 bit base64 character values to
	base64Codex := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";
	encodedString := ""

	/* Iterate through the string and convert 3 chars at a time for conversion from 
	   3 8bit numbers to 4 6bit numbers */
	for i := 0; i < len(msg); i+=3 {

		// Adding three 8 bit numbers into one 24 bit number
		n := (int(msg[i]) << 16) + (int(msg[i+1]) << 8) + int(msg[i+2])

		// Parsing the 24 bit number into four 6 bit numbers we can map to our codex
		n1 := (n >> 18) & 63
		n2 := (n >> 12) & 63
		n3 := (n >> 6) & 63
		n4 := (n) & 63

		// Mapping the 6 bit values to our codex string 
		encodedString += "" + string(base64Codex[n1]) + string(base64Codex[n2]) + string(base64Codex[n3]) + string(base64Codex[n4])
	}
	
	// Removing chars from end of string based on how much padding was needed
	encodedStringFinal := encodedString[:len(encodedString)-(padding)]

	// Adding '=' based on how many characters were removed
	for i := 0; i < padding; i++ {
		encodedStringFinal = encodedStringFinal + "="
	}

	fmt.Println("Implemented Base64 Encode: ")
	fmt.Println(encodedStringFinal)
	fmt.Println()
	return encodedStringFinal
}

func compareImplementations(implemented string, library string ) bool{
	flag := false
	if(implemented == library) {
		fmt.Println("Test against Golang library implementation passed")
		flag = true
	} else {
		fmt.Println("Implemented Encode failed in the check against Golang Library Encode.")
	}

	return flag
}

func AddCors(w *http.ResponseWriter) {
	//Allow CORS here By * or specific origin
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}