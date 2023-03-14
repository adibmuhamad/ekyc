package parser

import (
	"errors"
	"strings"
)

func ValidateNik(text string, nikMap AddressMap) (err error) {
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

func ValidateNpwp(text string) (err error) {
	a := strings.Split(text, ".")

	if len(a) == 5 {

		b := strings.Split(a[3], "-")
		if len(b) != 2 {
			return errors.New("Invalid NPWP number")
		}
		if len(a[0]) != 2 && len(a[1]) != 3 && len(a[2]) != 3 && len(b[0]) != 1 && len(b[1]) != 3 && len(a[4]) != 3 {
			return errors.New("Invalid NPWP number")
		}

		return nil

	}

	return errors.New("Invalid NPWP number")
}

func ValidateSim(text string) (err error) {
	a := strings.Split(text, "-")

	if len(a) == 3 {
		return nil

	}

	return errors.New("Invalid SIM number")
}
