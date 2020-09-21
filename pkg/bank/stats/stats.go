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

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	
	sum:=0
		for _, card := range payments {
			
			if (card.Category == category){
				sum = sum + int(card.Amount)
			}				
		} 
		return types.Money(sum)
}