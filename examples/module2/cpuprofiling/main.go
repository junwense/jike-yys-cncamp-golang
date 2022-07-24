package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
)

//go tool pprof /tmp/cpuprofile 查看文件
var cpuprofile = flag.String("cpuprofile", "/tmp/cpuprofile", "write cpu profile to file")

func main() {
	flag.Parse()
	f, err := os.Create(*cpuprofile)
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	var result int
	for i := 0; i < 100000000; i++ {
		result += i
	}
	log.Println("result:", result)
}
