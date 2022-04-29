package algorithm

import (
	"fmt"
	_ "github.com/schwarmco/go-cartesian-product"
	"sort"
	"strings"
	"unicode"
)

var (
	alphabet                = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	frequencyLettersEnglish = map[string]float64{
		"A": 0.082,
		"B": 0.014,
		"C": 0.028,
		"D": 0.038,
		"E": 0.131,
		"F": 0.029,
		"G": 0.020,
		"H": 0.053,
		"I": 0.064,
		"J": 0.001,
		"K": 0.004,
		"L": 0.034,
		"M": 0.025,
		"N": 0.071,
		"O": 0.080,
		"P": 0.020,
		"Q": 0.001,
		"R": 0.068,
		"S": 0.061,
		"T": 0.105,
		"U": 0.025,
		"V": 0.009,
		"W": 0.015,
		"X": 0.002,
		"Y": 0.020,
		"Z": 0.001,
	}
)

type Best3Hi struct {
	f1, f2, f3 int
}

func Encrypt(text string, key string, size int) string {
	var encrypted string
	var chKey byte

	text, key = strings.ToUpper(text), strings.ToUpper(key)

	for i, ch := range text {
		chKey = key[i%len(key)]
		encrypted += string((((int(ch) - 'A') + (int(chKey) - 'A')) % size) + 'A')
	}

	return encrypted

}

func Decrypt(text string, key string, size int) string {
	var decrypted string
	var chKey byte
	text, key = strings.ToUpper(text), strings.ToUpper(key)

	for i, ch := range text {
		chKey = key[i%len(key)]
		decrypted += string(((((int(ch) - 'A') - (int(chKey) - 'A')) + size) % size) + 'A')
	}

	return decrypted

}

func IndexIC(text string) (IC float64) {
	characters := make(map[string]float64)
	for _, symbol := range text {
		characters[string(symbol)]++
	}
	//fmt.Println(characters)
	IC = 0.0
	N := len(text)
	for _, value := range characters {
		IC += (value * (value - 1)) / float64(N*(N-1))
	}
	return
}

func FindKeyLength(text string) int {
	textSecond := ""
	//keyLength := 0
	avgICarr := make([]float64, 0)
	for i := 2; i < 10; i++ {
		cancelString := text
		avgIC := 0.0
		//fmt.Printf("if key were length %d\n", i)
		for kLen := 0; kLen < i; kLen++ {
			for j := 0; j < len(cancelString); j++ {
				if j%i == 0 {
					textSecond += string(cancelString[j])
				}
			}
			IC := IndexIC(textSecond)
			avgIC += IC
			cancelString = cancelString[1:] + string(cancelString[0])
			//fmt.Println(textSecond, cancelString)
			textSecond = ""
		}
		avgICarr = append(avgICarr, avgIC/float64(i))

	}
	fmt.Println("-------------------------------------------")
	fmt.Println("Length key  |  Average Index of Coincidence")
	fmt.Println("-------------------------------------------")
	for index, value := range avgICarr {
		fmt.Printf("key: %d      | %f\n", index+2, value)
	}
	fmt.Println("-------------------------------------------")
	return len(avgICarr)
}

func ShiftClass(shiftBase string) string {
	output := ""
	for _, value := range shiftBase {
		if value == 'A' {
			output += string('Z')
			continue
		}
		output += string((value - 1) % 'Z')
	}
	return output
}

func hiMethodAll(shift string) float64 {
	characters := make(map[string]float64)
	hi := 0.0
	for _, symbol := range shift {
		characters[string(symbol)]++
	}

	lenShift := float64(len(shift))

	for key, value := range frequencyLettersEnglish {
		value2 := 0.0
		if _, ok := characters[key]; ok {
			value2 = characters[key]
		}
		//fmt.Printf("key %s, value %g Fi: %g\n", key, value2/lenShift, value)
		//hiShift[key] = (value2- value*lenShift)*(value2-value*lenShift)/(value*lenShift)
		hi += ((value2/lenShift - value) * (value2/lenShift - value)) / (value)
		//((characters[key]/lenShift - value) * (characters[key]/lenShift - value)) / value
	}
	return hi
}

func printHI(hi [][]float64) {
	fmt.Println("----------------------------------------------------")
	fmt.Printf("Shift\t|Letter\t|")
	for i := 0; i < len(hi); i++ {
		fmt.Printf("Coset %d χ2\t| ", i+1)
	}
	fmt.Printf("\n---------------------------------------------------\n")
	for i := 0; i < 26; i++ {
		fmt.Printf("%d\t| %s\t|", i, string(rune('A'+i)))
		for j := 0; j < len(hi); j++ {
			a := hi[j]
			fmt.Printf("| %.4f \t|", a[i])
		}
		fmt.Println()
	}
	fmt.Println("----------------------------------------------------")
}

func minimumHI(arr []float64) string {
	min := arr[0]
	index := 0
	for i, value := range arr {
		if min > value {
			min = value
			index = i
		}
	}
	return string(rune('A' + index))
}

// find KEY
func FindKey(length int, text string) (string, [][]float64) {
	cancelString := text
	textSecond := ""
	table := make([][]float64, 0)
	keyWord := ""
	for kLen := 0; kLen < length; kLen++ {
		for j := 0; j < len(cancelString); j++ {
			if j%length == 0 {
				textSecond += string(cancelString[j])
			}
		}
		arrHI := make([]float64, 0)
		for symb := 0; symb < 26; symb++ {
			avgHI := hiMethodAll(textSecond)
			arrHI = append(arrHI, avgHI)
			textSecond = ShiftClass(textSecond)
		}
		table = append(table, arrHI)
		keyWord += minimumHI(arrHI)

		//fmt.Println("Average HI: ", avgHI)
		//fmt.Println(textSecond)

		textSecond = ""
		cancelString = cancelString[1:]

	}
	printHI(table)
	//fmt.Println(keyWord)
	// fmt.Println(table)
	return keyWord, table
}

func ChangeText(text string) string {
	output := ""
	for _, value := range text {
		if unicode.IsLetter(value) {
			output += string(value)
		}
	}
	return output
}

func ChangeKey(key string) string {
	output := ""
	for _, value := range key {
		if unicode.IsLetter(value) {
			output += string(value)
		}
	}
	return output
}

func IsCorrectText(text string) error {
	if text == "" {
		return fmt.Errorf("Error: text is empty")
	}
	return nil
}

func IsCorrectKey(key string) error {
	if key == "" {
		return fmt.Errorf("Error: key is empty or contain non letter symbol")
	}

	return nil
}

func CheckLengthKey(lengthKey, rangeLength int) error {
	if lengthKey < 2 || lengthKey > rangeLength+1 {
		return fmt.Errorf("key length invalid")
	}
	return nil
}

func Best3HiAction(a []float64) (float64, float64, float64) {
	sort.Float64s(a)
	return a[0], a[1], a[2]
}

func nextIndex(ix []int, lengthsFunc func(i int) int) {
	for j := len(ix) - 1; j >= 0; j-- {
		ix[j]++
		if j == 0 || ix[j] < lengthsFunc(j) {
			return
		}
		ix[j] = 0
	}
}

func Product(sets ...[]int) [][]int {
	lengths := func(i int) int { return len(sets[i]) }
	var product [][]int
	for ix := make([]int, len(sets)); ix[0] < lengths(0); nextIndex(ix, lengths) {
		var r []int

		for j, k := range ix {
			r = append(r, sets[j][k])
		}
		product = append(product, r)
	}
	return product
}

func AllKeys(table [][]float64) []string {
	tableMinHi := Min3HI(table)
	res := Product(tableMinHi...)

	keys := make([]string, 0)
	for _, value := range res {
		str := ""
		for _, shift := range value {
			str += string(rune('A' + shift))
		}
		keys = append(keys, str)
	}
	return keys
}

func Min3HI(table [][]float64) [][]int {
	arr := make([][]int, 0)
	for i := 0; i < len(table); i++ {

		a := table[i]
		b := make([]float64, len(a))
		copy(b, a)
		min1, min2, min3 := Best3HiAction(b)
		index1, index2, index3 := -1, -1, -1
		for j := 0; j < len(a); j++ {
			if a[j] == min1 {
				index1 = j
			} else if a[j] == min2 {
				index2 = j
			} else if a[j] == min3 {
				index3 = j
			}
		}
		arrSup := []int{index1, index2, index3}
		arr = append(arr, arrSup)
	}
	return arr
	// printPermutations(arr)
}

func transpose(A [][]int) [][]int {
	B := make([][]int, len(A[0]))
	for i := 0; i < len(A[0]); i++ {
		B[i] = make([]int, len(A))
		for j := 0; j < len(A); j++ {
			B[i][j] = A[j][i]
		}
	}
	return B
}

//func cartesianProduct
//
//func printPermutations(arr [][]int) {
//	//arr = transpose(arr)
//	a := arr[0]
//
//
//	for i := 1; i < len(arr)-1; i++ {
//		b := arr[i]
//		cartesian.Iter(a, b)
//	}
//}
