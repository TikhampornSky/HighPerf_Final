package mergesort

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func Write(result [][]Pair) {
	file, err := os.Create(InitInfo.OutputFile)
	if err != nil {
		panic(err)
	}
	file.Close()

	wg2 := sync.WaitGroup{}
	wg3 := sync.WaitGroup{}

	blockSize := make([]int64, InitInfo.CPU)
	for i := 0; i < int(InitInfo.CPU); i++ {
		wg2.Add(1)
		go func(i int) {
			for _, kv := range result[i] {
				blockSize[i] += int64(len([]byte(fmt.Sprintf("std-%d: %v\n", kv.Key, kv.Value))))
			}
			wg2.Done()
		}(i)
	}
	wg2.Wait()

	var offset int64
	for i := 0; i < int(InitInfo.CPU); i++ {
		wg3.Add(1)
		go func(offset int64, data []Pair) {
			writeOutput(offset, data)
			wg3.Done()
		}(offset, result[i])
		offset += blockSize[i]
	}
	wg3.Wait()
	
}

func writeOutput(offset int64, data []Pair) {
	file, err := os.OpenFile(InitInfo.OutputFile, os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Seek(offset, 0)
	writer := bufio.NewWriter(file)

	for _, kv := range data {
		writer.Write([]byte(fmt.Sprintf("std-%d: %v\n", kv.Key, kv.Value)))
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}
