package main

import (
	"runtime"
	"flag"
	"time"
	"net/http"
	"io/ioutil"
)

var cpuNumber = flag.Int("cpu", 0, "Process number")
var duration = flag.Int("time", 0, "Runing mintues")
var memory = flag.Int("memory", 1, "Use memory size(M)")
var callback_uri = flag.String("callback", "", "Call back uri")

const hog_deep = 38

func userInput() {
	flag.Parse()

	if *cpuNumber == 0 {
		*cpuNumber = runtime.NumCPU()
	}

	if *duration == 0 {
		*duration = 1
	}

}

func main() {
	userInput()

	duration := time.Duration(*duration) * time.Minute
	println("Start stress process on", *cpuNumber, "CPU(s)!")
	println("Processes will run", duration.String(), "!")

	for i := 0; i < *cpuNumber; i ++ {
		go threads()
	}

	buffer := memoryShuffle(*memory)
	time.Sleep(duration)
	buffer[0] = 0

}

func threads() {
	for i := 0; true; i ++ {
		cpuHog(hog_deep)
		callback()
	}
}

func memoryShuffle(size int) (buffer []byte) {
	if size < 5 {
		size = 5
	}else {
		size = size * (1 << 20) - (5064 << 10)
	}
	buffer = make([]byte, size)

	for i := 0; i < size; i++ {
		buffer[i] = 1
	}

	return buffer

}

func cpuHog(i int) (result uint64) {
	if i < 2 {
		return uint64(1)
	}
	if i < 5 {
		return cpuHog(i - 1) + cpuHog(i - 2)
	}
	return cpuHog(i - 1) + cpuHog(i - 2) - cpuHog(i - 5)
}

func callback() (bool) {

	if *callback_uri == ""{
		return false
	}
	res, err := http.Get(*callback_uri)
	if err != nil {
		return false
	}

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return false
	}
	return true
}

