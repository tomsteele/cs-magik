package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/satori/go.uuid"
	"os"
	"time"
)

func main() {
	serverID := flag.String("server-id", "", "server id of the teamserver")
	redisAddr := flag.String("redis-addr", "localhost", "address of redis")
	flag.Parse()

	id := uuid.Must(uuid.NewV4())
	encodedJob := fmt.Sprintf("%s|", base64.StdEncoding.EncodeToString([]byte(id.String())))
	for _, arg := range flag.Args() {
		fmt.Printf("DEBUG: encoding %s\n", arg)
		encodedJob += fmt.Sprintf("%s|",	base64.StdEncoding.EncodeToString([]byte(arg)))
	}

	fmt.Printf("DEBUG: Encoded job: %s\n", encodedJob)

	client := redis.NewClient(&redis.Options{
		Addr:     *redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if _, err := client.Ping().Result(); err != nil {
		fmt.Printf("Error pinging server: %s\n", err.Error())
		os.Exit(1)
	}

	if _, err := client.LPush(fmt.Sprintf("in:%s", *serverID), encodedJob).Result(); err != nil {
		fmt.Printf("Error during LPUSHX: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("DEBUG: polling for result\n")

	for {
		time.Sleep(1 * time.Second)
		res, err := client.Get(id.String()).Result()
		if err != nil {
			continue
		} else {
			fmt.Println("Result:")
			fmt.Println(res)
			break
		}
	}
}