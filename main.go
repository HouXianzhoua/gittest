package main

import (
	"flag"
	"log"
	"runtime"	
	"src"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()	
	log.Fatal(src.Run())
}

func init() {
	log.SetFlags(log.Lshortfile)
}

func jj() {
	log.Println("jj")
}