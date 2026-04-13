package main

import (
	"fmt"
	"net"
	"sync"
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

func faster_scanning(){ //faster but not accurate
	for i := 1; i <= 1024; i++ {
		go func(j int){
			webAddress := fmt.Sprintf("scanme.nmap.org:%d", i)
		con,err := net.Dial("tcp",webAddress)
		if err != nil {
			return
		}
		con.Close()
		fmt.Printf("port %d is open\n",i)
		}(i)
	}
}

func faster_accurate_scanning(){
	
}

func main(){
	faster_scanning()
}

