package usda

type FoundationFood struct {
	FoodClass                 string          `bson:"foodClass,omitempty" json:"foodClass,omitempty"`
	Description               string          `bson:"description,omitempty" json:"description,omitempty"`
	ScientificName            string          `bson:"scientificName,omitempty" json:"scientificName,omitempty"`
	FoodNutrients             []FoodNutrients `bson:"foodNutrients,omitempty" json:"foodNutrients,omitempty"`
	FoodAttributes            []interface{}   `bson:"foodAttributes,omitempty" json:"foodAttributes,omitempty"`
	NutrientConversionFactors []interface{}   `bson:"nutrientConversionFactors,omitempty" json:"nutrientConversionFactors,omitempty"`
	IsHistoricalReference     bool            `bson:"isHistoricalReference,omitempty" json:"isHistoricalReference,omitempty"`
	NdbNumber                 int32           `bson:"ndbNumber,omitempty" json:"ndbNumber,omitempty"`
	DataType                  string          `bson:"dataType,omitempty" json:"dataType,omitempty"`
	FoodCategory              FoodCategory    `bson:"foodCategory,omitempty" json:"foodCategory,omitempty"`
	FDCID                     int32           `bson:"fdcId,omitempty" json:"fdcId,omitempty"`
	FoodPortions              []FoodPortion   `bson:"foodPortions,omitempty" json:"foodPortions,omitempty"`
	PublicationDate           string          `bson:"publicationDate,omitempty" json:"publicationDate,omitempty"`
	InputFoods                []InputFood     `bson:"inputFoods,omitempty" json:"inputFoods,omitempty"`
}

type FoodNutrients struct {
	ID                     int32                  `bson:"id,omitempty" json:"id,omitempty"`
	Type                   string                 `bson:"type,omitempty" json:"type,omitempty"`
	Nutrients              []Nutrient             `bson:"nutrients,omitempty" json:"nutrients,omitempty"`
	DataPoints             int                    `bson:"dataPoints,omitempty" json:"dataPoints,omitempty"`
	FoodNutrientDerivation FoodNutrientDerivation `bson:"foodNutrientDerivation,omitempty" json:"foodNutrientDerivation,omitempty"`
	Median                 float64                `bson:"median,omitempty" json:"median,omitempty"`
	Amount                 float64                `bson:"amount,omitempty" json:"amount,omitempty"`
}

type Nutrient struct {
	ID       int32  `bson:"id,omitempty" json:"id,omitempty"`
	Number   string `bson:"number,omitempty" json:"number,omitempty"`
	Name     string `bson:"name,omitempty" json:"name,omitempty"`
	Rank     string `bson:"rank,omitempty" json:"rank,omitempty"`
	UnitName string `bson:"unitName,omitempty" json:"unitName,omitempty"`
}

type FoodNutrientDerivation struct {
	Code               string             `bson:"code,omitempty" json:"code,omitempty"`
	Description        string             `bson:"description,omitempty" json:"description,omitempty"`
	FoodNutrientSource FoodNutrientSource `bson:"foodNutrientSource,omitempty" json:"foodNutrientSource,omitempty"`
}

type FoodNutrientSource struct {
	ID          int    `bson:"id,omitempty" json:"id,omitempty"`
	Code        string `bson:"code,omitempty" json:"code,omitempty"`
	Description string `bson:"description,omitempty" json:"description,omitempty"`
}

type FoodCategory struct {
	ID          int    `bson:"id,omitempty" json:"id,omitempty"`
	Code        string `bson:"code,omitempty" json:"code,omitempty"`
	Description string `bson:"description,omitempty" json:"description,omitempty"`
}

type FoodPortion struct {
	ID              int32       `bson:"id,omitempty" json:"id,omitempty"`
	Value           float64     `bson:"value,omitempty" json:"value,omitempty"`
	MeasureUnit     MeasureUnit `bson:"measureUnit,omitempty" json:"measureUnit,omitempty"`
	Modifier        string      `bson:"modifier,omitempty" json:"modifier,omitempty"`
	GramWeight      float64     `bson:"gramWeight,omitempty" json:"gramWeight,omitempty"`
	MinYearAcquired int         `bson:"minYearAcquired,omitempty" json:"minYearAcquired,omitempty"`
	Amount          float64     `bson:"amount,omitempty" json:"amount,omitempty"`
	SequenceNumber  int         `bson:"sequenceNumber,omitempty" json:"sequenceNumber,omitempty"`
}

type MeasureUnit struct {
	ID           int    `bson:"id,omitempty" json:"id,omitempty"`
	Name         string `bson:"name,omitempty" json:"name,omitempty"`
	Abbreviation string `bson:"abbreviation,omitempty" json:"abbreviation,omitempty"`
}

type InputFood struct {
	ID              int            `bson:"id,omitempty" json:"id,omitempty"`
	FoodDescription string         `bson:"foodDescription,omitempty" json:"foodDescription,omitempty"`
	InputFood       FoundationFood `bson:"inputFood,omitempty" json:"inputFood,omitempty"`
}
