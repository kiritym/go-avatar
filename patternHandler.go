package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

func SavePNG(avatar image.Image, name string) {
	file, err := os.Create(name + ".png")
	err = png.Encode(file, avatar)

	if err != nil {
		panic(err)
	}

}

func DrawPattern(avatar *image.RGBA, nameBytes []byte, pixelColor color.RGBA) {

	// Random number based on the username.
	var nameSum int64
	for i := range nameBytes {
		nameSum += int64(nameBytes[i])*int64(i)
	}
	fmt.Println("NameSum: ", nameSum)
	// Use nameSum number to keep random-ness deterministic for a given name
	rand.Seed(nameSum)
	DrawTopPattern(avatar, pixelColor)
	DrawBottomPattern(avatar, pixelColor)
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

	var mutator = byte((len(nameBytes) * 6))

	bgColor.R = nameBytes[0] * mutator
	bgColor.G = nameBytes[1] * mutator
	bgColor.B = nameBytes[2] * mutator

	return
}

func DrawTopPattern(avatar *image.RGBA, pixelColor color.RGBA) {
	// Make the top left quadrant pattern
	for y := PixelSize; y < AvatarSize/2 - PixelSize; y=y+PixelSize {
		for x := PixelSize; x < AvatarSize/2 - PixelSize; x=x+PixelSize {
			if ((x + y) % 2) == 0 {
				randNumber := rand.Intn(3)
				if (randNumber == 1) || (randNumber == 2){
					//fmt.Printf("x= %d, y= %d\n", x,y)
					fillPixels(avatar, x, y, pixelColor)
				}
			}
		}
	}
	DrawTopRightPattern(avatar)
}

func DrawTopRightPattern(avatar *image.RGBA) {
	// Mirror top left quadrant to right top quadrant
	for y := 0; y < AvatarSize; y++ {
		for x := 0; x < AvatarSize; x++ {
			if x < AvatarSize/2 {
				avatar.Set(AvatarSize-x-1, y, avatar.At(x, y))
			}
		}
	}
}

func DrawBottomPattern(avatar *image.RGBA, pixelColor color.RGBA) {
	// Make the bottom left pattern
	for y := AvatarSize/2 + PixelSize; y < AvatarSize - PixelSize; y=y+PixelSize {
		for x := PixelSize; x < AvatarSize/2 - PixelSize; x=x+PixelSize {
			if ((x + y) % 2) == 0 {
				randNumber := rand.Intn(3)
				//fmt.Println(randNumber)
				if ((randNumber == 0) || (randNumber == 1)){
					//fmt.Printf("x= %d, y= %d\n", x,y)
					fillPixels(avatar, x, y, pixelColor)
				}
			}
		}
	}
	DrawBottomRightPattern(avatar)
}

func DrawBottomRightPattern(avatar *image.RGBA) {
	// Mirror bottom left quadrant to right bottom quadrant
	for y := AvatarSize/2; y < AvatarSize; y++ {
		for x := 0; x < AvatarSize; x++ {
			if (x < AvatarSize/2 ) {
				avatar.Set(AvatarSize-x-1, y, avatar.At(x, y))
			}
		}
	}
}
