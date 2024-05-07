package decorator

import (
	"fmt"
	"testing"
)

func Test_eDecorator(t *testing.T) {
	var c Component = &ConcreteComponent{}
	c = WarpAddDecorator(c, 10)
	c = WarpMulDecorator(c, 8)
	res := c.Calc()

	fmt.Printf("res %d\n", res)

	// Output:
	// res 80
}
