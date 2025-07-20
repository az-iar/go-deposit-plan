package deposit

import (
	"fmt"
	"sort"
)

const (
	PlanTypeOneTime = "one-time"
	PlanTypeMonthly = "monthly"
)

type DepositPlan struct {
	Type        string
	Allocations map[string]int // in cents
}

type Deposit struct {
	Amount int // in cents
}

func AllocateDeposits(plans []DepositPlan, deposits []Deposit) (map[string]int, error) {
	totalAllocations := make(map[string]int)
	totalDeposited := 0

	for _, d := range deposits {
		totalDeposited += d.Amount
	}

	var oneTime, monthly *DepositPlan

	for _, plan := range plans {
		if !isValidPlanType(plan.Type) {
			return nil, fmt.Errorf("invalid plan type: %s", plan.Type)
		}

		switch plan.Type {
		case PlanTypeOneTime:
			oneTime = &plan
		case PlanTypeMonthly:
			monthly = &plan
		}
	}

	if oneTime != nil {
		planTotal := sum(oneTime.Allocations)

		if totalDeposited >= planTotal {
			for k, v := range oneTime.Allocations {
				totalAllocations[k] += v
			}
			totalDeposited -= planTotal
		} else {
			proportionalDistribute(oneTime.Allocations, totalDeposited, totalAllocations)
			return totalAllocations, nil
		}
	}

	if monthly != nil && totalDeposited > 0 {
		planTotal := sum(monthly.Allocations)
		if planTotal > 0 {
			proportionalDistribute(monthly.Allocations, totalDeposited, totalAllocations)
		}
	}

	return totalAllocations, nil
}

func isValidPlanType(planType string) bool {
	return planType == PlanTypeOneTime || planType == PlanTypeMonthly
}

func sum(m map[string]int) int {
	total := 0
	for _, v := range m {
		total += v
	}
	return total
}

func proportionalDistribute(plan map[string]int, total int, result map[string]int) {
	var keys []string
	for k := range plan {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	planTotal := sum(plan)
	remaining := total

	for i, name := range keys {
		portion := plan[name]

		if i == len(keys)-1 {
			result[name] += remaining
		} else {
			alloc := portion * total / planTotal
			result[name] += alloc
			remaining -= alloc
		}
	}
}
