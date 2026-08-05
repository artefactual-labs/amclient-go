package amclient

import (
	"errors"
	"net/http"
	"testing"

	"gotest.tools/v3/assert"
)

func TestVersionComponentsAndComparison(t *testing.T) {
	version, err := parseVersion("1.19.2")
	assert.NilError(t, err)
	assert.Equal(t, version.Major(), uint64(1))
	assert.Equal(t, version.Minor(), uint64(19))
	assert.Equal(t, version.Patch(), uint64(2))
	assert.Equal(t, version.String(), "1.19.2")

	comparisons := []struct {
		other string
		want  int
	}{
		{other: "0.99.99", want: 1},
		{other: "1.18.99", want: 1},
		{other: "1.19.1", want: 1},
		{other: "1.19.2", want: 0},
		{other: "1.19.3", want: -1},
		{other: "1.20.0", want: -1},
		{other: "2.0.0", want: -1},
	}
	for _, tc := range comparisons {
		t.Run(tc.other, func(t *testing.T) {
			other, err := parseVersion(tc.other)
			assert.NilError(t, err)
			assert.Equal(t, version.Compare(other), tc.want)
		})
	}

	assert.Assert(t, version.AtLeast(1, 19, 2))
	assert.Assert(t, version.AtLeast(1, 18, 99))
	assert.Assert(t, !version.AtLeast(1, 19, 3))
}

func TestParseVersionRejectsInvalidValues(t *testing.T) {
	for _, value := range []string{
		"",
		"1.19",
		"1.19.0.1",
		"v1.19.0",
		"1.19.0-rc.1",
		"1.-19.0",
		"1.x.0",
		"18446744073709551616.0.0",
	} {
		t.Run(value, func(t *testing.T) {
			_, err := parseVersion(value)
			assert.ErrorContains(t, err, "invalid Archivematica version")
		})
	}
}

func TestResponse_ServerInfo(t *testing.T) {
	header := http.Header{}
	header.Set(archivematicaVersionHeader, "1.19.0")
	header.Set(archivematicaIDHeader, "5e557ca1-1b89-4bb9-a4b3-bdde4db1e3a2")
	resp := newResponse(&http.Response{Header: header})

	info, err := resp.ServerInfo()
	assert.NilError(t, err)
	assert.Equal(t, info.Version.String(), "1.19.0")
	assert.Equal(t, info.ID, "5e557ca1-1b89-4bb9-a4b3-bdde4db1e3a2")
}

func TestResponse_ServerInfoWithoutID(t *testing.T) {
	resp := newResponse(&http.Response{Header: http.Header{
		archivematicaVersionHeader: []string{"1.19.0"},
	}})

	info, err := resp.ServerInfo()
	assert.NilError(t, err)
	assert.Equal(t, info.ID, "")
}

func TestResponse_ServerInfoWithoutVersion(t *testing.T) {
	for _, resp := range []*Response{
		nil,
		newResponse(&http.Response{Header: http.Header{}}),
	} {
		_, err := resp.ServerInfo()
		assert.ErrorIs(t, err, ErrServerInfoUnavailable)
	}
}

func TestResponse_ServerInfoWithInvalidVersion(t *testing.T) {
	resp := newResponse(&http.Response{Header: http.Header{
		archivematicaVersionHeader: []string{"latest"},
	}})

	_, err := resp.ServerInfo()
	assert.ErrorContains(t, err, "invalid Archivematica version")
}

func TestClient_ServerInfo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2beta/package/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		assert.Equal(t, r.Header.Get("Authorization"), "ApiKey :")
		w.Header().Set(archivematicaVersionHeader, "1.19.0")
		w.Header().Set(archivematicaIDHeader, "5e557ca1-1b89-4bb9-a4b3-bdde4db1e3a2")
		w.WriteHeader(http.StatusNotImplemented)
	})

	info, resp, err := client.ServerInfo(ctx)
	assert.NilError(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusNotImplemented)
	assert.Equal(t, info.Version.String(), "1.19.0")
	assert.Equal(t, info.ID, "5e557ca1-1b89-4bb9-a4b3-bdde4db1e3a2")
}

func TestClient_ServerInfoWithSuccessfulResponse(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2beta/package/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(archivematicaVersionHeader, "1.20.0")
		w.WriteHeader(http.StatusOK)
	})

	info, resp, err := client.ServerInfo(ctx)
	assert.NilError(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusOK)
	assert.Equal(t, info.Version.String(), "1.20.0")
}

func TestClient_ServerInfoWithoutVersion(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2beta/package/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	})

	info, resp, err := client.ServerInfo(ctx)
	assert.Assert(t, info == nil)
	assert.Equal(t, resp.StatusCode, http.StatusNotImplemented)
	assert.ErrorIs(t, err, ErrServerInfoUnavailable)
}

func TestClient_ServerInfoWithAuthenticationFailure(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2beta/package/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	})

	info, resp, err := client.ServerInfo(ctx)
	assert.Assert(t, info == nil)
	assert.Equal(t, resp.StatusCode, http.StatusForbidden)
	assert.ErrorContains(t, err, "403")
}

func TestClient_ServerInfoWithUnexpectedFailure(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2beta/package/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(archivematicaVersionHeader, "1.19.0")
		w.WriteHeader(http.StatusInternalServerError)
	})

	info, resp, err := client.ServerInfo(ctx)
	assert.Equal(t, info.Version.String(), "1.19.0")
	assert.Equal(t, resp.StatusCode, http.StatusInternalServerError)
	assert.ErrorContains(t, err, "500")
}

func TestClient_ServerInfoWithNetworkFailure(t *testing.T) {
	want := errors.New("network failure")
	client := NewClient(&http.Client{
		Transport: roundTripperFunc(func(*http.Request) (*http.Response, error) {
			return nil, want
		}),
	}, "http://127.0.0.1", "", "")

	info, resp, err := client.ServerInfo(ctx)
	assert.Assert(t, info == nil)
	assert.Assert(t, resp == nil)
	assert.ErrorIs(t, err, want)
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}
