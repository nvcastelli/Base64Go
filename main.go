package main

import (
	"encoding/base64"
	"fmt"
	"bufio"
	"os"
	"strings"
	"time"
)

func main() {

	// Ingest data through console
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
