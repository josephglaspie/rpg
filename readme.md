# Role Playing Game

A simple command-line RPG game written in Go featuring turn-based combat, character progression, and multiple enemy types.

## How to Play

### Installation & Setup
1. Make sure you have Go installed on your system
2. Clone or download this repository
3. Navigate to the game directory

### Starting the Game
```bash
go run *.go
```

### Game Instructions

**Character Creation:**
- Enter your character's name when prompted
- Your hero starts at level 1 with base stats

**Main Menu Options:**
1. **Fight a monster** - Battle randomly spawned enemies
2. **Rest and heal** - Recover health points
3. **View stats** - Check your character's current status
4. **Quit game** - Exit the game

**Combat System:**
- Choose to **Attack** or **Run away** during battles
- Damage is calculated using your Attack vs enemy Defense
- Defeat enemies to gain experience points
- Level up automatically when you gain enough experience

**Character Progression:**
- Each level increases your Health, Attack, and Defense
- Higher levels unlock stronger enemies
- Experience required increases with each level

**Enemy Types:**
- **Goblin** (Level 1-2): Weak enemies, good for beginners
- **Orc** (Level 3-4): Moderate challenge
- **Troll** (Level 5-7): Strong opponents
- **Dragon** (Level 8+): Ultimate challenge

### Tips for Success
- Rest regularly to maintain full health
- Don't be afraid to run from tough battles early on
- Each victory makes you stronger for future encounters
- Enemy difficulty scales with your level

### Testing
Run the built-in test suite:
```bash
go run *.go test
```

Enjoy your adventure! 🗡️