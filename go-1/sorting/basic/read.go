package basic

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Read() []Pair {
	var data []Pair

	file, err := os.Open(InitInfo.InputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
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

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	
	return data

}
