package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type inputFile struct {
	filepath  string
	separator string
}

type questionAnswer struct {
	Question string
	Answer   string
}

func main() {
	if len(os.Args) < 2 {
		exitGracefully(errors.New("A filepath argument is required"))
	}

	var separator string

	flag.StringVar(&separator, "separator", "comma", "Column separator")

	flag.Parse()

	data, err := getFileData(separator)
	if err != nil {
		exitGracefully(err)
	}

	fmt.Println("Filepath:", data.filepath)
	fmt.Println("Separator:", data.separator)

	questions, err := readCSVFile(data)
	if err != nil {
		exitGracefully(err)
	}

	score := 0
	for _, qa := range questions {
		fmt.Println("Question:", qa.Question)
		var answer string
		_, err := fmt.Scanln(&answer)
		if err != nil {
			exitGracefully(err)
		}

		if answer == qa.Answer {
			score++
		}
	}

	fmt.Println("Final score:", score, "/", len(questions))
}

func getFileData(separator string) (inputFile, error) {
	if !(separator == "comma" || separator == "semicolon") {
		return inputFile{}, errors.New("only comma or semicolon separators are allowed")
	}

	fileLocation := flag.Arg(0)

	return inputFile{fileLocation, separator}, nil
}

func readCSVFile(data inputFile) ([]questionAnswer, error) {
	file, err := os.Open(data.filepath)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			exitGracefully(err)
		}
	}()

	reader := csv.NewReader(file)
	reader.Comma = getSeparatorRune(data.separator)

	var questions []questionAnswer
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if len(record) != 2 {
			return nil, errors.New("invalid CSV format: each row should contain exactly two columns")
		}

		qa := questionAnswer{
			Question: record[0],
			Answer:   record[1],
		}
		questions = append(questions, qa)
	}

	return questions, nil
}

func getSeparatorRune(separator string) rune {
	switch separator {
	case "comma":
		return ','
	case "semicolon":
		return ';'
	default:
		return ','
	}
}

func exitGracefully(err error) {
	fmt.Println("Error:", err)
	os.Exit(1)
}
