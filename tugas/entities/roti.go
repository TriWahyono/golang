package entities

type Roti struct {
	Id          int64
	Nama_roti   string `validate:"required"`
	Tittle      string `validate:"required"`
	Description string `validate:"required"`
	Rating      string `validate:"required"`
	Image       string `validate:"required"`
}
