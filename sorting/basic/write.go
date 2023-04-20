package basic

import (
	"bufio"
	"fmt"
	"os"
)

func Write(data []Pair) error {
	f, err := os.Create(InitInfo.OutputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	bw := bufio.NewWriter(f)

	for _, kv := range data {
		bw.Write([]byte(fmt.Sprintf("std-%d: %v\n", kv.Key, kv.Value)))
	}
	err = bw.Flush()
	if err != nil {
		return err
	}
	return nil
}
