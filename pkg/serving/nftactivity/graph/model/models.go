package model

import (
	"database/sql/driver"
	"encoding/json"

	"errors"
)

type Appraisal struct {
	N05      *float64 `json:"05"`
	N50      *float64 `json:"50"`
	N95      *float64 `json:"95"`
	Currency *string  `json:"currency"`
}

func (a Appraisal) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *Appraisal) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

type Discounts struct {
	N05 *float64 `json:"05"`
	N50 *float64 `json:"50"`
	N95 *float64 `json:"95"`
}

func (d Discounts) Value() (driver.Value, error) {
	return json.Marshal(d)
}

func (d *Discounts) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &d)
}

type AssetValuation struct {
	Confidence *float64   `json:"confidence"`
	Discounts  *Discounts `json:"discounts"`
	Appraisal  *Appraisal `json:"appraisal"`
}

func (t AssetValuation) Value() (driver.Value, error) {
	return json.Marshal(t)
}

func (t *AssetValuation) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &t)
}

type Price struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

func (p Price) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *Price) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &p)
}

type VenuePrice struct {
	Venue string `json:"venue"`
	Price *Price `json:"price"`
}

func (v VenuePrice) Value() (driver.Value, error) {
	return json.Marshal(v)
}

func (v *VenuePrice) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &v)
}

type AssetMarketability struct {
	Floor    *VenuePrice `json:"floor"`
	Ceiling  *VenuePrice `json:"ceiling"`
	LastSale *Price      `json:"last_sale"`
}

func (t AssetMarketability) Value() (driver.Value, error) {
	return json.Marshal(t)
}

func (t *AssetMarketability) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &t)
}
