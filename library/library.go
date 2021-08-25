package library

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Age   int
	Class int
}

func (s *Student) GetPropertyInfo() {
	// reflect berisi objek Struct s
	var reflectValue = reflect.ValueOf(s)

	// pengecekan variable reflectValue pointer  atau bukan,
	// jika iya, cari objek reflect aslinya  dengan cara memanggil method Elem() .
	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	// melakukan perulangan berdaasarkan tipe data yg berada di property struct student
	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama :", reflectType.Field(i).Name)
		fmt.Println("tipe data :", reflectType.Field(i).Type)
		fmt.Println("nilai :", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func (s *Student) SetName(name string) {
	s.Name = name
}
