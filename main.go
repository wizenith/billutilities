// todo [✓] add function abc
package billutilities

import (
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strings"
	"time"
	"unsafe"
)

func abc() {}

func StringToSliceBytes(str string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&str))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&bh))
}

func SliceByteToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{
		Data: bh.Data,
		Len:  bh.Len,
	}
	return *(*string)(unsafe.Pointer(&sh))
}

func CronFnEveryBy(t time.Duration, fn func(...interface{}), params ...interface{}) {
	for range time.Tick(t) {
		fn(params...)
	}
}

func Zip(params ...interface{}) [][]interface{} {
	boundary_length := 0

	for _, el := range params {
		curr_arr_len := reflect.ValueOf(el).Len()
		if boundary_length < curr_arr_len {
			boundary_length = curr_arr_len
		}
	}

	res_arr := make([][]interface{}, boundary_length)

	for i := 0; i < boundary_length; i++ {
		res_arr[i] = make([]interface{}, 0)
		for _, el := range params {
			ref_el := reflect.ValueOf(el)
			if ref_el.Len() > i {
				res_arr[i] = append(res_arr[i], ref_el.Index(i).Interface())
			}
		}

	}
	return res_arr
}

func Includes(params ...interface{}) bool {
	arr := reflect.ValueOf(params[0])
	target_element := reflect.ValueOf(params[1]).Interface()
	for index := 0; index < arr.Len(); index++ {
		if arr.Index(index).Interface() == target_element {
			return true
		}
	}
	return false
}

func Filter(arr interface{}, fn func(interface{}) bool) interface{} {
	ref_arr := reflect.ValueOf(arr)
	result_arr := []interface{}{}

	for i := 0; i < ref_arr.Len(); i++ {
		el := ref_arr.Index(i).Interface()
		if fn(el) {
			result_arr = append(result_arr, el)
		}
	}

	return result_arr
}

// "IndexOf | EXAMPLES"
// IndexOf([]interface{1,2,3,"Bill","Go","COOL"},"Bill")
// will reture index 3 as matched
func IndexOf(params ...interface{}) int {
	arr := reflect.ValueOf(params[0])
	target_element := reflect.ValueOf(params[1]).Interface()

	for index := 0; index < arr.Len(); index++ {
		if arr.Index(index).Interface() == target_element {
			return index
		}
	}

	return -1
}

// IndexOfAll([]int{1, 2, 1, 3, 1, 1}, 1) // [0 2 4 5]
func IndexOfAll(params ...interface{}) []int {
	arr := reflect.ValueOf(params[0])
	target_element := reflect.ValueOf(params[1]).Interface()
	result_all_match := []int{}
	for index := 0; index < arr.Len(); index++ {
		if arr.Index(index).Interface() == target_element {
			result_all_match = append(result_all_match, index)
		}
	}

	if len(result_all_match) != 0 {
		return result_all_match
	}

	return []int{-1}
}

// XProduct([]int{1,2,3},[]string{"a","b","c"}) //  [[1 a] [2 a] [3 a] [1 b] [2 b] [3 b] [1 c] [2 c] [3 c]]
func XProduct(params ...interface{}) [][]interface{} {
	arr_first := reflect.ValueOf(params[0])
	arr_second := reflect.ValueOf(params[1])
	total_compond_length := arr_first.Len() * arr_second.Len()

	result_arr := make([][]interface{}, total_compond_length)

	for i := 0; i < total_compond_length; i++ {
		result_arr[i] = []interface{}{
			arr_first.Index(i % arr_first.Len()).Interface(),
			arr_second.Index((i / arr_first.Len()) % arr_second.Len()).Interface(),
		}
	}
	return result_arr
}

// MAP
func MapGeneral(arr_i []interface{}, fn func(interface{}) interface{}) []interface{} {
	arm := make([]interface{}, len(arr_i))
	for i, el := range arr_i {
		arm[i] = fn(el)
	}

	return arm
}

func MapInt(arr []int, fn func(int) int) []int {
	arm := make([]int, len(arr))
	for i, v := range arr {
		arm[i] = fn(v)
	}

	return arm
}

func MapIntFloat64(arr []int, fn func(int) float64) []float64 {
	arm := make([]float64, len(arr))
	for i, v := range arr {
		arm[i] = fn(v)
	}

	return arm
}

func MapIntBool(arr []int, fn func(int) bool) []bool {
	arm := make([]bool, len(arr))
	for i, v := range arr {
		arm[i] = fn(v)
	}

	return arm
}

func MapIntString(arr []int, fn func(int) string) []string {
	arm := make([]string, len(arr))
	for i, v := range arr {
		arm[i] = fn(v)
	}

	return arm
}

// ints := []int{1, 2, 3}
// addTwo := func(x int) int { return x + 2 }
// dollarString := func(x int) string { return fmt.Sprintf("$%d", x) }
// MapInt(ints, addTwo) // [3 4 5]
// MapIntString(ints, dollarString) // ["$1" "$2" "$3"]

func MapFloat64(arr []float64, fn func(float64) float64) []float64 {
	arm := make([]float64, len(arr))
	for i, v := range arr {
		arm[i] = fn(v)
	}

	return arm
}

func MapFloat64Int(arr []float64, fn func(float64) int) []int {
	arm := make([]int, len(arr))
	for i, v := range arr {
		arm[i] = fn(v)
	}

	return arm
}

func MapFloat64Bool(arr []float64, fn func(float64) bool) []bool {
	arm := make([]bool, len(arr))
	for i, v := range arr {
		arm[i] = fn(v)
	}

	return arm
}

// floats := []float64{3.0, 2.0, 0.0}

// gtTwo := func(x float64) bool { return x > 2.0 }

// MapFloat64Bool(floats, gtTwo) // [true false false]

func MapFloat64String(arr []float64, fn func(float64) string) []string {
	arm := make([]string, len(arr))
	for i, v := range arr {
		arm[i] = fn(v)
	}

	return arm
}

func MapBool(arr []bool, fn func(bool) bool) []bool {
	arm := make([]bool, len(arr))
	for i, v := range arr {
		arm[i] = fn(v)
	}

	return arm
}

func MapBoolInt(arr []bool, fn func(bool) int) []int {
	arm := make([]int, len(arr))
	for i, v := range arr {
		arm[i] = fn(v)
	}

	return arm
}

func MapBoolFloat64(arr []bool, fn func(bool) float64) []float64 {
	arm := make([]float64, len(arr))
	for i, v := range arr {
		arm[i] = fn(v)
	}

	return arm
}

func MapBoolString(arr []bool, fn func(bool) string) []string {
	arm := make([]string, len(arr))
	for i, v := range arr {
		arm[i] = fn(v)
	}

	return arm
}

func MapString(arr []string, fn func(string) string) []string {
	arm := make([]string, len(arr))
	for i, v := range arr {
		arm[i] = fn(v)
	}

	return arm
}

func MapStringInt(arr []string, fn func(string) int) []int {
	arm := make([]int, len(arr))
	for i, v := range arr {
		arm[i] = fn(v)
	}

	return arm
}

// Example of MapStringInt
// strings := []string{"go", "is", "cool"}

// strLen := func(x string) int { return len(x) }

// MapStringInt(strings, strLen) // [2 2 4]

func MapStringFloat64(arr []string, fn func(string) float64) []float64 {
	arm := make([]float64, len(arr))
	for i, v := range arr {
		arm[i] = fn(v)
	}

	return arm
}

func MapStringBool(arr []string, fn func(string) bool) []bool {
	arm := make([]bool, len(arr))
	for i, v := range arr {
		arm[i] = fn(v)
	}

	return arm
}

// Dedupe
// EXAMPLES
// DedupeInts([]int{1, 2, 1, 2, 3, 3, 4}) // [1 2 3 4]
func DedupGeneral(arr []interface{}) []interface{} {
	m := make(map[interface{}]bool)
	// arr_uniq := []interface{}{}
	arr_uniq := make([]interface{}, 0)

	for _, interface_el := range arr {
		if _, el_occurs := m[interface_el]; !el_occurs {
			m[interface_el] = true
			arr_uniq = append(arr_uniq, interface_el)
		}
	}
	return arr_uniq
}

func DedupeInts(arr []int) []int {

	m, uniq := make(map[int]bool), make([]int, 0)

	for _, int_el := range arr {
		if _, el_occurs := m[int_el]; !el_occurs {
			m[int_el], uniq = true, append(uniq, int_el)
		}
	}
	return uniq
}

func DedupFloat64(arr []float64) []float64 {
	m := make(map[float64]bool)
	uniq := make([]float64, 0)

	for _, f64_el := range arr {
		if _, el_occurs := m[f64_el]; !el_occurs {
			m[f64_el], uniq = true, append(uniq, f64_el)

		}
	}
	return uniq
}

func DedupString(arr []string) []string {
	m := make(map[string]bool)
	arr_uniq := make([]string, 0)

	for _, string_el := range arr {
		if _, el_occurs := m[string_el]; !el_occurs {
			m[string_el], arr_uniq = true, append(arr_uniq, string_el)
		}
	}
	return arr_uniq
}

// CompactWhiteSpace | EXAMPLES
// CompactWhiteSpace("Lorem    Ipsum") // "Lorem Ipsum"
// CompactWhiteSpace("Lorem \n Ipsum") // "Lorem Ipsum"
func CompactWhiteSpace(str string) string {
	regex := regexp.MustCompile(`\s{2,}`)
	return regex.ReplaceAllString(str, " ")
}

func MaxOf(nums ...float64) float64 {
	max := math.Inf(-1)
	for _, num := range nums {
		max = math.Max(num, max)
	}
	return max
}

func ShowOnlyLast_N_Digit(str string, n_digit int, cover_with rune) string {
	return strings.Repeat(string(cover_with), len(str)-n_digit) + str[len(str)-n_digit:]
}

func Average(nums ...float64) float64 {
	var sum float64
	for _, num := range nums {
		sum += num
	}
	return (sum / float64(len(nums)))
}

func Capital(s string) string {
	return strings.ToUpper(s[0:1]) + s[1:]
}

func Stringify(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func IsEven(num int) bool {
	return num%2 == 0
}

func IsOdd(num int) bool {
	return num%2 == 1
}

func IsPowerOf2(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}
