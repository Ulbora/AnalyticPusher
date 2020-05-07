package pusher

import (
	"context"
	"fmt"
	"testing"
	"time"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/option"
)

type TestRecord struct {
	Name    string    `bigquery:"name"`
	Entered time.Time `bigquery:"entered"`
	Message string    `bigquery:"message"`
}

func TestPusher_Push(t *testing.T) {
	var p Pusher
	p.GcpProject = "august-gantry-192521"
	p.DatasetName = "ulboralabs"
	//fr.TableName = "test1234"
	ctx := context.Background()
	p.SetContext(ctx)
	client, err := bigquery.NewClient(ctx, p.GcpProject, option.WithCredentialsFile("../gcpCreds.json"))
	if err != nil {
		fmt.Println("bq err: ", err)
	} else {
		p.SetClient(client)
		var testObj TestRecord
		testObj.Entered = time.Now()
		testObj.Message = "this is a test"
		testObj.Name = "Some Tester"
		ap := p.GetNew()
		suc := ap.Push(testObj, "analytic_test")
		if !suc {
			t.Fail()
		}
	}
}
