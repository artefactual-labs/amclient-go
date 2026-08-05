package amclient

import (
	"fmt"
	"net/http"
	"testing"

	"gotest.tools/v3/assert"
)

func TestTransfer_Start(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/transfer/start_transfer/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"message": "Copy successful", "path": "/var/foobar"}`)
	})

	payload, _, err := client.Transfer.Start(ctx, &TransferStartRequest{
		Name:  "foobar",
		Paths: []string{"a.jpg", "b.jpg"},
		Type:  "standard",
	})
	assert.NilError(t, err)
	assert.Equal(t, payload.Message, "Copy successful")
}

func TestTransfer_Approve(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/transfer/approve/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		assert.NilError(t, r.ParseForm())
		assert.Equal(t, r.Form.Get("directory"), "Foobar")
		fmt.Fprint(w, `{"message": "Approval successful.", "uuid": "eaedbee3-2b02-4e40-baa0-3ef92c5fd17e"}`)
	})

	req := &TransferApproveRequest{
		Directory: "/var/archivematica/Foobar",
		Type:      "standard",
	}
	payload, _, err := client.Transfer.Approve(ctx, req)
	assert.NilError(t, err)
	assert.Equal(t, req.Directory, "/var/archivematica/Foobar")
	assert.Equal(t, payload.Message, "Approval successful.")
	assert.Equal(t, payload.UUID, "eaedbee3-2b02-4e40-baa0-3ef92c5fd17e")
}

func TestTransfer_Unapproved(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/transfer/unapproved/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
			"message": "Fetched unapproved transfers successfully.",
			"results": [
				{
					"type": "standard",
					"directory": "/var/foobar1",
					"uuid": "eaedbee3-2b02-4e40-baa0-3ef92c5fd17e"
				},
				{
					"type": "standard",
					"directory": "/var/foobar2",
					"uuid": "433f20e4-a0e4-484b-8fb4-ec9b3cda4cfc"
				}
			]
		}`)
	})

	payload, _, err := client.Transfer.Unapproved(ctx, &TransferUnapprovedRequest{})
	assert.NilError(t, err)
	assert.Equal(t, len(payload.Results), 2)
}

func TestTransfer_Status(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/transfer/status/52dd0c01-e803-423a-be5f-b592b5d5d61c/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
	"status": "COMPLETE",
	"name": "imgs",
	"sip_uuid": "41699e73-ec9e-4240-b153-71f4155e7da4",
	"microservice": "Microservice group name",
	"directory": "imgs-52dd0c01-e803-423a-be5f-b592b5d5d61c",
	"path": "/var/archivematica/sharedDirectory/watchedDirectories/SIPCreation/completedTransfers/imgs-52dd0c01-e803-423a-be5f-b592b5d5d61c/",
	"message": "Fetched status for 52dd0c01-e803-423a-be5f-b592b5d5d61c successfully.",
	"type": "transfer",
	"uuid": "52dd0c01-e803-423a-be5f-b592b5d5d61c"
}`)
	})

	payload, _, err := client.Transfer.Status(ctx, "52dd0c01-e803-423a-be5f-b592b5d5d61c")
	assert.NilError(t, err)
	assert.DeepEqual(t, &TransferStatusResponse{
		ID:           "52dd0c01-e803-423a-be5f-b592b5d5d61c",
		Status:       "COMPLETE",
		Name:         "imgs",
		SIPID:        "41699e73-ec9e-4240-b153-71f4155e7da4",
		Microservice: "Microservice group name",
		Directory:    "imgs-52dd0c01-e803-423a-be5f-b592b5d5d61c",
		Path:         "/var/archivematica/sharedDirectory/watchedDirectories/SIPCreation/completedTransfers/imgs-52dd0c01-e803-423a-be5f-b592b5d5d61c/",
		Message:      "Fetched status for 52dd0c01-e803-423a-be5f-b592b5d5d61c successfully.",
		Type:         "transfer",
	}, payload)
}

func TestTransfer_Hide(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/transfer/52dd0c01-e803-423a-be5f-b592b5d5d61c/delete/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		fmt.Fprint(w, `{
	"removed": true
}`)
	})

	payload, _, err := client.Transfer.Hide(ctx, "52dd0c01-e803-423a-be5f-b592b5d5d61c")
	assert.NilError(t, err)
	assert.DeepEqual(t, &TransferHideResponse{
		Removed: true,
	}, payload)
}
