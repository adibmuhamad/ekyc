package parser

import (
	"errors"
)

func ValidateNik(text string, nikMap NikMap) (err error) {
	if len(text) != 16 {
		return errors.New("Invalid NIK number")
	}

	provinsi := nikMap.ProvinceMap[text[0:2]]
	kabkot := nikMap.CityMap[text[0:4]]
	kec := nikMap.DistrictMap[text[0:6]]

	if provinsi != nil && kabkot != nil && kec != nil {
		return nil
	}

	return errors.New("Invalid NIK number")
}
