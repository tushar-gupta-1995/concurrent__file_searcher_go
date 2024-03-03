package main

import (
	concurrentsearch "conccurent_File_searcher/ConcurrentSearch"
	generatefileload "conccurent_File_searcher/GenerateFileLoad"
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	lg := generatefileload.LoadGenerator{
		FullFileName: "C:\\Users\\gupta\\concurrent_file_searcher_go\\LoadDirectory\\test.txt",
		RowNumbers:   10000,
	}

	lg.CreateFile()

	lg.WriteToFile()

	workerThreads := 100

	foundSignal := make(chan struct{}, workerThreads)
	lineChannel := make(chan concurrentsearch.LineDetails, workerThreads)

	wg.Add(workerThreads + 1)

	cs := concurrentsearch.ConcurrentSearchStruct{
		FileNameFull: lg.FullFileName,
		OutputChan:   lineChannel,
		FoundSignal:  foundSignal,
	}

	go cs.WriteToChannel(&wg)

	for i := 1; i <= workerThreads; i++ {
		go cs.ReadChannel(i, &wg)
	}

	wg.Wait()

	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("Time elapsed:", elapsed)
}
