package cook

type Cook struct {
	position string
	salary   float64
	address  string
}

func NewCook(position string, salary float64, address string) *Cook {
	return &Cook{
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (h *Cook) SetPosition(position string) {
	h.position = position
}

func (h *Cook) GetPosition() string {
	return h.position
}

func (h *Cook) SetSalary(salary float64) {
	h.salary = salary
}

func (h *Cook) GetSalary() float64 {
	return h.salary
}

func (h *Cook) SetAddress(address string) {
	h.address = address
}

func (h *Cook) GetAddress() string {
	return h.address
}
