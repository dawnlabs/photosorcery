package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) == 1 {
		logAndExit(errors.New("no command given"))
	}

	switch command := os.Args[1]; command {
	case "merge":
		files, outputPath := parseMergeInput(os.Args[2:])
		merge(files, outputPath)
	case "convert":
		files, outputDir, fileType := parseConvertInput(os.Args[2:])
		convert(files, outputDir, fileType)
	default:
		logAndExit(errors.New("invalid command"))
	}
}

func parseConvertInput(args []string) ([]string, string, FileType) {
	fset := flag.NewFlagSet("fset", flag.ContinueOnError)
	typePtr := fset.String("type", "", "Target image type")
	outPtr := fset.String("out", "", "Directory to write to")

	fmt.Printf("args %v\n", args)

	fset.Parse(args)

	if *typePtr == "" || *outPtr == "" {
		logAndExit(errors.New("type or output dir not provided"))
	}

	outputPath, err := filepath.Abs(*outPtr)
	if err != nil {
		logAndExit(errors.New("invalid file path"))
	}

	outputStat, err := os.Stat(outputPath)
	if err != nil {
		logAndExit(errors.New("error getting outputdir stats"))
	}
	if !outputStat.Mode().IsDir() {
		logAndExit(errors.New("output path is not directory"))
	}

	targetType := getFileType(*typePtr)
	if targetType == ERR {
		logAndExit(errors.New("invalid target file type"))
	}

	files := formatFiles(args[4:])

	return files, outputPath, targetType
}

func parseMergeInput(args []string) ([]string, string) {
	fset := flag.NewFlagSet("fset", flag.ContinueOnError)
	outPtr := fset.String("out", "", "Directory to write to")

	fset.Parse(args)

	if *outPtr == "" {
		logAndExit(errors.New("no output dir specified"))
	}

	outputPath, err := filepath.Abs(*outPtr)
	if err != nil {
		logAndExit(errors.New("invalid file path"))
	}

	files := formatFiles(args[2:])

	return files, outputPath
}

func formatFiles(args []string) []string {
	cleanedFiles := make([]string, len(args))

	for index, file := range args {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			logAndExit(errors.New("invalid file"))
		}

		absPath, err := filepath.Abs(file)
		if err != nil {
			logAndExit(errors.New("invalid file"))
		}
		fileStat, err := os.Stat(absPath)
		if err != nil {
			logAndExit(errors.New("error checking input file stats"))
		}

		if !fileStat.Mode().IsRegular() {
			logAndExit(errors.New("input file is not a regular file"))
		}

		cleanedFiles[index] = absPath
	}

	return cleanedFiles
}
