package main

import "fmt"
func counter(out chan<-int){
	for x:=0;x<5;x++{
		out<-x
	}
	close(out)
}
func squarer(out chan<-int,in <-chan int){
	for x:=range in{
		out<-x*x
	}
	close(out)
}
func printer(in <-chan int) {
	for v:=range in{
		fmt.Println(v)
	}
}
func mirroredQuery()string {
	responses:=make(chan string,3)
	go func(){responses<-request("1")}()
	go func(){responses<-request("2")}()
	go func(){responses<-request("3")}()
	return <-responses
}
func request(hostname string)(response string){
	return hostname
}
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
	go func(){worklist<-"baidu.com"}()
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
