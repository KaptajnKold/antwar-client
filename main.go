package main

import (
	"github.com/KaptajnKold/antwar"
	"github.com/KaptajnKold/clever_ant"
	"github.com/KaptajnKold/ghengis"
	"github.com/KaptajnKold/naive_ant"
	"github.com/KaptajnKold/random_ant"
)

func main() {
	antwar.NewGame([]*antwar.Team{
		antwar.NewTeam("ghengis", ghengis.Spawn),
		antwar.NewTeam("ghengis2", ghengis.Spawn),
		antwar.NewTeam("random", random_ant.Spawn),
		antwar.NewTeam("naive", naive_ant.Spawn),
		antwar.NewTeam("clever", clever_ant.Spawn),
	})
}
