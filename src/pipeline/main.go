package main

import "fmt"
import "links"
import (
	"log"
	"crypto/sha256"
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
	type currency int
	const(
		USD currency=iota
		EUR
		GBP
		RMB
		)
	symbol:=[]string{USD:"$",EUR:"￡"}

	fmt.Println(EUR,symbol[EUR])
	a:=[2]int{1,2}
	b:=[...]int{1,2}
	c:=[2]int{1,3}
	d:=[]int{1,2}
	fmt.Println(a==b,a==c,b==c,a==d)
	c1:=sha256.Sum256([]byte("x"))
	c2:=sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n",c1,c2,c1==c2,c1)
	
	// startCrawl()
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
