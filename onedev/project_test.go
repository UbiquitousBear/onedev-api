package onedev

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectService_List(t *testing.T) {
	// Arrange
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":1},{"id":2}]`)
	})

	// Act
	ctx := context.Background()
	got, _, err := client.Projects.List(ctx, 0, 100)

	// Assert
	assert.NoError(t, err)
	want := []Project{{Id: Int(1)}, {Id: Int(2)}}
	assert.Equal(t, &want, got)
}

func TestProjectService_Create(t *testing.T) {
	// Arrange
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		v := new(Project)
		json.NewDecoder(r.Body).Decode(v)
		fmt.Fprint(w, `2`)
	})
	input := &Project{Name: String("foo"), Description: String("bar")}

	// Act
	ctx := context.Background()
	got, _, err := client.Projects.Create(ctx, input)

	// Assert
	if err != nil {
		t.Errorf("Projects.Create returned error: %v", err)
	}

	want := Project{Id: Int(2), Name: String("foo"), Description: String("bar")}
	assert.Equal(t, want, *got)
}

func TestProjectService_Read(t *testing.T) {
	// Arrange
	client, mux, _, teardown := setup()
	defer teardown()
	id := 2
	want := Project{Id: Int(id), Name: String("foo"), Description: String("bar")}

	mux.HandleFunc("/projects/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		buf := new(bytes.Buffer)
		json.NewEncoder(buf).Encode(want)
		fmt.Fprint(w, buf.String())
	})

	ctx := context.Background()
	got, _, err := client.Projects.Read(ctx, id)

	// Assert
	if err != nil {
		t.Errorf("Projects.Read returned error: %v", err)
	}

	assert.Equal(t, want, *got)
}

func TestProjectService_Update(t *testing.T) {
	// Arrange
	client, mux, _, teardown := setup()
	defer teardown()
	id := 2
	want := Project{Id: Int(id), Name: String("foo"), Description: String("bar")}

	mux.HandleFunc("/projects/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		buf := new(bytes.Buffer)
		json.NewEncoder(buf).Encode(want)
		fmt.Fprint(w, buf.String())
	})

	ctx := context.Background()
	got, _, err := client.Projects.Update(ctx, &want)

	// Assert
	if err != nil {
		t.Errorf("Projects.Update returned error: %v", err)
	}

	assert.Equal(t, want, *got)
}

func TestProjectService_Delete(t *testing.T) {
	// Arrange
	client, mux, _, teardown := setup()
	defer teardown()
	id := 2

	mux.HandleFunc("/projects/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		fmt.Fprint(w, "")
	})

	ctx := context.Background()
	err := client.Projects.Delete(ctx, id)

	// Assert
	if err != nil {
		t.Errorf("Projects.Update returned error: %v", err)
	}

	assert.Nil(t, err)
}
