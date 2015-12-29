package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"mime"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/codegangsta/cli"
	"github.com/nfnt/resize"
	"golang.org/x/image/tiff"
)

var outputWidth string
var outputHeight string
var outputFormat string
var outputSubfolder string
var inputOnlyFormat string

var currentDirectory string
var outputDirectory string

var isVerbose bool

func main() {
	app := cli.NewApp()

	app.Name = "Rsz!"
	app.Usage = "A little CLI image resizer"
	app.Author = "Francesco Malatesta"
	app.Version = "0.3.0"

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
			Name:        "to",
			Value:       "jpeg",
			Usage:       "Desired output format (available: jpg, png - Default: jpg)",
			Destination: &outputFormat,
		},
		cli.StringFlag{
			Name:        "in",
			Value:       "resized",
			Usage:       "Desired subfolder name for resized images. (examples: 'subfolder' or 'sub/subfolder')",
			Destination: &outputSubfolder,
		},
		cli.StringFlag{
			Name:        "only",
			Value:       "",
			Usage:       "Desired input format. If specified, other images with different format are going to be ignored.",
			Destination: &inputOnlyFormat,
		},
		cli.BoolFlag{
			Name:        "verbose",
			Usage:       "If specified, more info will be given during the process.",
			Destination: &isVerbose,
		},
	}

	app.Action = resizeCommand

	app.Run(os.Args)
}

func resizeCommand(c *cli.Context) {
	startTime := time.Now()

	fmt.Println("Rsz! Simple Image Resizer 0.3.0")

	convertedCounter := 0

	currentDirectory, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	outputDirectory = currentDirectory

	if outputSubfolder != "" {
		outputDirectory = outputDirectory + "/" + outputSubfolder
		os.MkdirAll(outputDirectory, 0777)
	}

	files, err := ioutil.ReadDir(currentDirectory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileName := file.Name()
		fileMimeType := mime.TypeByExtension(filepath.Ext(fileName))

		if imageTypeIsValid(fileMimeType) {
			convertedCounter++

			if isVerbose {
				fmt.Println("- Resizing: " + fileName)
			}

			resizeImage(fileName)
		}
	}

	if convertedCounter > 0 {
		fmt.Println("... completed! :}")
	} else {
		fmt.Println("No images found here! :{")
		os.Remove(outputDirectory)
	}

	if isVerbose && convertedCounter > 0 {
		fmt.Println("--------------------------------------------------------------------------------")
		fmt.Println("Converted Images: " + strconv.Itoa(convertedCounter))
		fmt.Println("Destination Folder: " + outputDirectory)
		fmt.Println("Destination Format: " + outputFormat)
		fmt.Printf("Completed in: %.2f seconds\n", time.Since(startTime).Seconds())
		fmt.Println("--------------------------------------------------------------------------------")
	}
}

func resizeImage(fileName string) {
	imageFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	imageData := decodeInputImageFile(imageFile)

	oWidth, err := strconv.Atoi(outputWidth)
	oHeight, err := strconv.Atoi(outputHeight)
	if err != nil {
		log.Fatal(err)
	}

	if oWidth == 0 && oHeight == 0 {
		oWidth = 1024
	}

	resizedImageData := resize.Resize(uint(oWidth), uint(oHeight), imageData, resize.Lanczos2)

	encodeImageOnOutputFile(resizedImageData, fileName)
}

func decodeInputImageFile(imageFile *os.File) image.Image {
	imageData, _, err := image.Decode(imageFile)
	if err != nil {
		log.Fatal(err)
	}

	imageFile.Close()

	return imageData
}

func encodeImageOnOutputFile(imageData image.Image, fileName string) {
	resizedFileName := "resized-" + fileName[0:len(fileName)-len(filepath.Ext(fileName))]
	output, err := os.Create(outputDirectory + "/" + resizedFileName + "." + outputFormat)
	if err != nil {
		log.Fatal(err)
	}

	if outputFormat == "jpeg" {
		jpeg.Encode(output, imageData, nil)
	}

	if outputFormat == "png" {
		png.Encode(output, imageData)
	}

	if outputFormat == "tiff" {
		tiff.Encode(output, imageData, nil)
	}

	output.Close()
}

func imageTypeIsValid(mimeType string) bool {
	if inputOnlyFormat != "" {
		if mimeType == "image/"+inputOnlyFormat {
			return true
		}

		return false
	}

	if mimeType == "image/jpeg" {
		return true
	}

	if mimeType == "image/png" {
		return true
	}

	if mimeType == "image/tiff" {
		return true
	}

	return false
}
