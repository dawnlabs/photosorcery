package main

import (
        "os"
        "image"
        "github.com/jung-kurt/gofpdf"
)

func merge (files []string, outputPath string) {
        pdf := gofpdf.New("P", "mm", "A4", "")

        for _, file := range files {
                width, height := getImageDimension(file)
                size := gofpdf.SizeType{ Wd: width, Ht: height }
                pdf.AddPageFormat("P", size)

                pdf.ImageOptions(
                        file, 0, 0, width, height, false,
                        gofpdf.ImageOptions{ ReadDpi: true },
                        0, "")
        }

        err := pdf.OutputFileAndClose(outputPath)
        if err != nil {
                logAndExit(err)
        }
}

func getImageDimension(imagePath string) (float64, float64) {
        file, err := os.Open(imagePath)
        if err != nil {
                logAndExit(err)
        }

        image, _, err := image.DecodeConfig(file)
        if err != nil {
                logAndExit(err)
        }

        return float64(image.Width), float64(image.Height)
}
