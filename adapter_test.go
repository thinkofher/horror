package horror

import "testing"

func TestDefaultAdapter(t *testing.T) {
	a := NewAdapter(&AdapterBuilder{})

	// Expect these not to panic, because they're
	// doing nothing
	a.afterErr(nil, nil, nil)
	a.beforeErr(nil, nil, nil)

	// Check default fields
	if a.wrapInternal {
		t.Errorf("a.wrapInternal is %v, but should be false", a.wrapInternal)
	}
}
