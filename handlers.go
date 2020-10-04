package main

import (
	"github.com/bwmarrin/discordgo"
	"net/mail"
	"regexp"
	"strconv"
	"strings"
)

var nonNumericRx = regexp.MustCompile("[ ,.]+")

func handlePrivateText(s *discordgo.Session, m *discordgo.MessageCreate) {
	if getBackerIDByDiscordID(m.Author.ID) != 0 {
		s.ChannelMessageSend(m.ChannelID, "You've already confirmed your backer status.")
		return
	}

	_, err := s.GuildMember(cfg.GuildID, m.Author.ID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "You haven't joined our Discord server.\n"+
			"Join it first and then write me again.")
		return
	}

	id, email := extractBackerInfo(m.Content)
	if id == 0 || email == "" {
		s.ChannelMessageSend(m.ChannelID, "I can confirm your backer status and give you access to private backer-only channels.\n\n"+
			"Please send me your **backer number** and **Kickstarter email address** separated by whitespace, like this:\n"+
			"`1337 anna96@gmail.com`\n\n"+
			"Can't find your backer number? Press blue 'View pledge' button on the campaign page.")
		return
	}

	handleBackerInfo(s, m, id, email)
}

func handleBackerInfo(s *discordgo.Session, m *discordgo.MessageCreate, id int, email string) {
	foundEmail := getEmailByBackerID(id)
	if email != foundEmail {
		s.ChannelMessageSend(m.ChannelID, "I'm sorry, but I can't find this pair of backer number and email address.\n\n"+
			"Make sure you typed in the email you set in your Kickstarter profile.")
		return
	}

	s.GuildMemberRoleAdd(cfg.GuildID, m.Author.ID, cfg.RoleID)
	linkBackerIDAndDiscordID(id, m.Author.ID)
	s.ChannelMessageSend(m.ChannelID, "Yay! I've found you in our backers list and your backer status has been confirmed.")
}

func extractBackerInfo(text string) (id int, email string) {
	parts := strings.Split(text, " ")
	if len(parts) != 2 {
		return 0, ""
	}

	id, _ = strconv.Atoi(nonNumericRx.ReplaceAllString(parts[0], ""))
	if id <= 0 {
		return 0, ""
	}

	_, err := mail.ParseAddress(parts[1])
	if err != nil {
		return 0, ""
	}

	email = strings.ToLower(parts[1])
	return
}
