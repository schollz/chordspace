package chord

import (
	"fmt"
	"strings"
)

type Chord struct {
	Name  string
	Notes []string
}

var noteScale = map[string]string{
	"R":    "C",
	"I":    "C",
	"bII":  "Db",
	"II":   "D",
	"bIII": "Eb",
	"III":  "E",
	"bIV":  "E",
	"IV":   "F",
	"bV":   "Gb",
	"V":    "G",
	"bVI":  "Ab",
	"VI":   "A",
	"bVII": "Bb",
	"VII":  "B",
	"i":    "Cm",
	"bii":  "Dbm",
	"ii":   "Dm",
	"biii": "Ebm",
	"iii":  "Em",
	"biv":  "Em",
	"iv":   "Fm",
	"bv":   "Gbm",
	"v":    "Gm",
	"bvi":  "Abm",
	"vi":   "Am",
	"bvii": "Bbm",
	"vii":  "Bm",
}

var suffixSymbol = map[string]string{
	"64": "2nd inv",
	"65": "1st inv",
	"42": "3rd inv",
}

func Parse(symbol string) (chord Chord, err error) {
	longest := 0
	for n := range noteScale {
		if len(n) > longest && strings.HasPrefix(symbol, n) {
			longest = len(n)
			chord.Name = noteScale[n]
			chord.Name += strings.TrimPrefix(symbol, n)
		}
	}
	if chord.Name == "" {
		err = fmt.Errorf("could not parse '%s'", symbol)
		return
	}
	chord.Name = strings.Split(chord.Name, "65")[0]
	chord.Name = strings.Split(chord.Name, "43")[0]
	chord.Name = strings.Split(chord.Name, "42")[0]
	chord.Name = strings.Split(chord.Name, "11")[0]
	chord.Name = strings.Split(chord.Name, "64")[0]

	return
}
