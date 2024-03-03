package generatefileload

import (
	"fmt"
	"os"
)

type LoadGenerator struct {
	FullFileName string
	RowNumbers   int
}

func (lg *LoadGenerator) CreateFile() {
	if _, err := os.Stat(lg.FullFileName); os.IsNotExist(err) {
		file, err := os.Create(lg.FullFileName)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()
		fmt.Println("File created:", lg.FullFileName)
	} else if err != nil {
		fmt.Println("Error checking file existence:", err)
		return
	} else {
		fmt.Println("File already exists:", lg.FullFileName)
	}
}

func (lg *LoadGenerator) WriteToFile() {
	file, err := os.OpenFile(lg.FullFileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// first clear, then write
	err = file.Truncate(0)
	if err != nil {
		fmt.Println("Error truncating file:", err)
		return
	}

	// selectedRow := rand.Intn(lg.RowNumbers + 1)
	selectedRow := 8889
	fmt.Print("error will be on line: ")
	fmt.Println(selectedRow + 1)
	for i := 0; i < lg.RowNumbers; i++ {
		var err error
		if i == selectedRow {
			_, err = file.WriteString("quick brown fox jumps overr the lazy dog\n")
		} else {
			_, err = file.WriteString("quick brown fox jumps over the lazy dog\n")
		}
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
	fmt.Println("Content written to the file.")

}
