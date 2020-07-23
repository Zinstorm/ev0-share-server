package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var session *discordgo.Session

type Bot interface {
}

type DiscordSession struct {
	Session *discordgo.Session
}

func Init() (*DiscordSession, error) {
	var dg *discordgo.Session
	var err error
	if viper.GetBool("discordbot.useToken") {
		dg, err = discordgo.New("Bot " + viper.GetString("discordbot.token"))
	} else {
		// Create a new Discord session using email+password.
		dg, err = discordgo.New(viper.GetString("discordbot.email"), viper.GetString("discordbot.password"))
	}

	if err != nil {
		log.Error("error creating Discord session,", err)
		return nil, err
	}
	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)
	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Error("error opening connection,", err)
		return nil, err
	}
	return &DiscordSession{
		Session: dg,
	}, nil
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	isDm, _ := ComesFromDM(s, m)

	if isDm {
		log.Info(fmt.Sprintf("%s %s messaged the bot, content: %s", m.Author.Username, m.Author.ID, m.Content))
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s, fuck off", m.Author.Username))
	}
}

// ComesFromDM returns true if a message comes from a DM channel
func ComesFromDM(s *discordgo.Session, m *discordgo.MessageCreate) (bool, error) {
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		if channel, err = s.Channel(m.ChannelID); err != nil {
			return false, err
		}
	}

	return channel.Type == discordgo.ChannelTypeDM, nil
}
