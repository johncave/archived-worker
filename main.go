package main

import (
	"time" // temporary for simulating work.
	"fmt"
	"menteslibres.net/gosexy/redis"
)

var rediscon *redis.Client

func main(){

	//Print useless garbage to screen.
	fmt.Println("Now starting worker...")
	
	//Connect to the Redis server
	rediscon = redis.New()
	err := rediscon.Connect("127.0.0.1", 6379)
	if err != nil{
	panic("Could not connect to Redis")
	}

	rediscon.Select(7)

	msg, err := rediscon.Ping()
	if err != nil {
		panic("Could not ping Redis")
	} else {
	fmt.Println("Pinged Redis!:"+msg)
	}

	fmt.Println("Connected to Redis")
	
	dowork()
	
}

func dowork(){
	for true{
		
		job, err := rediscon.BRPop(0, "jobs")
		if err != nil {
			fmt.Println("No Job "+err.Error())
		}
		fmt.Println("Got a job: "+job)
		
		
		time.Sleep(5 * time.Second)
		
		
		
	}
}