package detect

import (
)

type DetectFace struct {
	Valid bool `json:"valid"`
}

func FormatDetectFace(text string) (formatFace DetectFace, err error) {
	data := DetectFace{}
	data.Valid = true
	
	return data, nil
}
