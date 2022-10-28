package commands

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func pingCmd(m string) (*exec.Cmd, error) {
	args := strings.Split(m, " ")[0:]
	var cmd *exec.Cmd
	var err error
	if len(args) != 2 {
		err := errors.New("badArgsLen")
		return cmd, err
	}
	host := args[1]
	switch args[0] {
	case "e!ping":
		cmd := exec.Command("ping", "-c3", host)
		return cmd, err
	case "e!ping4":
		cmd := exec.Command("ping", "-c3", "-4", host)
		return cmd, err
	case "e!ping6":
		cmd := exec.Command("ping", "-c3", "-6", host)
		return cmd, err
	default:
		err := errors.New("badCmd")
		return cmd, err
	}
}

func ping(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := m.Content

	cmd, err := pingCmd(message)

	if err != nil {
		out := fmt.Sprintf("Error: %v", err)
		s.ChannelMessageSend(m.ChannelID, out)
		return
	}

	timestamp := m.Timestamp.Format(time.RFC3339)
	command := strings.Join(cmd.Args, " ")
	user := fmt.Sprintf(m.Author.Username + "#" + m.Author.Discriminator)

	tool := "Ping"

	fmt.Println(user + " did " + command)

	s.ChannelMessageSend(m.ChannelID, "Running: "+strings.Join(cmd.Args, " "))
	out, err := cmd.CombinedOutput()

	resultEmbed := makeResultEmbed(err, out, tool, user, timestamp)
	_, embedErr := s.ChannelMessageSendEmbed(m.ChannelID, resultEmbed)
	if embedErr != nil {
		fmt.Println(embedErr)
	}
}
