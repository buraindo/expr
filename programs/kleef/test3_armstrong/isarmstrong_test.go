package armstrong

import (
	"testing"
)

func TestIsArmstrong(t *testing.T) {
	t.Log(IsArmstrong(153))
	t.Log(IsArmstrong(9800817))
	t.Log(IsArmstrong(146511208))
	t.Log(IsArmstrong(4210818))
	t.Log(IsArmstrong(370))
	t.Log(IsArmstrong(371))
	t.Log(IsArmstrong(93084))
	t.Log(IsArmstrong(2937676874))
}
