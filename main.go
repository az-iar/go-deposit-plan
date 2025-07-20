package main

import (
	"deposit/deposit"
	"fmt"
	"sort"
)

func runAllocateDeposits() {
	plans := []deposit.DepositPlan{
		{
			Type: deposit.PlanTypeOneTime,
			Allocations: map[string]int{
				"High risk":  1000000,
				"Retirement": 50000,
			},
		},
		{
			Type: deposit.PlanTypeMonthly,
			Allocations: map[string]int{
				"High risk":  0,
				"Retirement": 10000,
			},
		},
	}

	deposits := []deposit.Deposit{
		{Amount: 1050000},
		{Amount: 10000},
	}

	result, err := deposit.AllocateDeposits(plans, deposits)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Final allocations:")
	keys := make([]string, 0, len(result))
	for k := range result {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s: $%.2f\n", k, float64(result[k])/100)
	}
}

func main() {
	runAllocateDeposits()
}
