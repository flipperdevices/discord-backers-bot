package main

import (
	"context"
	"fmt"
	"strings"
)

var ctx = context.TODO()

func getEmailByBackerID(backerID int) string {
	return strings.ToLower(db.Get(ctx, formatEmailByBackerIDKey(backerID)).Val())
}

func getBackerIDByDiscordID(discordID string) int {
	id, _ := db.Get(ctx, formatBackerIDByDiscordIDKey(discordID)).Int()
	return id
}

func getDiscordIDByBackerID(backerID int) string {
	return db.Get(ctx, formatDiscordIDByBackerIDKey(backerID)).Val()
}

func linkBackerIDAndDiscordID(backerID int, discordID string) {
	db.Set(ctx, formatBackerIDByDiscordIDKey(discordID), backerID, 0)
	db.Set(ctx, formatDiscordIDByBackerIDKey(backerID), discordID, 0)
}

func formatEmailByBackerIDKey(backerID int) string {
	return fmt.Sprintf("backerEmail.%d", backerID)
}

func formatDiscordIDByBackerIDKey(backerID int) string {
	return fmt.Sprintf("backerDiscord.%d", backerID)
}

func formatBackerIDByDiscordIDKey(discordID string) string {
	return fmt.Sprintf("backerID.%s", discordID)
}
