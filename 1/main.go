package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	list1, list2 := getInputAsLists("./input.txt")
	intList1, err := StringsToUints16(list1)
	if err != nil {
		log.Fatal(err)
	}
	intList2, err := StringsToUints16(list2)
	if err != nil {
		log.Fatal(err)
	}
	sort.Slice(intList1, func(i, j int) bool {
		return intList1[i] < intList1[j]
	})

	sort.Slice(intList2, func(i, j int) bool {
		return intList2[i] < intList2[j]
	})

	var total int

	for i, _ := range intList1 {
		dist := absInt(intList1[i] - intList2[i])
		total += int(dist)
	}

	fmt.Println(total)
}

func absInt(x int) int{
    if x < 0 {
        return -x
    }
    return x
}

func sumint(numbers []int) int {
	var sum int
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func getInputAsLists(location string) ([]string, []string) {
	f, err := os.Open(location)

	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)

	var list1 []string
	var list2 []string
	for scanner.Scan() {
		line := scanner.Text()

		splitLine := strings.Split(line, "   ")
		// fmt.Println(splitLine[1])
		list1 = append(list1, splitLine[0])
		list2 = append(list2, splitLine[1])
	}

	if err := scanner.Err(); err != nil {
		// Handle error
		panic(err)
	}
	return list1, list2
}

func StringsToUints16(strings []string) ([]int, error) {
	uints := make([]int, 0, len(strings))
	for _, str := range strings {
		num, err := strconv.ParseUint(str, 10, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to convert '%s' to uint: %w", str, err)
		}
		uints = append(uints, int(num))
	}
	return uints, nil
}
