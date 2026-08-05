package amclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const (
	archivematicaVersionHeader = "X-Archivematica-Version"
	archivematicaIDHeader      = "X-Archivematica-ID"
	serverInfoPath             = packageBasePath + "/"
)

// ErrServerInfoUnavailable indicates that an Archivematica response does not
// include the server version header. Use [errors.Is] to test for this error.
var ErrServerInfoUnavailable = errors.New("archivematica server information unavailable")

// Version is an immutable numeric Archivematica release version parsed from the
// X-Archivematica-Version response header. Its zero value represents 0.0.0.
type Version struct {
	major uint64
	minor uint64
	patch uint64
}

// Major returns the major version component.
func (v Version) Major() uint64 {
	return v.major
}

// Minor returns the minor version component.
func (v Version) Minor() uint64 {
	return v.minor
}

// Patch returns the patch version component.
func (v Version) Patch() uint64 {
	return v.patch
}

// String returns the version in major.minor.patch form.
func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.patch)
}

// Compare compares v with other. It returns -1, 0, or 1 when v is less than,
// equal to, or greater than other, respectively.
func (v Version) Compare(other Version) int {
	for _, components := range [][2]uint64{
		{v.major, other.major},
		{v.minor, other.minor},
		{v.patch, other.patch},
	} {
		if components[0] < components[1] {
			return -1
		}
		if components[0] > components[1] {
			return 1
		}
	}

	return 0
}

// AtLeast reports whether v is greater than or equal to the version identified
// by major, minor, and patch. It is intended for release feature checks such as
// v.AtLeast(1, 19, 0).
func (v Version) AtLeast(major, minor, patch uint64) bool {
	return v.Compare(Version{major: major, minor: minor, patch: patch}) >= 0
}

func parseVersion(value string) (Version, error) {
	parts := strings.Split(value, ".")
	if len(parts) != 3 {
		return Version{}, fmt.Errorf("invalid Archivematica version %q: expected three numeric components", value)
	}

	components := [3]uint64{}
	for i, part := range parts {
		if part == "" || strings.Trim(part, "0123456789") != "" {
			return Version{}, fmt.Errorf("invalid Archivematica version %q: expected three numeric components", value)
		}

		component, err := strconv.ParseUint(part, 10, 64)
		if err != nil {
			return Version{}, fmt.Errorf("invalid Archivematica version %q: %w", value, err)
		}
		components[i] = component
	}

	return Version{
		major: components[0],
		minor: components[1],
		patch: components[2],
	}, nil
}

// ServerInfo describes the Archivematica server and pipeline that handled an
// API request.
type ServerInfo struct {
	// Version is the Archivematica release version.
	Version Version

	// ID is the pipeline identifier shared by Dashboard, MCPServer, MCPClient,
	// and other Archivematica subsystems. It is empty when the server does not
	// provide X-Archivematica-ID.
	ID string
}

// ServerInfo extracts Archivematica server information from response headers.
// It does not make a network request. It returns [ErrServerInfoUnavailable]
// when the version header is absent and an error when the version is malformed.
func (r *Response) ServerInfo() (*ServerInfo, error) {
	if r == nil || r.Response == nil {
		return nil, ErrServerInfoUnavailable
	}

	value := r.Header.Get(archivematicaVersionHeader)
	if value == "" {
		return nil, fmt.Errorf("%w: missing %s header", ErrServerInfoUnavailable, archivematicaVersionHeader)
	}

	version, err := parseVersion(value)
	if err != nil {
		return nil, err
	}

	return &ServerInfo{
		Version: version,
		ID:      r.Header.Get(archivematicaIDHeader),
	}, nil
}

// ServerInfo retrieves information about the Archivematica Dashboard server.
// Each call makes an authenticated, non-mutating GET request; results are not
// cached.
//
// Archivematica's package endpoint returns 501 Not Implemented for GET while
// including server information in its response headers. ServerInfo treats that
// expected status as successful when the headers are valid. In that case the
// returned Response retains its 501 status and err is nil.
func (c *Client) ServerInfo(ctx context.Context) (*ServerInfo, *Response, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, serverInfoPath, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, requestErr := c.Do(ctx, req, nil)
	if resp == nil {
		return nil, nil, requestErr
	}

	info, infoErr := resp.ServerInfo()
	if infoErr != nil {
		if requestErr != nil && resp.StatusCode != http.StatusNotImplemented {
			return nil, resp, requestErr
		}
		return nil, resp, infoErr
	}

	if requestErr == nil || resp.StatusCode == http.StatusNotImplemented {
		return info, resp, nil
	}

	return info, resp, requestErr
}
