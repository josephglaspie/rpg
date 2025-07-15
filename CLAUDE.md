# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a simple command-line RPG game written in Go. The game features a turn-based combat system with character progression, multiple enemy types, and basic RPG mechanics.

## Architecture

The game is structured into three main components:

- `character.go` - Player character system with stats, leveling, combat abilities, and experience gain
- `enemy.go` - Enemy types (Goblin, Orc, Troll, Dragon) with scaling difficulty and random spawning
- `main.go` - Game loop, user interface, and main game logic
- `test_game.go` - Test functions to validate game systems

## Key Features

- Character creation with customizable names
- Turn-based combat system with attack and defense mechanics
- Experience and leveling system with stat progression
- Multiple enemy types with difficulty scaling based on player level
- Rest and healing mechanics
- Interactive command-line interface

## Commands

### Build and Run
```bash
# Build the game
go build -o rpg

# Run console version
go run *.go

# Run web server version
go run *.go server
# or after building:
./rpg server

# Run system tests
go run *.go test
```

### Development
```bash
# Format code
go fmt ./...

# Check for issues
go vet ./...

# Build without running
go build
```

## Game Mechanics

- Players start at level 1 with base stats
- Combat is turn-based with damage calculations including defense
- Experience gained from defeating enemies triggers level-ups
- Enemy spawning scales with player level for balanced difficulty
- Healing restores 1/3 of max health per rest