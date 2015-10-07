package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

func SavePNG(avatar image.Image) {
	file, err := os.Create(os.Args[1] + ".png")
	err = png.Encode(file, avatar)

	if err != nil {
		panic(err)
	}

}

func DrawPattern(avatar *image.RGBA, nameBytes []byte, pixelColor color.RGBA) {

	// Random number based on the username.
	var nameSum int64
	for i := range nameBytes {
		nameSum += int64(nameBytes[i])
	}
	fmt.Println("nameSum:: ", nameSum)

	// Use nameSum number to keep random-ness deterministic for a given name
	rand.Seed(nameSum)

	// Make the "pattern"
	for y := PixelSize; y < AvatarSize/2 - PixelSize; y=y+PixelSize {
		for x := PixelSize; x < AvatarSize/2 - PixelSize; x=x+PixelSize {
			if ((x + y) % 2) == 0 {
				if (rand.Intn(3) == 1) || (rand.Intn(3) == 2){
					fmt.Printf("x= %d, y= %d\n", x,y)
					fillPixels(avatar, x, y, pixelColor)
				}
			}
		}
	}

	// Mirror top left half to right top half
	for y := 0; y < AvatarSize; y++ {
		for x := 0; x < AvatarSize; x++ {
			if x < AvatarSize/2 {
				avatar.Set(AvatarSize-x-1, y, avatar.At(x, y))
			}
		}
	}

	// Mirror top half to bottom half
	for y := 0; y < AvatarSize; y++ {
		for x := 0; x < AvatarSize; x++ {
			if y < AvatarSize/2 {
				avatar.Set(x, AvatarSize-y-1, avatar.At(x, y))
			}
		}
	}
}


func fillPixels(avatar *image.RGBA, x, y int, pixelColor color.RGBA){
	for i := x; i < x+PixelSize; i++ {
		for j := y; j < y+PixelSize; j++ {
			avatar.SetRGBA(i, j, pixelColor)
		}
	}
}

func PaintBackGround(avatar *image.RGBA, bgColor color.RGBA) {
	for y := 0; y < AvatarSize; y++ {
		for x := 0; x < AvatarSize; x++ {
			avatar.SetRGBA(x, y, bgColor)
		}
	}
}

func CalculatePixelColor(nameBytes []byte) (pixelColor color.RGBA) {
	pixelColor.A = 255

	var mutator = byte((len(nameBytes) * 4))

	pixelColor.R = nameBytes[0] * mutator
	pixelColor.G = nameBytes[1] * mutator
	pixelColor.B = nameBytes[2] * mutator

	return
}

func CalculateBGColor(nameBytes []byte) (bgColor color.RGBA) {
	bgColor.A = 255

	var mutator = byte((len(nameBytes) * 3))

	bgColor.R = nameBytes[0] * mutator
	bgColor.G = nameBytes[1] * mutator
	bgColor.B = nameBytes[2] * mutator

	return
}
