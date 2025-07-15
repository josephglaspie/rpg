# Role Playing Game

A simple RPG game written in Go featuring turn-based combat, character progression, and multiple enemy types. Play either in the console or in your web browser!

## How to Play

### Installation & Setup
1. Make sure you have Go installed on your system
2. Clone or download this repository
3. Navigate to the game directory

### Starting the Game

**Web Browser Version (Recommended):**
```bash
# Build the game
go build -o rpg

# Start the web server
./rpg server

# Open your browser and go to: http://localhost:8080
```

**Console Version:**
```bash
# Run the original console game
go run *.go
# or after building:
./rpg
```

### Game Instructions

**Character Creation:**
- Enter your character's name when prompted
- Your hero starts at level 1 with base stats

**Main Menu Options:**
1. **Fight a Monster** - Battle randomly spawned enemies
2. **Rest and Heal** - Recover health points  
3. **View Stats** - Check your character's current status

**Combat System:**
- Choose to **Attack** or **Run Away** during battles
- Damage is calculated using your Attack vs enemy Defense
- Defeat enemies to gain experience points
- Level up automatically when you gain enough experience

**Browser Interface:**
- Real-time stat display on the left panel
- Game log shows all actions and results
- Enemy stats appear on the right during combat
- Click buttons to perform actions
- Responsive design works on mobile devices

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
# or after building:
./rpg test
```

### Game Modes
- **Browser Mode**: Modern web interface with visual design and responsive layout
- **Console Mode**: Traditional text-based terminal gameplay
- **Test Mode**: Automated testing of game systems

Enjoy your adventure! 🗡️