package main

import (
	"fmt"
	"net"
)

func slow_scanning(){
	for i := 70; i <= 80; i++{
		webAddress := fmt.Sprintf("scanme.nmap.org:%d", i)
		con,err := net.Dial("tcp",webAddress)
		if err != nil {
			continue
		}
		con.Close()
		fmt.Printf("port %d is open\n",i)
	}
}

func faster_scanning(){
	
}

func main(){
	slow_scanning()
}

