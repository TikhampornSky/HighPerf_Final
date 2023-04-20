package mergesort

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var chunkSize int64

func Read() [][]Pair {
	file, err := os.Open(InitInfo.InputFile)
	if err != nil {
		panic(err)
	}
	fileSize, err := file.Stat()
	if err != nil {
		panic(err)
	}
	file.Close()

	blockData := make([][]Pair, InitInfo.CPU)
	var wg sync.WaitGroup
	wg.Add(int(InitInfo.CPU))

	chunkSize = fileSize.Size() / int64(InitInfo.CPU)
	var offset int64 = 0
	for i := 0; i < int(InitInfo.CPU); i++ {

		go func(offset int64, i int) {
			blockData[i] = readFile(offset)
			wg.Done()
		}(offset, i)

		offset += int64(chunkSize)
	}

	wg.Wait()

	return blockData
}

func readFile(offset int64) []Pair {
	var data []Pair

	file, err := os.Open(InitInfo.InputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err := file.Seek(offset, 0); err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(file)

	var current int64 = 0
	for current < chunkSize {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		if line[0] == 's' {
			group := strings.Split(line, ": ")
			index := group[0][4:]
			val := group[1][:len(group[1])-1]
			key, err := strconv.ParseInt(index, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			value, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Fatal(err)
			}
			data = append(data, Pair{Key: key, Value: value})
		}
		current += int64(len([]byte(line)))
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].Value < data[j].Value
	})

	return data
}
