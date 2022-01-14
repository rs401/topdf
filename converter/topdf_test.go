package converter

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"
)

func TestConvtopdf(t *testing.T) {
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "could not connect to Docker")

	resource, err := pool.Run("rs401/topdf", "1.1", []string{})
	require.NoError(t, err, "could not start container")

	t.Cleanup(func() {
		require.NoError(t, pool.Purge(resource), "failed to remove container")
	})

	input, err := os.CreateTemp("./", "input")
	require.NoError(t, err, "could not create temp input file")
	defer input.Close()
	defer os.Remove(input.Name())

	var resp *http.Response
	var respPost *http.Response

	err = pool.Retry(func() error {
		resp, err = http.Get(fmt.Sprint("http://localhost:", resource.GetPort("8888/tcp"), "/topdf"))
		if err != nil {
			t.Log("container not ready, waiting...")
			return err
		}
		return nil
	})
	require.NoError(t, err, "HTTP error")
	defer resp.Body.Close()

	err = pool.Retry(func() error {
		respPost, err = http.Post(fmt.Sprint("http://localhost:", resource.GetPort("8888/tcp"), "/topdf"), "text/plain", input)
		if err != nil {
			t.Log("container not ready, waiting...")
			return err
		}
		return nil
	})
	require.NoError(t, err, "HTTP error")
	defer respPost.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode, "HTTP status code")
	require.Equal(t, http.StatusOK, respPost.StatusCode, "HTTP status code")

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "failed to read HTTP body")

	require.Contains(t, string(body), "Select file to convert:", "does not contain template")

	// type args struct {
	// 	src string
	// }
	// tests := []struct {
	// 	name    string
	// 	args    args
	// 	want    string
	// 	wantErr bool
	// }{
	// 	{
	// 		name:    "Test with temp file",
	// 		args:    args{src: input.Name()},
	// 		want:    input.Name() + "pdf",
	// 		wantErr: false,
	// 	},
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		got, err := Convtopdf(tt.args.src)
	// 		if (err != nil) != tt.wantErr {
	// 			t.Errorf("Convtopdf() error = %v, wantErr %v", err, tt.wantErr)
	// 			return
	// 		}
	// 		if got != tt.want {
	// 			t.Errorf("Convtopdf() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}
