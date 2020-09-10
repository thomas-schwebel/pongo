package main

import (
	"image"
	"os"

	"github.com/faiface/pixel"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func loadSprite(filePath string) *pixel.Sprite {
	bgPic, err := loadPicture(filePath)
	if err != nil {
		panic(err)
	}

	return pixel.NewSprite(bgPic, bgPic.Bounds())
}
