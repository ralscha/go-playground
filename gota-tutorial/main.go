package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"log"
	"os"
	"strings"
)

func series1() {
	fmt.Println(series.New([]string{"z", "y", "d", "e"}, series.String, "col"))

	a := map[string]series.Type{
		"A": series.String,
		"D": series.Bool,
	}

	fmt.Println(a)
}

func dataframe1() {
	df := dataframe.New(
		series.New([]string{"a", "b", "c", "d", "e"}, series.String, "alphas"),
		series.New([]int{5, 4, 2, 3, 1}, series.Int, "numbers"),
		series.New([]string{"a1", "b2", "c3", "d4", "e5"}, series.String, "alnums"),
		series.New([]bool{true, false, true, true, false}, series.Bool, "state"),
	)

	fmt.Println(df)
}

func dataframe2() {
	type Dog struct {
		Name       string
		Color      string
		Height     int
		Vaccinated bool
	}

	dogs := []Dog{
		{"Buster", "Black", 56, false},
		{"Jake", "White", 61, false},
		{"Bingo", "Brown", 50, true},
		{"Gray", "Cream", 68, false},
	}

	dogsDf := dataframe.LoadStructs(dogs)

	fmt.Println(dogsDf)
}

func dataframe3() {
	type Dog struct {
		Name       string
		Color      string
		Height     int
		Vaccinated bool
	}

	dogs := []Dog{
		{"Buster", "Black", 56, false},
		{"Jake", "White", 61, false},
		{"Bingo", "Brown", 50, true},
		{"Gray", "Cream", 68, false},
	}

	dogsDf := dataframe.LoadStructs(dogs)

	fmt.Println(dogsDf.Dims())
	fmt.Println(dogsDf.Types())
	fmt.Println(dogsDf.Names())
	fmt.Println(dogsDf.Nrow())
	fmt.Println(dogsDf.Ncol())
}

func dataframe4() {
	type Dog struct {
		Name       string
		Color      string
		Height     int
		Vaccinated bool
	}

	dogs := []Dog{
		{"Buster", "Black", 56, false},
		{"Jake", "White", 61, false},
		{"Bingo", "Brown", 50, true},
		{"Gray", "Cream", 68, false},
	}

	dogsDf := dataframe.LoadStructs(dogs)

	col := dogsDf.Col("Height") // Selects a column
	fmt.Println(col.IsNaN())
	fmt.Println(col.Mean())
	fmt.Println(col.Copy())
	fmt.Println(col.HasNaN())
	fmt.Println(col.Records())
}

func dataframe5() {
	jsonString := `[
  {
    "Name": "John",
    "Age": 44,
    "Favorite Color": "Red",
    "Height(ft)": 6.7
  },
  {
    "Name": "Mary",
    "Age": 40,
    "Favorite Color": "Blue",
    "Height(ft)": 5.7
  }
]`

	jsonDf := dataframe.ReadJSON(strings.NewReader(jsonString))
	fmt.Println(jsonDf)
}

func dataframe6() {
	csvString := `
Name,Age,Favorite Color,Height(ft)
John,44,Red,6.7
Mary,40,Blue,5.7`

	csvDf := dataframe.ReadCSV(strings.NewReader(csvString))
	fmt.Println(csvDf)
}

func dataframe7() {
	file, err := os.Open("stats.csv")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	df := dataframe.ReadCSV(file)

	fmt.Println(df)
}

func dataframe8() {
	file, err := os.Open("stats.json")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	df := dataframe.ReadJSON(file)

	fmt.Println(df)
}

func dataframe9() {
	file, err := os.Open("stats.csv")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	df := dataframe.ReadCSV(file)

	// This selects the first two rows of the DataFrame
	rows := df.Subset([]int{0, 2})

	fmt.Println(rows)
}

func dataframe10() {
	file, err := os.Open("stats.csv")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	df := dataframe.ReadCSV(file)

	// This selects the first two columns of a DataFrame
	firstTwoColumns := df.Select([]int{0, 2})

	// This selects columns of a DataFrame by name
	namedColumns := df.Select([]string{"Name", "Favorite Color"})

	fmt.Println(firstTwoColumns)
	fmt.Println(namedColumns)
}

func dataframe11() {
	file, err := os.Open("stats.csv")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	df := dataframe.ReadCSV(file)

	updatedDf := df.Set(
		[]int{0, 3},
		dataframe.LoadRecords(
			[][]string{
				[]string{"Jenny", "23", "Purple", "2.2"},
				[]string{"Jesse", "34", "Indigo", "3.5"},
				[]string{"Peter", "33", "Violet", "3.3"},
			},
		),
	)

	fmt.Println(updatedDf)
}

func dataframe12() {
	df := dataframe.New(
		series.New([]string{"a", "b", "c", "d", "e"}, series.String, "alphas"),
		series.New([]int{5, 4, 2, 3, 1}, series.Int, "numbers"),
		series.New([]string{"a1", "b2", "c3", "d4", "e5"}, series.String, "alnums"),
		series.New([]bool{true, false, true, true, false}, series.Bool, "state"),
	)

	// Filter out b from the column named alphas
	fil := df.Filter(
		dataframe.F{Colname: "alphas", Comparator: series.Eq, Comparando: "b"},
	)

	fmt.Println(fil)
}

func dataframe13() {
	df := dataframe.New(
		series.New([]string{"a", "b", "c", "d", "e"}, series.String, "alphas"),
		series.New([]int{5, 4, 2, 3, 1}, series.Int, "numbers"),
		series.New([]string{"a1", "b2", "c3", "d4", "e5"}, series.String, "alnums"),
		series.New([]bool{true, false, true, true, false}, series.Bool, "state"),
	)

	sortedAscending := df.Arrange(
		dataframe.Sort("numbers"),
	)

	fmt.Println(sortedAscending)
}

func dataframe14() {
	file, err := os.Open("stats.csv")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	df := dataframe.ReadCSV(file)

	categorize := df.GroupBy("Name", "Age")

	fmt.Println(categorize)
}

func dataframe15() {
	dfA := dataframe.New(
		series.New([]string{"a", "b", "c", "d", "e"}, series.String, "alphas"),
		series.New([]int{5, 4, 2, 3, 1}, series.Int, "numbers"),
		series.New([]string{"a1", "b2", "c3", "d4", "e5"}, series.String, "alnums"),
		series.New([]bool{true, false, true, true, false}, series.Bool, "state"),
	)
	dfB := dataframe.New(
		series.New([]string{"f", "g", "h", "i", "j"}, series.String, "alphas"),
		series.New([]int{1, 2, 3, 4, 5}, series.Int, "numbers"),
		series.New([]string{"f6", "g7", "h8", "i9", "j10"}, series.String, "alnums"),
		series.New([]bool{false, true, false, false, true}, series.Bool, "state"),
	)

	leftJoin := dfA.RightJoin(dfB, "state")

	fmt.Println(leftJoin)
}

func dataframe16() {
	df := dataframe.New(
		series.New([]int{95, 74, 59, 82, 87}, series.Int, "score"),
	)

	mean := func(s series.Series) series.Series {
		floats := s.Float()
		sum := 0.0
		for _, f := range floats {
			sum += f
		}
		return series.Floats(sum / float64(len(floats)))
	}

	meanScore := df.Capply(mean)

	fmt.Println(meanScore)
}

func dataframe17() {
	df := dataframe.New(
		series.New([]string{"a", "b", "c", "d", "e"}, series.String, "alphas"),
		series.New([]int{5, 4, 2, 3, 1}, series.Int, "numbers"),
		series.New([]string{"a1", "b2", "c3", "d4", "e5"}, series.String, "alnums"),
		series.New([]bool{true, false, true, true, false}, series.Bool, "state"),
	)

	description := df.Describe()

	fmt.Println(description)
}

func dataframe18() {
	csvString := `
Name,Age,Favorite Color,Height(ft)
John,44,Red,6.7
Mary,40,Blue,5.7
Paul,27,green,5.6`

	df := dataframe.ReadCSV(strings.NewReader(csvString))

	updatedDf := df.Set(
		[]int{0, 2},
		dataframe.LoadRecords(
			[][]string{
				{"Jenny", "23", "Purple", "2.2"},
				{"Jesse", "34", "Indigo", "3.5"},
				{"Peter", "33", "Violet", "3.3"},
			},
		),
	)

	file, err := os.Create("output.csv")
	if err != nil {
		log.Fatal(err)
	}

	err = updatedDf.WriteCSV(file)
	if err != nil {
		log.Fatal(err)
	}
}

func dataframe19() {
	csvString := `
Name,Age,Favorite Color,Height(ft)
John,44,Red,6.7
Mary,40,Blue,5.7
Paul,27,green,5.6`

	df := dataframe.ReadCSV(strings.NewReader(csvString))

	updatedDf := df.Set(
		[]int{0, 2},
		dataframe.LoadRecords(
			[][]string{
				{"Jenny", "23", "Purple", "2.2"},
				{"Jesse", "34", "Indigo", "3.5"},
				{"Peter", "33", "Violet", "3.3"},
			},
		),
	)

	file, err := os.Create("output.json")
	if err != nil {
		log.Fatal(err)
	}

	err = updatedDf.WriteJSON(file)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// series1()
	// dataframe1()
	// dataframe2()
	// dataframe3()
	// dataframe4()
	// dataframe5()
	// dataframe6()
	// dataframe7()
	// dataframe8()
	// dataframe9()
	// dataframe10()
	// dataframe11()
	// dataframe12()
	// dataframe13()
	// dataframe14()
	// dataframe15()
	// dataframe16()
	// dataframe17()
	// dataframe18()
	dataframe19()
}
