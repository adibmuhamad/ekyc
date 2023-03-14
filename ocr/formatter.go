package ocr

import (
	"regexp"
	"strings"
)

type OcrKtp struct {
	NumberID      string `json:"idNumber"`
	Name          string `json:"name"`
	Religion      string `json:"religion"`
	Gender        string `json:"gender"`
	Province      string `json:"province"`
	City          string `json:"city"`
	District      string `json:"district"`
	Village       string `json:"village"`
	RtRw          string `json:"rtrw"`
	Occupation    string `json:"occupation"`
	ExpiryDate    string `json:"expiryDate"`
	Nationality   string `json:"nationality"`
	MaritalStatus string `json:"maritalStatus"`
	Address       string `json:"address"`
	PlaceOfBirth  string `json:"placeOfBirth"`
	DateOfBirth   string `json:"dateOfBirth"`
}

type OcrNpwp struct {
	NumberID string `json:"idNumber"`
	Name     string `json:"name"`
}

type OcrSim struct {
	NumberID     string `json:"idNumber"`
	SimType      string `json:"simType"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Province     string `json:"province"`
	DistrictCity     string `json:"districtCity"`
	PlaceOfBirth string `json:"placeOfBirth"`
	DateOfBirth  string `json:"dateOfBirth"`
	Gender       string `json:"gender"`
	Occupation   string `json:"occupation"`
	MB           string `json:"mb"`
}

func FormatDataKtp(text string) (formatKtp OcrKtp, err error) {

	var (
		datas     []string
		reg       *regexp.Regexp
		countData int
	)

	// Split string
	datas = strings.Split(strings.Replace(text, "\n\n", "\n", -1), "\n")

	// Regex config
	reg, _ = regexp.Compile("[^-A-Za-z0-9/.,'` ]+")

	//looping data split
	for _, data := range datas {

		obj := strings.Split(data, ":")

		if len(obj) > 1 {
			data = strings.TrimSpace(obj[1])
			data = reg.ReplaceAllString(data, "")
		} else {
			data = reg.ReplaceAllString(data, "")
		}

		if countData == 0 {
			replacer := strings.NewReplacer("PROVINSI", "", "PROVINS", "")
			formatKtp.Province = strings.TrimSpace(replacer.Replace(data))
		}

		if countData == 1 {
			replacer := strings.NewReplacer("KABUPATEN", "")
			formatKtp.City = strings.TrimSpace(replacer.Replace(data))
		}

		if countData == 2 {
			replacer := strings.NewReplacer("NIK", "")
			formatKtp.NumberID = strings.TrimSpace(replacer.Replace(data))
		}

		if countData == 3 {
			replacer := strings.NewReplacer("Nama", "")
			formatKtp.Name = strings.TrimSpace(replacer.Replace(data))
		}

		if countData == 4 {
			replacer := strings.NewReplacer("Tempat/Tgl-Lahir", "", "Tempat/Tgi-Lahir", "", "Tempat/Tgi Lahir", "")
			birthday := strings.Split(replacer.Replace(data), ",")

			formatKtp.PlaceOfBirth = strings.TrimSpace(birthday[0])
			if len(birthday) > 1 {
				formatKtp.DateOfBirth = strings.TrimSpace(birthday[1])
			}
		}

		if countData == 5 {

			replacer := strings.NewReplacer("Jenis Kelamin", "", "denis-Kelamin", "", "Gol-Darah", "", "Gol Darah", "")
			formatKtp.Gender = strings.TrimSpace(replacer.Replace(data))
		}

		if countData == 6 {
			replacer := strings.NewReplacer("Alamat", "")
			formatKtp.Address = strings.TrimSpace(replacer.Replace(data))
		}

		if countData == 7 {
			replacer := strings.NewReplacer("RT/RW", "", "RTIRW", "")
			formatKtp.RtRw = strings.TrimSpace(replacer.Replace(data))
		}

		if countData == 8 {

			formatKtp.Village = strings.TrimSpace(data)
		}

		if countData == 9 {
			replacer := strings.NewReplacer("Kecamatan", "")
			formatKtp.District = strings.TrimSpace(replacer.Replace(data))
		}

		if countData == 10 {
			replacer := strings.NewReplacer("Agama", "")
			formatKtp.Religion = strings.TrimSpace(replacer.Replace(data))
		}

		if countData == 11 {
			formatKtp.MaritalStatus = strings.TrimSpace(data)
		}

		if countData == 12 {
			replacer := strings.NewReplacer("Pekerjaan", "")
			formatKtp.Occupation = strings.TrimSpace(replacer.Replace(data))
		}

		if countData == 13 {
			formatKtp.Nationality = strings.TrimSpace(data)
		}

		if countData == 14 {
			formatKtp.ExpiryDate = strings.TrimSpace(data)
		}

		countData++
	}

	return formatKtp, nil
}

func FormatDataNpwp(text string) (formatNpwp OcrNpwp, err error) {

	var (
		datas     []string
		reg       *regexp.Regexp
		countData int
	)

	// Split string
	datas = strings.Split(strings.Replace(text, "\n\n", "\n", -1), "\n")

	// Regex config
	reg, _ = regexp.Compile("[^-A-Za-z0-9/.,'` ]+")

	//looping data split
	for _, data := range datas {

		obj := strings.Split(data, ":")

		if len(obj) > 1 {
			data = strings.TrimSpace(obj[1])
			data = reg.ReplaceAllString(data, "")
		} else {
			data = reg.ReplaceAllString(data, "")
		}

		if countData == 2 {
			replacer := strings.NewReplacer("NPWP", "")
			formatNpwp.NumberID = strings.TrimSpace(replacer.Replace(data))
		}

		if countData == 3 {
			// replacer := strings.NewReplacer("KABUPATEN", "")
			formatNpwp.Name = strings.TrimSpace(data)
		}

		countData++
	}

	return formatNpwp, nil
}

func FormatDataSim(text string) (formatSim OcrSim, err error) {
	var (
		datas     []string
		reg       *regexp.Regexp
		countData int
	)

	// Split string
	datas = strings.Split(strings.Replace(text, "\n\n", "\n", -1), "\n")

	// Regex config
	reg, _ = regexp.Compile("[^-A-Za-z0-9/.,'`: ]+")

	// looping data split
	for _, data := range datas {

		data = reg.ReplaceAllString(data, "")

		if countData == 1 {
			a := strings.Split(data ,"SURAT IZIN MENGEMUDI")
			formatSim.SimType = strings.TrimSpace(a[1])
		}

		if countData == 2 {
			formatSim.NumberID = strings.TrimSpace(data)
		}

		if countData == 3 {
			a := strings.Split(data ,".")
			formatSim.Name = strings.TrimSpace(a[1])
		}

		if countData == 4 {
			a := strings.Split(data ,".")
			b := strings.Split(a[1] ,",")
			formatSim.PlaceOfBirth = strings.TrimSpace(b[0])
			formatSim.DateOfBirth = strings.TrimSpace(b[1])
		}

		if countData == 5 {
			a := strings.Split(data ,"-")
			formatSim.Gender = strings.TrimSpace(a[1])
		}

		if countData == 6 {
			a := strings.Split(data ,".")
			formatSim.Address = strings.TrimSpace(a[1])
		}

		if countData == 7 {
			formatSim.DistrictCity = strings.TrimSpace(data)
		}

		if countData == 8 {
			formatSim.Province = strings.TrimSpace(data)
		}

		if countData == 9 {
			a := strings.Split(data ,".")
			formatSim.Occupation = strings.TrimSpace(a[1])
		}

		if countData == 10 {
			a := strings.Split(data ,".")
			formatSim.MB = strings.TrimSpace(a[1])
		}

		countData++
	}

	return formatSim, nil
}
