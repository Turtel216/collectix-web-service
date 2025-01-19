package models

import "github.com/google/uuid"

type Collection struct {
	CollectionId uint64    `json:"collection_id"`
	UserId       uuid.UUID `json:"user_id"`
	Items        []Item    `json:"items"`
}

type Item struct {
	ItemID   uuid.UUID `json:"item_id"`
	Quantity uint      `json:"quantity"`
	Price    uint      `json:"price"`
}
