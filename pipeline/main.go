package main

import "fmt"
func counter(out chan<-int){
	for x:=0;x<5;x++{
		out<-x
	}
	// close(out)
}
func squarer(out chan<-int,in <-chan int){
	for x:=range in{
		out<-x*x
	}
	// close(out)
}
func printer(in <-chan int) {
	for v:=range in{
		fmt.Println(v)
	}
}
func main() {
	naturals:=make(chan int)
	squares:=make(chan int)
	go counter(naturals)
	go squarer(squares,naturals)
	printer(squares)
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
