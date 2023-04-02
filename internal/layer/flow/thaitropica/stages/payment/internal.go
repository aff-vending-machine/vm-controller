package payment

import (
	"fmt"
)

func makeMerchantOrderIDFilter(id string) []string {
	return []string{
		fmt.Sprintf("merchant_order_id:=:%s", id),
	}
}
