package stats


import (
	
)


//Avg func return average amount from slice Payment  
func Avg(payments []types.Payment) types.Money  {
	var sum types.Money
	for _, v := range payments {
		sum+=v.Amount
	}
	return sum/types.Money(len(payments))
}

//TotalInCategory  returned total sum in one category
func TotalInCategory(payments []types.Payment, category types.Category) (total types.Money){

	for _, v := range payments {
		if category==v.Category{
			total+=v.Amount
		}
	}
	return
}  
