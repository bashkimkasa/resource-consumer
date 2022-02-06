# Resource Consumer

## Overview
Resource Consumer is a tool which allows to generate cpu/memory utilization in a container.
The reason why it was created is testing kubernetes autoscaling.
Resource Consumer can help with autoscaling tests for:
- cluster size autoscaling,
- horizontal autoscaling of pod - changing the size of replication controller,
- vertical autoscaling of pod - changing its resource limits.

This code references the following google resource:  
https://gitlab.cncf.ci/kubernetes/kubernetes/tree/9d0cbb7503b7070817b3ec08e76f3f3addf3675b/test/images/resource-consumer

## Usage
Resource Consumer starts an HTTP server and handles requests.
It listens on port given as a flag (default 8080).
Action of consuming resources is send to the container by a GET http request.

The container consumes specified amount of resources:
- CPU in millicores,
- Memory in megabytes,

Notes:
- Once you have started to consume a resource (Memory and/or CPU) you cannot consume it again until its done processing (duration in seconds).
- A "/health" endpoint is also exposed which return HTTP 200.
- An "/info" endpoint provides usage information for this tool
- Root ("/") endpoint is redirected by default to "/info" endpoint

### Consume CPU http request
- suffix "consume-cpu",
- parameters "millicores" and "durationSec".

Consumes specified amount of millicores for durationSec seconds.
When CPU consumption is too low this binary uses cpu by calculating math.sqrt(0) 10^7 times and if consumption is too high binary sleeps for 10 millisecond.
One replica of Resource Consumer cannot consume more that 1 cpu.

```console
Request Format:
HTTP GET http(s)://<host>:<port>/consume-cpu?millicores=<int>&durationSec=<int>

Curl Example:  
Consume 10 Millicores for a period of 10 minutes (600 seconds)  
curl http://localhost:8080/consume-cpu?millicores=10&durationSec=600
```

### Consume Memory http request
- suffix "consume-memory",
- parameters "megabytes" and "durationSec".

Consumes specified amount of megabytes for durationSec seconds.
Consume Memory uses stress-ng tool (stress-ng -m 1 --vm-bytes megabytes --vm-hang 0 -t durationSec).
Request leading to consuming more memory than container limit will be ignored.

```console
Request Format:  
HTTP GET http(s)://<host>:<port>/consume-memory?megabytes=<int>&durationSec=<int>

Curl Example:  
Consume 100 MB of memory for a period of 5 minutes (300 seconds)  
curl http://localhost:8080/consume-memory?megabytes=100&durationSec=300
```

## Image

Docker image of Resource Consumer can be found in DockerHub Container Registry as <fill-in-data>

## Use cases

### Cluster autoscaling
1. Consume more resources on each node that is specified for autoscaler
2. Observe that cluster size increased

### Horizontal pod autoscaling (HPA)
1. Create consuming RC and start consuming appropriate amount of resources
2. Observe that RC has been resized
3. Observe that usage on each replica decreased

### Vertical pod autoscaling (VPA)
1. Create consuming pod and start consuming appropriate amount of resources
2. Observed that limits has been increased
