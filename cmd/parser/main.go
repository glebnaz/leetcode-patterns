package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/gocarina/gocsv"
	"io"
	"os"
	"sort"
)

type Pattern struct {
	Id         int      `json:"id" csv:"id"`
	Name       string   `json:"name" csv:"name"`
	Url        string   `json:"url" csv:"url"`
	Pattern    []string `json:"pattern" csv:"pattern"`
	Difficulty string   `json:"difficulty" csv:"difficulty"`
	Premium    bool     `json:"premium" csv:"premium"`
	Companies  []string `json:"companies" csv:"companies"`
}

func (p Pattern) PatternToString() string {
	str := ""

	for i, v := range p.Pattern {
		str += str + v + ","
		if i == len(p.Pattern)-1 {
			str = str + v
		}
	}

	return str
}

func (p Pattern) CompaniesToString() string {
	str := ""

	for i, v := range p.Companies {
		str = str + v + ","
		if i == len(p.Companies)-1 {
			str = str + v
		}
	}

	return str
}

type PatternCSV struct {
	Id         int    `csv:"-"`
	Name       string `csv:"Name"`
	Url        string `csv:"URL"`
	Pattern    string `csv:"Pattern"`
	Difficulty string `csv:"Difficulty"`
	Premium    bool   `csv:"Premium"`
	Companies  string `csv:"Companies"`
}

func main() {
	var data []Pattern

	file, err := os.Open("cmd/parser/data.json")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].DifficultyToInt() < data[j].DifficultyToInt()
	})

	var csvData = make([]PatternCSV, len(data))

	for i := range data {
		d := data[i]

		csvData[i] = PatternCSV{
			Id:         d.Id,
			Name:       d.Name,
			Url:        d.Url,
			Pattern:    d.PatternToString(),
			Difficulty: d.Difficulty,
			Premium:    d.Premium,
			Companies:  d.CompaniesToString(),
		}
	}

	fileCSV, err := os.Create("cmd/parser/data.csv")
	if err != nil {
		panic(err)
	}

	defer fileCSV.Close()

	gocsv.SetCSVWriter(func(out io.Writer) *gocsv.SafeCSVWriter {
		writer := csv.NewWriter(out)
		writer.Comma = ';'
		return gocsv.NewSafeCSVWriter(writer)
	})

	csvBytes, err := gocsv.MarshalBytes(csvData)

	n, err := fileCSV.Write(csvBytes)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d bytes written\n", n)
}

func (p *Pattern) DifficultyToInt() int {
	switch p.Difficulty {
	case "Easy":
		return 1
	case "Medium":
		return 2
	case "Hard":
		return 3
	default:
		return 0
	}
}
