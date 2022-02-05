package main

import (
	"log"
	"math"
	"time"

	"bitbucket.org/bertimus9/systemstat"
)

const cpuSleep = time.Duration(10) * time.Millisecond //10 milliseconds

func ConsumeCPU(millicores, durationSec int) {
	//Set global variable to indicate that cpu resource consumer is now running
	CpuGoRoutineExecution = 1

  log.Printf("Consuming CPU millicores: %v, Duration period in seconds: %v", millicores, durationSec)

	// convert millicores to percentage
	millicoresPct := float64(millicores) / float64(10)
	startTime := time.Now()
	duration := time.Second * time.Duration(durationSec)

	first := systemstat.GetProcCPUSample()
	for time.Since(startTime) < duration {
		cpu := systemstat.GetProcCPUAverage(first, systemstat.GetProcCPUSample(), systemstat.GetUptime().Uptime)
		if cpu.TotalPct < millicoresPct {
			keepBusy()
		} else {
			time.Sleep(cpuSleep)
		}
	}
	log.Printf("CPU consuming process has now completed")
	
	//Set global variable to indicate that cpu resource consumer is done
	CpuGoRoutineExecution = 0
}

func keepBusy() {
	for i := 1; i < 10000000; i++ {
		x := float64(0)
		x += math.Sqrt(0)
	}
}