package commands

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func birdCmd(m string) (*exec.Cmd, error) {
	args := strings.Split(m, " ")[0:]
	var cmd *exec.Cmd
	var err error
	if len(args) != 2 {
		err := errors.New("badArgsLen")
		return cmd, err
	} else {
		cmd := exec.Command("birdc", "show", "route", "for", args[1], "all")
		return cmd, err
	}
}

func birdc(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := m.Content

	cmd, err := birdCmd(message)

	if err != nil {
		out := fmt.Sprintf("Error: %v", err)
		s.ChannelMessageSend(m.ChannelID, out)
		return
	}

	timestamp := m.Timestamp.String()
	command := strings.Join(cmd.Args, " ")
	user := fmt.Sprintf(m.Author.Username + "#" + m.Author.Discriminator)

	tool := "Birdc"

	fmt.Println(user + " did " + command)

	s.ChannelMessageSend(m.ChannelID, "Running: "+strings.Join(cmd.Args, " "))
	out, err := cmd.CombinedOutput()

	resultEmbed := makeResultEmbed(err, out, tool, user, timestamp)
	s.ChannelMessageSendEmbed(m.ChannelID, resultEmbed)
}
