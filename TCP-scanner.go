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
			webAddress := fmt.Sprintf("scanme.nmap.org:%d", j)
		con,err := net.Dial("tcp",webAddress)
		if err != nil {
			return
		}
		con.Close()
		fmt.Printf("port %d is open\n",i)
		}(i)
	}
}

func waitGroup_scanning(){
	var wg sync.WaitGroup
	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		go func(j int){
			defer wg.Done()
			webAddress := fmt.Sprintf("127.0.0.1:%d", j)
			conn,err := net.Dial("tcp",webAddress)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("port %d is open\n",j)
		}(i)
	}
	wg.Wait()
}

func main(){
	waitGroup_scanning()
}

