package plugin

import (
    "context"
    "encoding/json"
	"net/http"
    //"fmt"
    //"math/rand"
    //"time"


    "github.com/grafana/grafana-plugin-sdk-go/backend"
    "github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
    //"github.com/grafana/grafana-plugin-sdk-go/data"
)
// Make sure Datasource implements required interfaces. This is important to do
// since otherwise we will only get a not implemented error response from plugin in
// runtime. In this example datasource instance implements backend.QueryDataHandler,
// backend.CheckHealthHandler interfaces. Plugin should not implement all these
// interfaces- only those which are required for a particular task.
var (
    _ backend.QueryDataHandler      = (*Datasource)(nil)
    _ backend.CheckHealthHandler    = (*Datasource)(nil)
    _ instancemgmt.InstanceDisposer = (*Datasource)(nil)
)
// NewDatasource creates a new datasource instance.
func NewDatasource(_ backend.DataSourceInstanceSettings) (instancemgmt.Instance, error) {
    return &Datasource{}, nil
}
// Datasource is an example datasource which can respond to data queries, reports
// its health and has streaming skills.
type Datasource struct{}
// Dispose here tells plugin SDK that plugin wants to clean up resources when a new instance
// created. As soon as datasource settings change detected by SDK old datasource instance will
// be disposed and a new one will be created using NewSampleDatasource factory function.
func (d *Datasource) Dispose() {
    // Clean up datasource instance resources.
}
// QueryData handles multiple queries and returns multiple responses.
// req contains the queries []DataQuery (where each query contains RefID as a unique identifier).
// The QueryDataResponse contains a map of RefID to the response for each query, and each response
// contains Frames ([]*Frame).
func (d *Datasource) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
    // create response struct
    response := backend.NewQueryDataResponse()
    // loop over queries and execute them individually.
    for _, q := range req.Queries {
        res := d.query(ctx, req.PluginContext, q)
        // save the response in a hashmap
        // based on with RefID as identifier
        response.Responses[q.RefID] = res
    }
    return response, nil
}
func (d *Datasource) CheckHealth(_ context.Context, req *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
    // Your code to check metric goes here
    // ...
    return &backend.CheckHealthResult{
        Status:  backend.HealthStatusOk,
        Message: "Metric is healthy",
    }, nil

}

type PuppetDBResult struct {
    Name   string `json:"certname"`
    Status string `json:"status"`
}
