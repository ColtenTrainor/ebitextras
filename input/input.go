package input

import (
	"ebitextras/core"
	"ebitextras/objectpool"
	"github.com/hajimehoshi/ebiten/v2"
)

type Binds struct {
	bindMap *map[string]*objectpool.ObjectPool[ebiten.Key]
}

func NewBinds() *Binds {
	mp := make(map[string]*objectpool.ObjectPool[ebiten.Key])
	return &Binds{
		bindMap: &mp,
	}
}

func (binds *Binds) Add(action string, control ebiten.Key) {
	mp := *binds.bindMap
	if mp[action] == nil {
		mp[action] = objectpool.NewObjectPool[ebiten.Key]()
	}
	keys := mp[action]
	if keys.Contains(control) {
		return
	}
	keys.Add(control)
}

func (binds *Binds) Addr(action string, keys ...ebiten.Key) {
	for _, key := range keys {
		binds.Add(action, key)
	}
}

func (binds *Binds) Remove(action string, control ebiten.Key) {
	keys := (*binds.bindMap)[action]
	if keys.Contains(control) {
		_ = keys.FindAndRemove(control)
	}
}

func (binds *Binds) IsPressed(action string) bool {
	for _, key := range (*binds.bindMap)[action].Objects {
		if ebiten.IsKeyPressed(key) {
			return true
		}
	}
	return false
}

func (binds *Binds) ReadVectorComposite(xNegKey string, xPosKey string,
	yNegKey string, yPosKey string) core.Vector2 {
	var xPos float64
	if binds.IsPressed(xPosKey) {
		xPos = 1
	}
	var xNeg float64
	if binds.IsPressed(xNegKey) {
		xNeg = 1
	}
	var yPos float64
	if binds.IsPressed(yPosKey) {
		yPos = 1
	}
	var yNeg float64
	if binds.IsPressed(yNegKey) {
		yNeg = 1
	}

	return core.Vector2{
		X: xPos - xNeg,
		Y: yPos - yNeg,
	}
}
