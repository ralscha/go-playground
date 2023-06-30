package vehicle

type Type interface {
	isEnumValue()
}

var (
	Values = struct {
		Motorcycle *motorcycle
		Car        *car
		Bus        *bus
		Truck      *truck
	}{
		Motorcycle: &motorcycle{},
		Car:        &car{},
		Bus:        &bus{},
		Truck:      &truck{},
	}
)

type car struct {
	// potentially, add some fields here to keep some state
}

type motorcycle struct {
	// potentially, add some fields here to keep some state
}

type bus struct {
	// potentially, add some fields here to keep some state
}

type truck struct {
	// potentially, add some fields here to keep some state
}

func (_ *motorcycle) isEnumValue() {}
func (_ *car) isEnumValue()        {}
func (_ *bus) isEnumValue()        {}
func (_ *truck) isEnumValue()      {}

func (_ *truck) FetchSomeData() float64 {
	return 0.01
}
