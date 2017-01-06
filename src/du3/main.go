package main

import(
	"fmt"
	_"crypto/sha256"	
	"io/ioutil" 
	"os" 
	"path/filepath"
	"flag"
	"time"
	"sync"
)
var verbose =flag.Bool("v",false,"show verbose progress messages")
func walkDir(dir string,n *sync.WaitGroup,fileSizes chan<-int64) {
	defer n.Done()
	for _,entry:=range dirents(dir){
		if entry.IsDir(){
			n.Add(1)
			subdir:=filepath.Join(dir,entry.Name())
			go walkDir(subdir,n,fileSizes)
		}else{
			fileSizes<-entry.Size()
		}
	}
}
var sema=make(chan struct{},200)
func dirents(dir string)[]os.FileInfo {
	sema<-struct{}{}
	defer func(){<-sema}()
	entries,err:=ioutil.ReadDir(dir)
	if err!=nil{
		fmt.Fprintf(os.Stderr,"du1:%v\n",err)
		return nil
	}
	return entries
}
func printDiskUsage(nfiles,nbytes int64) {///1024/1024/1024
	fmt.Printf("%d files %.1f GB\n",nfiles,float64(nbytes)/1e9)
}

func start_du() {
	flag.Parse()
	roots:=flag.Args()
	if len(roots)==0{
		roots=[]string{"."}
	}
	fileSizes:=make(chan int64)
	var n sync.WaitGroup
	for _,root:=range roots{
		n.Add(1)
		go walkDir(root,&n,fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()
	var nfiles,nbytes int64
	var tick <-chan time.Time
	if *verbose{
		tick=time.Tick(500*time.Millisecond)
	}
	loop:
		for{
			select{
			case size,ok:=<-fileSizes:
				if !ok{
					break loop
				}
				nfiles++
				nbytes+=size
			case <-tick:
				printDiskUsage(nfiles,nbytes)
			}
		}
	printDiskUsage(nfiles, nbytes) // final totals
	
}

func main() {

	start_du()
	
}


