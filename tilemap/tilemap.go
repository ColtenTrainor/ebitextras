package tilemap

import (
	"github.com/coltentrainor/ebitextras/core"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled"
	"image"
	"io/fs"
	"log"
)

type Level struct {
	LevelData *tiled.Map
	Tiles     map[int]*ebiten.Image
}

func LoadTileImages(fs fs.FS, levelData *tiled.Map) map[int]*ebiten.Image {
	tiles := make(map[int]*ebiten.Image)

	for _, tileset := range levelData.Tilesets {
		tileImage, _, err := ebitenutil.NewImageFromFileSystem(fs, "assets/tilemap/"+tileset.Image.Source)
		if err != nil {
			log.Fatal("Failed to load tile image: ", err)
		}

		for i := 0; i < tileset.TileCount; i++ {
			id := tileset.FirstGID + uint32(i)

			tilesheetPos := core.Vector2i{
				X: i % tileset.Columns,
				Y: i / tileset.Columns,
			}

			subImgRect := image.Rect(
				tilesheetPos.X*tileset.TileWidth,
				tilesheetPos.Y*tileset.TileHeight,
				(tilesheetPos.X+1)*tileset.TileWidth,
				(tilesheetPos.Y+1)*tileset.TileHeight)

			tiles[int(id)] = tileImage.SubImage(subImgRect).(*ebiten.Image)
		}
	}

	return tiles
}

func (level *Level) DrawLevel(screen *ebiten.Image, drawOps ebiten.DrawImageOptions) {
	levelSize := core.Vector2i{X: level.LevelData.Width, Y: level.LevelData.Height}
	for _, layer := range level.LevelData.Layers {
		for y := 0; y < levelSize.Y; y++ {
			for x := 0; x < levelSize.X; x++ {
				ops := drawOps
				pos := core.Vector2i{
					X: level.LevelData.TileWidth * x,
					Y: level.LevelData.TileHeight * y,
				}
				ops.GeoM.Translate(float64(pos.X), float64(pos.Y))
				layerTile := layer.Tiles[y*levelSize.X+x]
				idOffset := uint32(0)
				if layerTile.Tileset != nil {
					idOffset = layerTile.Tileset.FirstGID
				}
				id := int(layerTile.ID + idOffset)
				if id == 0 {
					continue
				}
				tile := level.Tiles[id]
				screen.DrawImage(tile, &ops)
			}
		}
	}
}
