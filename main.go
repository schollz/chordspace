package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"

	wr "github.com/mroth/weightedrand"
	"github.com/schollz/chordspace/src/chord"
	log "github.com/schollz/logger"
)

func main() {

	err := run()
	if err != nil {
		log.Error(err)
	}
}

func run() (err error) {
	log.SetLevel("debug")

	b, err := ioutil.ReadFile("data/chords.json")
	if err != nil {
		return
	}
	songs := make(map[string]SongData)
	err = json.Unmarshal(b, &songs)
	if err != nil {
		return
	}

	var transitions = make(map[string]map[string]uint)
	var lastChord chord.Chord
	for song := range songs {
		for _, field := range strings.Fields(songs[song].Harmony[0]) {
			if strings.HasPrefix(field, "[") || strings.HasPrefix(field, "|") || strings.HasPrefix(field, ".") {
				continue
			}
			field = strings.Split(field, "/")[0]
			var c chord.Chord
			c, err = chord.Parse(field)
			if err != nil {
				log.Debug(err)
				continue
			}
			if lastChord.Name != "" {
				if _, ok := transitions[lastChord.Name]; !ok {
					transitions[lastChord.Name] = make(map[string]uint)
				}
				if _, ok := transitions[lastChord.Name][c.Name]; !ok {
					transitions[lastChord.Name][c.Name] = 0
				}
				transitions[lastChord.Name][c.Name]++
			}
			lastChord = chord.Chord{
				Name: c.Name,
			}
		}
	}

	b, err = json.MarshalIndent(transitions, "", "   ")
	if err != nil {
		return
	}
	err = ioutil.WriteFile("transitions.json", b, 0644)

	rand.Seed(time.Now().UTC().UnixNano()) // always seed random!

	for i := 0; i < 30; i++ {
		currentChord := "Am"
		fmt.Print(currentChord + " ")
		for j := 0; j < 3; j++ {
			choices := []wr.Choice{}
			for k2 := range transitions[currentChord] {
				choices = append(choices, wr.Choice{Item: k2, Weight: transitions[currentChord][k2]})
			}
			choicer := wr.NewChooser(choices...)
			result := choicer.Pick()
			fmt.Print(result.(string) + " ")
			currentChord = result.(string)
		}
		fmt.Println(" ")
	}
	return
}

type SongData struct {
	Title   string   `json:"title"`
	Harmony []string `json:"harmony"`
}
