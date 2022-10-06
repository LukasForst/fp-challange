package main

import (
	"fmt"
	"math"
)

// we have this global map as the assignment defines the `prioritize` args without it
// otherwise I'd pass it as an argument there
var latencies map[string]int

// function should return a subset (or full array)
// that will maximize the USD value and fit the transactions under 1 second
func prioritize(transactions []Transaction, totalTime int) []Transaction {
	// this 0-1 knapsack problem, I chose to solve it by using dynamic programming approach
	// as the total time (our capacity/max W) is quite small
	// final time complexity is O(N*W), in our case O(len(transactions) * totalTime)
	// space complexity is O(N*W), in our case O(len(transactions) * totalTime)

	// number of items
	elements := len(transactions)
	// maximum capacity fo the knapsack
	maximumCapacity := totalTime
	weight := make([]int, elements)
	value := make([]float64, elements)
	// now we fill data
	for i, transaction := range transactions {
		// in our case, weight of the element is the latency to the bank
		// as we have limit of how long can we process all transactions
		weight[i] = latencies[transaction.BankCountryCode]
		// value of each element, or how much we gain by including it
		value[i] = transaction.Amount
	}
	// let's make our dynamic programming cache, Golang by default sets integers to 0
	// so no need to zero weight with no elements
	m := make([][]float64, elements+1)
	for i := range m {
		m[i] = make([]float64, maximumCapacity+1)
	}
	// now we compute actual values by combining all weights
	// we go from 1 as we look on a row before
	// note: weight and values need to access i-1 instead of i as we have one row more in m
	for i := 1; i <= elements; i++ {
		for j := 0; j <= maximumCapacity; j++ {
			if weight[i-1] > j {
				// if including i-1 element exceeds current max capacity, we need to skip it
				m[i][j] = m[i-1][j]
			} else {
				// if not, we select maximal gain by comparing
				// current value skipping current AND new value with removing previous and adding this
				m[i][j] = math.Max(m[i-1][j], m[i-1][j-weight[i-1]]+value[i-1])
			}
		}
	}

	// and we backtrack to get the list of items there
	var filteredTransactions []Transaction
	remainingCapacity := maximumCapacity
	remainingValue := m[elements][maximumCapacity]
	for i := elements; i > 0; i-- {
		// let's see if the value came from the previous row
		if m[i-1][remainingCapacity] != remainingValue {
			// if it didn't then we found our element that created this value,
			// so we subtract the capacity that was filled by this element
			remainingCapacity -= weight[i-1]
			// and also the value
			remainingValue -= value[i-1]
			// finally we append the transaction to the result set
			filteredTransactions = append(filteredTransactions, transactions[i-1])
		}
		// now cut off if there's no remaining capacity or value
		if remainingCapacity == 0 || remainingValue == 0 {
			break
		}
	}
	return filteredTransactions
}

func main() {
	timeLimit := 1000
	latencies = loadLatencies("./latencies.json")
	transactions := loadTransactions("./transactions.csv")

	prioritized := prioritize(transactions, timeLimit)

	totalAmount := 0.0
	totalTime := 0
	for _, transaction := range prioritized {
		totalAmount += transaction.Amount
		totalTime += latencies[transaction.BankCountryCode]
	}

	fmt.Printf("Time Limit: %d\nTotal Time Needed: %d\nTotal Amount: %.2f", timeLimit, totalTime, totalAmount)
}
