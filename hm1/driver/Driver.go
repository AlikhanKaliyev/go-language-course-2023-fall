package driver

type Driver struct {
	position string
	salary   float64
	address  string
}

func NewDriver(position string, salary float64, address string) *Driver {
	return &Driver{
		position: position,
		salary:   salary,
		address:  address,
	}
}

func (h *Driver) SetPosition(position string) {
	h.position = position
}

func (h *Driver) GetPosition() string {
	return h.position
}

func (h *Driver) SetSalary(salary float64) {
	h.salary = salary
}

func (h *Driver) GetSalary() float64 {
	return h.salary
}

func (h *Driver) SetAddress(address string) {
	h.address = address
}

func (h *Driver) GetAddress() string {
	return h.address
}
