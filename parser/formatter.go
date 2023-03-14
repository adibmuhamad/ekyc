package parser

import (
	"id/projects/ekyc/helper"
	"strconv"
	"strings"
	"time"
)

type NikMap struct {
	ProvinceMap map[string]interface{}
	CityMap     map[string]interface{}
	DistrictMap map[string]interface{}
}

type Nik struct {
	NumberId    string `json:"idNumber"`
	DateOfBirth string `json:"dateOfBirth"`
	District    string `json:"district"`
	City        string `json:"city"`
	Province    string `json:"province"`
	PostalCode  string `json:"postalCode"`
	IsFemale    bool   `json:"isFemale"`
	Unicode     string `json:"unicode"`
}

type Npwp struct {
	NumberId           string `json:"idNumber"`
	TaxPayerCode       string `json:"taxPayerCode"`
	Serial             string `json:"serial"`
	CheckDigit         string `json:"checkDigit"`
	LocalTaxOfficeCode string `json:"localTaxOfficeCode"`
	BranchCode         string `json:"branchCode"`
}

func FormatDataNik(text string, nikMap NikMap) (Nik, error) {
	format := Nik{}
	format.NumberId = text

	format.Province = nikMap.ProvinceMap[text[0:2]].(string)
	format.City = nikMap.CityMap[text[0:4]].(string)
	district := nikMap.DistrictMap[text[0:6]].(string)
	format.District = GetDistrict(district)
	format.PostalCode = GetPostalCode(district)

	format.Unicode = text[12:16]

	atoi, _ := strconv.Atoi(text[6:8])
	format.IsFemale = atoi > 40
	format.DateOfBirth = GetDateOfBirth(text)

	return format, nil

}

func FormatDataNpwp(text string) (Npwp, error) {
	format := Npwp{}
	format.NumberId = text

	a := strings.Split(text, ".")

	format.TaxPayerCode = a[0]
	format.Serial = a[1] + a[2]
	b := strings.Split(a[3], "-")
	format.CheckDigit = b[0]

	format.LocalTaxOfficeCode = b[1]
	format.BranchCode = a[4]

	return format, nil

}

func GetBornDay(text string) int {
	atoi, _ := strconv.Atoi(text[6:8])

	if atoi > 40 {
		print(text[6:8])
		atoi -= 40
	}
	return atoi
}

func GetBornMonth(text string) int {
	atoi, _ := strconv.Atoi(text[8:10])
	return atoi
}

func GetBornYear(text string) int {
	t, _ := strconv.Atoi(time.Now().Format("06"))
	nikT, _ := strconv.Atoi(text[10:12])
	if nikT < t {
		nikT += 2000
	} else {
		nikT += 1900
	}
	return nikT
}

func GetDateOfBirth(text string) string {
	l := strconv.Itoa(GetBornDay(text)) + " " + helper.ConvertMonth(GetBornMonth(text)) + " " + strconv.Itoa(GetBornYear(text))

	return l
}

func GetDistrict(text string) string {
	a := strings.Split(text, "--")
	return strings.Trim(a[0], " ")
}

func GetPostalCode(text string) string {
	a := strings.Split(text, "--")
	return strings.Trim(a[1], " ")
}
