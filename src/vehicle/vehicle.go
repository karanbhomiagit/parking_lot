package vehicle

type Vehicle struct {
	RegistrationNumber string
	Color              string
}

func (v Vehicle) GetColor() string {
	return v.Color
}

func (v Vehicle) GetRegistrationNumber() string {
	return v.RegistrationNumber
}

func (v Vehicle) SetColor(c string) {
	v.Color = c
}

func (v Vehicle) SetRegistrationNumber(r string) {
	v.RegistrationNumber = r
}
