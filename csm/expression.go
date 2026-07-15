package csm

import "fmt"

type Expression string

func (e Expression) String() string {
	return fmt.Sprintf("expression=%v", string(e))
}
