package stats

import (
	"github.com/nkomiljon/hwbank/pkg/types"
	"fmt"

)


func ExampleAvg(){
	payments := []types.Payment{
		{
			ID: 1,
			Amount: 53_00,
			Category: "Cat",
			Status: types.StatusOK,
		},
		{
			ID: 2,
			Amount: 51_00,
			Category: "Cat",
			Status: types.StatusOK,
		},
		{
			ID: 3,
			Amount: 52_00,
			Category: "Cat",
			Status: types.StatusFail,
		},
	}

	fmt.Println(Avg(payments))

	//Output: 5200
}

func ExampleTotalInCategory(){
	payments := []types.Payment{
		{
			ID: 1,
			Amount: 10_000_00,
			Category: "Auto",
			Status: types.StatusOK,
		},
		{
			ID: 2,
			Amount: 20_000_00,
			Category: "Cafe",
			Status: types.StatusOK,
		},
		{
			ID: 3,
			Amount: 30_000_00,
			Category: "Restaurant",
			Status: types.StatusFail,
		},
	}
	inCategory := types.Category("Auto")
	totalInCategory := TotalInCategory(payments, inCategory)
	fmt.Println(totalInCategory)

	//Output: 1000000
}

