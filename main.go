package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	channel, err := os.ReadFile("channel.txt")
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf(
		"https://api.streamelements.com/kappa/v2/songrequest/%s/playing",
		strings.TrimSpace(string(channel)),
	)

	file, err := os.Create("now_playing.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	spacer := strings.Repeat(" ", 10)

	delay := 5 * time.Second
	ticker := time.NewTicker(delay)
	defer ticker.Stop()

	for range ticker.C {
		response, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer response.Body.Close()

		if response.StatusCode == http.StatusOK {
			body, err := io.ReadAll(response.Body)
			if err != nil {
				fmt.Println(err)
				continue
			}

			var data map[string]any
			err = json.Unmarshal(body, &data)
			if err != nil {
				fmt.Println(err)
				continue
			}

			title := data["title"].(string)
			file.WriteString(title + spacer)
			fmt.Printf("Now playing: %s\n", title)
		} else {
			file.WriteString("")
			fmt.Printf("Error: %s\n", response.Status)
		}
	}
}
