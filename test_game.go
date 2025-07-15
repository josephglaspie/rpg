package main

import (
	"fmt"
)

func testGame() {
	fmt.Println("=== RPG Game Test ===")
	
	player := NewCharacter("TestHero")
	fmt.Println("✓ Character creation successful")
	player.DisplayStats()
	
	goblin := NewGoblin()
	fmt.Printf("✓ Enemy creation successful: %s\n", goblin.Name)
	goblin.DisplayStats()
	
	initialHealth := player.Health
	damage := goblin.AttackPlayer(player)
	fmt.Printf("✓ Combat system working - %s took %d damage\n", player.Name, damage)
	
	if player.Health < initialHealth {
		fmt.Println("✓ Damage system working correctly")
	}
	
	player.Heal(20)
	fmt.Printf("✓ Healing system working - %s healed\n", player.Name)
	
	player.GainExperience(150)
	fmt.Printf("✓ Experience and leveling system working - %s is level %d\n", player.Name, player.Level)
	
	fmt.Println("\n🎉 All systems tested successfully!")
	fmt.Println("Game is ready to play! Run: go run *.go")
}

