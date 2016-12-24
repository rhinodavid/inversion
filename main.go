package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// See: https://stackoverflow.com/questions/9862443/golang-is-there-a-better-way-read-a-file-of-integers-into-an-array
func readFile(fname string) (nums []int, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	// Assign cap to avoid resize on every append.
	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}

func countInversions(arr []int, inv int) ([]int, int) {
	if len(arr) == 0 || len(arr) == 1 {
		return arr, inv
	}
	halfway := len(arr) / 2
	leftArr, leftI := countInversions(arr[:halfway], inv)
	rightArr, rightI := countInversions(arr[halfway:], inv)
	return mergeAndCountInversions(leftArr, leftI, rightArr, rightI)
}

func mergeAndCountInversions(a []int, aI int, b []int, bI int) ([]int, int) {
	result := make([]int, len(a)+len(b))
	i := 0
	j := 0
	inversions := aI + bI
	for k := 0; k < len(a)+len(b); k++ {
		// used up all of a
		if i == len(a) {
			result = append(result[:k], b[j:]...)
			break
		}
		// used up all of b
		if j == len(b) {
			result = append(result[:k], a[i:]...)
			break
		}
		if a[i] < b[j] {
			result[k] = a[i]
			i++
		} else {
			inversions += len(a) - i
			result[k] = b[j]
			j++
		}
	}
	return result, inversions
}

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		panic("Enter the name of the file with the integer list you'd like inversions counted on.")
	}
	nums, err := readFile(flag.Args()[0])
	if err != nil {
		panic(err)
	}
	_, inversions := countInversions(nums, 0)
	fmt.Println(inversions)
}
