package stats

import (
	"github.com/sulton1995/bank/pkg/bank/types"
	"fmt"
)

func ExampleAvg() {
	payment:=[]types.Payment{
		{
			ID: 1000,
			Amount: 10_000_00,
			Category: "gold",
		},
		{
			ID: 2000,
			Amount: 20_000_00,
			Category: "gold",
		},
				
	}
	max:=Avg(payment)
	fmt.Println(max)
	//Output:
	//1500000
}
