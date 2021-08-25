package main

import (
	"fmt"
	"pribadi/reflect/library"
	"reflect"
)

/*
	Reflect
	teknik untuk inspeksi sebuah variabel,
	mengambil informasi dari variabel tersebut atau bahkan memanipulasinya.
	Cakupan informasi yang bisa didapatkan lewat reflection sangat luas,
	seperti melihat struktur variabel, tipe, nilai pointer, dan banyak lagi.

	reflect.ValueOf() => mengembalikan objek dalam tipe reflect.Value , yang berisikan informasi
		yang berhubungan dengan nilai pada variabel yang dicari
	reflect.TypeOf() => mengembalikan objek dalam tipe reflect.Type . Objek tersebut berisikan informasi
		yang berhubungan dengan tipe data variabel yang dicari

	Berikut adalah konstanta tipe data dan method yang bisa digunakan dalam refleksi di Golang:
		Bool, Int, Int8, Int16, Int32, Int64, Uint, Uint8, Uint16, Uint32, Uint64, Uintptr,
		Float32, Float64, Complex64, Complex128, Array, Chan, Func, Interface, Map, Ptr, Slice, String,
		Struct, UnsafePointer
*/

func main() {

	number := 23
	reflectValue := reflect.ValueOf(number)
	var tipeVariableNumber = reflectValue.Type() // mengembalikan tipe data dalam bentuk string

	fmt.Println("Tipe Variable :", tipeVariableNumber)

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("Nilai Variable :", reflectValue.Int()) // reflectValue.Int() mengembalikan nilai variable
	}

	/*
		PengaksesaN nilai dalam bentuk interface{}
	*/

	fmt.Println()
	fmt.Println(reflectValue.Interface())
	fmt.Println(reflectValue.Interface().(int))

	var s1 = &library.Student{Name: "Rahmatulah Sidik", Age: 27, Class: 12}
	s1.GetPropertyInfo()

	s2 := library.Student{}
	// cara memanggil method
	s2.SetName("John Druce")

	var reflectValue2 = reflect.ValueOf(s1)
	// cara memanggil method dengan reflect
	var methodSetName = reflectValue2.MethodByName("SetName")
	methodSetName.Call([]reflect.Value{
		reflect.ValueOf("John Wick"),
	})

	fmt.Println(s1.Name)
	fmt.Println(s2.Name)

}
