package fetcher

import "testing"

func TestFetch(t *testing.T) {
	cases := []struct {
		arg     string
		data    string
		success bool
	}{
		{"does-not-exist", "", false},
		{"testdata/test", "success\n", true},
		{"https://raw.githubusercontent.com/juhovuori/builder/master/fetcher/testdata/does-not-exist", "", false},
		{"https://raw.githubusercontent.com/juhovuori/builder/master/fetcher/testdata/test", "success\n", true},
	}
	for i, c := range cases {
		data, err := Fetch(c.arg)
		if !c.success {
			if err == nil {
				t.Errorf("%d: Expected error.\n", i)
			}
			continue
		}
		if err != nil {
			t.Errorf("%d: Got error %v.\n", i, err)
		}
		if string(data) != c.data {
			t.Errorf("%d: Got %s, expected %s", i, string(data)[:200], c.data)
		}
	}
}
