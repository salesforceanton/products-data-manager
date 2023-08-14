package products_manager

import "errors"

const (
	ASCENDING_ORDER  = "ASC"
	DESCENDING_ORDER = "DESC"
)

type SortingOrder string

const (
	ASC  SortingOrder = "ASC"
	DESC SortingOrder = "DESC"
)

var (
	SortOrderMapping = map[SortingOrder]int{
		ASCENDING_ORDER:  1,
		DESCENDING_ORDER: -1,
	}

	sortOrders = map[SortingOrder]SortingOptions_Order{
		ASCENDING_ORDER:  SortingOptions_ASC,
		DESCENDING_ORDER: SortingOptions_DESC,
	}
)

type Product struct {
	Name  string  `bson:"name" json:"name"`
	Price float64 `bson:"price" json:"price"`
}

type PaginOptions struct {
	Limit  int
	Offset int
}

type SortOptions struct {
	Field string
	Order SortingOrder
}

func ToPbSortOrder(order SortingOrder) (SortingOptions_Order, error) {
	val, ok := sortOrders[order]
	if !ok {
		return 0, errors.New("invalid order param")
	}

	return val, nil
}
