package parser

import (
	"encoding/json"
	"id/projects/ekyc/helper"
	"io/ioutil"
)

type service struct {
}

type Service interface {
	ParseDataKtp(input ParserInput) (Nik, error)
	ParseDataNpwp(input ParserInput) (Npwp, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) ParseDataKtp(input ParserInput) (Nik, error) {
	nikMap := NikMap{}

	// Read file wilayah
	err := nikMap.ReadFileWilayah()
	if err != nil {
		return Nik{}, err
	}

	// Validate NIK number
	err = ValidateNik(input.NumberID, nikMap)
	if err != nil {
		return Nik{}, err
	}

	nik := Nik{}
	nik, err = FormatDataNik(input.NumberID, nikMap)
	if err != nil {
		return Nik{}, err
	}

	return nik, nil

}

func (s *service) ParseDataNpwp(input ParserInput) (Npwp, error) {
	npwp := Npwp{}

	// Validate NPWP number
	err := ValidateNpwp(input.NumberID)
	if err != nil {
		return Npwp{}, err
	}

	npwp, err = FormatDataNpwp(input.NumberID)
	if err != nil {
		return Npwp{}, err
	}

	return npwp, nil

}

func (n *NikMap) ReadFileWilayah() error {
	file, err := ioutil.ReadFile("assets/wilayah.json")
	if err != nil {
		return err
	}
	var a map[string]interface{}
	err = json.Unmarshal(file, &a)
	if err != nil {
		return err
	}

	province, err := helper.Marshaler(a["province"])
	if err != nil {
		return err
	}
	n.ProvinceMap = province

	city, err := helper.Marshaler(a["city"])
	if err != nil {
		return err
	}
	n.CityMap = city

	district, err := helper.Marshaler(a["district"])
	if err != nil {
		return err
	}
	n.DistrictMap = district

	return nil
}
