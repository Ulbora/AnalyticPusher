package pusher

import (
	"context"
	"log"

	"cloud.google.com/go/bigquery"
)

//AnalyticPusher AnalyticPusher
type AnalyticPusher interface {
	Push(recs interface{}, table string) bool
	SetClient(clt *bigquery.Client)
	SetContext(ctx context.Context)
}

//Pusher Pusher
type Pusher struct {
	Ctx         context.Context
	Client      *bigquery.Client
	GcpProject  string
	DatasetName string
}

//GetNew GetNew
func (p *Pusher) GetNew() AnalyticPusher {
	return p
}

//Push Push
func (p *Pusher) Push(recs interface{}, table string) bool {
	log.Println("in push--------:")
	log.Println("p in push--------:", p)
	var rtn bool
	u := p.Client.Dataset(p.DatasetName).Table(table).Inserter()
	err := u.Put(p.Ctx, recs)
	if err != nil {
		log.Println("big query put err:", err)
	} else {
		rtn = true
	}
	return rtn
}

//SetClient SetClient
func (p *Pusher) SetClient(clt *bigquery.Client) {
	p.Client = clt
}

//SetContext SetContext
func (p *Pusher) SetContext(ctx context.Context) {
	p.Ctx = ctx
}

//go mod init github.com/Ulbora/AnalyticPusher
