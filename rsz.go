package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/nfnt/resize"
)

var outputWidth string
var outputHeight string
var outputFormat string

func main() {
	app := cli.NewApp()

	app.Name = "Rsz!"
	app.Usage = "A little CLI image resizer"
	app.Author = "Francesco Malatesta"
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "width",
			Value:       "0",
			Usage:       "Desired width for resize. Optional.",
			Destination: &outputWidth,
		},
		cli.StringFlag{
			Name:        "height",
			Value:       "0",
			Usage:       "Desired height for resize. Optional.",
			Destination: &outputHeight,
		},
		cli.StringFlag{
			Name:        "format",
			Value:       "jpg",
			Usage:       "Desired output format (available: jpg, png - Default: jpg)",
			Destination: &outputFormat,
		},
	}

	app.Action = func(c *cli.Context) {

		fmt.Println("Rsz!")

		currentDirectory, err := filepath.Abs(filepath.Dir(os.Args[0]))

		if err != nil {
			log.Fatal(err)
		}

		files, err := ioutil.ReadDir(currentDirectory)

		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			fileName := file.Name()
			fileExtension := filepath.Ext(fileName)

			if fileExtension == ".jpg" || fileExtension == ".png" {
				resizeImage(fileName, fileExtension, currentDirectory)
			}
		}

		fmt.Println("... complete! :}")
	}

	app.Run(os.Args)
}

func resizeImage(fileName string, format string, outputDirectory string) {
	imageFile, error := os.Open(fileName)

	if error != nil {
		log.Fatal(error)
	}

	var imageStream image.Image

	if format == ".jpg" {
		imageStream = openJPGImage(imageFile)
	} else {
		imageStream = openPNGImage(imageFile)
	}

	imageFile.Close()

	oWidth, error := strconv.Atoi(outputWidth)
	oHeight, error := strconv.Atoi(outputHeight)

	if error != nil {
		log.Fatal(error)
	}

	if oWidth == 0 && oHeight == 0 {
		oWidth = 1024
	}

	resizedImageData := resize.Resize(uint(oWidth), uint(oHeight), imageStream, resize.Lanczos2)
	resizedFileName := "resized-" + strconv.Itoa(resizedImageData.Bounds().Dx()) + "x" + strconv.Itoa(resizedImageData.Bounds().Dy()) + "-" + fileName[0:len(fileName)-len(format)]

	output, error := os.Create(outputDirectory + "/" + resizedFileName + "." + outputFormat)
	if error != nil {
		log.Fatal(error)
	}

	if outputFormat == "jpg" {
		jpeg.Encode(output, resizedImageData, nil)
	} else {
		png.Encode(output, resizedImageData)
	}

	output.Close()

	fmt.Println("- Resized: " + resizedFileName + "." + outputFormat)
}

func openJPGImage(imageFile *os.File) image.Image {
	imageStream, error := jpeg.Decode(imageFile)
	if error != nil {
		log.Fatal(error)
	}

	return imageStream
}

func openPNGImage(imageFile *os.File) image.Image {
	imageStream, error := png.Decode(imageFile)
	if error != nil {
		log.Fatal(error)
	}

	return imageStream
}
