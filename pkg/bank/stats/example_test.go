package stats

import (
	"fmt"

	"github.com/Shah-2021/bank/pkg/bank/types")


func ExampleAvg() {
	payments := []types.Payment{
		{
			Amount: 100,
		},
		{
			Amount: 80,
		},
		{
			Amount: 99,
		},
	}

    fmt.Println(Avg(payments))

	//Output: 93
}
