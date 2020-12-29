package model

type student struct {
	name string
	age  byte
}

// GetDistance  *student
func GetDistance(n string, a byte) *student {
	return &student{
		name: n,
		age:  a,
	}
}

func (stu *student) GetName() string {
	return stu.name
}

func (stu *student) GetAge() byte {
	return stu.age
}
