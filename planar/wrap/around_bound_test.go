package wrap

import (
	"reflect"
	"testing"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/planar"
)

func TestRing(t *testing.T) {
	cases := []struct {
		name   string
		bound  planar.Bound
		input  planar.Ring
		output planar.Ring
		orient orb.Orientation
	}{
		{
			name:   "wrap around whole box ccw",
			bound:  planar.Bound{{-1, -1}, {1, 1}},
			input:  planar.Ring{{-2, 0.5}, {0, 0.5}, {0, -0.5}, {-2, -0.5}},
			output: planar.Ring{{-2, 0.5}, {0, 0.5}, {0, -0.5}, {-2, -0.5}, {-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-2, 0.5}},
			orient: orb.CCW,
		},
		{
			name:   "just close the ring",
			bound:  planar.Bound{{-1, -1}, {1, 1}},
			input:  planar.Ring{{-2, 0.5}, {0, 0.5}, {0, -0.5}, {-2, -0.5}},
			output: planar.Ring{{-2, 0.5}, {0, 0.5}, {0, -0.5}, {-2, -0.5}, {-2, 0.5}},
			orient: orb.CW,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			out := Ring(tc.bound, tc.input, tc.orient)
			if !reflect.DeepEqual(out, tc.output) {
				t.Errorf("does not match")
				t.Logf("%v", out)
				t.Logf("%v", tc.output)
			}
		})
	}
}
