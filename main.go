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