package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Adeithe/go-twitch"
	"github.com/Adeithe/go-twitch/irc"
)

func onShardMessage(shardID int, msg irc.ChatMessage) {
	fmt.Printf("#%s %s: %s\n", msg.Channel, msg.Sender.DisplayName, msg.Text)

	args := strings.Split(msg.Text, ",")
	if len(args) == 3 {
		//
	}
}

func connectTwitch() {
	writer := &irc.Conn{}

	//Get aiuth
	auth, err := os.ReadFile("auth.txt")
	if err != nil {
		log.Fatal(err)
	}

	//Connect

	writer.SetLogin("xboxtv81", "oauth:"+string(auth))
	if err := writer.Connect(); err != nil {
		panic("failed to start writer")
	}

	reader := twitch.IRC()
	reader.OnShardReconnect(onShardReconnect)
	reader.OnShardLatencyUpdate(onShardLatencyUpdate)
	reader.OnShardMessage(onShardMessage)

	if err := reader.Join("xboxtv81"); err != nil {
		panic(err)
	}
	fmt.Println("Connected to IRC!")
}

func onShardReconnect(shardID int) {
	fmt.Printf("Shard #%d reconnected\n", shardID)
}

func onShardLatencyUpdate(shardID int, latency time.Duration) {
	fmt.Printf("Shard #%d has %dms ping\n", shardID, latency.Milliseconds())
}
