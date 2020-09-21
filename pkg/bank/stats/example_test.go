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


func ExampleTotalInCategory() {
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
		{
			ID: 2000,
			Amount: 30_000_00,
			Category: "silver",
		},
		{
			ID: 2000,
			Amount: 40_000_00,
			Category: "silver",
		},
			
	}
	a:="silver"
	max:=TotalInCategory(payment,types.Category(a))
	fmt.Println(max)
	//Output:
	//7000000
}