package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"

	"github.com/taowu99/webservice/controllers"
	"github.com/taowu99/webservice/leetcode"
	"github.com/taowu99/webservice/models"
)

func testReturn(d1 float64, d2 float64) (res float64, err error) {
	if d2 == 0 {
		err = errors.New("Can't devide by zero")
	}
	res = d1 / d2
	return
}

func main3() {
	mp := map[string]int{"tao": 1, "wu": 2}
	for k, v := range mp {
		fmt.Println(k, v)
	}
	for k, _ := range mp {
		fmt.Println(k)
	}
	panic("Something wrong here")
	slice := []int{1, 2, 3}
	slice = append(slice, 4)
	for _, v := range slice {
		fmt.Println(v)
	}
}

func main_http() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}

var gc = 9

func main5() {
	// solution := leetcode.Constructor(4, []int{0, 2})
	// println(solution.Pick(), solution.Pick(), solution.Pick(), solution.Pick(), solution.Pick(), solution.Pick())
	solution := leetcode.Constructor(4, []int{2, 1})
	fmt.Println(solution.Pick(), solution.Pick(), solution.Pick(), solution.Pick(), solution.Pick(), solution.Pick())
	// fmt.Println(leetcode2.TestHi())
}

func main4() {
	slc := []int{11, 12, 13}
	fmt.Println(slc)
	changeSlc2(slc)
	changeSlc(&slc)
	fmt.Println(slc)

	func() {
		fmt.Println("Hello World, I am anonymous funciton")
	}()

	fa := func(name string) {
		fmt.Printf("Hello World, I am anonymous funciton [%s]\n", name)
	}

	fa("function 2")

	// fmt.Printf("%f", double(3, 2, mathExpression(SubtractExpr)))
	testDefer()
}

// func mathExpress(expr MathExpr) {
// 	switch expr {
// 	case AddExpr:
// 		return simplemath.Add
// 	case SubtractExpr:
// 		return simplemath.Subtract
// 	case MultiplyExpr:
// 		return simplemath.Multiple
// 	default:
// 		return func(f, f2 float64) float64 {
// 			return 0
// 		}
// 	}
// }

func testDefer() {
	defer func() {
		fmt.Println("This is 1st defer")
	}()

	for i := 1; i <= 10; i = i + 1 {
		cc := i
		defer func() {
			fmt.Printf("This is %d defer function\n", cc)
		}()
		fmt.Println(i)
		if i == 2 {
			return
		}
	}
	defer func() {
		fmt.Println("this is 2nd defer")
	}()
}

func double(f1, f2 float64, mathExpr func(float64, float64) float64) float64 {
	return 2 * mathExpr(f1, f2)
}

func changeSlc2(data []int) {
	data[0] = 101
}

func changeSlc(data *[]int) {
	*data = append(*data, 14)
}

func main1() {
	const ca = 9.000
	var b = 45
	type tmp int
	a := tmp(10)
	a = a + tmp(b)
	fmt.Println(reflect.TypeOf(ca), ca, len(os.Args), os.Args)
	fmt.Println(reflect.TypeOf(&a), a)

	slc := []string{"a", "b", "d"}
	fmt.Println(len(slc), cap(slc))
	slc = append(slc, "hah")
	fmt.Println(len(slc), cap(slc))
}

func main2() {
	u := models.User{
		ID:        2,
		FirstName: "Tao",
		LastName:  "Wu",
	}
	fmt.Println(u)
	port := 3000
	port, error := startWebServer(port, 8) //_, error := startWebServer(port, 8), //if don't want to use the port var
	fmt.Println(port, error)
}

func startWebServer(port int, numberOfRetries int) (int, error) {
	fmt.Println("Starting Server on port", port)

	fmt.Println("Server started")
	return port, nil
	// return port, errors.New("Something went wrong")
}

func test() {
	var i int = 78
	fmt.Println(i)
	var f float32 = 3.2
	fmt.Println(f)
	firstName := "Tao"
	fmt.Println(firstName)
	var lastName *string = new(string)
	*lastName = "Wu"
	fmt.Println(*lastName)
	ptr := &firstName
	fmt.Println(*ptr, firstName)

	var arr [3]int
	arr[0] = 1
	ar := [5]int{1, 2, 3}
	fmt.Println(arr, ar)

	slc := arr[:]
	slc2 := []int{1, 2, 31}
	arr[1] = 999
	slc = append(slc, 90)
	arr[0] = 0
	fmt.Println(arr, slc, slc2)
	slc2 = append(slc2, 4, 22)
	fmt.Println(slc2)
	slc3 := slc2[1:]
	fmt.Println(slc3)

	v := map[string]int{"a": 1, "b": 2}
	v["c"] = 3
	delete(v, "a")
	delete(v, "b")
	fmt.Println(v)

	type user struct {
		ID   int
		Name string
	}
	var u user
	u.ID = 1234
	u.Name = "Tao Wu"
	u2 := user{ID: 1234, Name: "Wu"}
	fmt.Println(u, u2)
}
