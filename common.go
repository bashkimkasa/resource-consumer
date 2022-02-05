package main

import (
	"flag"
	"runtime"
)

//various constants and variables used in the application

var (
	Port                     = flag.Int("port", 8080, "Port number.")
	MemoryGoRoutineExecution = 0 //default to 0 - works as a singleton
	CpuGoRoutineExecution    = 0 //default to 0 - works as a singleton
)

const (
	ConsumeMemoryAddress  = "/consume-memory"
	ConsumeCpuAddress     = "/consume-cpu"
	GetHealthCheckAddress = "/health"
	GetInfoAddress        = "/info"
	GetRootAddress        = "/"

	MillicoresQuery  = "millicores"
	MegabytesQuery   = "megabytes"
	durationSecQuery = "durationSec"

	ChunkSize = 1048576 // 1MB

	BadRequest               = "Bad request"
	MethodNotAllowed         = "Method not allowed"
	NotFound                 = "Not Found"
	InternalServerError      = "Internal Server Error"
	InvalidQueryParamsMemory = "Invalid input query params and/or values. Required query params are: [megabytes, durationSec]"
	InvalidQueryParamsCpu    = "Invalid input query params and/or values. Required query params are: [millicores, durationSec]"
	MemoryConsumeError       = "Memory consuming process is already running. Try again later..."
	CpuConsumeError          = "CPU consuming process is already running. Try again later..."
	UnsupportedSystemError   = "This operating system is not supported"
)

const InfoMessage = `Resource consumer service for memory and CPU
		
Usage:
HTTP GET http(s)://<host>:<port>/consume-memory?megabytes=<int>&durationSec=<int>
HTTP GET http(s)://<host>:<port>/consume-cpu?millicores=<int>&durationSec=<int>

Examples:
Consume 100 MB of memory for a period of 5 minutes (300 seconds)
HTTP GET http://localhost:8080/consume-memory?megabytes=100&durationSec=300

Consume 10 Millicores for a period of 10 minutes (600 seconds)
HTTP GET http://localhost:8080/consume-cpu?millicores=10&durationSec=600

Notes:
- Once you have started to consume a resource (Memory and/or CPU) you cannot consume it again until its done processing (duration in seconds).
- A "/health" endpoint is also exposed which return HTTP 200.
`

// Supported OSs for resource consumtion are below and this function checks accordingly.
func CheckSupportedSystem() bool {
	var supported = []string{"linux", "aix", "freebsd", "solaris"}
	for _, v := range supported {
		if v == runtime.GOOS {
			return true
		}
	}
	return false
}
