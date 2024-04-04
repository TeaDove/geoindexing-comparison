package ds_supplier

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"geoindexing_comparison/core/utils"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type Supplier struct {
	client   *http.Client
	basePath string
}

func New() (*Supplier, error) {
	r := Supplier{}
	r.client = http.DefaultClient
	r.basePath = "http://localhost:8000"

	return &r, nil
}

func (r *Supplier) sendRequest(ctx context.Context, path string, input any) ([]byte, error) {
	t0 := time.Now()

	reqBody, err := json.Marshal(&input)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal request body")
	}

	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		fmt.Sprintf("%s/%s", r.basePath, path),
		bytes.NewReader(reqBody),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request")
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to do request")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read body request")
	}

	utils.CloseOrLog(resp.Body)

	log.Info().
		Str("status", "ds.request.done").
		Dur("elapsed", time.Since(t0)).
		Str("path", path).Send()

	if resp.StatusCode >= 400 {
		return nil, errors.Errorf("bad status code: %d, content: %s", resp.StatusCode, string(body))
	}

	return body, nil
}
