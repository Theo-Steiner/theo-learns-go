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

		t.Run("maps", func(t *testing.T) {
			// map might fail, because go does not guarantee order for maps
			myMap := map[string]string{
				"cow":   "moo",
				"sheep": "baa",
			}
			var got []string
			walk(myMap, func(input string) {
				got = append(got, input)
			})
			assertContains(t, got, "moo")
			assertContains(t, got, "baa")
		})

		t.Run("channels", func(t *testing.T) {
			aChannel := make(chan Profile)
			go func() {
				aChannel <- Profile{"Berlin", 33}
				aChannel <- Profile{"Katowice", 34}
				close(aChannel)
			}()
			var got []string
			want := []string{"Berlin", "Katowice"}
			err := walk(aChannel, func(input string) {
				got = append(got, input)
			})

			if err != nil {
				t.Fatalf("Did not expect error, but received %q", err)
			}

			if !slices.Equal(got, want) {
				t.Errorf("wrong number of function calls, got %v want %v", got, test.ExpectedCalls)
			}
		})

		t.Run("func", func(t *testing.T) {
			aFunction := func() (a, b Profile) {
				return Profile{"Berlin", 34}, Profile{"Katowice", 34}
			}
			var got []string
			want := []string{"Berlin", "Katowice"}
			err := walk(aFunction, func(input string) {
				got = append(got, input)
			})

			if err != nil {
				t.Fatalf("Did not expect error, but received %q", err)
			}

			if !slices.Equal(got, want) {
				t.Errorf("wrong number of function calls, got %v want %v", got, test.ExpectedCalls)
			}
		})
	}
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()
	didFindNeedle := false
	for _, element := range haystack {
		if element == needle {
			didFindNeedle = true
		}
	}
	if !didFindNeedle {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
