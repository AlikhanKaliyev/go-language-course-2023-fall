package president

type President struct {
	position string
	salary   float64
	address  string
}

func NewPresident(position string, salary float64, address string) *President {
	return &President{
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (h *President) SetPosition(position string) {
	h.position = position
}

func (h *President) GetPosition() string {
	return h.position
}

func (h *President) SetSalary(salary float64) {
	h.salary = salary
}

func (h *President) GetSalary() float64 {
	return h.salary
}

func (h *President) SetAddress(address string) {
	h.address = address
}

func (h *President) GetAddress() string {
	return h.address
}
