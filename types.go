package main

type User struct {
	ID        int    `json: "id"`
	DeviceID  string `json: "device_id"`
	Email     string `json: "email"`
	CreatedAt string `json: "created_at"`
	UpdatedAt string `json: "updated_at"`
}

type Item struct {
	ID        int    `json: "id"`
	UserID    int    `json: "user_id"`
	Name      string `json: "name"`
	CreatedAt string `json: "created_at"`
}

type ItemForm struct {
	UserID int    `json: "user_id"`
	Name   string `json: "name"`
}

type Payment struct {
	ID        int    `json: "id"`
	UserID    int    `json: "user_id"`
	ItemID    int    `json: "item_id"`
	StoreName string `json: "store_name"`
	Amount    int    `json: "amount"`
	PaidAt    string `json: "paid_at"`
	CreatedAt string `json: "created_at"`
}
