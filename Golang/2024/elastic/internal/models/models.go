package models

type Category struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type Product struct {
	CategoryID int 		 `json:"category_id"`
	ID         int       `json:"id"`
	Name       string    `json:"name"`
}

// TestProduct структура для продукта
type TestProduct struct {
	ProdName string `json:"prod_name"`
}

// TestCategory структура для категории
type TestCategory struct {
	Name     string       `json:"name"`
	Products []TestProduct `json:"products"`
}

// TestData структура для внешнего объекта данных
type TestData struct {
	Category TestCategory `json:"category"`
}

// InputData тип для данных, содержащих массив TestData
type InputData struct {
	Data []TestData `json:"data"`
}