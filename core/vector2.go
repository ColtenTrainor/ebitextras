package core

import "math"

// Vector2 Utility struct to hold 2 floats.
type Vector2 struct {
	X float64
	Y float64
}

// Vector2i Utility struct to hold 2 ints.
type Vector2i struct {
	X int
	Y int
}

func (vec Vector2) Add(vec2 Vector2) Vector2 {
	vec.X += vec2.X
	vec.Y += vec2.Y
	return vec
}

func (vec Vector2i) Add(vec2 Vector2i) Vector2i {
	vec.X += vec2.X
	vec.Y += vec2.Y
	return vec
}

func (vec Vector2) Sub(vec2 Vector2) Vector2 {
	vec.X -= vec2.X
	vec.Y -= vec2.Y
	return vec
}

func (vec Vector2i) Sub(vec2 Vector2i) Vector2i {
	vec.X -= vec2.X
	vec.Y -= vec2.Y
	return vec
}

func (vec Vector2) Mul(vec2 Vector2) Vector2 {
	vec.X *= vec2.X
	vec.Y *= vec2.Y
	return vec
}

func (vec Vector2i) Mul(vec2 Vector2i) Vector2i {
	vec.X *= vec2.X
	vec.Y *= vec2.Y
	return vec
}

func (vec Vector2) Mulf(f float64) Vector2 {
	vec.X *= f
	vec.Y *= f
	return vec
}

func (vec Vector2i) Mulf(i int) Vector2i {
	vec.X *= i
	vec.Y *= i
	return vec
}

func (vec Vector2) Magnitude() float64 {
	return math.Sqrt(vec.X*vec.X + vec.Y*vec.Y)
}

func (vec Vector2i) Magnitude() float64 {
	return math.Sqrt(float64(vec.X*vec.X + vec.Y*vec.Y))
}

func (vec Vector2) Normalized() Vector2 {
	mag := vec.Magnitude()
	if vec.X != 0 {
		vec.X /= mag
	}
	if vec.Y != 0 {
		vec.Y /= mag
	}
	return vec
}

func (vec Vector2i) Normalized() Vector2i {
	mag := vec.Magnitude()
	if vec.X != 0 {
		vec.X = int(float64(vec.X) / mag)
	}
	if vec.Y != 0 {
		vec.Y = int(float64(vec.Y) / mag)
	}
	return vec
}

func (vec Vector2) Equals(vec2 Vector2) bool {
	if vec.X == vec2.X && vec.Y == vec2.Y {
		return true
	}
	return false
}

func (vec Vector2i) Equals(vec2 Vector2i) bool {
	if vec.X == vec2.X && vec.Y == vec2.Y {
		return true
	}
	return false
}

func (vec Vector2) ToVector2i() Vector2i {
	return Vector2i{X: int(vec.X), Y: int(vec.Y)}
}

func (vec Vector2i) ToVector2() Vector2 {
	return Vector2{X: float64(vec.X), Y: float64(vec.Y)}
}
