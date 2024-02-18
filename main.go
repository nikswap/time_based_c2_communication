//Decode output from the server on *nix:
//for i in $(seq 0 15); do curl -s -w "%{time_total}" -o /dev/null 'http://localhost:3333?bitNumber='"$i" | awk -F '.' '{print $1}'; done | tr '31' '01' | tr -d '\n'

//Decode using powershell
//0..15 | % {Measure-Command -Expression { iwr "http://localhost:3333?bitNumber=$_"} | select -Property TotalSeconds } | % {[math]::Round($_.TotalSeconds) -replace "3","0" -replace "1", "1"} | join-string 

/*
TODO:
   * Implement checksum
   * Implement more time delay for sending different letters
*/

package main

import (
	"io"
	"log"
	"net/http"
	"time"
	"strconv"
	"fmt"
)


func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func testSendMessage(w http.ResponseWriter, r *http.Request) {
	bitVal := r.URL.Query().Get("bitNumber")
	message := "0100100001001001"
	bitCount,err := strconv.Atoi(bitVal)
	CheckError(err)
	timeToWait := 3
	if message[bitCount] == 49 {
		timeToWait = 1
	}
	fmt.Printf("Sending %d\n", message[bitCount])
	time.Sleep(time.Duration(timeToWait) * time.Second)
	io.WriteString(w, "OK")
}

func main() {
	http.HandleFunc("/", testSendMessage)
	err := http.ListenAndServe(":3333", nil)
	CheckError(err)
}
