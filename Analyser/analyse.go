package Analyser

import (
	"fmt"

	"github.com/icza/s2prot/rep"
)

func Analyse() error {
	pathName := "../../../../Library/Application Support/Blizzard/StarCraft II/Accounts/620084517/5-S2-1-11471102/Replays/Multiplayer/大气2000-天梯版 (271).SC2Replay"

	r, err := rep.NewFromFile(pathName)
	if err != nil {
		return fmt.Errorf("error reading from file, error is %w", err)
	}
	defer r.Close()
	fmt.Println("Players:")
	for _, p := range r.Details.Players() {
		fmt.Printf("\tName: %-20s, Race: %c, Result: %v\n",
			p.Name, p.Race().Letter, p.Result())
	}
	return nil
}
