package main

type config struct {
	DBAddr string `env:"DB_ADDR" envDefault:"localhost:6379"`

	DiscordToken string `env:"DISCORD_TOKEN,required"`
	GuildID      string `env:"GUILD_ID,required"`
	RoleID       string `env:"ROLE_ID,required"`
}
