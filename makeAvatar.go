package main

import (
	"fmt"
	"image"
	//"os"
	"net/http"
	"image/jpeg"
	"bytes"
	"log"
	"strconv"
	"encoding/base64"
	"html/template"
)

var (
	AvatarSize int
	PixelSize int
)


func main() {

	http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
	log.Println("Listening on 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

  nameBytes := []byte(r.URL.Path[1:])
  //fmt.Println(r.URL.Path[1:])

  //fmt.Println(nameBytes)
	if (len(nameBytes) < 3){
		fmt.Fprintf(w, "Please provide a valid username(atleast 3 characters).")
	}else{
		setSize(w, r)
		avatar := image.NewRGBA(image.Rect(0, 0, AvatarSize, AvatarSize))
		PaintBackGround(avatar, CalculateBGColor(nameBytes))
		DrawPattern(avatar, nameBytes, CalculatePixelColor(nameBytes))
		SavePNG(avatar, r.URL.Path[1:], r.URL.Query().Get("size"))
		var img image.Image = avatar
		writeImageWithTemplate(w, &img)

	}

}

var ImageTemplate string = `<!DOCTYPE html>
<html lang="en"><head></head>
<body><img src="data:image/jpg;base64,{{.Image}}"></body>`

// encodes an image 'img' in jpeg format and writes it into ResponseWriter using a template.
func writeImageWithTemplate(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Fatalln("unable to encode image.")
	}

	str := base64.StdEncoding.EncodeToString(buffer.Bytes())
	if tmpl, err := template.New("image").Parse(ImageTemplate); err != nil {
		log.Println("unable to parse image template.")
	} else {
		data := map[string]interface{}{"Image": str}
		if err = tmpl.Execute(w, data); err != nil {
			log.Println("unable to execute template.")
		}
	}
}

// encodes an image 'img' in jpeg format and writes it into ResponseWriter.
func writeImage(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

//set avatar size
func setAvatarSize(pixelSize int) {
	AvatarSize = pixelSize*10
	PixelSize = AvatarSize/10
}

//read size parameter
func setSize(w http.ResponseWriter, r *http.Request) {
	size := r.URL.Query().Get("size")
	fmt.Println(size)
	if (len(size) > 0){
		i, err := strconv.Atoi(size)
		fmt.Println(i)
		if (err != nil)  {
				// handle error
				fmt.Println(err)
				fmt.Fprintf(w, "Size should be integer.")
		}
		setAvatarSize(i)
	}else{
		setAvatarSize(20)
	}
}
