package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-gota/gota/dataframe"
	"github.com/machinebox/sdk-go/textbox"
)


func main() {
	textboxIP := "http://157.230.171.113:8080"
	textboxClient := textbox.New(textboxIP)
	testStatement := "I've seen things you people wouldn't believe. Attack ships on fire off the shoulder of Orion. I watched C-beams glitter in the dark near the Tannhäuser Gate. All those moments will be lost in time, like tears in rain. Time to die."

	file, err := os.Open("./data/results.csv")
	if err != nil {
		fmt.Println("An error occured loading the CSV data: ", err)
	}
	defer file.Close()

	frame := dataframe.ReadCSV(file)
	comments := frame.Col("text").String()

	for _, comment := range comments {
		fmt.Println("COMMENT: ", comment)
	}

	analysis, err := textboxClient.Check(strings.NewReader(testStatement))
	if err != nil {
		fmt.Println("An error occured determining the string statement: ", err)
	}

	// Calculate the sentiment.
	sentimentTotal := 0.0
	for _, sentence := range analysis.Sentences {
		sentimentTotal += sentence.Sentiment
	}

	fmt.Printf("\nSentiment: %0.2f\n\n", sentimentTotal/float64(len(analysis.Sentences)))
}