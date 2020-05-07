package pusher

import (
	"context"
	"fmt"
	"testing"
	"time"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/option"
)

func TestMockPusher_Push(t *testing.T) {
	var p MockPusher
	p.MockPushSuccess = true
	p.GcpProject = "august-gantry-192521"
	p.DatasetName = "ulboralabs"
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
