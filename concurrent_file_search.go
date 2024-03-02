package main

import generatefileload "conccurent_File_searcher/GenerateFileLoad"

func main() {

	lg := generatefileload.LoadGenerator{
		FullFileName: "C:\\Users\\gupta\\concurrent_file_searcher_go\\LoadDirectory\\test.txt",
		RowNumbers:   100000,
	}

	lg.CreateFile()

	lg.WriteToFile()

}
