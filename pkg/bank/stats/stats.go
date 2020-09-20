package stats

import "github.com/sulton1995/bank/pkg/bank/types"

func Avg(payments []types.Payment) types.Money  {
	// total := []types.Payment{}
	 sum :=0
	 
	for _, card := range payments {
	
	 sum=int(card.Amount) + sum
	 	
	}
	return   types.Money(sum/len(payments))
}