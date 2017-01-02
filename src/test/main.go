package main

import(
	"os"
	"log"
	"strings"
)
var cwd string
func init() {
	cwd,err:=os.Getwd()
	if err!=nil{
		log.Fatalf("os.Getwd failed:%v",err)
	}
	log.Printf("working directory=%s",cwd)
}
func main() {
	
}

