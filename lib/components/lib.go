package components

import (
	ecs "github.com/x-hgg-x/goecs"
	"github.com/x-hgg-x/goecsengine/math"
)

// Components contains references to all game components
type Components struct {
	Player       *ecs.Component
	Enemy        *ecs.Component
	Controllable *ecs.Component
	Alien        *ecs.Component
	AlienMaster  *ecs.Component
	Bunker       *ecs.Component
	Bullet       *ecs.Component
	PlayerLine   *ecs.Component
	Deleted      *ecs.Component
}

// Player component
type Player struct{}

// Enemy component
type Enemy struct{}

// Controllable component
type Controllable struct {
	Width  float64
	Height float64
}

// Alien component
type Alien struct {
	Width       float64
	Height      float64
	Translation math.Vector2
}

// AlienMaster component
type AlienMaster struct {
	Direction float64
}

// Bunker component
type Bunker struct {
	PixelSize int `toml:"pixel_size"`
}

// Bullet component
type Bullet struct {
	Width    float64
	Height   float64
	Velocity float64
	Health   float64
}

// PlayerLine component
type PlayerLine struct{}

// Deleted component
type Deleted struct{}
