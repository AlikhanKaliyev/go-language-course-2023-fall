package teacher

type Teacher struct {
	position string
	salary   float64
	address  string
}

func NewTeacher(position string, salary float64, address string) *Teacher {
	return &Teacher{
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (h *Teacher) SetPosition(position string) {
	h.position = position
}

func (h *Teacher) GetPosition() string {
	return h.position
}

func (h *Teacher) SetSalary(salary float64) {
	h.salary = salary
}

func (h *Teacher) GetSalary() float64 {
	return h.salary
}

func (h *Teacher) SetAddress(address string) {
	h.address = address
}

func (h *Teacher) GetAddress() string {
	return h.address
}
