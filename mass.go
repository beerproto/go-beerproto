package beerprotov1

import "fmt"

func (x *MassType) DisplayName() string {
	return fmt.Sprintf("%v %v", x.Value, x.Unit.DisplayName())
}
