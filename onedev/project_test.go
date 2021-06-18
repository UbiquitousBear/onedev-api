package onedev

import (
	"context"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"testing"
)

func TestProjectService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects?offset=0&count=100", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	ctx := context.Background()
	got, _, err := client.Projects.List(ctx, 0, 100)
	if err != nil {
		t.Errorf("Projects.List returned error: %v", err)
	}

	want := []*Project{{Id: Int(1)}, {Id: Int(2)}}
	if !cmp.Equal(got, want) {
		t.Errorf("Repositories.List returned %+v, want %+v", got, want)
	}
}