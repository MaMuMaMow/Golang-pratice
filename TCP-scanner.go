package main

import (
	"fmt"
	"net"
	"sort"
	"sync"
)

//var v[]int

func slow_scanning() {
	for i := 70; i <= 80; i++ {
		webAddress := fmt.Sprintf("scanme.nmap.org:%d", i)
		con, err := net.Dial("tcp", webAddress)
		if err != nil {
			continue
		}
		con.Close()
		fmt.Printf("port %d is open\n", i)
	}
}

func faster_scanning() { //faster but not accurate
	for i := 1; i <= 1024; i++ {
		go func(j int) {
			webAddress := fmt.Sprintf("scanme.nmap.org:%d", j)
			con, err := net.Dial("tcp", webAddress)
			if err != nil {
				return
			}
			con.Close()
			fmt.Printf("port %d is open\n", i)
		}(i)
	}
}

func waitGroup_scanning() {
	var wg sync.WaitGroup
	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			webAddress := fmt.Sprintf("127.0.0.1:%d", j)
			conn, err := net.Dial("tcp", webAddress)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("port %d is open\n", j)
		}(i)
	}
	wg.Wait()
}

func worker(ports chan int, results chan int) {
	for p := range ports {
		webAddress := fmt.Sprintf("127.0.0.1:%d",p)
		conn, err := net.Dial("tcp",webAddress)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func the_best_scaning_one() {
	ports := make(chan int, 100)
	results := make(chan int)
	var openPorts[] int
	
	go func(){
		for i := 1; i <= 65535; i++ {
			ports <- i
		}
	}()

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	for i := 0 ; i < 65535; i++{
		port := <- results
		if port != 0{
			openPorts = append(openPorts, port)
		}
	}
	close(ports)
	close(results)
	sort.Ints(openPorts)
	for _, port := range openPorts {
		fmt.Printf("%d is open\n",port)
	}
	/*
	slices.Sort(v)
	for i := 0; i < len(v); i++ {
		fmt.Println(v[i])
	}
	*/

}


func main() {
	the_best_scaning_one()
}
