package models

type Class struct {
	ID        int64  `json:"id"`
	DateTime  string `json:"datetime"`
	Location  string `json:"location"`
	Price     int64  `json:"price"`
	TrainerID int64  `json:"trainer_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAT string `json:"updated_at"`
}

type CreateClassRequest struct {
	DateTime string `json:"datetime"`
	Location string `json:"location"`
	Price    int64  `json:"price"`
}

type UpdateClassRequest struct {
	DateTime  string `json:"datetime"`
	Location  string `json:"location"`
	Price     int64  `json:"price"`
	TrainerID int64  `json:"trainer_id"`
}

type Trainer struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type ClassResponse struct {
	ID       int64   `json:"id"`
	DateTime string  `json:"datetime"`
	Location string  `json:"location"`
	Price    int64   `json:"price"`
	Trainer  Trainer `json:"trainer"`
}
