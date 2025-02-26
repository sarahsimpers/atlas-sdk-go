// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
	"fmt"
)

// FTSAnalyzersTokenFiltersInner struct for FTSAnalyzersTokenFiltersInner
type FTSAnalyzersTokenFiltersInner struct {
	TokenFilterasciiFolding          *TokenFilterasciiFolding
	TokenFilterdaitchMokotoffSoundex *TokenFilterdaitchMokotoffSoundex
	TokenFilteredgeGram              *TokenFilteredgeGram
	TokenFiltericuFolding            *TokenFiltericuFolding
	TokenFiltericuNormalizer         *TokenFiltericuNormalizer
	TokenFilterlength                *TokenFilterlength
	TokenFilterlowercase             *TokenFilterlowercase
	TokenFilternGram                 *TokenFilternGram
	TokenFilterregex                 *TokenFilterregex
	TokenFilterreverse               *TokenFilterreverse
	TokenFiltershingle               *TokenFiltershingle
	TokenFiltersnowballStemming      *TokenFiltersnowballStemming
	TokenFilterstopword              *TokenFilterstopword
	TokenFiltertrim                  *TokenFiltertrim
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *FTSAnalyzersTokenFiltersInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into TokenFilterasciiFolding
	err = json.Unmarshal(data, &dst.TokenFilterasciiFolding)
	if err == nil {
		jsonTokenFilterasciiFolding, _ := json.Marshal(dst.TokenFilterasciiFolding)
		if string(jsonTokenFilterasciiFolding) == "{}" { // empty struct
			dst.TokenFilterasciiFolding = nil
		} else {
			return nil // data stored in dst.TokenFilterasciiFolding, return on the first match
		}
	} else {
		dst.TokenFilterasciiFolding = nil
	}

	// try to unmarshal JSON data into TokenFilterdaitchMokotoffSoundex
	err = json.Unmarshal(data, &dst.TokenFilterdaitchMokotoffSoundex)
	if err == nil {
		jsonTokenFilterdaitchMokotoffSoundex, _ := json.Marshal(dst.TokenFilterdaitchMokotoffSoundex)
		if string(jsonTokenFilterdaitchMokotoffSoundex) == "{}" { // empty struct
			dst.TokenFilterdaitchMokotoffSoundex = nil
		} else {
			return nil // data stored in dst.TokenFilterdaitchMokotoffSoundex, return on the first match
		}
	} else {
		dst.TokenFilterdaitchMokotoffSoundex = nil
	}

	// try to unmarshal JSON data into TokenFilteredgeGram
	err = json.Unmarshal(data, &dst.TokenFilteredgeGram)
	if err == nil {
		jsonTokenFilteredgeGram, _ := json.Marshal(dst.TokenFilteredgeGram)
		if string(jsonTokenFilteredgeGram) == "{}" { // empty struct
			dst.TokenFilteredgeGram = nil
		} else {
			return nil // data stored in dst.TokenFilteredgeGram, return on the first match
		}
	} else {
		dst.TokenFilteredgeGram = nil
	}

	// try to unmarshal JSON data into TokenFiltericuFolding
	err = json.Unmarshal(data, &dst.TokenFiltericuFolding)
	if err == nil {
		jsonTokenFiltericuFolding, _ := json.Marshal(dst.TokenFiltericuFolding)
		if string(jsonTokenFiltericuFolding) == "{}" { // empty struct
			dst.TokenFiltericuFolding = nil
		} else {
			return nil // data stored in dst.TokenFiltericuFolding, return on the first match
		}
	} else {
		dst.TokenFiltericuFolding = nil
	}

	// try to unmarshal JSON data into TokenFiltericuNormalizer
	err = json.Unmarshal(data, &dst.TokenFiltericuNormalizer)
	if err == nil {
		jsonTokenFiltericuNormalizer, _ := json.Marshal(dst.TokenFiltericuNormalizer)
		if string(jsonTokenFiltericuNormalizer) == "{}" { // empty struct
			dst.TokenFiltericuNormalizer = nil
		} else {
			return nil // data stored in dst.TokenFiltericuNormalizer, return on the first match
		}
	} else {
		dst.TokenFiltericuNormalizer = nil
	}

	// try to unmarshal JSON data into TokenFilterlength
	err = json.Unmarshal(data, &dst.TokenFilterlength)
	if err == nil {
		jsonTokenFilterlength, _ := json.Marshal(dst.TokenFilterlength)
		if string(jsonTokenFilterlength) == "{}" { // empty struct
			dst.TokenFilterlength = nil
		} else {
			return nil // data stored in dst.TokenFilterlength, return on the first match
		}
	} else {
		dst.TokenFilterlength = nil
	}

	// try to unmarshal JSON data into TokenFilterlowercase
	err = json.Unmarshal(data, &dst.TokenFilterlowercase)
	if err == nil {
		jsonTokenFilterlowercase, _ := json.Marshal(dst.TokenFilterlowercase)
		if string(jsonTokenFilterlowercase) == "{}" { // empty struct
			dst.TokenFilterlowercase = nil
		} else {
			return nil // data stored in dst.TokenFilterlowercase, return on the first match
		}
	} else {
		dst.TokenFilterlowercase = nil
	}

	// try to unmarshal JSON data into TokenFilternGram
	err = json.Unmarshal(data, &dst.TokenFilternGram)
	if err == nil {
		jsonTokenFilternGram, _ := json.Marshal(dst.TokenFilternGram)
		if string(jsonTokenFilternGram) == "{}" { // empty struct
			dst.TokenFilternGram = nil
		} else {
			return nil // data stored in dst.TokenFilternGram, return on the first match
		}
	} else {
		dst.TokenFilternGram = nil
	}

	// try to unmarshal JSON data into TokenFilterregex
	err = json.Unmarshal(data, &dst.TokenFilterregex)
	if err == nil {
		jsonTokenFilterregex, _ := json.Marshal(dst.TokenFilterregex)
		if string(jsonTokenFilterregex) == "{}" { // empty struct
			dst.TokenFilterregex = nil
		} else {
			return nil // data stored in dst.TokenFilterregex, return on the first match
		}
	} else {
		dst.TokenFilterregex = nil
	}

	// try to unmarshal JSON data into TokenFilterreverse
	err = json.Unmarshal(data, &dst.TokenFilterreverse)
	if err == nil {
		jsonTokenFilterreverse, _ := json.Marshal(dst.TokenFilterreverse)
		if string(jsonTokenFilterreverse) == "{}" { // empty struct
			dst.TokenFilterreverse = nil
		} else {
			return nil // data stored in dst.TokenFilterreverse, return on the first match
		}
	} else {
		dst.TokenFilterreverse = nil
	}

	// try to unmarshal JSON data into TokenFiltershingle
	err = json.Unmarshal(data, &dst.TokenFiltershingle)
	if err == nil {
		jsonTokenFiltershingle, _ := json.Marshal(dst.TokenFiltershingle)
		if string(jsonTokenFiltershingle) == "{}" { // empty struct
			dst.TokenFiltershingle = nil
		} else {
			return nil // data stored in dst.TokenFiltershingle, return on the first match
		}
	} else {
		dst.TokenFiltershingle = nil
	}

	// try to unmarshal JSON data into TokenFiltersnowballStemming
	err = json.Unmarshal(data, &dst.TokenFiltersnowballStemming)
	if err == nil {
		jsonTokenFiltersnowballStemming, _ := json.Marshal(dst.TokenFiltersnowballStemming)
		if string(jsonTokenFiltersnowballStemming) == "{}" { // empty struct
			dst.TokenFiltersnowballStemming = nil
		} else {
			return nil // data stored in dst.TokenFiltersnowballStemming, return on the first match
		}
	} else {
		dst.TokenFiltersnowballStemming = nil
	}

	// try to unmarshal JSON data into TokenFilterstopword
	err = json.Unmarshal(data, &dst.TokenFilterstopword)
	if err == nil {
		jsonTokenFilterstopword, _ := json.Marshal(dst.TokenFilterstopword)
		if string(jsonTokenFilterstopword) == "{}" { // empty struct
			dst.TokenFilterstopword = nil
		} else {
			return nil // data stored in dst.TokenFilterstopword, return on the first match
		}
	} else {
		dst.TokenFilterstopword = nil
	}

	// try to unmarshal JSON data into TokenFiltertrim
	err = json.Unmarshal(data, &dst.TokenFiltertrim)
	if err == nil {
		jsonTokenFiltertrim, _ := json.Marshal(dst.TokenFiltertrim)
		if string(jsonTokenFiltertrim) == "{}" { // empty struct
			dst.TokenFiltertrim = nil
		} else {
			return nil // data stored in dst.TokenFiltertrim, return on the first match
		}
	} else {
		dst.TokenFiltertrim = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(FTSAnalyzersTokenFiltersInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *FTSAnalyzersTokenFiltersInner) MarshalJSON() ([]byte, error) {
	if src.TokenFilterasciiFolding != nil {
		return json.Marshal(&src.TokenFilterasciiFolding)
	}

	if src.TokenFilterdaitchMokotoffSoundex != nil {
		return json.Marshal(&src.TokenFilterdaitchMokotoffSoundex)
	}

	if src.TokenFilteredgeGram != nil {
		return json.Marshal(&src.TokenFilteredgeGram)
	}

	if src.TokenFiltericuFolding != nil {
		return json.Marshal(&src.TokenFiltericuFolding)
	}

	if src.TokenFiltericuNormalizer != nil {
		return json.Marshal(&src.TokenFiltericuNormalizer)
	}

	if src.TokenFilterlength != nil {
		return json.Marshal(&src.TokenFilterlength)
	}

	if src.TokenFilterlowercase != nil {
		return json.Marshal(&src.TokenFilterlowercase)
	}

	if src.TokenFilternGram != nil {
		return json.Marshal(&src.TokenFilternGram)
	}

	if src.TokenFilterregex != nil {
		return json.Marshal(&src.TokenFilterregex)
	}

	if src.TokenFilterreverse != nil {
		return json.Marshal(&src.TokenFilterreverse)
	}

	if src.TokenFiltershingle != nil {
		return json.Marshal(&src.TokenFiltershingle)
	}

	if src.TokenFiltersnowballStemming != nil {
		return json.Marshal(&src.TokenFiltersnowballStemming)
	}

	if src.TokenFilterstopword != nil {
		return json.Marshal(&src.TokenFilterstopword)
	}

	if src.TokenFiltertrim != nil {
		return json.Marshal(&src.TokenFiltertrim)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableFTSAnalyzersTokenFiltersInner struct {
	value *FTSAnalyzersTokenFiltersInner
	isSet bool
}

func (v NullableFTSAnalyzersTokenFiltersInner) Get() *FTSAnalyzersTokenFiltersInner {
	return v.value
}

func (v *NullableFTSAnalyzersTokenFiltersInner) Set(val *FTSAnalyzersTokenFiltersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFTSAnalyzersTokenFiltersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFTSAnalyzersTokenFiltersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFTSAnalyzersTokenFiltersInner(val *FTSAnalyzersTokenFiltersInner) *NullableFTSAnalyzersTokenFiltersInner {
	return &NullableFTSAnalyzersTokenFiltersInner{value: val, isSet: true}
}

func (v NullableFTSAnalyzersTokenFiltersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFTSAnalyzersTokenFiltersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
