package main
import . "fmt"

type MyStruct struct {
	On bool
	Ammo, Power int
}

func (m *MyStruct) Shoot()(bool){
	if m.On == false {
		return false
	} else if m.Ammo > 0 {
		m.Ammo--
		Println("piu-piu nahui, ostatok ammo:", m.Ammo)
		return true
	} else {
		return false
	}

}

func (m *MyStruct) RideBike()(bool){
	if m.On == false {
		return false
	} else if m.Power > 0 {
		m.Power--
		Println("yarik huyarit, ostatok stamini:", m.Power)
		return true
	} else {
		return false
	}
	
}

func main(){
	asd := MyStruct{
		On: true,
		Ammo: 50,
		Power: 50,
	}
	testStruct := asd
	// testStruct.Shoot()
	// testStruct.RideBike()
	// Println(testStruct.Shoot(), testStruct.RideBike())
}


// func main() {

	// testStruct :=
	/*
	 * Экземпляр созданной вами структуры необходимо передать в качестве
	 * аргумента функции testStruct, которая выполнит проверку соблюдения
	 * всех условий задания/
	 */

// }