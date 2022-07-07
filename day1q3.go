package main

import "fmt"
type Salary interface {
	sal() int
}
type employee struct {
	basic int
	duration int
}
type freelancer employee
type fulltime employee
type contractor employee



func (e freelancer)sal() int {
	return e.basic*e.duration
}
func (e fulltime)sal() int {
	return e.basic*e.duration
}
func (e contractor)sal() int {
	return e.basic*e.duration
}
func main() {
	f:=freelancer{
		basic: 100,
		duration: 28}
	ft:=fulltime{500,28}
	c:=contractor{10,20}
	fmt.Println(f.sal())
	fmt.Println(ft.sal())
	fmt.Println(c.sal())

}
