package main

import "fmt"
import "links"
import (
	"log"
)

var tokens=make(chan struct{},200)
func crawl(url string)[]string{
	fmt.Println(url)
	tokens<-struct{}{}
	list,err:=links.Extract(url)
	<-tokens
	if err!=nil{
		log.Print(err)
	}
	return list
}
func startCrawl() {
	worklist:=make(chan[]string)
	var n int
	n++
	ll:=[]string{"https://baidu.com","http://gopl.io"}
	go func(){worklist<-ll}()
	seen:=make(map[string]bool)
	for ;n>0;n--{
		list:=<-worklist
		for _,link:=range list{
			if!seen[link]{
				seen[link]=true
				n++
				go func(link string) {
					worklist<-crawl(link)
				}(link)
			}

		}
	}
}
func main() {
	startCrawl()
	// naturals:=make(chan int)
	// squares:=make(chan int)
	// go counter(naturals)
	// go squarer(squares,naturals)
	// printer(squares)
	// ret:=mirroredQuery()
	// fmt.Println(ret)
	// go func(){
	// 	for x:=0;x<5;x++{
	// 		naturals<-x
	// 	}
	// 	close(naturals)
	// }()
	// go func(){
	// 	for x:=range naturals{
	// 		squares<-x*x
	// 	}
	// 	close(squares)
	// }()
	// for x:=range squares{
	// 	fmt.Println(x)
	// }
	// fmt.Println("vim-go")
}
