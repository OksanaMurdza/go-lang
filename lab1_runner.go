package main

import (
	"./lab1"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"runtime/pprof"
)

var profiler = flag.String("cpuprofile", "", "Profiler output file")
var inputFile = flag.String("inputFile", "", "File to parse")
var template = flag.String("template", "", "Template string")
var cycles = flag.Int("cycles", 1, "Loop cycles")

const MaxLenToPrint = 1000

func main() {
	flag.Parse()
	if *profiler != "" {
		f, err := os.Create(*profiler)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	args := flag.Args()
	if *inputFile == "" && len(args) > 0 {
		*inputFile = args[0]
	}
	if *template == "" && len(args) > 1 {
		*template = args[1]
	}
	bytes, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	input := string(bytes)
	var result []string
	for i := 0; i < *cycles; i++ {
		result = lab1.TemplateMatches(input, *template)
	}
	if len(input) > MaxLenToPrint {
		input = "Too long, won`t be printed"
	}
	println("Input: ", input)
	println("Template: ", *template)
	println("Results:")
	for _, e := range result {
		println(e)
	}
	println("Total: ", len(result))
}
