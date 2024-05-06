package main

import "fmt"

type Player struct {
	health int
}

func playerDamageByExplosion(player *Player) {
	player.health -= 10
}

// a function not a method
func (player *Player) takeDamageByExplosion(dmg int) {
	player.health -= dmg
}

func main() {
	player := &Player{
		health: 100,
	}
	fmt.Printf("Player health before explosion %+v\n", player)
	//player = nil
	//playerDamageByExplosion(player)
	player.takeDamageByExplosion(50)
	playerDamageByExplosion(player)
	fmt.Printf("Player health after explosion %+v\n", player)

}
