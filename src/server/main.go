package main

import(
	"fmt"
	"bufio"
	"log"
	"net"
	"time"
	"strings"
)
func main() {
	listener,err:=net.Listen("tcp","localhost:8000")
	if err!=nil{
		log.Fatal(err)
	}
	for{
		conn,err:=listener.Accept()
		if err!=nil{
			log.Print(err)
			continue;
		}
		go handleConn(conn)
	}
}
func handleConn(c net.Conn){
	defer c.Close()
	input:=bufio.NewScanner(c) 
	for input.Scan(){
		echo(c,input.Text(),1*time.Second)
	}
}
func echo(c net.Conn,shout string,deley time.Duration){
	fmt.Fprintln(c,"\t",strings.ToUpper(shout))
	time.Sleep(deley)
	fmt.Fprintln(c,"\t",shout)
	time.Sleep(deley)
	fmt.Fprintln(c,"\t",strings.ToLower(shout))
}
