package main

import "fmt"

func calcSalary(f func(int, int) string, j, k int) {
	fmt.Println(f(j, k))
}
func manager(j, k int) string {
	f := func() string {
		netSalary := j + k
		return fmt.Sprintf("Manager\nBasic salary:$%d | Bonus:$%d\nNet salary:$%d\n\nDeveloper", j, k, netSalary)
	}
	return f()
}
func devSalary(j, k int) string {
	f := func() string {
		netsalary := j * k
		return fmt.Sprintf("Hourly rate:$%d | Hours worked:%d\nNet salary:$%d", j, k, netsalary)
	}
	return f()

}
func main() {
	calcSalary(manager, 150000, 50000)
	calcSalary(devSalary, 1000, 250)

}
