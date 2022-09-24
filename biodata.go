package main

import (
	"fmt"
	"log"
	"reflect"
)

type IStudent interface {
	AddStudent(...Student)
	Display(...int)
}

type Student struct {
	No                              int
	Nama, Alamat, Pekerjaan, Alasan string
}

type DataFGA struct {
	Students []Student
}

// Add new Student to existing DataFGA
func (d *DataFGA) AddStudent(Student ...Student) {
	d.Students = append(d.Students, Student...)
}

// Find Student struct field with the longest value
func (d DataFGA) getLongestStudentsValue(n int) int {
	reflectValue := reflect.ValueOf(d.Students[n])

	longest := len(reflectValue.Field(0).String())

	for i := 1; i < reflectValue.NumField(); i++ {
		if len(reflectValue.Field(i).String()) > longest {
			longest = len(reflectValue.Field(i).String())
		}
	}

	if longest%2 != 0 {
		longest += 1
	}

	return longest
}

// display student(s) bio for given No(s)
func (d DataFGA) Display(search ...int) {
	for i := 0; i < len(search); i++ {
		for j := 0; j < len(d.Students); j++ {
			if j == len(d.Students)-1 && d.Students[j].No != search[i] {
				fmt.Printf("!!Error : Student dengan nomor absen %d tidak ditemukan.\n", search[i])
			}
			if d.Students[j].No == search[i] {
				fmt.Printf(`
~~~~~%s~~~~~~
%s Student

Nama      : %s
Alamat    : %s
Pekerjaan : %s
Alasan    : %s
~~~~~%s~~~~~~
`, RepeatChar("~", d.getLongestStudentsValue(j)), RepeatChar(" ", d.getLongestStudentsValue(j)/2), d.Students[j].Nama, d.Students[j].Alamat, d.Students[j].Pekerjaan, d.Students[j].Alasan, RepeatChar("~", d.getLongestStudentsValue(j)))

				break
			}
		}
	}
}

var Biodata IStudent

func init() {
	Biodata = CreateNewDataFGA(Student{
		No:   1,
		Nama: "Rafli Abi Kusuma",
	}, Student{
		No:        2,
		Nama:      "M Fitrah Ramadhan",
		Alamat:    "Palembang",
		Pekerjaan: "Software Developer",
		Alasan:    "Belajar",
	}, Student{
		No:   3,
		Nama: "Agung chumaidi",
	}, Student{
		No:   4,
		Nama: "Adrian Metanoia Gawa",
	}, Student{
		No:   5,
		Nama: "M DIYANSYAH",
	}, Student{
		No:   6,
		Nama: "Risda Tamam Aljava",
	}, Student{
		No:   7,
		Nama: "Dikky setiawan",
	}, Student{
		No:   8,
		Nama: "Simon Aditia",
	}, Student{
		No:   9,
		Nama: "Farrel Athaillah Putra",
	}, Student{
		No:   10,
		Nama: "Arif Rahmadi",
	})
}

// Create new Student, returning Student interface
func CreateNewDataFGA(Student ...Student) IStudent {
	return &DataFGA{
		Students: Student,
	}
}

func main() {
	args, err := GetArgs()
	if err != nil {
		log.Fatalln(err.Error())
	}

	Biodata.Display(args...)
}
