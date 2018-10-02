/*
Author: John Lahut
Project: Sorting
Date: 9/18/2018
*/
package main

import (
	"encoding/json" // JSON parsing
	"fmt"           // needed for printing, returning json
	"math/rand"     // used for sort testing
	"net/http"      // to serve requests
	"os"            // needed by AWS
	"reflect"       // compare slices
	"sort"          // used for sort testing
	"sorting"       // where MergeSort is located
)

// JSONIn struct - receiver structs
type JSONIn struct {
	List *[]int `json:"inList"`
}

// JSONOut struct - returner struct
type JSONOut struct {
	List   []int  `json:"outList"`
	Time   int64  `json:"timeMS"`
	Method string `json:"algorithm"`
}

// JSONError struct - error struct
type JSONError struct {
	Message string `json:"message"`
}

// main entry point of program
func main() {

	// AWS EBS Stuff
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// testing endpoints
	http.HandleFunc("/mergesort", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			w.Header().Set("Content-Type", "application/json")

			// handle incoming JSON, try and decode it
			var data *JSONIn
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&data)

			// if error is encountered, or if list cannot be mapped, return error message
			if err != nil || data.List == nil {
				var e JSONError
				e.Message = "Malformed JSON."
				b, _ := json.Marshal(e)
				fmt.Fprint(w, string(b))
			} else {

				// time is returned by mergesort, list is passed by ref
				time := sorting.MergeSort(*data.List)

				// output json object
				var outJSON JSONOut
				outJSON.Method = "MergeSort"
				outJSON.Time = time
				outJSON.List = *data.List
				b, err := json.Marshal(outJSON)
				if err != nil {
					panic(err)
				}

				// return to caller
				fmt.Fprint(w, string(b))
			}
		} else {

			b, err := json.Marshal("Hello! Please send a post request.")

			if err != nil {
				panic(err)
			}

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, string(b))
		}
	})

	// serve requests
	http.ListenAndServe(":"+port, nil)

	// testSorting(10000, 1)

}

// this function uses the builtin sort library to validate sorting algo
// int arrLen: length of array to sort through
// int iters: number of test iterations
// returns: true if all tests passed, false otherwise
func testSorting(arrLen, iters int) bool {
	for k := 0; k < iters; k++ {
		arr, valid := make([]int, arrLen), make([]int, arrLen)
		for i := 0; i < len(arr); i++ {
			j := rand.Intn(20)
			arr[i] = j
			valid[i] = j
		}
		fmt.Printf("%v\n", arr)
		t := sorting.MergeSort(arr)
		sort.Ints(valid)

		if !reflect.DeepEqual(arr, valid) {
			fmt.Printf("%v\n", "**TEST FAILED!!**")
			return false
		}

		fmt.Printf("Elapsed time (ms): %v\n", t)
	}
	return true
}
