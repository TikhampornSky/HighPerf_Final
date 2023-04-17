package usemergeconcurrent

import (
	"fmt"
	"math"
	"os"
	"sync"
)

var initInfo InitInfo

func Init(input string, output string, core int64) {
	initInfo.InputFile = input
	initInfo.OutputFile = output
	initInfo.CPU = core
}

func Run() {
	dataBlock := ReadInput()
	min, max := MinMaxOfAllBlock(dataBlock)
	numMergeBlock := (max - min) / float64(initInfo.CPU)
	var start, stop float64
	start = min
	stop = min + numMergeBlock
	result := make([][]Pair, initInfo.CPU)

	wg1 := sync.WaitGroup{}
	for i := 0; i < int(initInfo.CPU); i++ {
		if i == int(initInfo.CPU)-1 {
			stop = max
		}
		wg1.Add(1)
		go func(i int, start float64, stop float64) {
			result[i] = MergeInBlock(start, stop, dataBlock)
			wg1.Done()
		}(i, start, stop)
		start = stop + math.SmallestNonzeroFloat64
		stop = start + numMergeBlock
	}

	wg1.Wait()

	wg3 := sync.WaitGroup{}
	blockSize := make([]int64, initInfo.CPU)
	for i := 0; i < int(initInfo.CPU); i++ {
		wg3.Add(1)
		go func(i int) {
			for _, kv := range result[i] {
				blockSize[i] += int64(len([]byte(fmt.Sprintf("std-%d: %v\n", kv.Key, kv.Value))))
			}
			wg3.Done()
		}(i)
	}

	wg3.Wait()

	// write 
	file, err := os.Create(initInfo.OutputFile)
	if err != nil {
		panic(err)
	}
	file.Close()

	wg2 := sync.WaitGroup{}
	var offset int64
	for i := 0; i < int(initInfo.CPU); i++ {
		wg2.Add(1)
		go func(offset int64, data []Pair) {
			write(offset, data)
			wg2.Done()
		}(offset, result[i])
		offset += blockSize[i]
	}

	wg2.Wait()
}
