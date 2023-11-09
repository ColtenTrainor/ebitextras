package core

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/png"
	"log"
)

// SpriteLayout defines whether a sprite sheet is formatted horizontally (rows are animations), or
// vertically (columns are animations).
type SpriteLayout int

const (
	// Horizontal declares that a sprite sheet is formatted horizontally (rows are animations).
	Horizontal SpriteLayout = iota
	// Vertical declares that a sprite sheet is formatted vertically (columns are animations).
	Vertical
)

// SpriteSheet stores the image of a sprite sheet and all necessary information to render and animate it.
// To use a non-animated static image, set Animated to false and do not call UpdateAnim on it.
// There is currently no support for animations composed of separate image files
type SpriteSheet struct {
	// Image the ebiten.Image of the sprite sheet, or just the sprite if it is not animated.
	Image *ebiten.Image
	// Animated is whether the SpriteSheet is animated or just a single static image.
	Animated bool
	// Cells defines how many rows and columns the sprite sheet contains.
	Cells Vector2i
	// Frame is the current sub-image of the animation.
	Frame        int
	delayCounter int
	// DelayFrames is the number of frames each sub-image of the animation is shown before changing to
	// the next one.
	DelayFrames int
	// Layout specifies whether a sprite sheet is formatted in Horizontal (rows are animations) or
	// Vertical (columns are animations) form.
	Layout SpriteLayout
	// SelectedAnim specifies the row or column (based on the Layout) that is currently selected.
	SelectedAnim int
	// Flipped flips the image when drawn to the screen
	Flipped bool
}

// DrawGameObject draws the sub-image of the SpriteSheet for the current frame at the GameObject's position.
func (gameObject *GameObject) DrawGameObject(screen *ebiten.Image, ops ebiten.DrawImageOptions) {
	anim := gameObject.CurrentSheet

	var finalImg *ebiten.Image
	if anim.Animated {
		finalImg = anim.getAnimSubImage()
	} else {
		finalImg = anim.Image
	}

	if anim.Flipped {
		ops.GeoM.Scale(-1, 1)
		ops.GeoM.Translate(float64(finalImg.Bounds().Dx()), 0)
	}

	ops.GeoM.Translate(gameObject.Position.X, gameObject.Position.Y)
	screen.DrawImage(finalImg, &ops)
}

func (spriteSheet *SpriteSheet) getAnimSubImage() *ebiten.Image {
	cellSize := Vector2i{
		X: spriteSheet.Image.Bounds().Dx() / spriteSheet.Cells.X,
		Y: spriteSheet.Image.Bounds().Dy() / spriteSheet.Cells.Y,
	}

	var spriteLoc Vector2i
	if spriteSheet.Layout == Horizontal {
		spriteLoc = Vector2i{X: spriteSheet.Frame, Y: spriteSheet.SelectedAnim}
	} else {
		spriteLoc = Vector2i{X: spriteSheet.SelectedAnim, Y: spriteSheet.Frame}
	}

	rect := image.Rect(
		spriteLoc.X*cellSize.X,
		spriteLoc.Y*cellSize.Y,
		(spriteLoc.X+1)*cellSize.X,
		(spriteLoc.Y+1)*cellSize.Y)

	return spriteSheet.Image.SubImage(rect).(*ebiten.Image)
}

func (gameObject *GameObject) ChangeSheet(newSheet *SpriteSheet) {
	gameObject.CurrentSheet = newSheet
	gameObject.CurrentSheet.Frame = 0
	gameObject.CurrentSheet.delayCounter = 0
}

// UpdateAnim manages the timing and updating of a SpriteSheet. Call this during ebiten's Update function
// on each SpriteSheet that needs its animation to play.
func (spriteSheet *SpriteSheet) UpdateAnim() {
	anim := spriteSheet
	anim.delayCounter++
	if anim.delayCounter%anim.DelayFrames == 0 {
		anim.Frame++
		if (anim.Layout == Horizontal && anim.Frame >= anim.Cells.X) ||
			(anim.Layout == Vertical && anim.Frame >= anim.Cells.Y) {
			anim.Frame = 0
		}
	}
}

// LoadEmbeddedImage loads a png from the embedded filesystem
func LoadEmbeddedImage(assets embed.FS, path string) *ebiten.Image {
	file, err := assets.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	ebitImg := ebiten.NewImageFromImage(img)
	return ebitImg
}
