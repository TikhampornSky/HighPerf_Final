package mergesort

import (
	"math"
	"sync"
)

var InitInfo InitInfoS

func Init(input string, output string, core int64) {
	InitInfo.InputFile = input
	InitInfo.OutputFile = output
	InitInfo.CPU = core
}

func Run() {
	result := make([][]Pair, InitInfo.CPU)

	dataBlock := Read()

	min, max := MinMaxOfAllBlock(dataBlock)
	numMergeBlock := (max - min) / float64(InitInfo.CPU)

	var start float64 = min
	var stop float64 = min + numMergeBlock

	wg1 := sync.WaitGroup{}
	for i := 0; i < int(InitInfo.CPU); i++ {
		if i == int(InitInfo.CPU)-1 {
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

	Write(result)
}
