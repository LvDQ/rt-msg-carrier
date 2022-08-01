package biz

import (
	"context"
	"encoding/json"
	"fmt"
	"rt-msg-carrier/models/gen/dal"
	"rt-msg-carrier/models/gen/dal/query"
)

func GetCustomer() {
	ctx := context.Background()
	db := dal.GetDB()
	c := query.Use(db).Customer
	cust, err := c.WithContext(ctx).First()
	if err != nil {
		return
	}
	j, _ := json.Marshal(cust)
	fmt.Println(j)
	fmt.Printf("j: %T\n", j)
	fmt.Println(string(j))
	fmt.Printf("string(j): %T\n", string(j))
}
