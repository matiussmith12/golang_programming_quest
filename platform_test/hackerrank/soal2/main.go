package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)
func Abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func closestNumbers(arr []int32) []int32 {
	minValue := int32(math.MaxInt32)
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
    // Write your code here
	output := make([]int32, 0)
	var temp int32
	for i := 0; i < len(arr) - 1; i++ {
		temp := Abs(arr[i] -  arr[i + 1])

		if temp < minValue {
			minValue = temp
		}
	}



	// Display pair
	for i := 0; i < len(arr) - 1; i++ {
		r1, r2 := arr[i], arr[i + 1]
		temp = Abs(r1 - r2)
		var pair []int32

		if temp  == minValue {
            pair = append(pair,r1)
            pair = append(pair,r2)
        }
		output = append(output, pair...)
	}
	return output
}

func main() {

	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    result := closestNumbers(arr)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%d", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, " ")
        }
    }

    fmt.Fprintf(writer, "\n")

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }

}