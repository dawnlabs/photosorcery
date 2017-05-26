package main

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"sync"
	"strings"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

type FileType int

const (
	PNG FileType = iota
	JPG
	GIF
	BMP
	TIFF
	ERR
)

func getFileType(input string) FileType {
	switch input {
	case "jpg":
		fallthrough
	case "jpeg":
		return JPG
	case "png":
		return PNG
	case "gif":
		return GIF
	case "bmp":
		return BMP
	case "tiff":
		return TIFF
	default:
		return ERR
	}
}

func getFileExtension(input FileType) string {
	switch input {
	case JPG:
		return "jpg"
	case PNG:
		return "png"
	case GIF:
		return "gif"
	case BMP:
		return "bmp"
	case TIFF:
		return "tiff"
	default:
		return ""
	}
}

func convert(files []string, outputDir string, fileType FileType) {
	var wg sync.WaitGroup
	for _, currPath := range files {
		wg.Add(1)
		go convertFile(&wg, currPath, outputDir, fileType)
	}

	wg.Wait()
}

func convertFile(wg *sync.WaitGroup, currPath string, outputDir string, fileType FileType) {
	// call done when finished
	defer wg.Done()

	ext := strings.ToLower(filepath.Ext(currPath))
	newExt := getFileExtension(fileType)

	_, filename := filepath.Split(currPath)
	filenameNoExt := filename[0 : len(filename)-len(ext)]
	newFileName := filenameNoExt + "." + newExt
	newFilePath := outputDir + "/" + newFileName

	// validate file type
	startType := getFileType(ext[1:])
	if startType == ERR {
		logAndExit(errors.New("input file type not valid"))
	}

	// open files
	file, err := os.Open(currPath)
	if err != nil {
		logAndExit(err)
	}
	defer file.Close()

	outFile := openOrCreate(newFilePath)
	defer outFile.Close()

	// decode
	imageData, _, err := image.Decode(file)
	if err != nil {
		logAndExit(errors.New("error decoding image"))
	}

	// encode in new type
	switch fileType {
	case JPG:
		err := jpeg.Encode(outFile, imageData, nil)
		if err != nil {
			logAndExit(errors.New("error converting to jpeg"))
		}
	case PNG:
		err := png.Encode(outFile, imageData)
		if err != nil {
			logAndExit(err)
		}
	case GIF:
		err := gif.Encode(outFile, imageData, nil)
		if err != nil {
			logAndExit(errors.New("error converting to gif"))
		}
	case BMP:
		err := bmp.Encode(outFile, imageData)
		if err != nil {
			logAndExit(errors.New("error converting to bmp"))
		}
	case TIFF:
		err := tiff.Encode(outFile, imageData, nil)
		if err != nil {
			logAndExit(errors.New("error converting to tiff"))
		}
	}
}

func openOrCreate(filename string) *os.File {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			logAndExit(errors.New("error creating output file"))
		}
		return file
	} else {
		file, err := os.Open(filename)
		if err != nil {
			logAndExit(errors.New("error opening output file"))
		}
		return file
	}
}
