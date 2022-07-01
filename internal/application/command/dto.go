package command

type Menu struct {
	Menu  string  `json:"menu" from:"menu" validate:"required,min=3,max=32"`
	Price float32 `json:"price" from:"price" validate:"required,number"`
}
