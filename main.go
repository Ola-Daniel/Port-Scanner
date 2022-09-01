//a multi thread tcp port scanner that takes a url or an ip address and scans all ports
// asks for url or ip address as user input

package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var host string
	var port int
	var start int
	var end int

	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")

	fmt.Println("Enter a url or ip address: ")
	fmt.Scanln(&host)

	fmt.Println("Enter a starting port: ")
	fmt.Scanln(&start)

	fmt.Println("Enter a ending port: ")
	fmt.Scanln(&end)

	for port = start; port <= end; port++ {
		wg.Add(1)
		go scan(host, port)
	}
	wg.Wait()
}

func scan(host string, port int) {
	defer wg.Done()
	address := host + ":" + strconv.Itoa(port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return
	}
	conn.Close()
	fmt.Println("Port", port, "is open")
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}


