package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"sort"
	"strings"

	utils "github.com/adhityapp/go-starterkit/pkg/utils"

	"github.com/labstack/echo"
)

func (uh HandlerClient) GetSmithActiveEmployee(ectx echo.Context) error {
	// var resp Response
	ctx := ectx.Request().Context()
	data, err := uh.service.GetSmithActiveEmployee(ctx)
	if err != nil {
		resp := Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    data,
		}
		return ectx.JSON(http.StatusInternalServerError, resp)
	}

	resp := Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    data,
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		resp := Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return ectx.JSON(http.StatusInternalServerError, resp)
	}

	err = utils.ToTxt(jsonData, "txt/contoh2.txt")
	if err != nil {
		resp := Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return ectx.JSON(http.StatusInternalServerError, resp)
	}

	return ectx.JSON(http.StatusOK, resp)
}

func (uh HandlerClient) GetEmployeeNoReview(ectx echo.Context) error {
	ctx := ectx.Request().Context()
	data, err := uh.service.GetEmployeeNoReview(ctx)
	if err != nil {
		resp := Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    data,
		}
		return ectx.JSON(http.StatusInternalServerError, resp)
	}

	resp := Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    data,
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		resp := Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    data,
		}
		return ectx.JSON(http.StatusInternalServerError, resp)
	}

	err = utils.ToTxt(jsonData, "txt/contoh3.txt")
	if err != nil {
		resp := Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    data,
		}
		return ectx.JSON(http.StatusInternalServerError, resp)
	}

	return ectx.JSON(http.StatusOK, resp)
}

type DiffDay struct {
	DifferentDay int `json:"different_day"`
}

func (uh HandlerClient) GetEmployeeDifferentDay(ectx echo.Context) error {
	ctx := ectx.Request().Context()
	data, err := uh.service.GetEmployeeDifferentDay(ctx)
	if err != nil {
		resp := Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    data,
		}
		return ectx.JSON(http.StatusInternalServerError, resp)
	}

	diff := DiffDay{
		DifferentDay: *data,
	}

	resp := Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    diff,
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		resp := Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    data,
		}
		return ectx.JSON(http.StatusInternalServerError, resp)
	}

	err = utils.ToTxt(jsonData, "txt/contoh4.txt")
	if err != nil {
		resp := Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    data,
		}
		return ectx.JSON(http.StatusInternalServerError, resp)
	}

	return ectx.JSON(http.StatusOK, resp)
}

func (uh HandlerClient) GetSalary(ectx echo.Context) error {
	ctx := ectx.Request().Context()
	data, err := uh.service.GetSalary(ctx)
	if err != nil {
		resp := Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    data,
		}
		return ectx.JSON(http.StatusInternalServerError, resp)
	}

	resp := Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    data,
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		resp := Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    data,
		}
		return ectx.JSON(http.StatusInternalServerError, resp)
	}

	err = utils.ToTxt(jsonData, "txt/contoh5.txt")
	if err != nil {
		resp := Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    data,
		}
		return ectx.JSON(http.StatusInternalServerError, resp)
	}

	return ectx.JSON(http.StatusOK, resp)
}

type FileRequest struct {
	Filename string `json:"filename"`
}

func (uh HandlerClient) GetTxt(ectx echo.Context) error {
	var req FileRequest
	err := ectx.Bind(&req)
	if err != nil {
		resp := Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}
		return ectx.JSON(http.StatusInternalServerError, resp)
	}

	file, err := os.Open(req.Filename)
	if err != nil {
		resp := Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}
		return ectx.JSON(http.StatusInternalServerError, resp)
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		resp := Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}
		return ectx.JSON(http.StatusInternalServerError, resp)
	}

	var resp Response
	err = json.Unmarshal(content, &resp)
	if err != nil {
		resp := Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}
		return ectx.JSON(http.StatusInternalServerError, resp)
	}

	return ectx.JSON(http.StatusOK, resp)
}

type City struct {
	City string `json:"kota"`
}

type CityResp struct {
	Result    bool   `json:"result"`
	Recommend string `json:"saran_kota"`
}

func (uh HandlerClient) GetCity(ectx echo.Context) error {
	var req City
	var resp CityResp
	result := false

	err := ectx.Bind(&req)
	if err != nil {
		resp := Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}
		return ectx.JSON(http.StatusInternalServerError, resp)
	}

	array := []string{"Bandung", "Cimahi", "Ambon", "Jayapura", "Makasar"}

	for _, c := range array {
		if strings.EqualFold(c, req.City) {
			result = true
			return ectx.JSON(http.StatusOK, result)
		}
	}

	var f, l string
	inputLower := strings.ToLower(req.City)
	for _, c := range array {
		cLower := strings.ToLower(c)
		if strings.HasPrefix(cLower, inputLower[:1]) {
			f = f + c + ", "
		}
		if strings.HasSuffix(cLower, inputLower[len(inputLower)-1:]) {
			l = l + c + ", "
		}
	}

	resp.Result = result
	resp.Recommend = f + l

	response := Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    resp,
	}

	return ectx.JSON(http.StatusOK, response)
}

type Array struct {
	Type  string
	Array []int
	Value int
}

func (uh HandlerClient) GetArray(ectx echo.Context) error {
	var req Array
	err := ectx.Bind(&req)
	if err != nil {
		resp := Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		}
		return ectx.JSON(http.StatusInternalServerError, resp)
	}

	array := []int{9, 1, 6, 4, 8, 6, 6, 3, 8, 2, 3, 3, 4, 1, 8, 2}

	if req.Type == "a" {
		uniqueValue := make(map[int]bool)
		for _, value := range array {
			uniqueValue[value] = true
		}

		var result []int
		for key := range uniqueValue {
			result = append(result, key)
		}
		sort.Ints(result)

		response := Response{
			Status:  http.StatusOK,
			Message: "Success",
			Data:    result,
		}

		return ectx.JSON(http.StatusOK, response)

	}
	if req.Type == "b" {

		frequency := make(map[int]int)
		for _, value := range array {
			frequency[value] = frequency[value] + 1
		}

		var keys []int
		for key := range frequency {
			keys = append(keys, key)
		}
		sort.Ints(keys)

		result := ""
		for _, key := range keys {
			result += fmt.Sprintf("%d[%d],", key, frequency[key])
		}

		response := Response{
			Status:  http.StatusOK,
			Message: "Success",
			Data:    result,
		}

		return ectx.JSON(http.StatusOK, response)

	}
	if req.Type == "c" {

		removeIndex := make(map[int]bool)
		for _, value := range req.Array {
			removeIndex[value] = true
		}

		var result []int
		for _, value := range array {
			if !removeIndex[value] {
				result = append(result, value)
			}
		}
		response := Response{
			Status:  http.StatusOK,
			Message: "Success",
			Data:    result,
		}

		return ectx.JSON(http.StatusOK, response)

	}
	if req.Type == "d" {

		val := req.Value

		for i, v := range array {
			newValue := v + val
			if newValue > 10 {
				array[i] = 10
				val = newValue - 10
			} else {
				array[i] = newValue
				break
			}
		}

		response := Response{
			Status:  http.StatusOK,
			Message: "Success",
			Data:    array,
		}

		return ectx.JSON(http.StatusOK, response)

	}

	response := Response{
		Status:  http.StatusBadRequest,
		Message: "Type Not found",
	}

	return ectx.JSON(http.StatusBadRequest, response)

}

func randomChar(isLetter bool) byte {
	if isLetter {
		if rand.Intn(2) == 0 {
			return byte('A' + rand.Intn(26))
		} else {
			return byte('a' + rand.Intn(26))
		}
	} else {
		return byte('0' + rand.Intn(10))
	}
}

func contains(str string, char rune) bool {
	for _, c := range str {
		if c == char {
			return true
		}
	}
	return false
}

type RandString struct {
	TotalLetters    int
	TotalVowels     int
	TotalDigits     int
	TotalEvenDigits int
	UrutanB         string
	UrutanC         string
}

func (uh HandlerClient) GetRandomString(ectx echo.Context) error {

	letters := make([]byte, 50)
	numbers := make([]byte, 50)

	// Menghasilkan 50 huruf acak
	for i := range letters {
		letters[i] = randomChar(true)
	}

	// Menghasilkan 50 angka acak
	for i := range numbers {
		numbers[i] = randomChar(false)
	}

	// Menggabungkan huruf dan angka
	result := append(letters, numbers...)

	vowels := "AEIOUaeiou"

	var m RandString
	var num []int
	var let []string

	for _, char := range result {
		if char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z' {
			let = append(let, string(char))
			m.TotalLetters++
			if contains(vowels, rune(char)) {
				m.TotalVowels++
			}
		} else if char >= '0' && char <= '9' {
			num = append(num, int(char-'0'))
			m.TotalDigits++
			if (char-'0')%2 == 0 {
				m.TotalEvenDigits++
			}
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(num)))
	sort.Strings(let)

	numnotdup := removeDuplicates(num)
	letnotdup := removeDuplicatesString(let)

	var sortedResultB []string
	for _, numval := range numnotdup {
		sortedResultB = append(sortedResultB, fmt.Sprintf("%d", numval))
	}
	sortedResultB = append(sortedResultB, letnotdup...)
	m.UrutanB = strings.Join(sortedResultB, ",")

	var sortedResultC []string
	for _, num := range num {
		for _, letval := range let {
			sortedResultC = append(sortedResultC, fmt.Sprintf("%d%s", num, letval))
		}
	}

	m.UrutanC = strings.Join(sortedResultC, ",")

	resp := Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    m,
	}
	return ectx.JSON(http.StatusOK, resp)
}

func removeDuplicates(numbers []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range numbers {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func removeDuplicatesString(letters []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range letters {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
