class Character {
    constructor(name) {
        this.name = name || "Hero";
        this.level = 1;
        this.health = 100;
        this.maxHealth = 100;
        this.attack = 15;
        this.defense = 5;
        this.experience = 0;
        this.experienceToNext = 100;
    }

    isAlive() {
        return this.health > 0;
    }

    takeDamage(damage) {
        const actualDamage = Math.max(1, damage - this.defense);
        this.health = Math.max(0, this.health - actualDamage);
        return actualDamage;
    }

    heal(amount) {
        this.health = Math.min(this.maxHealth, this.health + amount);
    }

    attackEnemy(enemy) {
        const baseDamage = this.attack + Math.floor(Math.random() * 10);
        const actualDamage = Math.max(1, baseDamage - enemy.defense);
        enemy.takeDamage(actualDamage);
        return actualDamage;
    }

    gainExperience(exp) {
        this.experience += exp;
        if (this.experience >= this.experienceToNext) {
            this.levelUp();
        }
    }

    levelUp() {
        this.level++;
        this.experience -= this.experienceToNext;
        this.experienceToNext = this.level * 100;

        const healthIncrease = 20 + Math.floor(Math.random() * 10);
        const attackIncrease = 3 + Math.floor(Math.random() * 3);
        const defenseIncrease = 2 + Math.floor(Math.random() * 2);

        this.maxHealth += healthIncrease;
        this.health = this.maxHealth;
        this.attack += attackIncrease;
        this.defense += defenseIncrease;

        return {
            healthIncrease,
            attackIncrease,
            defenseIncrease
        };
    }
}

class Enemy {
    constructor(name, health, attack, defense, expReward) {
        this.name = name;
        this.health = health;
        this.maxHealth = health;
        this.attack = attack;
        this.defense = defense;
        this.expReward = expReward;
    }

    isAlive() {
        return this.health > 0;
    }

    takeDamage(damage) {
        this.health = Math.max(0, this.health - damage);
    }

    attackPlayer(player) {
        const baseDamage = this.attack + Math.floor(Math.random() * 5);
        return player.takeDamage(baseDamage);
    }
}

function createGoblin() {
    return new Enemy("Goblin", 30, 8, 2, 25);
}

function createOrc() {
    return new Enemy("Orc", 60, 12, 4, 50);
}

function createTroll() {
    return new Enemy("Troll", 120, 20, 8, 100);
}

function createDragon() {
    return new Enemy("Dragon", 200, 35, 15, 250);
}

function spawnRandomEnemy(playerLevel) {
    const rand = Math.random();
    
    if (playerLevel <= 2) {
        return createGoblin();
    } else if (playerLevel <= 4) {
        return rand < 0.5 ? createGoblin() : createOrc();
    } else if (playerLevel <= 7) {
        const choice = Math.floor(Math.random() * 3);
        switch (choice) {
            case 0: return createOrc();
            case 1: return createTroll();
            default: return createGoblin();
        }
    } else {
        const choice = Math.floor(Math.random() * 4);
        switch (choice) {
            case 0: return createOrc();
            case 1: return createTroll();
            case 2: return createDragon();
            default: return createGoblin();
        }
    }
}

class Game {
    constructor() {
        this.player = null;
        this.currentEnemy = null;
        this.inCombat = false;
        this.initializeEventListeners();
    }

    initializeEventListeners() {
        document.getElementById('start-game').addEventListener('click', () => this.startGame());
        document.getElementById('fight-btn').addEventListener('click', () => this.startCombat());
        document.getElementById('rest-btn').addEventListener('click', () => this.restAndHeal());
        document.getElementById('stats-btn').addEventListener('click', () => this.showStats());
        document.getElementById('attack-btn').addEventListener('click', () => this.playerAttack());
        document.getElementById('run-btn').addEventListener('click', () => this.runAway());
        document.getElementById('restart-btn').addEventListener('click', () => this.restart());

        document.getElementById('character-name').addEventListener('keypress', (e) => {
            if (e.key === 'Enter') {
                this.startGame();
            }
        });
    }

    startGame() {
        const nameInput = document.getElementById('character-name');
        const playerName = nameInput.value.trim() || 'Hero';
        
        this.player = new Character(playerName);
        this.showScreen('game-screen');
        this.updateUI();
        this.logMessage(`Welcome, ${this.player.name}! Your adventure begins now.`, 'info-msg');
    }

    showScreen(screenId) {
        document.querySelectorAll('.screen').forEach(screen => {
            screen.classList.remove('active');
        });
        document.getElementById(screenId).classList.add('active');
    }

    updateUI() {
        if (!this.player) return;

        document.getElementById('character-title').textContent = 
            `${this.player.name} (Level ${this.player.level})`;
        document.getElementById('health-display').textContent = 
            `${this.player.health}/${this.player.maxHealth}`;
        document.getElementById('attack-display').textContent = this.player.attack;
        document.getElementById('defense-display').textContent = this.player.defense;
        document.getElementById('exp-display').textContent = 
            `${this.player.experience}/${this.player.experienceToNext}`;
    }

    updateEnemyUI() {
        if (!this.currentEnemy) {
            document.getElementById('enemy-panel').style.display = 'none';
            return;
        }

        document.getElementById('enemy-panel').style.display = 'block';
        document.getElementById('enemy-name').textContent = this.currentEnemy.name;
        document.getElementById('enemy-health').textContent = 
            `${this.currentEnemy.health}/${this.currentEnemy.maxHealth}`;
        document.getElementById('enemy-attack').textContent = this.currentEnemy.attack;
        document.getElementById('enemy-defense').textContent = this.currentEnemy.defense;
    }

    logMessage(message, className = '') {
        const gameLog = document.getElementById('game-log');
        const p = document.createElement('p');
        p.textContent = message;
        if (className) {
            p.className = className;
        }
        gameLog.appendChild(p);
        gameLog.scrollTop = gameLog.scrollHeight;
    }

    startCombat() {
        this.currentEnemy = spawnRandomEnemy(this.player.level);
        this.inCombat = true;
        
        this.logMessage(`⚔️ A wild ${this.currentEnemy.name} appears!`, 'combat-msg');
        this.updateEnemyUI();
        this.showCombatActions();
    }

    showCombatActions() {
        document.getElementById('main-actions').classList.add('hidden');
        document.getElementById('combat-actions').classList.remove('hidden');
    }

    showMainActions() {
        document.getElementById('main-actions').classList.remove('hidden');
        document.getElementById('combat-actions').classList.add('hidden');
    }

    playerAttack() {
        if (!this.inCombat || !this.currentEnemy) return;

        const damage = this.player.attackEnemy(this.currentEnemy);
        this.logMessage(`You attack the ${this.currentEnemy.name} for ${damage} damage!`, 'combat-msg');
        
        this.updateEnemyUI();

        if (!this.currentEnemy.isAlive()) {
            this.logMessage(`🎉 You defeated the ${this.currentEnemy.name}!`, 'success-msg');
            this.logMessage(`You gained ${this.currentEnemy.expReward} experience points!`, 'success-msg');
            
            const oldLevel = this.player.level;
            this.player.gainExperience(this.currentEnemy.expReward);
            
            if (this.player.level > oldLevel) {
                const increases = this.player.levelUp();
                this.logMessage(`🎉 Level Up! ${this.player.name} is now level ${this.player.level}!`, 'success-msg');
                this.logMessage(`Health: +${increases.healthIncrease} (now ${this.player.maxHealth})`, 'success-msg');
                this.logMessage(`Attack: +${increases.attackIncrease} (now ${this.player.attack})`, 'success-msg');
                this.logMessage(`Defense: +${increases.defenseIncrease} (now ${this.player.defense})`, 'success-msg');
            }
            
            this.endCombat();
            return;
        }

        const enemyDamage = this.currentEnemy.attackPlayer(this.player);
        this.logMessage(`The ${this.currentEnemy.name} attacks you for ${enemyDamage} damage!`, 'combat-msg');
        this.logMessage(`Your health: ${this.player.health}/${this.player.maxHealth}`, 'info-msg');
        
        this.updateUI();

        if (!this.player.isAlive()) {
            this.gameOver();
        }
    }

    runAway() {
        this.logMessage('You successfully run away from the battle!', 'warning-msg');
        this.endCombat();
    }

    endCombat() {
        this.inCombat = false;
        this.currentEnemy = null;
        this.updateEnemyUI();
        this.showMainActions();
        this.updateUI();
    }

    restAndHeal() {
        const healAmount = Math.max(10, Math.floor(this.player.maxHealth / 3));
        this.player.heal(healAmount);
        this.logMessage(`💚 You rest and recover ${healAmount} health points.`, 'success-msg');
        this.logMessage(`Current health: ${this.player.health}/${this.player.maxHealth}`, 'info-msg');
        this.updateUI();
    }

    showStats() {
        this.logMessage(`=== ${this.player.name} (Level ${this.player.level}) ===`, 'info-msg');
        this.logMessage(`Health: ${this.player.health}/${this.player.maxHealth}`, 'info-msg');
        this.logMessage(`Attack: ${this.player.attack}`, 'info-msg');
        this.logMessage(`Defense: ${this.player.defense}`, 'info-msg');
        this.logMessage(`Experience: ${this.player.experience}/${this.player.experienceToNext}`, 'info-msg');
    }

    gameOver() {
        this.logMessage(`💀 ${this.player.name} has fallen in battle! Game Over! 💀`, 'combat-msg');
        document.getElementById('final-stats').textContent = 
            `${this.player.name} reached level ${this.player.level} before falling in battle.`;
        this.showScreen('game-over');
    }

    restart() {
        this.player = null;
        this.currentEnemy = null;
        this.inCombat = false;
        document.getElementById('character-name').value = '';
        document.getElementById('game-log').innerHTML = '';
        this.showScreen('character-creation');
    }
}

document.addEventListener('DOMContentLoaded', () => {
    new Game();
});