package diet

import (
	"github.com/phonghaido/healthy-habits/internal/usda"
)

type MealPlan struct {
	ID             string                    `bson:"id,omitempty" json:"id,omitempty"`
	Name           string                    `bson:"name" json:"name"`
	Items          []usda.FoundationFood     `bson:"items" json:"items"`
	Type           string                    `bson:"type" json:"type"`
	Description    string                    `bson:"description,omitempty" json:"description,omitempty"`
	TotalNutrients map[string]TotalNutrients `bson:"totalNutrients" json:"totalNutrients"`
}

type DietPlan struct {
	ID          string     `bson:"id,omitempty" json:"id,omitempty"`
	Name        string     `bson:"name" json:"name"`
	Description string     `bson:"description,omitempty" json:"description,omitempty"`
	Meals       []MealPlan `bson:"meals" json:"meals"`
}

type TotalNutrients struct {
	Amount float64
	Unit   string
}

func (m MealPlan) CalculateTotalNutrients() map[string]TotalNutrients {
	result := make(map[string]TotalNutrients)

	for _, item := range m.Items {
		for _, nutrient := range item.FoodNutrients {
			nutrientName := nutrient.Nutrient.Name
			amount := nutrient.Amount

			val, ok := result[nutrientName]
			if ok {
				val.Amount += amount
			} else {
				result[nutrientName] = TotalNutrients{
					Amount: amount,
					Unit:   nutrient.Nutrient.UnitName,
				}
			}
		}
	}

	return result
}
