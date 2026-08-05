package amclient

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"testing"

	"gotest.tools/v3/assert"
)

func TestPackage_Create(t *testing.T) {
	setup()
	defer teardown()

	var (
		path        = "<uuid>:<path>"
		pathb64     = base64.StdEncoding.EncodeToString([]byte(path))
		autoApprove = true
	)

	mux.HandleFunc("/api/v2beta/package/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, ok := r.Header["Idempotency-Key"]
		assert.Assert(t, !ok)

		blob, err := io.ReadAll(r.Body)
		assert.NilError(t, err)
		defer r.Body.Close()

		assert.DeepEqual(t,
			string(bytes.TrimSpace(blob)),
			fmt.Sprintf(`{"name":"Foobar","type":"standard","path":"%s","accession":"12345","access_system_id":"fig-123","processing_config":"automated","auto_approve":true}`, pathb64))

		fmt.Fprint(w, `{"id": "096a284d-5067-4de0-a0a4-a684018cd6df"}`)
	})

	req := &PackageCreateRequest{
		Name:             "Foobar",
		Type:             "standard",
		Path:             path,
		Accession:        "12345",
		AccessSystemID:   "fig-123",
		ProcessingConfig: "automated",
		AutoApprove:      &autoApprove,
	}
	payload, _, err := client.Package.Create(ctx, req)
	assert.NilError(t, err)
	assert.Equal(t, req.Path, path)
	assert.Equal(t, payload.ID, "096a284d-5067-4de0-a0a4-a684018cd6df")
}

func TestPackage_CreateWithIdempotencyKey(t *testing.T) {
	setup()
	defer teardown()

	var (
		path    = "<uuid>:<path>"
		pathb64 = base64.StdEncoding.EncodeToString([]byte(path))
	)

	mux.HandleFunc("/api/v2beta/package/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		assert.Equal(t, r.Header.Get("Idempotency-Key"), "transfer-submission-123")

		blob, err := io.ReadAll(r.Body)
		assert.NilError(t, err)
		defer r.Body.Close()

		assert.DeepEqual(t,
			string(bytes.TrimSpace(blob)),
			fmt.Sprintf(`{"name":"Foobar","type":"standard","path":"%s"}`, pathb64))

		fmt.Fprint(w, `{"id": "096a284d-5067-4de0-a0a4-a684018cd6df"}`)
	})

	req := &PackageCreateRequest{
		Name:           "Foobar",
		Type:           "standard",
		Path:           path,
		IdempotencyKey: "transfer-submission-123",
	}
	for i := 0; i < 2; i++ {
		payload, _, err := client.Package.Create(ctx, req)
		assert.NilError(t, err)
		assert.Equal(t, payload.ID, "096a284d-5067-4de0-a0a4-a684018cd6df")
		assert.Equal(t, req.Path, path)
	}
}

func TestPackage_CreateWithoutAutoApprove(t *testing.T) {
	setup()
	defer teardown()

	var (
		path        = "<uuid>:<path>"
		pathb64     = base64.StdEncoding.EncodeToString([]byte(path))
		autoApprove = false
	)

	mux.HandleFunc("/api/v2beta/package/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		blob, err := io.ReadAll(r.Body)
		assert.NilError(t, err)
		defer r.Body.Close()

		assert.DeepEqual(t,
			string(bytes.TrimSpace(blob)),
			fmt.Sprintf(`{"name":"Foobar","type":"standard","path":"%s","auto_approve":false}`, pathb64))

		fmt.Fprint(w, `{"id": "096a284d-5067-4de0-a0a4-a684018cd6df"}`)
	})

	req := &PackageCreateRequest{
		Name:        "Foobar",
		Type:        "standard",
		Path:        path,
		AutoApprove: &autoApprove,
	}

	payload, _, err := client.Package.Create(ctx, req)
	assert.NilError(t, err)
	assert.Equal(t, payload.ID, "096a284d-5067-4de0-a0a4-a684018cd6df")
	assert.Equal(t, req.Path, path)
}
