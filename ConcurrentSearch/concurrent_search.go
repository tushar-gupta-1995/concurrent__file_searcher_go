package concurrentsearch

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type LineDetails struct {
	LineNo   int
	RowValue string
}

type ConcurrentSearchStruct struct {
	FoundSignal  chan struct{}
	FileNameFull string
	OutputChan   chan LineDetails
}

func (css *ConcurrentSearchStruct) WriteToChannel(wg *sync.WaitGroup) {

	defer wg.Done()
	defer close(css.OutputChan)
	file, err := os.Open(css.FileNameFull)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		select {
		case <-css.FoundSignal:
			fmt.Println("stopping writing to channel")
			return
		default:
			// fmt.Println("writing to channel")
			v := LineDetails{
				LineNo:   i,
				RowValue: scanner.Text(),
			}
			// fmt.Println(v)
			css.OutputChan <- v
			i = i + 1
		}
	}

}

func (css *ConcurrentSearchStruct) ReadChannel(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for rowDetail := range css.OutputChan {
		// fmt.Println("worker at your service: ")
		// fmt.Println(id)
		rowValue := rowDetail.RowValue
		if strings.Contains(rowValue, "overr") {
			fmt.Print("found at line: ")
			fmt.Println(rowDetail.LineNo)
			css.FoundSignal <- struct{}{}
		}
	}
}
