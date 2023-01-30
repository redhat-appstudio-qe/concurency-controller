package main

import (
	"errors"
	"log"
	"testing"
	"time"

	controller "github.com/redhat-appstudio-qe/concurency-controller/controller"
)


func doreqstgeterr() error {
	return errors.New("Simple error")
}
func runner(id int, active int) error {
	log.Println("Making Request: ", id, "of batch: ", active)
	if id == 3 { return doreqstgeterr() }
	return nil
	
}

func testFunction(counter *int64) error {
    // Replace this with your actual function to test
    log.Println("Making request: ", *counter)
    *counter++
    if *counter%3 == 0 { return errors.New("temp error")}
    return nil
}


func TestController(t *testing.T){
	// Execute in Batches 
	// This takes two parameters i.e max no of requests to make and number of batches it has to execute those requests 
	// RPS is automatically calculated based on the above params
	// example: assume MAX_REQ = 50 , BATCHES = 5 then RPS = 10 
	// if you want to capture/send metrics please  provide the third parameter i.e MonitoringURL
	// Monitoring URL should point to hosted/self hosted instance of https://github.com/redhat-appstudio-qe/perf-monitoring
	// if you dont want to push metrics then just pass an empty string
	// MAX_REQ := 500
	// BATCHES := 10
	// controller.NewLoadController(MAX_REQ,BATCHES, "http://localhost:8000/").ConcurentlyExecute(runner)



	// Execute infinitely untill a timeout is met 
	// This takes two parameters i.e timeout duration and RPS (Requests Per Second)
	// if you want to capture/send metrics please  provide the third parameter i.e MonitoringURL
	// Monitoring URL should point to hosted/self hosted instance of https://github.com/redhat-appstudio-qe/perf-monitoring
	// if you dont want to push metrics then just pass an empty string
	TIMEOUT := 1 * time.Second
	RPS := 10
	controller.NewInfiniteLoadController(TIMEOUT, RPS, "").ExecuteInfinite(runner)



	// Executes infinitely untill a timeout is met 
	// This takes three parameters i.e timeout duration and MAX_RPS (Requests Per Second) and error threshold 
	// Runs in a way like locust it keeps on callinf the runner function, if the execution is error free it doulbles the RPS
	// if the execution has errors then it decrements the RPS gives more info on latency
	// if you want to capture/send metrics please  provide the third parameter i.e MonitoringURL
	// Monitoring URL should point to hosted/self hosted instance of https://github.com/redhat-appstudio-qe/perf-monitoring
	// if you dont want to push metrics then just pass an empty string
	TIMEOUT = 5 * time.Second
	maxRPS := 30
	errorThresholdRate := 0.5
	controller.NewSpikeLoadController(TIMEOUT, maxRPS, errorThresholdRate, "").CuncurentSpikeExecutor(testFunction)
}


