package service_domain

type serviceDomain struct {
	id              int
	name            string
	price           float32
	durationMinutes int
	available       bool
}

func NewServiceDomain(name string, price float32, durationMinutes int, available bool) *serviceDomain {
	return &serviceDomain{
		name:            name,
		price:           price,
		durationMinutes: durationMinutes,
		available:       available,
	}
}

func (sd *serviceDomain) GetPrice() float32 {
	return sd.price
}

func (sd *serviceDomain) GetID() int {
	return sd.id
}

func (sd *serviceDomain) GetName() string {
	return sd.name
}

func (sd serviceDomain) GetDurationMinutes() int {
	return sd.durationMinutes
}

func (sd *serviceDomain) GetAvailable() bool { return sd.available }

func (sd *serviceDomain) SetID(id int) {
	sd.id = id
}
