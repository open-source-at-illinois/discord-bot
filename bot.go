package main

import (
    "fmt"
    "github.com/bwmarrin/discordgo"
    "os"
    "os/signal"
    "strings"
    "syscall"
    "math/rand"
    crypto_rand "crypto/rand"
    "encoding/base64"
)


func GenerateToken() string {
    b := make([]byte, 64)
    _, err := crypto_rand.Read(b)
    if err != nil {
        fmt.Println("Couldn't generate a token!")
        os.Exit(1)
        return "NO TOKEN"
    }
    return base64.URLEncoding.EncodeToString(b)
}


func main() {
    TOKEN := os.Getenv("TOKEN")
    dg, err := discordgo.New("Bot " + TOKEN)
    if err != nil {
        fmt.Println("Something bad happened when creating the discord session!", err)
        return
    }
    dg.AddHandler(messageCreate)

    err = dg.Open()
    if err != nil {
        fmt.Println("Something spoopy happened when opening the connection! ", err)
        return
    }
    fmt.Println("Bot is running. CTRL-C to exit.")
    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <-sc

    dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
    if m.Author.ID == s.State.User.ID {
        return
    }

    if strings.Contains(strings.ToLower(m.Content), "token") {
        tokenMessages := []string{
            "Did someone say token? I got one!",
            "A token is on it's way!",
            "Want a token? Come and get it.",
            "toe can",
            "toe can",
            "Sooooo mannnnyyyyyy tokennnns",
            "tOkEn",
            "Generating cryptographically secure token in\n\n3\n\n2\n\n1",
            "Tokin'... heh",
            "Listen, I'm running out of token-related one-liners.",
            "Just take it.",
            "Decrypt this, Yml0Y2gK!",
            "T25lIG9mIHRoZXNlIGlzIG5vdCBsaWtlIHRoZSBvdGhlci4K",
            "Web safe.",
            "Use this the next time you log in:",
            "Multi.\nLine.\nString.",
            "Does discord compile?\n// Hmm, I guess not",
            "Does discord compile?\n# Hmm, I guess not",
            "Can't you just generate your own token?",
            "How many bits does it take to get to the center of discord bot?\n\n||It's 512.||",
            "加密安全令牌，現在為中文", // Cryptographically secure tokens, now in Chinese!

        }

        i := rand.Intn(len(tokenMessages))

        token := GenerateToken()

        fmt.Println("Sending a token: " + token)
        s.ChannelMessageSend(m.ChannelID, tokenMessages[i] + "\n\n`" + token + "`")
    }
}
