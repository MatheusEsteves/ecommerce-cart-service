package models

import "strings"

type CartProductData struct {
	ID   string `bson:"_id" json:"id"`
	Name string `bson:"name" json:"name"`
}

func (c *CartProductData) IsValid() bool {
	return strings.TrimSpace(c.ID) != "" && strings.TrimSpace(c.Name) != ""
}
