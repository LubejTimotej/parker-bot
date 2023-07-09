package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	discordToken := goDotEnvVariable("DISCORD_TOKEN")
	sess, err := discordgo.New(discordToken)
	if err != nil {
		log.Fatal(err)
	}

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "what time is it?" {
			msg := &discordgo.MessageSend{
				Content: "Pizza time!",
				Embed: &discordgo.MessageEmbed{
					Image: &discordgo.MessageEmbedImage{
						URL: "https://i.imgflip.com/52kdxc.png?a469080",
					},
				},
			}
			_, err := s.ChannelMessageSendComplex(m.ChannelID, msg)
			if err != nil {
				log.Println("Failed to send message with image:", err)
			}
		}
	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	fmt.Println("the bot is online!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
