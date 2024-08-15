package main

import (
	"fmt"
	_ "net/http/pprof"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
)

func addBot(StopChan chan struct{}, UUID string, token string, wg *sync.WaitGroup) {
	defer wg.Done()
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error initializing bot with Token= ", token, "\n", err)
		return
	}
	discord.StateEnabled = false

	discord.AddHandler(messageCreate)
	discord.AddHandler(func(Session *discordgo.Session, Interaction *discordgo.InteractionCreate) {
		handleCommand(Session, Interaction)
	})

	//discord.Identify.Intents = discordgo.IntentsGuildMessages
	//discord.Identify.Intents |= discordgo.IntentMessageContent
	err = discord.Open()
	if err != nil {
		fmt.Print("error opening websocket with token = ", token, "\n", err)
		close(StopChan)
		return
	}
	fmt.Println("started bot")

	//discord.Close()

	for {
		select {
		case <-StopChan:
			fmt.Println("stopping bot with uuid", UUID)
			discord.Close()
			return
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}

}

func messageCreate(Session *discordgo.Session, Message *discordgo.MessageCreate) {
	fmt.Println("received message: ", Message.Content)

	if Message.Author.ID == Session.State.User.ID {
		return
	}
	if Message.Content == "ping" {
		fmt.Println("Received ping")
		Session.ChannelMessageSend(Message.ChannelID, "Pong!")
	}

	if Message.Content == "pong" {
		Session.ChannelMessageSend(Message.ChannelID, "Ping!")
	}
}

func handleCommand(Session *discordgo.Session, Interaction *discordgo.InteractionCreate) {
	Session.InteractionRespond(Interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Hey there! Congratulations, you just executed your first slash command",
		},
	})
}
