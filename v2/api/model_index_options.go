/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// IndexOptions One or more settings that determine how the MongoDB Cloud creates this MongoDB index.
type IndexOptions struct {
	// Index version number applied to the 2dsphere index. MongoDB 3.2 and later use version 3. Use this option to override the default version number. This option applies to the **2dsphere** index type only.
	Var2dsphereIndexVersion int32 `json:"2dsphereIndexVersion"`
	// Flag that indicates whether MongoDB should build the index in the background. This applies to MongoDB databases running feature compatibility version 4.0 or earlier. MongoDB databases running FCV 4.2 or later build indexes using an optimized build process. This process holds the exclusive lock only at the beginning and end of the build process. The rest of the build process yields to interleaving read and write operations. MongoDB databases running FCV 4.2 or later ignore this option. This option applies to all index types.
	Background *bool `json:"background,omitempty"`
	// Number of precision applied to the stored geohash value of the location data. This option applies to the **2d** index type only.
	Bits *int32 `json:"bits,omitempty"`
	// Number of units within which to group the location values. You could group in the same bucket those location values within the specified number of units to each other. This option applies to the geoHaystack index type only.  MongoDB 5.0 removed geoHaystack Indexes and the `geoSearch` command.
	BucketSize *int32 `json:"bucketSize,omitempty"`
	// Human language that determines the list of stop words and the rules for the stemmer and tokenizer. This option accepts the supported languages using its name in lowercase english or the ISO 639-2 code. If you set this parameter to `\"none\"`, then the text search uses simple tokenization with no list of stop words and no stemming. This option applies to the **text** index type only.
	DefaultLanguage *string `json:"default_language,omitempty"`
	// Number of seconds that MongoDB retains documents in a Time To Live (TTL) index.
	ExpireAfterSeconds *int32 `json:"expireAfterSeconds,omitempty"`
	// Flag that determines whether the index is hidden from the query planner. A hidden index is not evaluated as part of the query plan selection.
	Hidden *bool `json:"hidden,omitempty"`
	// Human-readable label that identifies the document parameter that contains the override language for the document. This option applies to the **text** index type only.
	LanguageOverride *string `json:"language_override,omitempty"`
	// Upper inclusive boundary to limit the longitude and latitude values. This option applies to the 2d index type only.
	Max *int32 `json:"max,omitempty"`
	// Lower inclusive boundary to limit the longitude and latitude values. This option applies to the 2d index type only.
	Min *int32 `json:"min,omitempty"`
	// Human-readable label that identifies this index. This option applies to all index types.
	Name *string `json:"name,omitempty"`
	// Rules that limit the documents that the index references to a filter expression. All MongoDB index types accept a **partialFilterExpression** option. **partialFilterExpression** can include following expressions:  - equality (`\"parameter\" : \"value\"` or using the `$eq` operator) - `\"$exists\": true` , maximum: `$gt`, `$gte`, `$lt`, `$lte` comparisons - `$type` - `$and` (top-level only)  This option applies to all index types.
	PartialFilterExpression map[string]interface{} `json:"partialFilterExpression,omitempty"`
	// Flag that indicates whether the index references documents that only have the specified parameter. These indexes use less space but behave differently in some situations like when sorting. The following index types default to sparse and ignore this option: `2dsphere`, `2d`, `geoHaystack`, `text`.  Compound indexes that includes one or more indexes with `2dsphere` keys alongside other key types, only the `2dsphere` index parameters determine which documents the index references. If you run MongoDB 3.2 or later, use partial indexes. This option applies to all index types.
	Sparse *bool `json:"sparse,omitempty"`
	// Storage engine set for the specific index. This value can be set only at creation. This option uses the following format: `\"storageEngine\" : { \"<storage-engine-name>\" : \"<options>\" }` MongoDB validates storage engine configuration options when creating indexes. To support replica sets with members with different storage engines, MongoDB logs these options to the oplog during replication. This option applies to all index types.
	StorageEngine map[string]interface{} `json:"storageEngine,omitempty"`
	// Version applied to this text index. MongoDB 3.2 and later use version `3`. Use this option to override the default version number. This option applies to the **text** index type only.
	TextIndexVersion *int32 `json:"textIndexVersion,omitempty"`
	// Flag that indicates whether this index can accept insertion or update of documents when the index key value matches an existing index key value. Set `\"unique\" : true` to set this index as unique. You can't set a hashed index to be unique. This option applies to all index types.
	Unique *bool `json:"unique,omitempty"`
	// Relative importance to place upon provided index parameters. This object expresses this as key/value pairs of index parameter and weight to apply to that parameter. You can specify weights for some or all the indexed parameters. The weight must be an integer between 1 and 99,999. MongoDB 5.0 and later can apply **weights** to **text** indexes only.
	Weights map[string]interface{} `json:"weights,omitempty"`
}

// NewIndexOptions instantiates a new IndexOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndexOptions() *IndexOptions {
	this := IndexOptions{}
	var var2dsphereIndexVersion int32 = 3
	this.Var2dsphereIndexVersion = var2dsphereIndexVersion
	var background bool = false
	this.Background = &background
	var bits int32 = 26
	this.Bits = &bits
	var defaultLanguage string = "english"
	this.DefaultLanguage = &defaultLanguage
	var hidden bool = false
	this.Hidden = &hidden
	var languageOverride string = "language"
	this.LanguageOverride = &languageOverride
	var max int32 = 180
	this.Max = &max
	var min int32 = -180
	this.Min = &min
	var sparse bool = false
	this.Sparse = &sparse
	var textIndexVersion int32 = 3
	this.TextIndexVersion = &textIndexVersion
	var unique bool = false
	this.Unique = &unique
	return &this
}

// NewIndexOptionsWithDefaults instantiates a new IndexOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndexOptionsWithDefaults() *IndexOptions {
	this := IndexOptions{}
	var var2dsphereIndexVersion int32 = 3
	this.Var2dsphereIndexVersion = var2dsphereIndexVersion
	var background bool = false
	this.Background = &background
	var bits int32 = 26
	this.Bits = &bits
	var defaultLanguage string = "english"
	this.DefaultLanguage = &defaultLanguage
	var hidden bool = false
	this.Hidden = &hidden
	var languageOverride string = "language"
	this.LanguageOverride = &languageOverride
	var max int32 = 180
	this.Max = &max
	var min int32 = -180
	this.Min = &min
	var sparse bool = false
	this.Sparse = &sparse
	var textIndexVersion int32 = 3
	this.TextIndexVersion = &textIndexVersion
	var unique bool = false
	this.Unique = &unique
	return &this
}

// GetVar2dsphereIndexVersion returns the Var2dsphereIndexVersion field value
func (o *IndexOptions) GetVar2dsphereIndexVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Var2dsphereIndexVersion
}

// GetVar2dsphereIndexVersionOk returns a tuple with the Var2dsphereIndexVersion field value
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetVar2dsphereIndexVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Var2dsphereIndexVersion, true
}

// SetVar2dsphereIndexVersion sets field value
func (o *IndexOptions) SetVar2dsphereIndexVersion(v int32) {
	o.Var2dsphereIndexVersion = v
}

// GetBackground returns the Background field value if set, zero value otherwise.
func (o *IndexOptions) GetBackground() bool {
	if o == nil || o.Background == nil {
		var ret bool
		return ret
	}
	return *o.Background
}

// GetBackgroundOk returns a tuple with the Background field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetBackgroundOk() (*bool, bool) {
	if o == nil || o.Background == nil {
		return nil, false
	}
	return o.Background, true
}

// HasBackground returns a boolean if a field has been set.
func (o *IndexOptions) HasBackground() bool {
	if o != nil && o.Background != nil {
		return true
	}

	return false
}

// SetBackground gets a reference to the given bool and assigns it to the Background field.
func (o *IndexOptions) SetBackground(v bool) {
	o.Background = &v
}

// GetBits returns the Bits field value if set, zero value otherwise.
func (o *IndexOptions) GetBits() int32 {
	if o == nil || o.Bits == nil {
		var ret int32
		return ret
	}
	return *o.Bits
}

// GetBitsOk returns a tuple with the Bits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetBitsOk() (*int32, bool) {
	if o == nil || o.Bits == nil {
		return nil, false
	}
	return o.Bits, true
}

// HasBits returns a boolean if a field has been set.
func (o *IndexOptions) HasBits() bool {
	if o != nil && o.Bits != nil {
		return true
	}

	return false
}

// SetBits gets a reference to the given int32 and assigns it to the Bits field.
func (o *IndexOptions) SetBits(v int32) {
	o.Bits = &v
}

// GetBucketSize returns the BucketSize field value if set, zero value otherwise.
func (o *IndexOptions) GetBucketSize() int32 {
	if o == nil || o.BucketSize == nil {
		var ret int32
		return ret
	}
	return *o.BucketSize
}

// GetBucketSizeOk returns a tuple with the BucketSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetBucketSizeOk() (*int32, bool) {
	if o == nil || o.BucketSize == nil {
		return nil, false
	}
	return o.BucketSize, true
}

// HasBucketSize returns a boolean if a field has been set.
func (o *IndexOptions) HasBucketSize() bool {
	if o != nil && o.BucketSize != nil {
		return true
	}

	return false
}

// SetBucketSize gets a reference to the given int32 and assigns it to the BucketSize field.
func (o *IndexOptions) SetBucketSize(v int32) {
	o.BucketSize = &v
}

// GetDefaultLanguage returns the DefaultLanguage field value if set, zero value otherwise.
func (o *IndexOptions) GetDefaultLanguage() string {
	if o == nil || o.DefaultLanguage == nil {
		var ret string
		return ret
	}
	return *o.DefaultLanguage
}

// GetDefaultLanguageOk returns a tuple with the DefaultLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetDefaultLanguageOk() (*string, bool) {
	if o == nil || o.DefaultLanguage == nil {
		return nil, false
	}
	return o.DefaultLanguage, true
}

// HasDefaultLanguage returns a boolean if a field has been set.
func (o *IndexOptions) HasDefaultLanguage() bool {
	if o != nil && o.DefaultLanguage != nil {
		return true
	}

	return false
}

// SetDefaultLanguage gets a reference to the given string and assigns it to the DefaultLanguage field.
func (o *IndexOptions) SetDefaultLanguage(v string) {
	o.DefaultLanguage = &v
}

// GetExpireAfterSeconds returns the ExpireAfterSeconds field value if set, zero value otherwise.
func (o *IndexOptions) GetExpireAfterSeconds() int32 {
	if o == nil || o.ExpireAfterSeconds == nil {
		var ret int32
		return ret
	}
	return *o.ExpireAfterSeconds
}

// GetExpireAfterSecondsOk returns a tuple with the ExpireAfterSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetExpireAfterSecondsOk() (*int32, bool) {
	if o == nil || o.ExpireAfterSeconds == nil {
		return nil, false
	}
	return o.ExpireAfterSeconds, true
}

// HasExpireAfterSeconds returns a boolean if a field has been set.
func (o *IndexOptions) HasExpireAfterSeconds() bool {
	if o != nil && o.ExpireAfterSeconds != nil {
		return true
	}

	return false
}

// SetExpireAfterSeconds gets a reference to the given int32 and assigns it to the ExpireAfterSeconds field.
func (o *IndexOptions) SetExpireAfterSeconds(v int32) {
	o.ExpireAfterSeconds = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *IndexOptions) GetHidden() bool {
	if o == nil || o.Hidden == nil {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetHiddenOk() (*bool, bool) {
	if o == nil || o.Hidden == nil {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *IndexOptions) HasHidden() bool {
	if o != nil && o.Hidden != nil {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *IndexOptions) SetHidden(v bool) {
	o.Hidden = &v
}

// GetLanguageOverride returns the LanguageOverride field value if set, zero value otherwise.
func (o *IndexOptions) GetLanguageOverride() string {
	if o == nil || o.LanguageOverride == nil {
		var ret string
		return ret
	}
	return *o.LanguageOverride
}

// GetLanguageOverrideOk returns a tuple with the LanguageOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetLanguageOverrideOk() (*string, bool) {
	if o == nil || o.LanguageOverride == nil {
		return nil, false
	}
	return o.LanguageOverride, true
}

// HasLanguageOverride returns a boolean if a field has been set.
func (o *IndexOptions) HasLanguageOverride() bool {
	if o != nil && o.LanguageOverride != nil {
		return true
	}

	return false
}

// SetLanguageOverride gets a reference to the given string and assigns it to the LanguageOverride field.
func (o *IndexOptions) SetLanguageOverride(v string) {
	o.LanguageOverride = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *IndexOptions) GetMax() int32 {
	if o == nil || o.Max == nil {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetMaxOk() (*int32, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *IndexOptions) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *IndexOptions) SetMax(v int32) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *IndexOptions) GetMin() int32 {
	if o == nil || o.Min == nil {
		var ret int32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetMinOk() (*int32, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *IndexOptions) HasMin() bool {
	if o != nil && o.Min != nil {
		return true
	}

	return false
}

// SetMin gets a reference to the given int32 and assigns it to the Min field.
func (o *IndexOptions) SetMin(v int32) {
	o.Min = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IndexOptions) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IndexOptions) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IndexOptions) SetName(v string) {
	o.Name = &v
}

// GetPartialFilterExpression returns the PartialFilterExpression field value if set, zero value otherwise.
func (o *IndexOptions) GetPartialFilterExpression() map[string]interface{} {
	if o == nil || o.PartialFilterExpression == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.PartialFilterExpression
}

// GetPartialFilterExpressionOk returns a tuple with the PartialFilterExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetPartialFilterExpressionOk() (map[string]interface{}, bool) {
	if o == nil || o.PartialFilterExpression == nil {
		return nil, false
	}
	return o.PartialFilterExpression, true
}

// HasPartialFilterExpression returns a boolean if a field has been set.
func (o *IndexOptions) HasPartialFilterExpression() bool {
	if o != nil && o.PartialFilterExpression != nil {
		return true
	}

	return false
}

// SetPartialFilterExpression gets a reference to the given map[string]interface{} and assigns it to the PartialFilterExpression field.
func (o *IndexOptions) SetPartialFilterExpression(v map[string]interface{}) {
	o.PartialFilterExpression = v
}

// GetSparse returns the Sparse field value if set, zero value otherwise.
func (o *IndexOptions) GetSparse() bool {
	if o == nil || o.Sparse == nil {
		var ret bool
		return ret
	}
	return *o.Sparse
}

// GetSparseOk returns a tuple with the Sparse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetSparseOk() (*bool, bool) {
	if o == nil || o.Sparse == nil {
		return nil, false
	}
	return o.Sparse, true
}

// HasSparse returns a boolean if a field has been set.
func (o *IndexOptions) HasSparse() bool {
	if o != nil && o.Sparse != nil {
		return true
	}

	return false
}

// SetSparse gets a reference to the given bool and assigns it to the Sparse field.
func (o *IndexOptions) SetSparse(v bool) {
	o.Sparse = &v
}

// GetStorageEngine returns the StorageEngine field value if set, zero value otherwise.
func (o *IndexOptions) GetStorageEngine() map[string]interface{} {
	if o == nil || o.StorageEngine == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.StorageEngine
}

// GetStorageEngineOk returns a tuple with the StorageEngine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetStorageEngineOk() (map[string]interface{}, bool) {
	if o == nil || o.StorageEngine == nil {
		return nil, false
	}
	return o.StorageEngine, true
}

// HasStorageEngine returns a boolean if a field has been set.
func (o *IndexOptions) HasStorageEngine() bool {
	if o != nil && o.StorageEngine != nil {
		return true
	}

	return false
}

// SetStorageEngine gets a reference to the given map[string]interface{} and assigns it to the StorageEngine field.
func (o *IndexOptions) SetStorageEngine(v map[string]interface{}) {
	o.StorageEngine = v
}

// GetTextIndexVersion returns the TextIndexVersion field value if set, zero value otherwise.
func (o *IndexOptions) GetTextIndexVersion() int32 {
	if o == nil || o.TextIndexVersion == nil {
		var ret int32
		return ret
	}
	return *o.TextIndexVersion
}

// GetTextIndexVersionOk returns a tuple with the TextIndexVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetTextIndexVersionOk() (*int32, bool) {
	if o == nil || o.TextIndexVersion == nil {
		return nil, false
	}
	return o.TextIndexVersion, true
}

// HasTextIndexVersion returns a boolean if a field has been set.
func (o *IndexOptions) HasTextIndexVersion() bool {
	if o != nil && o.TextIndexVersion != nil {
		return true
	}

	return false
}

// SetTextIndexVersion gets a reference to the given int32 and assigns it to the TextIndexVersion field.
func (o *IndexOptions) SetTextIndexVersion(v int32) {
	o.TextIndexVersion = &v
}

// GetUnique returns the Unique field value if set, zero value otherwise.
func (o *IndexOptions) GetUnique() bool {
	if o == nil || o.Unique == nil {
		var ret bool
		return ret
	}
	return *o.Unique
}

// GetUniqueOk returns a tuple with the Unique field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetUniqueOk() (*bool, bool) {
	if o == nil || o.Unique == nil {
		return nil, false
	}
	return o.Unique, true
}

// HasUnique returns a boolean if a field has been set.
func (o *IndexOptions) HasUnique() bool {
	if o != nil && o.Unique != nil {
		return true
	}

	return false
}

// SetUnique gets a reference to the given bool and assigns it to the Unique field.
func (o *IndexOptions) SetUnique(v bool) {
	o.Unique = &v
}

// GetWeights returns the Weights field value if set, zero value otherwise.
func (o *IndexOptions) GetWeights() map[string]interface{} {
	if o == nil || o.Weights == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Weights
}

// GetWeightsOk returns a tuple with the Weights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndexOptions) GetWeightsOk() (map[string]interface{}, bool) {
	if o == nil || o.Weights == nil {
		return nil, false
	}
	return o.Weights, true
}

// HasWeights returns a boolean if a field has been set.
func (o *IndexOptions) HasWeights() bool {
	if o != nil && o.Weights != nil {
		return true
	}

	return false
}

// SetWeights gets a reference to the given map[string]interface{} and assigns it to the Weights field.
func (o *IndexOptions) SetWeights(v map[string]interface{}) {
	o.Weights = v
}

func (o IndexOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["2dsphereIndexVersion"] = o.Var2dsphereIndexVersion
	}
	if o.Background != nil {
		toSerialize["background"] = o.Background
	}
	if o.Bits != nil {
		toSerialize["bits"] = o.Bits
	}
	if o.BucketSize != nil {
		toSerialize["bucketSize"] = o.BucketSize
	}
	if o.DefaultLanguage != nil {
		toSerialize["default_language"] = o.DefaultLanguage
	}
	if o.ExpireAfterSeconds != nil {
		toSerialize["expireAfterSeconds"] = o.ExpireAfterSeconds
	}
	if o.Hidden != nil {
		toSerialize["hidden"] = o.Hidden
	}
	if o.LanguageOverride != nil {
		toSerialize["language_override"] = o.LanguageOverride
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PartialFilterExpression != nil {
		toSerialize["partialFilterExpression"] = o.PartialFilterExpression
	}
	if o.Sparse != nil {
		toSerialize["sparse"] = o.Sparse
	}
	if o.StorageEngine != nil {
		toSerialize["storageEngine"] = o.StorageEngine
	}
	if o.TextIndexVersion != nil {
		toSerialize["textIndexVersion"] = o.TextIndexVersion
	}
	if o.Unique != nil {
		toSerialize["unique"] = o.Unique
	}
	if o.Weights != nil {
		toSerialize["weights"] = o.Weights
	}
	return json.Marshal(toSerialize)
}

type NullableIndexOptions struct {
	value *IndexOptions
	isSet bool
}

func (v NullableIndexOptions) Get() *IndexOptions {
	return v.value
}

func (v *NullableIndexOptions) Set(val *IndexOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableIndexOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableIndexOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndexOptions(val *IndexOptions) *NullableIndexOptions {
	return &NullableIndexOptions{value: val, isSet: true}
}

func (v NullableIndexOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndexOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


