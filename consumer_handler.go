package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type ResourceConsumerHandler struct{}

// NewResourceConsumerHandler creates and initializes a ResourceConsumerHandler to defaults.
func NewResourceConsumerHandler() *ResourceConsumerHandler {
	return &ResourceConsumerHandler{}
}

func (handler *ResourceConsumerHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//handle healthcheck address
	if req.URL.Path == GetHealthCheckAddress {
		if req.Method != "GET" {
			http.Error(w, MethodNotAllowed, http.StatusMethodNotAllowed)
			return
		}
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		return
	}

	//handle info address
	if req.URL.Path == GetInfoAddress {
		if req.Method != "GET" {
			http.Error(w, MethodNotAllowed, http.StatusMethodNotAllowed)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, InfoMessage)
		return
	}

	//handle consume-memory address
	if req.URL.Path == ConsumeMemoryAddress {
		if req.Method != "GET" {
			http.Error(w, MethodNotAllowed, http.StatusMethodNotAllowed)
			return
		}
		//Check if the operating system is supported
		if !CheckSupportedSystem() {
			http.Error(w, UnsupportedSystemError, http.StatusConflict)
			return
		}
		//check if memory resource consumer is already running
		if MemoryGoRoutineExecution == 1 {
			http.Error(w, MemoryConsumeError, http.StatusConflict)
			return
		}
		durationSecString := req.URL.Query().Get(durationSecQuery)
		megabytesString := req.URL.Query().Get(MegabytesQuery)

		// convert data (strings to ints) for consume-mem
		durationSec, durationSecError := strconv.Atoi(durationSecString)
		megabytes, megabytesError := strconv.Atoi(megabytesString)
		if durationSecError != nil || megabytesError != nil {
			http.Error(w, InvalidQueryParamsMemory, http.StatusBadRequest)
			return
		}

		go ConsumeMemory(megabytes, durationSec)

		startTime := time.Now()
		endTime := startTime.Add(time.Second * time.Duration(durationSec))
		fmt.Fprintln(w, "Consuming memory in megabytes: ", megabytes)
		fmt.Fprintln(w, "Duration period in seconds: ", durationSec)
		fmt.Fprintln(w, "StartTime: ", startTime.Format(time.RFC3339))
		fmt.Fprintln(w, "Endtime: ", endTime.Format(time.RFC3339))
		return
	}

	//handle consume-cpu address
	if req.URL.Path == ConsumeCpuAddress {
		if req.Method != "GET" {
			http.Error(w, MethodNotAllowed, http.StatusMethodNotAllowed)
			return
		}
		//Check if the operating system is supported
		if !CheckSupportedSystem() {
			http.Error(w, UnsupportedSystemError, http.StatusConflict)
			return
		}
		//check if cpu resource consumer is already running
		if CpuGoRoutineExecution == 1 {
			http.Error(w, CpuConsumeError, http.StatusConflict)
			return
		}

		durationSecString := req.URL.Query().Get(durationSecQuery)
		millicoresString := req.URL.Query().Get(MillicoresQuery)

		// convert data (strings to ints) for consume-cpu
		durationSec, durationSecError := strconv.Atoi(durationSecString)
		millicores, millicoresError := strconv.Atoi(millicoresString)
		if durationSecError != nil || millicoresError != nil {
			http.Error(w, InvalidQueryParamsCpu, http.StatusBadRequest)
			return
		}

		go ConsumeCPU(millicores, durationSec)

		startTime := time.Now()
		endTime := startTime.Add(time.Second * time.Duration(durationSec))
		fmt.Fprintln(w, "Consuming CPU millicores: ", millicores)
		fmt.Fprintln(w, "Duration period in seconds: ", durationSec)
		fmt.Fprintln(w, "StartTime: ", startTime.Format(time.RFC3339))
		fmt.Fprintln(w, "Endtime: ", endTime.Format(time.RFC3339))
		return
	}

	//handle root address
	if req.URL.Path == GetRootAddress {
		http.Redirect(w, req, GetInfoAddress, http.StatusMovedPermanently)
		return
	}

	http.Error(w, fmt.Sprintf("%s: %s", NotFound, req.URL.Path), http.StatusNotFound)
}
