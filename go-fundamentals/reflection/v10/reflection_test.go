package v10

import (
	"reflect"
	"slices"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	// table test
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "struct with one string field",
			Input: struct {
				Name string
			}{"Chris"},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "struct with two string fields",
			Input: struct {
				Name string
				City string
			}{Name: "Chris", City: "London"},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "struct with non string field",
			Input: struct {
				Name string
				Age  int
			}{"Chris", 33},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "nested fields",
			Input: Person{
				Name:    "Chris",
				Profile: Profile{Age: 33, City: "London"},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "pointers to things",
			Input: &Person{
				Name:    "Chris",
				Profile: Profile{Age: 33, City: "London"},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "slices",
			Input: []Profile{
				{Age: 33, City: "London"},
				{Age: 34, City: "California"},
			},
			ExpectedCalls: []string{"London", "California"},
		},
		{
			Name: "arrays",
			Input: [2]Profile{
				{Age: 33, City: "London"},
				{Age: 34, City: "California"},
			},
			ExpectedCalls: []string{"London", "California"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %q, want %q", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps, because it does not guarantee order", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with chan", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "London"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "London"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := slices.Contains(haystack, needle)
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
