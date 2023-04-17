package usemergeconcurrent

import (
	"bufio"
	"fmt"
	"os"
)

func write(offset int64, data []Pair) {
	file, err := os.OpenFile(initInfo.OutputFile, os.O_WRONLY, 0644)
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
