package metadata

import (
	"fmt"
	"homework/pkg/vm/types"
)

func (a Attribute) String() string {
	return fmt.Sprintf("%s(%s)", a.Name, &types.T{a.Type})
}
