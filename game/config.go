package game

type EnemyConfig struct {
	InitialLife int
	UpgradeLife int
}

type GameConfig struct {
	IterationDuration    int
	MoneyPerIteration    int
	MaximumIterations    int
	MoneyPerKill         int
	MoneyPerTowerBuild   int
	MoneyPerTowerUpgrade int
}

type PlayerConfig struct {
	InitialLifes int
	InitialMoney int
}

type TowerConfig struct {
	InitialStrength int
	UpgradeStrength int
	InitialCooldown int
	UpgradeCooldown int
}

type WaveConfig struct {
	StartDecay int
}

var enemyConfig = EnemyConfig{
	InitialLife: 1,
	UpgradeLife: 1,
}

var gameConfig = GameConfig{
	IterationDuration:    60,
	MoneyPerIteration:    1,
	MoneyPerKill:         1,
	MoneyPerTowerBuild:   1,
	MoneyPerTowerUpgrade: 1,
	MaximumIterations:    100000,
}

var playerConfig = PlayerConfig{
	InitialLifes: 10,
	InitialMoney: 1,
}

var towerConfig = TowerConfig{
	InitialStrength: 1,
	UpgradeStrength: 1,
	InitialCooldown: 15,
	UpgradeCooldown: 1,
}

var waveConfig = WaveConfig{
	StartDecay: -5,
}

var Config = struct {
	Enemy  EnemyConfig
	Game   GameConfig
	Player PlayerConfig
	Tower  TowerConfig
	Wave   WaveConfig
}{
	Enemy:  enemyConfig,
	Game:   gameConfig,
	Player: playerConfig,
	Tower:  towerConfig,
	Wave:   waveConfig,
}
