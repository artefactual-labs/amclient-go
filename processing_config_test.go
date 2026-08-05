package amclient

import (
	"fmt"
	"net/http"
	"testing"

	"gotest.tools/v3/assert"
)

func TestProcessingConfig_Get(t *testing.T) {
	setup()
	defer teardown()

	const document = `<processingMCP>
  <preconfiguredChoices>
    <preconfiguredChoice>
      <appliesTo>56eebd45-5600-4768-a8c2-ec0114555a3d</appliesTo>
      <goToChain>e9eaef1e-c2e0-4e3b-b942-bfb537162795</goToChain>
    </preconfiguredChoice>
  </preconfiguredChoices>
</processingMCP>`

	mux.HandleFunc("/api/processing-configuration/default/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, document)
	})

	payload, _, err := client.ProcessingConfig.Get(ctx, "default")
	assert.NilError(t, err)
	assert.Equal(t, payload.String(), document)
}

func TestProcessingConfig_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/processing-configuration/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"processing_configurations": ["automated", "default"]}`)
	})

	payload, _, err := client.ProcessingConfig.List(ctx)
	assert.NilError(t, err)
	assert.DeepEqual(t, payload, []string{"automated", "default"})
}
