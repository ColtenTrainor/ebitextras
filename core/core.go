package core

// GameObject is the core component of ebitextras. It contains common information most objects in a game
// would need to have.
type GameObject struct {
	Position     Vector2
	CurrentSheet *SpriteSheet
	Speed        float64
	Collider     *BoxCollider
}
