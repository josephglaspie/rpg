package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Character struct {
	Name      string
	Level     int
	Health    int
	MaxHealth int
	Attack    int
	Defense   int
	Experience int
	ExperienceToNext int
}

func NewCharacter(name string) *Character {
	return &Character{
		Name:      name,
		Level:     1,
		Health:    100,
		MaxHealth: 100,
		Attack:    15,
		Defense:   5,
		Experience: 0,
		ExperienceToNext: 100,
	}
}

func (c *Character) IsAlive() bool {
	return c.Health > 0
}

func (c *Character) TakeDamage(damage int) {
	actualDamage := damage - c.Defense
	if actualDamage < 1 {
		actualDamage = 1
	}
	c.Health -= actualDamage
	if c.Health < 0 {
		c.Health = 0
	}
}

func (c *Character) Heal(amount int) {
	c.Health += amount
	if c.Health > c.MaxHealth {
		c.Health = c.MaxHealth
	}
}

func (c *Character) AttackEnemy(enemy *Enemy) int {
	rand.Seed(time.Now().UnixNano())
	baseDamage := c.Attack + rand.Intn(10)
	actualDamage := baseDamage - enemy.Defense
	if actualDamage < 1 {
		actualDamage = 1
	}
	enemy.TakeDamage(actualDamage)
	return actualDamage
}

func (c *Character) GainExperience(exp int) {
	c.Experience += exp
	if c.Experience >= c.ExperienceToNext {
		c.LevelUp()
	}
}

func (c *Character) LevelUp() {
	c.Level++
	c.Experience -= c.ExperienceToNext
	c.ExperienceToNext = c.Level * 100
	
	healthIncrease := 20 + rand.Intn(10)
	attackIncrease := 3 + rand.Intn(3)
	defenseIncrease := 2 + rand.Intn(2)
	
	c.MaxHealth += healthIncrease
	c.Health = c.MaxHealth
	c.Attack += attackIncrease
	c.Defense += defenseIncrease
	
	fmt.Printf("\n🎉 Level Up! %s is now level %d!\n", c.Name, c.Level)
	fmt.Printf("Health: +%d (now %d)\n", healthIncrease, c.MaxHealth)
	fmt.Printf("Attack: +%d (now %d)\n", attackIncrease, c.Attack)
	fmt.Printf("Defense: +%d (now %d)\n", defenseIncrease, c.Defense)
}

func (c *Character) DisplayStats() {
	fmt.Printf("\n=== %s (Level %d) ===\n", c.Name, c.Level)
	fmt.Printf("Health: %d/%d\n", c.Health, c.MaxHealth)
	fmt.Printf("Attack: %d\n", c.Attack)
	fmt.Printf("Defense: %d\n", c.Defense)
	fmt.Printf("Experience: %d/%d\n", c.Experience, c.ExperienceToNext)
}