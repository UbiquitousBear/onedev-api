package onedev

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.NoError(t, err)

	want := []Project{{Id: Int(1)}, {Id: Int(2)}}
	assert.Equal(t, &want, got)
}
