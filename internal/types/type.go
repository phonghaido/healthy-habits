package internal_type

type FindFoodReqBody struct {
	Category    string `json:"category,omitempty"`
	Description string `json:"description,omitempty"`
}

type FindMealReqBody struct {
	Name string `json:"string"`
}
