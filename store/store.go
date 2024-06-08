package store

type Car struct {
	ID    string
	Make  string
	Model string
	Year  int
}

type InMemoryStore struct {
	cars map[string]Car
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		cars: make(map[string]Car),
	}
}

func (s *InMemoryStore) ListCars() []Car {
	cars := []Car{}
	for _, car := range s.cars {
		cars = append(cars, car)
	}
	return cars
}

func (s *InMemoryStore) AddCar(car Car) {
	s.cars[car.ID] = car
}

func (s *InMemoryStore) DeleteCar(id string) {
	delete(s.cars, id)
}

func (s *InMemoryStore) SearchCar(make string) []Car {
	cars := []Car{}
	for _, car := range s.cars {
		if car.Make == make {
			cars = append(cars, car)
		}
	}
	return cars
}
