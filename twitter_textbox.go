package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-gota/gota/dataframe"
	"github.com/machinebox/sdk-go/textbox"
)


func main() {
	textboxIP := "http://localhost:8080"
	textboxClient := textbox.New(textboxIP)

	file, err := os.Open("./data/results.csv")
	if err != nil {
		fmt.Println("An error occured loading the CSV data: ", err)
	}
	defer file.Close()
	frame := dataframe.ReadCSV(file)
	comments := frame.Col("text").Records()

	sentiments := []float64{}
	var total float64
	for _, comment := range comments {
		analysis, err := textboxClient.Check(strings.NewReader(comment))
		if err != nil {
			fmt.Println("we hit an error: ", err)
			continue
		}

		sentimentTotal := 0.0
		for _, sentence := range analysis.Sentences {
			sentimentTotal += sentence.Sentiment
		}
		sentiment := sentimentTotal/float64(len(analysis.Sentences))

		total = total + sentiment
		sentiments = append(sentiments, sentiment)
	}

	// fmt.Println("Sentiments: ", sentiments)
	fmt.Printf("\nSentiment Average: %0.2f\n\n", total/float64(len(sentiments)))
}