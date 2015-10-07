package main

import (
	"fmt"
	"image"
	"os"
)

const (
	AvatarSize = 120
	PixelSize = AvatarSize/12
)

func main() {
	fmt.Println(os.Args[1])
	if len(os.Args) != 2 {
		fmt.Println("Usage: go-avatar username")
		return
	}

	nameBytes := []byte(os.Args[1])
	fmt.Println(nameBytes)
	avatar := image.NewRGBA(image.Rect(0, 0, AvatarSize, AvatarSize))
	PaintBackGround(avatar, CalculateBGColor(nameBytes))
	DrawPattern(avatar, nameBytes, CalculatePixelColor(nameBytes))
	SavePNG(avatar)
}
