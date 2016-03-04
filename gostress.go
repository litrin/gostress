/*

	Copyright (c) 2016, Litrin Jiang
	All rights reserved.

	Redistribution and use in source and binary forms, with or without
	modification, are permitted provided that the following conditions are met:

	* Redistributions of source code must retain the above copyright notice, this
	  list of conditions and the following disclaimer.

	* Redistributions in binary form must reproduce the above copyright notice,
	  this list of conditions and the following disclaimer in the documentation
	  and/or other materials provided with the distribution.

	THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
	AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
	IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
	DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
	FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
	DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
	SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
	CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
	OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
	OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

*/

package main

import (
	"runtime"
	"flag"
	"time"
)

var cpuNumber = flag.Int("cpu", 0, "Process number")
var duration = flag.Int("time", 0, "Runing mintues")
var memory = flag.Int("memory", 1, "Use memory size(M)")

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
		cpuHog(38)
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