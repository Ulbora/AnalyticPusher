package pusher

import (
	"context"

	"cloud.google.com/go/bigquery"
)

//MockPusher MockPusher
type MockPusher struct {
	Ctx             context.Context
	Client          *bigquery.Client
	GcpProject      string
	DatasetName     string
	MockPushSuccess bool
}

//GetNew GetNew
func (p *MockPusher) GetNew() AnalyticPusher {
	return p
}

//Push Push
func (p *MockPusher) Push(recs interface{}, table string) bool {
	return p.MockPushSuccess
}

//SetClient SetClient
func (p *MockPusher) SetClient(clt *bigquery.Client) {

}

//SetContext SetContext
func (p *MockPusher) SetContext(ctx context.Context) {

}
