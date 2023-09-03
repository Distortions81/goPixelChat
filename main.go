package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Adeithe/go-twitch"
	"github.com/Adeithe/go-twitch/irc"
)

func main() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)

	writer := &irc.Conn{}
	auth, err := os.ReadFile("auth.txt")
	if err != nil {
		log.Fatal(err)
	}
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

	<-sc
	fmt.Println("Stopping...")
	reader.Close()
	writer.Close()
}

func onShardReconnect(shardID int) {
	fmt.Printf("Shard #%d reconnected\n", shardID)
}

func onShardLatencyUpdate(shardID int, latency time.Duration) {
	fmt.Printf("Shard #%d has %dms ping\n", shardID, latency.Milliseconds())
}

func onShardMessage(shardID int, msg irc.ChatMessage) {
	fmt.Printf("#%s %s: %s\n", msg.Channel, msg.Sender.DisplayName, msg.Text)
}
