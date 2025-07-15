package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Check if running in server mode
	if len(os.Args) > 1 && os.Args[1] == "server" {
		startWebServer()
		return
	}

	// Check if running tests
	if len(os.Args) > 1 && os.Args[1] == "test" {
		testGame()
		return
	}

	// Default: run console version
	fmt.Println("🗡️  Welcome to the RPG Adventure! 🗡️")
	fmt.Println("=====================================")
	
	player := createPlayer()
	gameLoop(player)
}

func startWebServer() {
	fmt.Println("🌐 Starting RPG Web Server...")
	fmt.Println("🎮 Game available at: http://localhost:8081")
	fmt.Println("Press Ctrl+C to stop the server")
	
	// Serve static files (HTML, CSS, JS)
	http.Handle("/", http.FileServer(http.Dir("./")))
	
	// Start server on port 8080
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func createPlayer() *Character {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Print("Enter your character's name: ")
	scanner.Scan()
	name := scanner.Text()
	
	if name == "" {
		name = "Hero"
	}
	
	player := NewCharacter(name)
	fmt.Printf("\nWelcome, %s! Your adventure begins now.\n", player.Name)
	player.DisplayStats()
	
	return player
}

func gameLoop(player *Character) {
	scanner := bufio.NewScanner(os.Stdin)
	
	for player.IsAlive() {
		fmt.Println("\n" + strings.Repeat("=", 50))
		fmt.Println("What would you like to do?")
		fmt.Println("1. Fight a monster")
		fmt.Println("2. Rest and heal")
		fmt.Println("3. View stats")
		fmt.Println("4. Quit game")
		fmt.Print("Choose an option (1-4): ")
		
		if !scanner.Scan() {
			fmt.Println("\nInput error or EOF. Exiting game.")
			break
		}
		choice := strings.TrimSpace(scanner.Text())
		
		switch choice {
		case "1":
			fightMonster(player)
		case "2":
			restAndHeal(player)
		case "3":
			player.DisplayStats()
		case "4":
			fmt.Println("Thanks for playing! Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please enter 1, 2, 3, or 4.")
		}
		
		if !player.IsAlive() {
			fmt.Printf("\n💀 %s has fallen in battle! Game Over! 💀\n", player.Name)
			fmt.Printf("Final Level: %d\n", player.Level)
			break
		}
	}
}

func fightMonster(player *Character) {
	enemy := SpawnRandomEnemy(player.Level)
	fmt.Printf("\n⚔️  A wild %s appears!\n", enemy.Name)
	enemy.DisplayStats()
	
	for player.IsAlive() && enemy.IsAlive() {
		fmt.Println("\nCombat Options:")
		fmt.Println("1. Attack")
		fmt.Println("2. Run away")
		fmt.Print("Choose your action: ")
		
		scanner := bufio.NewScanner(os.Stdin)
		if !scanner.Scan() {
			fmt.Println("\nInput error. Running away from battle!")
			return
		}
		choice := strings.TrimSpace(scanner.Text())
		
		switch choice {
		case "1":
			damage := player.AttackEnemy(enemy)
			fmt.Printf("You attack the %s for %d damage!\n", enemy.Name, damage)
			
			if enemy.IsAlive() {
				damage = enemy.AttackPlayer(player)
				fmt.Printf("The %s attacks you for %d damage!\n", enemy.Name, damage)
				fmt.Printf("Your health: %d/%d\n", player.Health, player.MaxHealth)
			} else {
				fmt.Printf("🎉 You defeated the %s!\n", enemy.Name)
				fmt.Printf("You gained %d experience points!\n", enemy.ExpReward)
				player.GainExperience(enemy.ExpReward)
				return
			}
		case "2":
			fmt.Println("You successfully run away from the battle!")
			return
		default:
			fmt.Println("Invalid choice. Please enter 1 or 2.")
		}
	}
}

func restAndHeal(player *Character) {
	healAmount := player.MaxHealth / 3
	if healAmount < 10 {
		healAmount = 10
	}
	
	player.Heal(healAmount)
	fmt.Printf("💚 You rest and recover %d health points.\n", healAmount)
	fmt.Printf("Current health: %d/%d\n", player.Health, player.MaxHealth)
}