package main

import (
	"os"
	"log"
	"image"
	"fmt"
    "github.com/oliamb/cutter"
			"gopkg.in/alecthomas/kingpin.v2"
	"path"
	"image/jpeg"
	"image/png"
)

var (
	app = kingpin.New("Tiler", "A tool for making tiles from an image.")
	verbose = app.Flag("verbose", "Verbose mode.").Short('v').Bool()
	tileSize = app.Flag("tile", "Tile size (default 256).").Default("256").Int()
	outputFormat = app.Flag("format", "Image output format.").Short('f').Default("png").HintOptions("jpeg", "png").String()
	inputFile = app.Arg("input", "Input image file path (accepted jpeg, png).").Required().ExistingFile()
	outputDir = app.Arg("output", "Output directory.").Default(".").ExistingDir()

)

func main() {

	kingpin.MustParse(app.Parse(os.Args[1:]))

	switch *outputFormat {
	case "png","jpeg":
		break
	default:
		app.Fatalf("Format not valid (possible values: jpeg, png).")
	}

	f, err := os.Open(*inputFile)
	if err != nil {
		app.Fatalf("Cannot open file", err)
	}
	defer f.Close()

	if *verbose {
		fmt.Printf("Decoding image %s\n", *inputFile)
	}

	img, _, err := image.Decode(f)
	if err != nil {
		app.Fatalf("Cannot decode image:", err)
	}

	width, height := img.Bounds().Max.X, img.Bounds().Max.Y

	if *verbose {
		fmt.Printf("Cropping image (size %dx%d)\n", width, height)
	}

	nColumns := width/ *tileSize
	nRows := height/ *tileSize

	var fileName string

	for x:= 0; x < nColumns; x++ {
		for y := 0; y < nRows; y++ {
			img, err := cutter.Crop(img, cutter.Config{
				Height:  *tileSize,
				Width:   *tileSize,
				Mode:    cutter.TopLeft,
				Anchor:  image.Point{x* *tileSize, y* *tileSize},
				Options: 0,
			})

			fileName = path.Join(*outputDir, fmt.Sprintf("arazzo_%d_%d.%s", x, y, *outputFormat))

			outImgFile, err := os.Create(fileName)
			if err != nil {
				app.Fatalf("Cannot create file %s: %s",fileName, err)
			}

			switch *outputFormat {
			case "png":
				err = png.Encode(outImgFile, img)
			case "jpeg":
				err = jpeg.Encode(outImgFile, img, &jpeg.Options{jpeg.DefaultQuality})
			default:
				app.Fatalf("Format not valid (possible values: jpeg, png).")
			}
			if err != nil {
				app.Fatalf("Could not save image to %s: %s", fileName, err)
			}

			outImgFile.Close()
		}
	}

	if *verbose {
		log.Print("Done. Saved tiles to ", *outputDir)
	}
}
