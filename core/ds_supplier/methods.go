package ds_supplier

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

type DrawInput struct {
	Title       string `json:"title,omitempty"`
	XLabel      string `json:"xlabel,omitempty"`
	YLabel      string `json:"ylabel,omitempty"`
	FigSize     []int  `json:"figsize,omitempty"`
	ImageFormat string `json:"image_format,omitempty"`
}

type DrawBarInput struct {
	DrawInput

	Values map[string]float64 `json:"values"`
	Limit  int                `json:"limit,omitempty"`
	Asc    bool               `json:"asc"`
}

func (r *Supplier) DrawBar(ctx context.Context, input *DrawBarInput) ([]byte, error) {
	body, err := r.sendRequest(ctx, "plot/histogram", input)
	if err != nil {
		return nil, errors.Wrap(err, "failed to draw")
	}

	return body, nil
}

type DrawTimeseriesInput struct {
	DrawInput

	Values   map[string]map[time.Time]float64 `json:"values"`
	OnlyTime bool                             `json:"only_time"`
}

func (r *Supplier) DrawTimeseries(ctx context.Context, input *DrawTimeseriesInput) ([]byte, error) {
	body, err := r.sendRequest(ctx, "plot/timeseries", input)
	if err != nil {
		return nil, errors.Wrap(err, "failed to draw")
	}

	return body, nil
}

type DrawLinePlotInput struct {
	DrawInput

	Values [][]any `json:"values"`
}

func (r *Supplier) DrawLinePlot(ctx context.Context, input *DrawLinePlotInput) ([]byte, error) {
	body, err := r.sendRequest(ctx, "plot/lineplot", input)
	if err != nil {
		return nil, errors.Wrap(err, "failed to draw")
	}

	return body, nil
}

type GraphEdge struct {
	First  string  `json:"first"`
	Second string  `json:"second"`
	Weight float64 `json:"weight"`
}

type GraphNode struct {
	Image  []byte  `json:"image,omitempty"`
	Weight float64 `json:"weight,omitempty"`
}

type DrawGraphInput struct {
	DrawInput

	Edges         []GraphEdge          `json:"edges,omitempty"`
	Layout        string               `json:"layout,omitempty"`
	WeightedEdges bool                 `json:"weighted_edges"`
	Nodes         map[string]GraphNode `json:"nodes,omitempty"`
	RootNode      string               `json:"root_node,omitempty"`
}

func (r *Supplier) DrawGraph(ctx context.Context, input *DrawGraphInput) ([]byte, error) {
	body, err := r.sendRequest(ctx, "plot/graph", input)
	if err != nil {
		return nil, errors.Wrap(err, "failed to draw")
	}

	return body, nil
}

func (r *Supplier) DrawGraphAsHeatpmap(ctx context.Context, input *DrawGraphInput) ([]byte, error) {
	body, err := r.sendRequest(ctx, "plot/graph-as-heatmap", input)
	if err != nil {
		return nil, errors.Wrap(err, "failed to draw")
	}

	return body, nil
}
