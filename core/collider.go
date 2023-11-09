package core

import "github.com/co0p/tankism/lib/collision"

type BoxCollider struct {
	Enabled  bool
	Offset   Vector2i
	Position Vector2
	Size     Vector2i
}

type CollisionData struct {
	GameObject1 *GameObject
	GameObject2 *GameObject
}

func (col BoxCollider) getBoundingBox() collision.BoundingBox {
	return collision.BoundingBox{
		X:      col.Position.X + float64(col.Offset.X),
		Y:      col.Position.Y + float64(col.Offset.Y),
		Width:  float64(col.Size.X),
		Height: float64(col.Size.Y),
	}
}

func CheckCollision(col1 BoxCollider, col2 BoxCollider) bool {
	b1 := col1.getBoundingBox()
	b2 := col2.getBoundingBox()
	return collision.AABBCollision(b1, b2)
}

func CheckOffsetCollision(col1 BoxCollider, col2 BoxCollider, offset Vector2) bool {
	col1.Position = col1.Position.Add(offset)
	return CheckCollision(col1, col2)
}

//func (col BoxCollider) vectorToCenter(col2 BoxCollider) Vector2 {
//	point1 := col.Position.Add(col.Offset.ToVector2())
//	point2 := col2.Position.Add(col2.Offset.ToVector2())
//	return point2.Sub(point1)
//}
//
//func (col BoxCollider) vectorEdgeToEdge(col2 BoxCollider) Vector2 {
//	centerDiff := col.vectorToCenter(col2)
//	if centerDiff.X > 0 {
//		centerDiff.X = centerDiff.X - float64(col.Size.X)/2 + float64(col2.Size.X)/2
//	} else {
//		centerDiff.X = centerDiff.X + float64(col.Size.X)/2 - float64(col2.Size.X)/2
//	}
//	if centerDiff.Y > 0 {
//		centerDiff.Y = centerDiff.Y - float64(col.Size.Y)/2 + float64(col2.Size.Y)/2
//	} else {
//		centerDiff.Y = centerDiff.Y + float64(col.Size.Y)/2 - float64(col2.Size.Y)/2
//	}
//	return centerDiff
//}
