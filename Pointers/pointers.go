// POINTERS

package main

import "fmt"

type BigData struct {
	// 500 MB
}

// 8 bytes => pointer (64bit machines)
func processBigData(bd *BigData) {

}

type Player struct {
	health int
}

func (player *Player) takeDamageFromExplosion(dmg int) {
	fmt.Println("PLAYER IS TAKING DAMAGE FROM EXPLOSION")
	player.health -= dmg
}

func takeDamageFromExplosion(player Player, dmg int) {
	fmt.Println("PLAYER IS TAKING DAMAGE FROM EXPLOSION")
	player.health -= dmg
}

func main() {
	player := &Player{
		health: 100,
	}
	// 8 byte long integer pointer

	fmt.Printf("BEFORE EXPLOSION %+v\n", player)
	player.takeDamageFromExplosion((50))
	fmt.Printf("AFTER EXPLOSION %+v\n", player)

}
