package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)

const (
	inputFile     = "measurements.txt"
	cpuProfileOut = "cpuProfile.out"
)

func main() {
	cpuProfileFile, err := os.Create(cpuProfileOut)
	if err != nil {
		fmt.Printf("error creating cpu Profile file: %v", err)
	}
	defer cpuProfileFile.Close()

	pprof.StartCPUProfile(cpuProfileFile)
	defer pprof.StopCPUProfile()

	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("error opening input file: %s, err: %v", inputFile, err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	regionMaps := make(map[string][]float64)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("error reading lines from the input file: %v", err)
		}
		station, weather, err := parseStationAndWeather(line)
		if err != nil {
			fmt.Printf("error processing file: %v", err)
			return
		}

		regionMaps[station] = append(regionMaps[station], weather)
	}

	fmt.Println(regionMaps)

}

func parseStationAndWeather(line string) (string, float64, error) {
	line = strings.TrimRight(line, "\n")
	tokens := strings.Split(line, ";")
	if len(tokens) < 2 {
		return "", 0.0, fmt.Errorf("error parsing line: %s", line)
	}

	weather, err := strconv.ParseFloat(tokens[1], 64)
	if err != nil {
		return "", 0.0, fmt.Errorf("error parsing float from line: %s", line)
	}

	return tokens[0], weather, nil
}
