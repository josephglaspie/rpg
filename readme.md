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

## Kubernetes Deployment (Minikube)

Deploy the RPG game to your local minikube cluster and access it through your browser.

### Prerequisites
- [Minikube](https://minikube.sigs.k8s.io/docs/start/) installed and running
- [kubectl](https://kubernetes.io/docs/tasks/tools/) configured to work with minikube
- Docker (for building the image locally)

### Quick Start

1. **Start minikube with ingress addon:**
   ```bash
   minikube start
   minikube addons enable ingress
   ```

2. **Build and load the Docker image into minikube:**
   ```bash
   # Build the Docker image
   docker build -t josephglaspie/rpg-game:latest .
   
   # Push the image to Dockerhub
    docker push josephglaspie/rpg-game:latest
   ```

3. **Deploy to Kubernetes:**
   ```bash
   # Apply all Kubernetes manifests
   kubectl apply -f k8s/
   ```

4. **Set up local DNS (add to /etc/hosts):**
   ```bash
   # Get minikube IP
   minikube ip
   
   # Add this line to your /etc/hosts file (replace <MINIKUBE_IP> with actual IP):
   # <MINIKUBE_IP> rpg-game.local
   echo "$(minikube ip) rpg-game.local" | sudo tee -a /etc/hosts
   ```

5. **Access the game:**
   - Open your browser and go to: **http://rpg-game.local**
   - The game should load and be ready to play!

### Alternative Access Methods

**Option 1: Port Forwarding (no DNS setup needed):**
```bash
kubectl port-forward -n rpg-game service/rpg-game-service 8081:80
# Then visit: http://localhost:8081
```

**Option 2: Minikube Service (opens automatically):**
```bash
minikube service rpg-game-service -n rpg-game
```

### Deployment Management

**Check deployment status:**
```bash
kubectl get pods -n rpg-game
kubectl get services -n rpg-game
kubectl get ingress -n rpg-game
```

**View logs:**
```bash
kubectl logs -n rpg-game -l app=rpg-game -f
```

**Scale the deployment:**
```bash
kubectl scale deployment rpg-game -n rpg-game --replicas=3
```

**Clean up:**
```bash
# Remove all resources
kubectl delete -f k8s/

# Remove from /etc/hosts
sudo sed -i '/rpg-game.local/d' /etc/hosts
```

### Troubleshooting

- **Ingress not working?** Wait a few minutes for the ingress controller to start, or try `minikube addons disable ingress && minikube addons enable ingress`
- **Image not found?** Make sure you ran `minikube image load josephglaspie/rpg-game:latest`
- **Can't access the game?** Check that minikube is running with `minikube status`

Enjoy your adventure! 🗡️