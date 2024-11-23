package intensity_segments

import (
	"testing"
)

type call struct {
	// args for Add or Set method
	from   int
	to     int
	amount int
	// after call Add or Set method, ToString method is expected to return this value
	expectToString string
}

type testcase struct {
	name  string
	calls []call
}

func TestIntensitySegments_Add(t *testing.T) {
	tests := []*testcase{
		{
			name: "an example",
			calls: []call{
				{from: 10, to: 30, amount: 1, expectToString: "[[10,1],[30,0]]"},
				{from: 20, to: 40, amount: 1, expectToString: "[[10,1],[20,2],[30,1],[40,0]]"},
				{from: 10, to: 40, amount: -2, expectToString: "[[10,-1],[20,0],[30,-1],[40,0]]"},
			},
		},
		{
			name: "another example ",
			calls: []call{
				{from: 10, to: 30, amount: 1, expectToString: "[[10,1],[30,0]]"},
				{from: 20, to: 40, amount: 1, expectToString: "[[10,1],[20,2],[30,1],[40,0]]"},
				{from: 10, to: 40, amount: -1, expectToString: "[[20,1],[30,0]]"},
				{from: 10, to: 40, amount: -1, expectToString: "[[10,-1],[20,0],[30,-1],[40,0]]"},
			},
		},
		{
			name: "continuous segments",
			calls: []call{
				{from: 20, to: 30, amount: 1, expectToString: "[[20,1],[30,0]]"},
				{from: 10, to: 20, amount: 1, expectToString: "[[10,1],[30,0]]"},
				{from: 30, to: 40, amount: 1, expectToString: "[[10,1],[40,0]]"},
				{from: 10, to: 20, amount: -1, expectToString: "[[20,1],[40,0]]"},
			},
		},
		{
			name: "no overlapping segments",
			calls: []call{
				{from: 30, to: 40, amount: 1, expectToString: "[[30,1],[40,0]]"},
				{from: 10, to: 20, amount: 1, expectToString: "[[10,1],[20,0],[30,1],[40,0]]"},
				{from: 50, to: 60, amount: -1, expectToString: "[[10,1],[20,0],[30,1],[40,0],[50,-1],[60,0]]"},
			},
		},
		{
			name: "split segment",
			calls: []call{
				{from: 10, to: 60, amount: 1, expectToString: "[[10,1],[60,0]]"},
				{from: 20, to: 30, amount: -1, expectToString: "[[10,1],[20,0],[30,1],[60,0]]"},
				{from: 40, to: 50, amount: -1, expectToString: "[[10,1],[20,0],[30,1],[40,0],[50,1],[60,0]]"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewIntensitySegments()
			if str := s.ToString(); str != "[]" {
				t.Fatalf("expect:%s, actual:%s", "[]", str)
			}
			for _, c := range tt.calls {
				s.Add(c.from, c.to, c.amount)
				if str := s.ToString(); str != c.expectToString {
					t.Fatalf("expect:%s, actual:%s", c.expectToString, str)
				}
			}
		})
	}
}

func TestIntensitySegments_Set(t *testing.T) {
	tests := []*testcase{
		{
			name: "continuous segments",
			calls: []call{
				{from: 10, to: 20, amount: 1, expectToString: "[[10,1],[20,0]]"},
				{from: 20, to: 30, amount: 1, expectToString: "[[10,1],[30,0]]"},
				{from: 30, to: 40, amount: 1, expectToString: "[[10,1],[40,0]]"},
			},
		},
		{
			name: "no overlapping segments",
			calls: []call{
				{from: 30, to: 40, amount: 1, expectToString: "[[30,1],[40,0]]"},
				{from: 10, to: 20, amount: 1, expectToString: "[[10,1],[20,0],[30,1],[40,0]]"},
				{from: 50, to: 60, amount: 1, expectToString: "[[10,1],[20,0],[30,1],[40,0],[50,1],[60,0]]"},
			},
		},
		{
			name: "split segment",
			calls: []call{
				{from: 10, to: 60, amount: 1, expectToString: "[[10,1],[60,0]]"},
				{from: 20, to: 30, amount: 0, expectToString: "[[10,1],[20,0],[30,1],[60,0]]"},
				{from: 40, to: 50, amount: 0, expectToString: "[[10,1],[20,0],[30,1],[40,0],[50,1],[60,0]]"},
			},
		},
		{
			name: "overlapping segments",
			calls: []call{
				{from: 10, to: 30, amount: 1, expectToString: "[[10,1],[30,0]]"},
				{from: 20, to: 40, amount: 3, expectToString: "[[10,1],[20,3],[40,0]]"},
				{from: 40, to: 50, amount: 1, expectToString: "[[10,1],[20,3],[40,1],[50,0]]"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewIntensitySegments()
			if str := s.ToString(); str != "[]" {
				t.Fatalf("expect:%s, actual:%s", "[]", str)
			}
			for _, c := range tt.calls {
				s.Set(c.from, c.to, c.amount)
				if str := s.ToString(); str != c.expectToString {
					t.Fatalf("expect:%s, actual:%s", c.expectToString, str)
				}
			}
		})
	}
}
