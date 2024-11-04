package nutrient

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/phonghaido/healthy-habits/internal/config"
)

type FoodItem struct {
	FDCID           int32          `json:"fdcId" bson:"fdcId"`
	Description     string         `json:"description" bson:"description,omitempty"`
	DataType        string         `json:"dataType" bson:"dataType,omitempty"`
	PublicationDate string         `json:"publicationDate" bson:"publicationDate,omitempty"`
	FoodNutrients   []FoodNutrient `json:"foodNutrients" bson:"foodNutrients,omitempty"`
}

type FoodNutrient struct {
	Number                string  `json:"number" bson:"number,omitempty"`
	Name                  string  `json:"name" bson:"name,omitempty"`
	Amount                float32 `json:"amount" bson:"amount,omitempty"`
	UnitName              string  `json:"unitName" bson:"unitName,omitempty"`
	DerivationCode        string  `json:"derivationCode" bson:"derivationCode,omitempty"`
	DerivationDescription string  `json:"derivationDescription" bson:"derivationDescription,omitempty"`
}

type USDAClient struct {
	USDAConfig config.USDAConfig
	Client     http.Client
}

func NewUSDAClient() (*USDAClient, error) {
	usdaConfig, err := config.GetUSDAConfig()
	if err != nil {
		return nil, err
	}

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	return &USDAClient{
		USDAConfig: usdaConfig,
		Client:     client,
	}, nil
}

func (c USDAClient) GetAllFood() ([]FoodItem, error) {
	ctx := context.Background()
	endpoint := "/v1/foods/list"
	req, err := http.NewRequestWithContext(ctx, "GET", c.USDAConfig.URL+endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Api-Key", c.USDAConfig.APIKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var foodItems []FoodItem
	err = json.Unmarshal(body, &foodItems)
	if err != nil {
		return nil, err
	}

	return foodItems, nil
}
