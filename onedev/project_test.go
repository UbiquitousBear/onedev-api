package onedev

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	ctx := context.Background()
	got, _, err := client.Projects.List(ctx, 0, 100)
	assert.NoError(t, err)

	want := []Project{{Id: Int(1)}, {Id: Int(2)}}
	assert.Equal(t, &want, got)
}

func TestProjectService_Create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(Project)
		json.NewDecoder(r.Body).Decode(v)
		fmt.Fprint(w, `2`)
	})

	ctx := context.Background()

	input := &Project{Name: String("foo"), Description: String("bar")}
	got, _, err := client.Projects.Create(ctx, input)
	if err != nil {
		t.Errorf("Projects.Create returned error: %v", err)
	}

	want := Project{Id: Int(2), Name: String("foo"), Description: String("bar")}
	assert.Equal(t, want, *got)
}