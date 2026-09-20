package main

import (
	"fmt"
	"math/rand"
)

type FootballPlayer struct {
	stamina int
	power   int
}

type CR7 struct {
	stamina int
	power   int
	sui     int
}

type MESSI struct {
	stamina int
	power   int
	sui     int
}

func (f CR7) kickBall() {
	shot := f.stamina + f.power*f.sui
	fmt.Println("CR7 is kicking the ball", shot)
}

func (f MESSI) kickBall() {
	shot := f.stamina + f.power*f.sui
	fmt.Println("MESSI is kicking the ball", shot)
}

type Player interface {
	kickBall()
}

func (f FootballPlayer) kickBall() {
	shot := f.stamina + f.power
	fmt.Println("I'm kicking the ball", shot)
}

func main() {
	team := make([]Player, 11)
	for i := 0; i < len(team)-2; i++ {
		team[i] = FootballPlayer{
			stamina: rand.Intn(10),
			power:   rand.Intn(10),
		}
	}
	team[len(team)-2] = CR7{
		stamina: 10,
		power:   10,
		sui:     10,
	}
	team[len(team)-1] = MESSI{
		stamina: 10,
		power:   10,
		sui:     5,
	}
	for i := 0; i < len(team); i++ {
		team[i].kickBall()
	}
}
