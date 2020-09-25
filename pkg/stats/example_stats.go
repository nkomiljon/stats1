package stats



import (
	"github.com/nkomiljon/hwbank/pkg/bank/types"
	"fmt"
)


func ExampleAvg(){
	payments := []types.Payment{
		{
			ID:2,
			Amount:53_00,
			Category: "Cat",
		},
		{
			ID:1,
			Amount:51_00,
			Category: "Cat",
		},
		{
			ID:3,
			Amount:52_00,
			Category: "Cat",
		},
	}

	fmt.Println(Avg(payments))

	//Output: 5200
}

func ExampleTotalInCategory(){
	payments := []types.Payment{
		{
			ID:2,
			Amount:53_00,
			Category: "Cafe",
		},
		{
			ID:1,
			Amount:51_00,
			Category: "Cafe",
		},
		{
			ID:3,
			Amount:52_00,
			Category: "Restaurant",
		},
	}
	inCategory := types.Category("auto")
	totalInCategory := TotalInCategory(payments, inCategory)
	fmt.Println(totalInCategory)

	//Output: 10400
}

