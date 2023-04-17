package mergesort

import (
	"bufio"
	"fmt"
	"os"
)

func Write(data [][]Pair) error {
	f, err := os.Create(InitInfo.OutputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	bw := bufio.NewWriter(f)

	for _, d := range data {
		for _, kv := range d {
			bw.Write([]byte(fmt.Sprintf("std-%d: %v\n", kv.Key, kv.Value)))
		}
	}
	err = bw.Flush()
	if err != nil {
		return err
	}
	return nil
}
