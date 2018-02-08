package surfliner

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestDefaultClient(t *testing.T) {
	ua := "SurflineR Client"
	bu, err := url.Parse("http://surfliner.maxworld.tech")
	if err != nil {
		t.Fatal(err)
	}
	c, err := DefaultClient()
	if err != nil {
		t.Fatal(err)
	}

	if c.BaseURL.String() != bu.String() {
		t.Errorf("Got %s, expected %s", c.BaseURL.String(), bu.String())
	}
	if c.UserAgent != ua {
		t.Errorf("Got %s, expected %s", c.UserAgent, ua)
	}
}

func TestListAreas(t *testing.T) {
	ua := "SurflineR Test"
	d, err := ioutil.ReadFile("fixtures/areas.json")
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(d)
	}))
	defer ts.Close()

	bu, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	c := Client{BaseURL: bu, UserAgent: ua, httpClient: http.DefaultClient}

	as, err := c.ListAreas()
	if err != nil {
		t.Fatal(err)
	}

	a := as[0]

	id := "4716"
	if a.ID != id {
		t.Errorf("Got %s, expected %s", a.ID, id)
	}

	name := "North America"
	if a.Name != name {
		t.Errorf("Got %s, expected %s", a.Name, name)
	}
}

func TestListRegions(t *testing.T) {
	ua := "SurflineR Test"
	d, err := ioutil.ReadFile("fixtures/regions.json")
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(d)
	}))
	defer ts.Close()

	bu, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	c := Client{BaseURL: bu, UserAgent: ua, httpClient: http.DefaultClient}

	rs, err := c.ListRegions("1")
	if err != nil {
		t.Fatal(err)
	}

	r := rs[0]

	id := "2081"
	if r.ID != id {
		t.Errorf("Got %s, expected %s", r.ID, id)
	}

	name := "Southern California"
	if r.Name != name {
		t.Errorf("Got %s, expected %s", r.Name, name)
	}
}

func TestListSubRegions(t *testing.T) {
	ua := "SurflineR Test"
	d, err := ioutil.ReadFile("fixtures/subregions.json")
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(d)
	}))
	defer ts.Close()

	bu, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	c := Client{BaseURL: bu, UserAgent: ua, httpClient: http.DefaultClient}

	srs, err := c.ListSubRegions("1", "1")
	if err != nil {
		t.Fatal(err)
	}

	sr := srs[0]

	id := "2141"
	if sr.ID != id {
		t.Errorf("Got %s, expected %s", sr.ID, id)
	}

	name := "Santa Barbara"
	if sr.Name != name {
		t.Errorf("Got %s, expected %s", sr.Name, name)
	}
}
