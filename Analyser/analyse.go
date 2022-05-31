package Analyser

import (
	"fmt"
	"strings"

	"github.com/SC2Analyser/Player"
	"github.com/icza/s2prot/rep"
)

type repPlayer struct {
	name   string
	race   string
	result string
}

func Analyse(pathName string, p0, p1 *Player.Player) error {

	r, err := rep.NewFromFile(pathName)
	if err != nil {
		fmt.Printf("error reading from file %v, error is %v\n", pathName, err.Error())
		return nil
	}
	defer r.Close()
	players := r.Details.Players()
	if len(players) != 2 {
		return nil
	}
	repP0 := repPlayer{
		name:   players[0].Name,
		race:   string(players[0].Race().Letter),
		result: string(players[0].Result().Letter),
	}
	repP1 := repPlayer{
		name:   players[1].Name,
		race:   string(players[1].Race().Letter),
		result: string(players[1].Result().Letter),
	}
	if match(repP0, p0) && match(repP1, p1) {
		if repP0.result == "V" {
			p0.Winning += 1
		}
		if repP1.result == "V" {
			p1.Winning += 1
		}
	} else if match(repP0, p1) && match(repP1, p0) {
		if repP0.result == "V" {
			p1.Winning += 1
		}
		if repP1.result == "V" {
			p0.Winning += 1
		}
	}

	return nil
}

func match(repP repPlayer, player *Player.Player) bool {
	ok1, ok2 := false, false
	for _, n := range player.Names {
		if strings.Contains(repP.name, n) {
			ok1 = true
			break
		}
	}
	for _, n := range player.Race {
		if n == repP.race {
			ok2 = true
			break
		}
	}
	return ok1 && ok2
}
