package main

import (
	"github.com/KaptajnKold/antwar"
	"github.com/KaptajnKold/ghengis"
	"github.com/KaptajnKold/naive_ant"
)

func main() {
	antwar.NewGame([]*antwar.Team{
		antwar.NewTeam("ghengis", ghengis.Spawn),
		antwar.NewTeam("ghengis2", ghengis.Spawn),
		antwar.NewTeam("naive", naive_ant.Spawn),
	})
}
