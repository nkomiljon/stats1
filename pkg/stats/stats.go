package stats


import (
	"github.com/nkomiljon/hwbank/pkg/bank/types"
)


//Avg func return average amount from slice Payment  
func Avg(payments []types.Payment) types.Money  {
	countPayments := types.Money(len(payments))
	sumPayments := types.Money(0)
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			continue
		}
		moneyPayments := payment.Amount
		sumPayments += moneyPayments
	}
	return sumPayments / countPayments
}

//TotalInCategory  returned total sum in one category
func TotalInCategory(payments []types.Payment, category types.Category) (total types.Money){
  sumPayments := types.Money(0)
  for _, payment := range payments {
	  if payment.Category != category {
		  continue
	  }
	  if payment.Status != types.StatusFail {
		  continue
	  }
	  moneyPayments := payment.Amount
	  sumPayments += moneyPayments
  }
  return sumPayments
}  
