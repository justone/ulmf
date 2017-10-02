package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/justone/simpleslack"
)

func main() {
	token := flag.String("t", os.Getenv("SLACK_TOKEN"), "slack token")
	undo := flag.Bool("u", false, "undo the away")
	minutes := flag.String("d", "120", "how long to DND")
	emoji := flag.String("e", ":computer:", "emoji to set for status")
	message := flag.String("m", "Busy, head's down.", "message to set for status")

	flag.Parse()
	sc := simpleslack.Client{*token}

	if *undo {
		errCheck(setStatus(sc, "", ""))
		errCheck(setPresence(sc, "auto"))
		errCheck(clearDND(sc))
	} else {
		errCheck(setStatus(sc, *emoji, *message))
		errCheck(setPresence(sc, "away"))
		errCheck(setDND(sc, *minutes))
	}
}

func errCheck(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func setStatus(sc simpleslack.Client, emoji, message string) error {
	encodedStatus, _ := json.Marshal(map[string]string{
		"status_emoji": emoji,
		"status_text":  message,
	})

	_, err := sc.Post("users.profile.set", url.Values{"profile": {string(encodedStatus)}})
	return err
}

func setPresence(sc simpleslack.Client, presence string) error {
	_, err := sc.Post("users.setPresence", url.Values{"presence": {presence}})
	return err
}

func setDND(sc simpleslack.Client, minutes string) error {
	_, err := sc.Post("dnd.setSnooze", url.Values{"num_minutes": {minutes}})
	return err
}

func clearDND(sc simpleslack.Client) error {
	_, err := sc.Post("dnd.endSnooze", url.Values{})
	return err
}
