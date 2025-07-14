package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Enemy struct {
	Name      string
	Health    int
	MaxHealth int
	Attack    int
	Defense   int
	ExpReward int
}

func NewGoblin() *Enemy {
	return &Enemy{
		Name:      "Goblin",
		Health:    30,
		MaxHealth: 30,
		Attack:    8,
		Defense:   2,
		ExpReward: 25,
	}
}

func NewOrc() *Enemy {
	return &Enemy{
		Name:      "Orc",
		Health:    60,
		MaxHealth: 60,
		Attack:    12,
		Defense:   4,
		ExpReward: 50,
	}
}

func NewTroll() *Enemy {
	return &Enemy{
		Name:      "Troll",
		Health:    120,
		MaxHealth: 120,
		Attack:    20,
		Defense:   8,
		ExpReward: 100,
	}
}

func NewDragon() *Enemy {
	return &Enemy{
		Name:      "Dragon",
		Health:    200,
		MaxHealth: 200,
		Attack:    35,
		Defense:   15,
		ExpReward: 250,
	}
}

func (e *Enemy) IsAlive() bool {
	return e.Health > 0
}

func (e *Enemy) TakeDamage(damage int) {
	e.Health -= damage
	if e.Health < 0 {
		e.Health = 0
	}
}

func (e *Enemy) AttackPlayer(player *Character) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	baseDamage := e.Attack + r.Intn(5)
	actualDamage := baseDamage - player.Defense
	if actualDamage < 1 {
		actualDamage = 1
	}
	player.TakeDamage(actualDamage)
	return actualDamage
}

func (e *Enemy) DisplayStats() {
	fmt.Printf("%s - Health: %d/%d, Attack: %d, Defense: %d\n", 
		e.Name, e.Health, e.MaxHealth, e.Attack, e.Defense)
}

func SpawnRandomEnemy(playerLevel int) *Enemy {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	var enemy *Enemy
	switch {
	case playerLevel <= 2:
		enemy = NewGoblin()
	case playerLevel <= 4:
		if r.Intn(2) == 0 {
			enemy = NewGoblin()
		} else {
			enemy = NewOrc()
		}
	case playerLevel <= 7:
		switch r.Intn(3) {
		case 0:
			enemy = NewOrc()
		case 1:
			enemy = NewTroll()
		default:
			enemy = NewGoblin()
		}
	default:
		switch r.Intn(4) {
		case 0:
			enemy = NewOrc()
		case 1:
			enemy = NewTroll()
		case 2:
			enemy = NewDragon()
		default:
			enemy = NewGoblin()
		}
	}
	
	return enemy
}