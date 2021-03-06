package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const NumOfBanks = 16

func main() {
	inp := bufio.NewScanner(os.Stdin)
	banks := [NumOfBanks]int{}
	// Keep track of previous configurations and the cycle they occur in
	prev := make(map[string]int)

	inp.Scan()
	for i, a := range strings.Fields(inp.Text()) {
		blocks, _ := strconv.Atoi(a)
		banks[i] = blocks
	}

	cycles := 0
	diff := 0
	for {
		// Validate and store bank configuration
		var config bytes.Buffer
		for _, blocks := range banks {
			a := strconv.Itoa(blocks)
			config.WriteString(a)
		}

		if o, ok := prev[config.String()]; ok {
			diff = cycles - o
			break
		}
		prev[config.String()] = cycles

		index := getIndexOfBankWithMostBlocks(banks)
		blocks := banks[index]
		banks[index] = 0

		// Distribute blocks
		for blocks != 0 {
			index++
			if index == NumOfBanks {
				index = 0
			}

			banks[index]++
			blocks--
		}
		cycles++
	}

	fmt.Println(diff)
}

func getIndexOfBankWithMostBlocks(banks [NumOfBanks]int) int {
	max := -1
	index := -1
	for i, blocks := range banks {
		if blocks > max {
			max = blocks
			index = i
		}
	}

	return index
}
