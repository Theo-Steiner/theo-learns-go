package main

import (
	"slices"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Profession string
	Age        int
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Theo"},
			[]string{"Theo"},
		},
		{
			"struct with two string fields",
			struct {
				Name       string
				Profession string
			}{"Theo", "Engineer"},
			[]string{"Theo", "Engineer"},
		},
		{
			"struct with two string fields",
			struct {
				Name       string
				Profession string
				Age        int
			}{"Theo", "Engineer", 27},
			[]string{"Theo", "Engineer"},
		},
		{
			"nested fields",
			Person{
				"Theo",
				Profile{"Engineer", 27},
			},
			[]string{"Theo", "Engineer"},
		},
		{"pointers to things",
			&Profile{"Engineer", 27},
			[]string{"Engineer"},
		},
		{"slices",
			[]Profile{
				{"Engineer", 27},
				{"Merchandise Coordinator", 26},
			},
			[]string{"Engineer", "Merchandise Coordinator"},
		},
		{"arrays",
			[2]Profile{
				{"Engineer", 27},
				{"Merchandise Coordinator", 26},
			},
			[]string{"Engineer", "Merchandise Coordinator"},
		},
		{"maps",
			map[string]string{
				"cow":   "moo",
				"sheep": "baa",
			},
			[]string{"moo", "baa"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			err := walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if err != nil {
				t.Fatalf("Did not expect error, but received %q", err)
			}

			if !slices.Equal(got, test.ExpectedCalls) {
				t.Errorf("wrong number of function calls, got %v want %v", got, test.ExpectedCalls)
			}
		})
	}
}
