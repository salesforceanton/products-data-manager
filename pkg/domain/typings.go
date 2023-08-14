package products_manager

const (
	ASCENDING_ORDER  = "ASC"
	DESCENDING_ORDER = "DESC"
)

var (
	sortOrderMapping = map[string]int{
		ASCENDING_ORDER:  1,
		DESCENDING_ORDER: -1,
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
	Order int
}
