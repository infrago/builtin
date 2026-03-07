package builtin

import (
	"testing"

	"github.com/infrago/base"
	"github.com/infrago/infra"
)

func TestNumericBuiltinTypeValueWrappers(t *testing.T) {
	builtin.Setup()

	types := infra.Types()
	getType := func(name string) infra.Type {
		tp, ok := types[name]
		if !ok || tp.Value == nil {
			t.Fatalf("missing builtin type value handler: %s", name)
		}
		return tp
	}

	if got, ok := getType("int").Value(12, base.Var{}).(int64); !ok || got != 12 {
		t.Fatalf("int should wrap to int64, got %#v", getType("int").Value(12, base.Var{}))
	}
	if got, ok := getType("uint").Value(uint(15), base.Var{}).(int64); !ok || got != 15 {
		t.Fatalf("uint should wrap to int64, got %#v", getType("uint").Value(uint(15), base.Var{}))
	}
	if got, ok := getType("float").Value(float32(1.25), base.Var{}).(float64); !ok || got != 1.25 {
		t.Fatalf("float should wrap to float64, got %#v", getType("float").Value(float32(1.25), base.Var{}))
	}

	if got, ok := getType("[int]").Value([]int{1, 2, 3}, base.Var{}).([]int64); !ok || len(got) != 3 || got[0] != 1 || got[1] != 2 || got[2] != 3 {
		t.Fatalf("[int] should wrap to []int64, got %#v", getType("[int]").Value([]int{1, 2, 3}, base.Var{}))
	}
	if got, ok := getType("[uint]").Value([]uint{4, 5}, base.Var{}).([]int64); !ok || len(got) != 2 || got[0] != 4 || got[1] != 5 {
		t.Fatalf("[uint] should wrap to []int64, got %#v", getType("[uint]").Value([]uint{4, 5}, base.Var{}))
	}
	if got, ok := getType("[float]").Value([]float32{1.5, 2.5}, base.Var{}).([]float64); !ok || len(got) != 2 || got[0] != 1.5 || got[1] != 2.5 {
		t.Fatalf("[float] should wrap to []float64, got %#v", getType("[float]").Value([]float32{1.5, 2.5}, base.Var{}))
	}
}
