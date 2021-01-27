package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"

    client "github.com/bwmarrin/discordgo"
)

func main() {
    session, err := client.New(
        "Bot " +
            "Token")

    fmt.Print("If you recieve this you're good")
    if err = session.Open(); err != nil {
        fmt.Println(err)
        return
    }

    discord := make(chan os.Signal, 1)
    signal.Notify(discord, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSEGV, syscall.SIGHUP)
    <-discord
}
