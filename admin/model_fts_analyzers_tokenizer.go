// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
	"fmt"
)

// FTSAnalyzersTokenizer - Tokenizer that you want to use to create tokens. Tokens determine how Atlas Search splits up text into discrete chunks for indexing.
type FTSAnalyzersTokenizer struct {
	TokenizeredgeGram          *TokenizeredgeGram
	Tokenizerkeyword           *Tokenizerkeyword
	TokenizernGram             *TokenizernGram
	TokenizerregexCaptureGroup *TokenizerregexCaptureGroup
	TokenizerregexSplit        *TokenizerregexSplit
	Tokenizerstandard          *Tokenizerstandard
	TokenizeruaxUrlEmail       *TokenizeruaxUrlEmail
	Tokenizerwhitespace        *Tokenizerwhitespace
}

// TokenizeredgeGramAsFTSAnalyzersTokenizer is a convenience function that returns TokenizeredgeGram wrapped in FTSAnalyzersTokenizer
func TokenizeredgeGramAsFTSAnalyzersTokenizer(v *TokenizeredgeGram) FTSAnalyzersTokenizer {
	return FTSAnalyzersTokenizer{
		TokenizeredgeGram: v,
	}
}

// TokenizerkeywordAsFTSAnalyzersTokenizer is a convenience function that returns Tokenizerkeyword wrapped in FTSAnalyzersTokenizer
func TokenizerkeywordAsFTSAnalyzersTokenizer(v *Tokenizerkeyword) FTSAnalyzersTokenizer {
	return FTSAnalyzersTokenizer{
		Tokenizerkeyword: v,
	}
}

// TokenizernGramAsFTSAnalyzersTokenizer is a convenience function that returns TokenizernGram wrapped in FTSAnalyzersTokenizer
func TokenizernGramAsFTSAnalyzersTokenizer(v *TokenizernGram) FTSAnalyzersTokenizer {
	return FTSAnalyzersTokenizer{
		TokenizernGram: v,
	}
}

// TokenizerregexCaptureGroupAsFTSAnalyzersTokenizer is a convenience function that returns TokenizerregexCaptureGroup wrapped in FTSAnalyzersTokenizer
func TokenizerregexCaptureGroupAsFTSAnalyzersTokenizer(v *TokenizerregexCaptureGroup) FTSAnalyzersTokenizer {
	return FTSAnalyzersTokenizer{
		TokenizerregexCaptureGroup: v,
	}
}

// TokenizerregexSplitAsFTSAnalyzersTokenizer is a convenience function that returns TokenizerregexSplit wrapped in FTSAnalyzersTokenizer
func TokenizerregexSplitAsFTSAnalyzersTokenizer(v *TokenizerregexSplit) FTSAnalyzersTokenizer {
	return FTSAnalyzersTokenizer{
		TokenizerregexSplit: v,
	}
}

// TokenizerstandardAsFTSAnalyzersTokenizer is a convenience function that returns Tokenizerstandard wrapped in FTSAnalyzersTokenizer
func TokenizerstandardAsFTSAnalyzersTokenizer(v *Tokenizerstandard) FTSAnalyzersTokenizer {
	return FTSAnalyzersTokenizer{
		Tokenizerstandard: v,
	}
}

// TokenizeruaxUrlEmailAsFTSAnalyzersTokenizer is a convenience function that returns TokenizeruaxUrlEmail wrapped in FTSAnalyzersTokenizer
func TokenizeruaxUrlEmailAsFTSAnalyzersTokenizer(v *TokenizeruaxUrlEmail) FTSAnalyzersTokenizer {
	return FTSAnalyzersTokenizer{
		TokenizeruaxUrlEmail: v,
	}
}

// TokenizerwhitespaceAsFTSAnalyzersTokenizer is a convenience function that returns Tokenizerwhitespace wrapped in FTSAnalyzersTokenizer
func TokenizerwhitespaceAsFTSAnalyzersTokenizer(v *Tokenizerwhitespace) FTSAnalyzersTokenizer {
	return FTSAnalyzersTokenizer{
		Tokenizerwhitespace: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *FTSAnalyzersTokenizer) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TokenizeredgeGram
	err = json.Unmarshal(data, &dst.TokenizeredgeGram)
	if err == nil {
		jsonTokenizeredgeGram, _ := json.Marshal(dst.TokenizeredgeGram)
		if string(jsonTokenizeredgeGram) == "{}" { // empty struct
			dst.TokenizeredgeGram = nil
		} else {
			match++
		}
	} else {
		dst.TokenizeredgeGram = nil
	}

	// try to unmarshal data into Tokenizerkeyword
	err = json.Unmarshal(data, &dst.Tokenizerkeyword)
	if err == nil {
		jsonTokenizerkeyword, _ := json.Marshal(dst.Tokenizerkeyword)
		if string(jsonTokenizerkeyword) == "{}" { // empty struct
			dst.Tokenizerkeyword = nil
		} else {
			match++
		}
	} else {
		dst.Tokenizerkeyword = nil
	}

	// try to unmarshal data into TokenizernGram
	err = json.Unmarshal(data, &dst.TokenizernGram)
	if err == nil {
		jsonTokenizernGram, _ := json.Marshal(dst.TokenizernGram)
		if string(jsonTokenizernGram) == "{}" { // empty struct
			dst.TokenizernGram = nil
		} else {
			match++
		}
	} else {
		dst.TokenizernGram = nil
	}

	// try to unmarshal data into TokenizerregexCaptureGroup
	err = json.Unmarshal(data, &dst.TokenizerregexCaptureGroup)
	if err == nil {
		jsonTokenizerregexCaptureGroup, _ := json.Marshal(dst.TokenizerregexCaptureGroup)
		if string(jsonTokenizerregexCaptureGroup) == "{}" { // empty struct
			dst.TokenizerregexCaptureGroup = nil
		} else {
			match++
		}
	} else {
		dst.TokenizerregexCaptureGroup = nil
	}

	// try to unmarshal data into TokenizerregexSplit
	err = json.Unmarshal(data, &dst.TokenizerregexSplit)
	if err == nil {
		jsonTokenizerregexSplit, _ := json.Marshal(dst.TokenizerregexSplit)
		if string(jsonTokenizerregexSplit) == "{}" { // empty struct
			dst.TokenizerregexSplit = nil
		} else {
			match++
		}
	} else {
		dst.TokenizerregexSplit = nil
	}

	// try to unmarshal data into Tokenizerstandard
	err = json.Unmarshal(data, &dst.Tokenizerstandard)
	if err == nil {
		jsonTokenizerstandard, _ := json.Marshal(dst.Tokenizerstandard)
		if string(jsonTokenizerstandard) == "{}" { // empty struct
			dst.Tokenizerstandard = nil
		} else {
			match++
		}
	} else {
		dst.Tokenizerstandard = nil
	}

	// try to unmarshal data into TokenizeruaxUrlEmail
	err = json.Unmarshal(data, &dst.TokenizeruaxUrlEmail)
	if err == nil {
		jsonTokenizeruaxUrlEmail, _ := json.Marshal(dst.TokenizeruaxUrlEmail)
		if string(jsonTokenizeruaxUrlEmail) == "{}" { // empty struct
			dst.TokenizeruaxUrlEmail = nil
		} else {
			match++
		}
	} else {
		dst.TokenizeruaxUrlEmail = nil
	}

	// try to unmarshal data into Tokenizerwhitespace
	err = json.Unmarshal(data, &dst.Tokenizerwhitespace)
	if err == nil {
		jsonTokenizerwhitespace, _ := json.Marshal(dst.Tokenizerwhitespace)
		if string(jsonTokenizerwhitespace) == "{}" { // empty struct
			dst.Tokenizerwhitespace = nil
		} else {
			match++
		}
	} else {
		dst.Tokenizerwhitespace = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.TokenizeredgeGram = nil
		dst.Tokenizerkeyword = nil
		dst.TokenizernGram = nil
		dst.TokenizerregexCaptureGroup = nil
		dst.TokenizerregexSplit = nil
		dst.Tokenizerstandard = nil
		dst.TokenizeruaxUrlEmail = nil
		dst.Tokenizerwhitespace = nil

		return fmt.Errorf("data matches more than one schema in oneOf(FTSAnalyzersTokenizer)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(FTSAnalyzersTokenizer)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FTSAnalyzersTokenizer) MarshalJSON() ([]byte, error) {
	if src.TokenizeredgeGram != nil {
		return json.Marshal(&src.TokenizeredgeGram)
	}

	if src.Tokenizerkeyword != nil {
		return json.Marshal(&src.Tokenizerkeyword)
	}

	if src.TokenizernGram != nil {
		return json.Marshal(&src.TokenizernGram)
	}

	if src.TokenizerregexCaptureGroup != nil {
		return json.Marshal(&src.TokenizerregexCaptureGroup)
	}

	if src.TokenizerregexSplit != nil {
		return json.Marshal(&src.TokenizerregexSplit)
	}

	if src.Tokenizerstandard != nil {
		return json.Marshal(&src.Tokenizerstandard)
	}

	if src.TokenizeruaxUrlEmail != nil {
		return json.Marshal(&src.TokenizeruaxUrlEmail)
	}

	if src.Tokenizerwhitespace != nil {
		return json.Marshal(&src.Tokenizerwhitespace)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *FTSAnalyzersTokenizer) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.TokenizeredgeGram != nil {
		return obj.TokenizeredgeGram
	}

	if obj.Tokenizerkeyword != nil {
		return obj.Tokenizerkeyword
	}

	if obj.TokenizernGram != nil {
		return obj.TokenizernGram
	}

	if obj.TokenizerregexCaptureGroup != nil {
		return obj.TokenizerregexCaptureGroup
	}

	if obj.TokenizerregexSplit != nil {
		return obj.TokenizerregexSplit
	}

	if obj.Tokenizerstandard != nil {
		return obj.Tokenizerstandard
	}

	if obj.TokenizeruaxUrlEmail != nil {
		return obj.TokenizeruaxUrlEmail
	}

	if obj.Tokenizerwhitespace != nil {
		return obj.Tokenizerwhitespace
	}

	// all schemas are nil
	return nil
}

type NullableFTSAnalyzersTokenizer struct {
	value *FTSAnalyzersTokenizer
	isSet bool
}

func (v NullableFTSAnalyzersTokenizer) Get() *FTSAnalyzersTokenizer {
	return v.value
}

func (v *NullableFTSAnalyzersTokenizer) Set(val *FTSAnalyzersTokenizer) {
	v.value = val
	v.isSet = true
}

func (v NullableFTSAnalyzersTokenizer) IsSet() bool {
	return v.isSet
}

func (v *NullableFTSAnalyzersTokenizer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFTSAnalyzersTokenizer(val *FTSAnalyzersTokenizer) *NullableFTSAnalyzersTokenizer {
	return &NullableFTSAnalyzersTokenizer{value: val, isSet: true}
}

func (v NullableFTSAnalyzersTokenizer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFTSAnalyzersTokenizer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
