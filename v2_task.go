package amclient

import (
	"bytes"
	"context"
	"fmt"
	"strconv"
	"time"
)

const taskBasePath = "api/v2beta/task"

type TaskService interface {
	Read(context.Context, string) (*Task, *Response, error)
}

type TaskServiceOp struct {
	client *Client
}

var _ TaskService = &TaskServiceOp{}

type TaskDateTime struct {
	time.Time
}

func (t *TaskDateTime) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		t.Time = time.Time{}
		return nil
	}

	s := string(bytes.Trim(data, `"`))

	const layout = "2006-01-02T15:04:05"
	td, err := time.Parse(layout, s)
	if err != nil {
		return err
	}

	t.Time = td
	return nil
}

// TaskDuration extends the standard time.Duration type to provide specialized
// unmarshaling from JSON. Unlike time.Duration, which typically requires
// duration values to be specified in string format (like "300ms", "1h30m"),
// TaskDuration is designed to unmarshal from two distinct JSON value types:
//
//  1. Integer values representing the duration in seconds.
//  2. A specific string literal "< 1", which we interpret as half a second.
//
// Examples:
//   - `{"duration": 60}` unmarshals to 60 seconds,
//   - `{"duration": "< 1"}` unmarshals to 500 milliseconds.
//
// Python function: https://github.com/artefactual/archivematica/blob/2c03a79e710d655c3740ddbc4bc62eebd4e4b4fe/src/dashboard/src/components/helpers.py#L163-L170.
type TaskDuration time.Duration

func (d *TaskDuration) UnmarshalJSON(data []byte) error {
	unquoted := bytes.Trim(data, `"`)

	// Check for the special "< 1" case.
	if bytes.Equal(unquoted, []byte("< 1")) {
		*d = TaskDuration(time.Second / 2)
		return nil
	}

	// Check for the empty string case.
	if len(unquoted) == 0 {
		*d = TaskDuration(0)
		return nil
	}

	// Parse the integer value.
	seconds, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}

	*d = TaskDuration(time.Duration(seconds) * time.Second)

	return nil
}

type Task struct {
	ID          string       `json:"uuid"`
	ExitCode    int          `json:"exit_code"`
	FileID      string       `json:"file_uuid"`
	Filename    string       `json:"file_name"`
	CreatedAt   TaskDateTime `json:"time_created"`
	StartedAt   TaskDateTime `json:"time_started"`
	CompletedAt TaskDateTime `json:"time_ended"`
	Duration    TaskDuration `json:"duration"`
}

func (s *TaskServiceOp) Read(ctx context.Context, ID string) (*Task, *Response, error) {
	path := fmt.Sprintf("%s/%s", taskBasePath, ID)

	req, err := s.client.NewRequestJSON(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	payload := &Task{}
	resp, err := s.client.Do(ctx, req, payload)

	return payload, resp, err
}
