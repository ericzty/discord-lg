package commands

import (
	"regexp"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	isPrefix, _ := regexp.MatchString(`^[e!]{2}`, m.Content)
	if isPrefix == false {
		return
	}

	isPing, _ := regexp.MatchString(`^[e!ping]{6}`, m.Content)
	isMtr, _ := regexp.MatchString(`^[e!mtr]{5}`, m.Content)
	isDig, _ := regexp.MatchString(`^[e!dig]{5}`, m.Content)
	isBirdc, _ := regexp.MatchString(`^[e!birdc]{7}`, m.Content)

	switch {
	case isPing:
		ping(s, m)
	case isMtr:
		mtr(s, m)
	case isDig:
		dig(s, m)
	case isBirdc:
		birdc(s, m)
	default:
		s.ChannelMessageSend(m.ChannelID, "e!ping(4/6), e!mtr(4/6)")
	}

}
