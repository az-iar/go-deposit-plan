package deposit

import (
	"reflect"
	"testing"
)

func TestAllocateDeposits(t *testing.T) {
	tests := []struct {
		name        string
		plans       []DepositPlan
		deposits    []Deposit
		expected    map[string]int
		expectError bool
	}{
		{
			name: "happy path - one-time and monthly",
			plans: []DepositPlan{
				{
					Type: PlanTypeOneTime,
					Allocations: map[string]int{
						"High risk":  1000000,
						"Retirement": 50000,
					},
				},
				{
					Type: PlanTypeMonthly,
					Allocations: map[string]int{
						"High risk":  0,
						"Retirement": 10000,
					},
				},
			},
			deposits: []Deposit{
				{Amount: 1050000},
				{Amount: 10000},
			},
			expected: map[string]int{
				"High risk":  1000000,
				"Retirement": 60000,
			},
		},
		{
			name: "invalid plan type should return error",
			plans: []DepositPlan{
				{
					Type: "invalid",
					Allocations: map[string]int{
						"Play fund": 100000,
					},
				},
			},
			deposits:    []Deposit{{Amount: 100000}},
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := AllocateDeposits(tt.plans, tt.deposits)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected: %v, got: %v", tt.expected, result)
			}
		})
	}
}
