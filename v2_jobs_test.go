package amclient

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

func TestJobs_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/v2beta/jobs/e99afef7-90c5-4fd9-bf8f-bed13b3bd4ba/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		assert.Assert(t, r.URL.Query().Has("detailed") && r.URL.Query().Get("detailed") == "")

		fmt.Fprint(w, `[
	{
		"uuid": "624581dc-ec01-4195-9da3-db0ab0ad1cc3",
		"name": "Check transfer directory for objects",
		"status": "COMPLETE",
		"microservice": "Create SIP from Transfer",
		"link_uuid": "6ce08b95-1b3b-498a-8baa-e595e2ae7466",
		"tasks": [
			{
				"uuid": "491aebbd-457b-4a6e-adf6-87a3a9ee951a",
				"exit_code": 1,
				"file_name": "Images-94ade01c-49ce-49e0-9cc3-805575c676d0",
				"time_created": "2024-01-18T01:27:49",
				"time_started": "2024-01-18T01:27:49",
				"time_ended": "2024-01-18T01:27:49",
				"duration": "< 1"
			}
		]
	}
]`)
	})

	payload, _, err := client.Jobs.List(ctx, "e99afef7-90c5-4fd9-bf8f-bed13b3bd4ba", &JobsListRequest{
		LinkID:   "6ce08b95-1b3b-498a-8baa-e595e2ae7466",
		Detailed: true,
	})

	assert.NilError(t, err)
	assert.DeepEqual(t, []Job{
		{
			ID:           "624581dc-ec01-4195-9da3-db0ab0ad1cc3",
			Name:         "Check transfer directory for objects",
			Status:       JobStatusComplete,
			Microservice: "Create SIP from Transfer",
			LinkID:       "6ce08b95-1b3b-498a-8baa-e595e2ae7466",
			Tasks: []Task{
				{
					ID:          "491aebbd-457b-4a6e-adf6-87a3a9ee951a",
					ExitCode:    1,
					Filename:    "Images-94ade01c-49ce-49e0-9cc3-805575c676d0",
					CreatedAt:   TaskDateTime{time.Date(2024, time.January, 18, 1, 27, 49, 0, time.UTC)},
					StartedAt:   TaskDateTime{time.Date(2024, time.January, 18, 1, 27, 49, 0, time.UTC)},
					CompletedAt: TaskDateTime{time.Date(2024, time.January, 18, 1, 27, 49, 0, time.UTC)},
					Duration:    TaskDuration(time.Second / 2),
				},
			},
		},
	}, payload)
}
