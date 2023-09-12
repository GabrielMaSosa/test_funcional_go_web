package handlers

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/GabrielMaSosa/test_funcional/internal/domain"
	product "github.com/GabrielMaSosa/test_funcional/internal/products"
)

var (
	ErrIdNoInteger = errors.New("Out of range ID")
)

// este es el mapa donde equivale a struct
// que va a representar el body que corresponde
var Mymap = map[string]any{
	"name":         "",
	"quantity":     0,
	"code_value":   "",
	"is_published": false,
	"expiration":   "",
	"price":        0.0,
}

func ValidateData(data domain.Product) (err error) {

	if data.Quantity <= 0 {
		err = ErrIdNoInteger
	}

	return
}

// para metodo Patch vamos a validar que sea del tipo a la estructura ProductPatch viene sin id
// este caso
// se valida que vengan datos y sean compatibles
func ValidateRequest(dta map[string]interface{}) (err error) {
	//flagNoMatch := false
	src := product.DataProductRepository{}

	st := reflect.TypeOf(src)
	fmt.Println(st)
	field0 := st.Field(0)
	field1 := st.Field(1)
	field2 := st.Field(2)
	field3 := st.Field(3)
	field4 := st.Field(4)
	field5 := st.Field(5)
	field6 := st.Field(6)

	fmt.Println("file", field1)
	// use of Get method
	fmt.Println(string(field0.Tag.Get("val")))
	fmt.Println(string(field1.Tag.Get("val")))
	fmt.Println(string(field2.Tag.Get("val")))
	fmt.Println(string(field3.Tag.Get("val")))
	fmt.Println(string(field4.Tag.Get("val")))
	fmt.Println(string(field5.Tag.Get("val")))
	fmt.Println(string(field6.Tag.Get("val")))
	for k1, v1 := range dta {

		for kp, _ := range Mymap {

			if k1 == kp {
				fmt.Println(k1, kp)
				switch {
				case k1 == "name":
					fmt.Println("estoy en ", k1)
					val, ok := v1.(string)
					if ok == false {
						err = errors.New(fmt.Sprintf("No support %s  for PATCH", v1))
						fmt.Printf("No support %s  for PATCH", v1)
					} else {
						fmt.Println(val)
						if val == "" {
							err = errors.New(fmt.Sprintf("Out of range %s  for PATCH", v1))
							fmt.Printf("Out of range %s  for PATCH", v1)
						}
					}

				case k1 == "quantity":
					fmt.Println("estoy en ", k1)
					val7, ok := v1.(float64)

					if ok == false {
						err = errors.New(fmt.Sprintf("No support   for PATCH"))
						fmt.Println("No support for PATCH", val7)
					} else {

						val8 := int(val7)
						if val8 <= 0 {

							err = errors.New(fmt.Sprintf("Out of range %s  for PATCH", v1))
							fmt.Printf("Out of range %s  for PATCH", v1)
						} else {
							dta[k1] = val8
							fmt.Println(k1, v1)
						}
					}
				case k1 == "code_value":
					fmt.Println("estoy en ", k1)
					val, ok := v1.(string)
					if ok == false {
						err = errors.New(fmt.Sprintf("No support %s  for PATCH", v1))
						fmt.Printf("No support %s  for PATCH", v1)
						fmt.Println("")
					} else {
						fmt.Println(val)
						if val == "" {
							err = errors.New(fmt.Sprintf("Out of range %s  for PATCH", v1))
							fmt.Printf("Out of range %s  for PATCH", v1)
							fmt.Println("")
						}
					}
				case k1 == "is_published":
					fmt.Println("estoy en ", k1)
					val, ok := v1.(bool)

					if ok == false {
						err = errors.New(fmt.Sprintf("No support %s  for PATCH", v1))
						fmt.Printf("No support %s  for PATCH", v1)
					} else {
						fmt.Println(val)
					}
				case k1 == "expiration":
					fmt.Println("estoy en ", k1)
					val, ok := v1.(string)
					if ok == false {
						err = errors.New(fmt.Sprintf("No support %s  for PATCH", v1))
						fmt.Printf("No support %s  for PATCH", v1)
					} else {
						fmt.Println(val)
						if val == "" {
							err = errors.New(fmt.Sprintf("Out of range %s  for PATCH", v1))
							fmt.Printf("Out of range %s  for PATCH", v1)
						} else {
							if errx := MyValidateDate(val); errx != nil {
								err = errx
							}
						}
					}
				case k1 == "price":
					fmt.Println("estoy en ", k1)
					val, ok := v1.(float64)
					if ok == false {
						err = errors.New(fmt.Sprintf("No support %s  for PATCH", v1))
						fmt.Printf("No support %s  for PATCH", v1)
					} else {
						fmt.Println(val)
						if val <= 0.0 {
							err = errors.New(fmt.Sprintf("Out of range %s  for PATCH", v1))
							fmt.Printf("Out of range %s  for PATCH", v1)
						}
					}

				}

			}
		}
	}
	fmt.Println(dta)
	fmt.Println(err)
	return
}

func MyValidateDate(data string) (err error) {

	if len(data) != 10 {
		err = errors.New("DATE NO FORMAT dd/mm/yyyy")
		return
	}
	v := data
	fmt.Println(v)
	fmt.Println(v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7], v[8], v[9])

	if v[2] != 47 && v[5] != 47 {
		err = errors.New("separate parameters invalid please use /")
	}

	dd, err2 := strconv.Atoi(v[0:2])
	if err2 != nil {
		err = errors.New("Days no valid")
	}
	fmt.Println(dd)
	mm, err3 := strconv.Atoi(v[3:5])
	if err3 != nil {
		err = errors.New("Mouth no valid")
		return
	}
	fmt.Println(mm)
	yyyy, err4 := strconv.Atoi(v[6:])
	if err4 != nil {
		err = errors.New("Year no valid")
		return
	}
	fmt.Println(yyyy)
	switch {
	case dd > 31 || dd <= 0:
		err = errors.New("days out of range ")
		return
	case mm > 12 || mm <= 0:
		err = errors.New("mouth out of range ")
		return
	case yyyy < 2023:
		err = errors.New("year out of range ")
		return
	default:
	}

	return
}

func ValidateEmpty(data *domain.Product) (ok bool, err error) {

	switch {

	case (data).Price <= 0.0:
		err = errors.New("Price empty or negative")
		ok = false
	case (data).Expiration == "":
		err = errors.New("Expiration empty")
		ok = false
	case (data).Name == "":
		err = errors.New("Name empty")
		ok = false
	case (data).Code_value == "":
		err = errors.New("Codevalue empty")
		ok = false
	case (data).Quantity <= 0:
		err = errors.New("Quantity empty o negative number")
		ok = false
	default:
		err = nil
		ok = false
		return

	}

	return
}

func ValidateCodeValue(data *domain.Product) (err error) {
	//tiene que ser todo mayuscula almenos 1 letra y un numero
	var (
		cletra = 0
		cnum   = 0
		flag   = false
	)
	for i, v := range (*data).Code_value {
		fmt.Println(v, i)
		switch {
		case i == 0 && v == 48:
			err = errors.New("El primer caracter no puede ser cero")
			return
		case v >= 65 && v <= 90:
			cletra++
			flag = true
			err = nil
		case i != 0 && v >= 48 && v <= 57:
			flag = true
			cnum++
			err = nil

		default:

		}

	}
	if cletra == 0 {
		err = errors.New("Debe haber almenos una letra")
	}
	if flag == false && cnum == 0 {
		err = errors.New("Debe haber almenos un numero")
	}
	return
}

func ValidateDate(data *domain.Product) (err error) {

	if len((*data).Expiration) != 10 {
		err = errors.New("DATE NO FORMAT dd/mm/yyyy")
		return
	}
	v := (*data).Expiration
	fmt.Println(v)
	fmt.Println(v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7], v[8], v[9])

	if v[2] != 47 && v[5] != 47 {
		err = errors.New("separate parameters invalid please use /")
	}

	dd, err2 := strconv.Atoi(v[0:2])
	if err2 != nil {
		err = errors.New("Days no valid")
	}
	fmt.Println(dd)
	mm, err3 := strconv.Atoi(v[3:5])
	if err3 != nil {
		err = errors.New("Mouth no valid")
		return
	}
	fmt.Println(mm)
	yyyy, err4 := strconv.Atoi(v[6:])
	if err4 != nil {
		err = errors.New("Year no valid")
		return
	}
	fmt.Println(yyyy)
	switch {
	case dd > 31 || dd <= 0:
		err = errors.New("days out of range ")
		return
	case mm > 12 || mm <= 0:
		err = errors.New("mouth out of range ")
		return
	case yyyy < 2023:
		err = errors.New("year out of range ")
		return
	default:
	}

	return
}
