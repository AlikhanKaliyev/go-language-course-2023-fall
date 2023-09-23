package builder

type Builder struct {
	position string
	salary   float64
	address  string
}

func NewBuilder(position string, salary float64, address string) *Builder {
	return &Builder{
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (h *Builder) SetPosition(position string) {
	h.position = position
}

func (h *Builder) GetPosition() string {
	return h.position
}

func (h *Builder) SetSalary(salary float64) {
	h.salary = salary
}

func (h *Builder) GetSalary() float64 {
	return h.salary
}

func (h *Builder) SetAddress(address string) {
	h.address = address
}

func (h *Builder) GetAddress() string {
	return h.address
}
