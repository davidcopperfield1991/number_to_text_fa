package ntt

import (
	"fmt"
	"strconv"
	"strings"
)

var Va = " و "

func chandragham(i int) int {
	count := 0
	l := i
	for l > 0 {
		l = l / 10
		count++
	}
	fmt.Println(count)
	return count
}

func makearray(i int) []int {

	stringesh := strconv.Itoa(i)
	addadstr := strings.Split(stringesh, "")
	addadint := []int{}
	for _, i := range addadstr {
		j, _ := strconv.Atoi(i)
		addadint = append(addadint, j)
	}
	// fmt.Println(addadint)

	return addadint

}

func yekraghami(i int) string {
	if i > 10 {
		fmt.Println("addad eshtebah")
	}
	list := [10]string{"صفر", "یک", "دو", "سه", "چهار", "پنج", "شش", "هفت", "هشت", "نه"}
	// fmt.Println(list[i])
	return list[i]
}

func doraghami(i int) string {

	list := [10]string{"ده", "یازده", "دوازده", "سیزده", "چهارده", "پانزده", "شانزده", "هفده", "هجده", "نوزده"}
	list2 := [8]string{"بیست", "سی", "چهل", "پنجاه", "شصت", "هفتاد", "هشتاد", "نود"}
	va := " و "
	array := makearray(i)

	if i < 10 || i > 99 {
		fmt.Println("dadash do raghami balad nisti?")
	}
	if i < 20 {
		fmt.Println(list[i-10])
		return list[i-10]
	}

	if i > 20 || i < 100 {

		if array[1] == 0 {
			fmt.Println("dastan")
			return list2[array[0]-2]
		}

		return list2[array[0]-2] + va + yekraghami(array[1])
	}
	return ""

}

func seraghami(i int) string {
	list := [9]string{"یکصد", "دویست", "سیصد", "چهارصد", "پانصد", "ششصد", "هفتصد", "هشتصد", "نهصد"}

	if i < 100 || i > 999 {
		fmt.Println("adad eshtebah hastesh")
	}
	array := makearray(i)
	dahgan := i % 100
	fmt.Println("se raghami ro bebin :")
	fmt.Println(list[array[0]-1])

	if dahgan == 0 {
		return list[array[0]-1]
	}

	if dahgan < 10 {
		yekraghami(dahgan)
		return list[array[0]-1] + Va + yekraghami(dahgan)
	}

	return list[array[0]-1] + Va + doraghami(dahgan)
}

// func render(i int) string {
// 	if i > 0 && i < 10 {
// 		return yekraghami(i)
// 	}
// 	if i > 9 && i < 100 {
// 		return doraghami(i)
// 	}
// 	if i > 99 && i < 1000 {
// 		return seraghami(i)
// 	}
// 	return ""
// }

func hezaregan(i int) string {
	l := chandragham(i)
	adadmain := i / 1000
	baghiadad := i % 1000

	if baghiadad == 0 {
		Va = ""
	}

	switch l {
	case 4:
		char := yekraghami(adadmain)
		baghye := dana(baghiadad)
		return char + " هزار " + Va + baghye

	case 5:
		char := doraghami(adadmain)
		baghye := dana(baghiadad)
		return char + " هزار " + Va + baghye
	case 6:
		char := seraghami(adadmain)
		baghye := dana(baghiadad)
		return char + " هزار " + Va + baghye
	}
	return ""
}

func millions(i int) string {

	l := chandragham(i)
	adadmain := i / 1000000
	baghiadad := i % 1000000
	switch l {

	case 7:
		char := yekraghami(adadmain)
		baghye := dana(baghiadad)
		return char + " میلیون " + Va + baghye
	case 8:
		char := doraghami(adadmain)
		baghye := dana(baghiadad)
		return char + " میلیون " + Va + baghye
	case 9:
		char := seraghami(adadmain)
		baghye := dana(baghiadad)
		return char + " میلیون " + Va + baghye
	}
	return ""

}

func billions(i int) string {
	l := chandragham(i)
	adadmain := i / 1000000000
	baghiadad := i % 1000000000

	switch l {

	case 10:
		char := yekraghami(adadmain)
		baghye := dana(baghiadad)
		return char + " میلیارد " + Va + baghye
	case 11:
		char := doraghami(adadmain)
		baghye := dana(baghiadad)
		return char + " میلیارد " + Va + baghye
	case 12:
		char := seraghami(adadmain)
		baghye := dana(baghiadad)
		return char + " میلیارد " + Va + baghye
	}
	return "alaki"
}

func dana(i int) string {
	l := chandragham(i)
	fmt.Println(l)
	if l == 0 {
		return ""
	}
	if l == 1 {
		return yekraghami(i)
	}
	if l == 2 {
		return doraghami(i)
	}
	if l == 3 {
		return seraghami(i)
	}
	if l > 3 && l < 7 {
		return hezaregan(i)
	}
	if l > 6 && l < 10 {
		return millions(i)
	}
	if l > 9 && l < 13 {
		return billions(i)
	}
	return "dadash waghan be in adad niaz dari ?"
}

// 1001 eshtebas
func Call(i int) string {
	return dana(i)
}
