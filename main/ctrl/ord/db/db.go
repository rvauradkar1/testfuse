package db

type IDBService interface {
	saveOrder(order string)
}

type DBService struct {
	Counts map[string]int
	Status string
}

func (db *DBService) Add(product string) bool {
	if count, ok := db.Counts[product]; ok {
		count++
		return true
	}
	db.Counts[product] = 1
	return true
}
