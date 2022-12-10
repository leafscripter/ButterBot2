package bot

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

// Bot parameters
var (
	Token string
	GuildID string
)

var session *discordgo.Session

// We wanna parse those flags and create a new session before anything else runs
func init() {

	flag.StringVar(&Token, "token", "", "Bot access token")
	flag.StringVar(&GuildID, "guild_id", "", "Server ID")

	flag.Parse()
}

func init() {
	var err error
	session, err = discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatalf("Could not create session: %v", err)
	}
}

// Commands and their handlers
var (
	commands = []*discordgo.ApplicationCommand{episodes, characters}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		"episodes": episodesHandler,
		"characters": charactersHandler,
		"help": helpHandler,
	}
)

// Register all the handlers we've created
func init() {
	session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s,i)
		}
	})
}

func handleLogin(s *discordgo.Session, r *discordgo.Ready) {
	log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
}

func Run() {
	
	session.AddHandler(handleLogin)

	// Open a new session and handle any errors
	err := session.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	// Register all commands
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i,v := range commands {
		cmd, err := session.ApplicationCommandCreate(session.State.User.ID, GuildID, v) 
		// Check for errors
		if err != nil { 
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}

		registeredCommands[i] = cmd
	}

	defer session.Close()

	// Wait for interruption signal
	stop_signal := make(chan os.Signal, 1)
	signal.Notify(stop_signal, os.Interrupt)
	<-stop_signal	
	
	// Delete all commands before shutdown
	for _,v := range registeredCommands { 
		err := session.ApplicationCommandDelete(session.State.User.ID, GuildID, v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}

	log.Println("Shutting down")
}