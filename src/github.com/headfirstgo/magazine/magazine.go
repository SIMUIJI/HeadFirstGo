package magazine

//type Subscriber struct {
//	Name        string
//	Rate        float64
//	Active      bool
//	HomeAddress Address
//}
//
//type Employee struct {
//	Name        string
//	Salary      float64
//	HomeAddress Address
//}

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address
}

type Employee struct {
	Name   string
	Salary float64
	Address
}

// Employee.Address.Street
// Employee.Street

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
