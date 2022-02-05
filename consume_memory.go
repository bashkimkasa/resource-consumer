package main

import (
	"log"
	"os/exec"
	"strconv"
)

// ConsumeMemory consumes a given number of megabytes for the specified duration.
func ConsumeMemory(megabytes, durationSec int) {
	//Set global variable to indicate that memory resource consumer is now running
	MemoryGoRoutineExecution = 1

	log.Printf("Consuming memory in megabytes: %v, Duration period in seconds: %v", megabytes, durationSec)

	megabytesString := strconv.Itoa(megabytes) + "M"
	durationSecString := strconv.Itoa(durationSec)

	// creating new consume memory process
	consumeMem := exec.Command("stress-ng", "-m", "1", "--vm-bytes", megabytesString, "--vm-hang", "0", "-t", durationSecString)
	err := consumeMem.Run()
	if err != nil {
		log.Println(err.Error())
		MemoryGoRoutineExecution = 0
		return
	}

	// no need to sleep here as 'stress-ng' command above is synchronous
	// time.Sleep(time.Second * time.Duration(durationSec))
	log.Println("Memory consuming process has now completed")

	//Set global variable to indicate that memory resource consumer is done
	MemoryGoRoutineExecution = 0
}
