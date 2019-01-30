package vehicle

type Vehicle struct {
	RegistrationNumber string
	Color              string
}

func (v Vehicle) getColor() string {
	return v.Color
}

func (v Vehicle) getRegistrationNumber() string {
	return v.RegistrationNumber
}
