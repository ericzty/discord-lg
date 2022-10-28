package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func makeResultEmbed(err error, out []byte, command, user, timestamp string) *discordgo.MessageEmbed {
	color := 2227217
	result := "```\nOutput: " + string(out) + "\n```"
	if err != nil {
		color = 16723502
		result = fmt.Sprintf("Output: "+string(out)+"\n Error: %v", err)
	}
	return &discordgo.MessageEmbed{
		Title:       "EricNet " + command,
		Color:       color,
		Description: result,
		Timestamp:   timestamp,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Requested by " + user,
		},
	}
}
