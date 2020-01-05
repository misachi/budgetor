package budgets

type BudgetPeriodRepo interface {
	Items(items []interface{}) []interface{}
	Item(item interface{}) interface{}
	AddItem(item interface{})
	UpdateItem(item interface{})
	Deleteitem(item interface{})
}
