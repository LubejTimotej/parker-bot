﻿# Parker-bot
This project is a simple Discord bot that responds to a specific message with a predefined image. It utilizes the DiscordGo library and interacts with the Discord API.

### Getting Started:

To run this bot, follow the steps below:
1. Clone the repository or download the source code.
2. Make sure you have Go installed on your system.
3. Create a '.env' file in the project's root directory and add the following line:

        DISCORD_TOKEN=your_discord_token

replace 'your_discord_token' with your actual Discord bot token.
4. Install the required dependencies by executing the following command in the projects's root directory:

        go get -u github.com/bwmarrin/discordgo
        go get -u github.com/joho/godotenv

5. build and run the bot using the following command:

        go run main.go

6. The bot should now be online and ready to respond to messages.

### Usage:

Once the bot is online and connected to your Discord server, it will respond with a predefined image whenever it receives a message with the content "what time is it?". The response message will contain an embedded image.

### Customization:

to customize the bot's response, you can modify the following code block in the 'main' function:

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

You can change the trigger message from "what time is it?" to any other phrase, and modify the response content and image URL accordingly.

### Dependencies:

This project relies on the following third-party libraries:

    github.com/bwmarrin/discordgo: DiscordGo is a Go library for interacting with the Discord API.
    github.com/joho/godotenv: GoDotEnv is a Go library for loading environment variables from a '.env' file.
