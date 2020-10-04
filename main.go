package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/caarlos0/env"
	"github.com/go-redis/redis"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var db *redis.Client
var cfg config
var dg *discordgo.Session

func main() {
	if err := env.Parse(&cfg); err != nil {
		log.Fatalln("Config", err)
	}

	var err error

	db = redis.NewClient(&redis.Options{
		Addr: cfg.DBAddr,
	})

	dg, err = discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		log.Fatalln("Discord", err)
	}
	dg.AddHandler(handleMessage)
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAllWithoutPrivileged)

	err = dg.Open()
	if err != nil {
		log.Fatalln("Discord", err)
	}

	log.Println("Started")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return //handle only incoming messages
	}
	if m.GuildID != "" {
		return //handle only private messages
	}
	if m.Type != discordgo.MessageTypeDefault {
		return //handle only normal messages
	}
	if m.Content == "" {
		return //handle only text
	}
	handlePrivateText(s, m)
}
