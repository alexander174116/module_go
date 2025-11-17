// Программа для управления небольшим складом товаров
package main

import (
	"fmt"
	"time"
)

// Глобальные константы для категорий товаров
const (
	CategoryElectronics = "Электроника"
	CategoryFood        = "Продукты"
	CategoryClothes     = "Одежда"
	MaxItems            = 100 // Максимальное количество товаров на складе
)

// Глобальная переменная для подсчета товаров
var totalItems int

// структура для товаров, в дальнейшем удобнее текущего варианта
type someItem struct {
	Name string
	Price float64
	MinQuan, MaxQuan int
	Available bool
	Quan int
	ID int64
	Color string
	Weight float64
	WhenAdded time.Time
}

func main() {

	var itemName string
	itemName = "Смартфон"
	var itemPrice float64 = 999.99
	var minQuantity, maxQuantity int = 5, 20
	var isAvailable = true
	quantity := 15

	var (
		itemID     int64  = 12345
		itemColor  string = "Черный"
		itemWeight float32
		dateAdded  time.Time = time.Now()
	)

	// var asd time.Time 
	// asd = time.Now()
	// fmt.Printf("%s\n", asd.Format("01-01-2006"))
	// Не трогал стандартные переменные. добавим структуру для нового товара

	Apples := someItem {
		Name: "Яблоко",
		Price: 23.77,
		MinQuan: 10,
		MaxQuan: 40,
		Available: true,
		Quan: 33,
		ID: 65487,
		Color: "Зеленый",
		Weight: 0.1,
		WhenAdded: time.Now(),
	}
	// calculateDiscount(20, itemName, itemID, itemPrice, itemColor)
	calculateDiscount(Apples, 20)
	addNewItem(Apples)
	itemWeight = 0.3

	// Приведение типов
	percentInStock := float64(quantity) / float64(MaxItems) * 100

	// Использование констант
	category := CategoryElectronics

	// Вызов функции для вывода информации о товаре
	displayItemInfo(itemID, itemName, quantity, itemPrice, isAvailable, category)

	// Демонстрация работы с базовыми типами
	fmt.Println("\nДополнительная информация:")
	fmt.Printf("Цвет товара: %s\n", itemColor)
	fmt.Printf("Вес товара: %.2f кг\n", itemWeight)
	fmt.Printf("Минимальное количество: %d, Максимальное количество: %d\n", minQuantity, maxQuantity)
	fmt.Printf("Процент от максимального количества на складе: %.1f%%\n", percentInStock)
	fmt.Printf("Дата добавления: %s\n", dateAdded.Format("02-01-2006"))

	const (
		Small  = 1
		Medium = 5
		Large  = 10
	)

	// Работа со строками и байтами
	productCode := "ЯЩИК-12345"
	fmt.Printf("\nКод товара: %s, длина: %d символов, длина: %d байт\n", productCode, len([]rune(productCode)), len(productCode))

	// Обновление общего количества товаров
	updateTotalItems(quantity)
	fmt.Printf("\nОбщее количество товаров на складе: %d\n", totalItems)

}

// Функция для отображения информации о товаре
func displayItemInfo(id int64, name string, qty int, price float64, isAvailable bool, category string) {
	fmt.Println("=== Информация о товаре ===")
	fmt.Printf("ID: %d\n", id)
	fmt.Printf("Название: %s\n", name)
	fmt.Printf("Категория: %s\n", category)
	fmt.Printf("Количество: %d\n", qty)
	fmt.Printf("Цена: %.2f руб.\n", price)

	// Использование условного оператора с логическим типом
	if isAvailable {
		fmt.Println("Статус: В наличии")
	} else {
		fmt.Println("Статус: Нет в наличии")
	}
}

// Функция для обновления общего количества товаров
func updateTotalItems(qty int) {
	totalItems += qty

	// Проверка на превышение максимального количества
	if totalItems > MaxItems {
		fmt.Println("Предупреждение: Превышено максимальное количество товаров на складе!")
		totalItems = MaxItems
	}
}

// Функция добавления товара старым методом на склад и расчета заполнения


// Функция добавления товара новым методом на склад и расчета заполнения
func addNewItem (s someItem)(id int64){

	if totalItems+s.Quan > MaxItems{
		fmt.Println("Предупреждение: На складе нет места для такого количества товаров!")
		return 0
	}
	totalItems += s.Quan
	fmt.Printf("\nТовар %q, цвет: %q в количестве %d добавлен на склад\n\n", s.Name, s.Color, s.Quan)
	return s.ID

}


// Функция для расчета и вывода цены со скидкой
//для поддержки с вариантом старых товаров
// func calculateDiscount (dicsount int, name string, id int64, price float64, color string){
// 	res := float64(price) - (float64(price) * float64(dicsount) / 100.0)
// 	fmt.Printf("\nНовая цена товара %q, c цветом: %q, со скидкой %d%% - %.2f. Его ID: %v.\n", name, color, dicsount, res, id)
// 	fmt.Printf("Скидка на товар составила %.2f.\n", price - res)
// }

//при условии перезда на структуры другая реализация
func calculateDiscount (s someItem, dicsount int){
	res := float64(s.Price) - (float64(s.Price) * float64(dicsount) / 100.0)
	fmt.Printf("\nНовая цена товара %q, со скидкой %d%% - %.2f, его ID: %v.\n", s.Name, dicsount, res, s.ID)
	fmt.Printf("Скидка на товар составила %.2f.\n", s.Price - res)
}
