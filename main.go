package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"discord-lg/commands"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	enverr := godotenv.Load()
	if enverr != nil {
		log.Fatal("Error loading .env file")
	}
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	dg.AddHandler(commands.MessageCreate)
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()
}
