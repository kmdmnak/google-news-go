package googlenews

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/mmcdole/gofeed/rss"
)

func Test_addLanguageParams(t *testing.T) {
	type args struct {
		val  url.Values
		lang *langProperty
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addLanguageParams(tt.args.val, tt.args.lang)
		})
	}
}

func Test_client_newRequest(t *testing.T) {
	type fields struct {
		base    string
		lang    Language
		hclient *http.Client
	}
	type args struct {
		ctx context.Context
		url string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Request
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				base:    tt.fields.base,
				lang:    tt.fields.lang,
				hclient: tt.fields.hclient,
			}
			got, err := c.newRequest(tt.args.ctx, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.newRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.newRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_do(t *testing.T) {
	type fields struct {
		base    string
		lang    Language
		hclient *http.Client
	}
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *rss.Feed
		want1   *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				base:    tt.fields.base,
				lang:    tt.fields.lang,
				hclient: tt.fields.hclient,
			}
			got, got1, err := c.do(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.do() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Client.do() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_client_SearchByTopic(t *testing.T) {
	type fields struct {
		lang Language
	}
	type args struct {
		ctx   context.Context
		topic NewsTopic
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *rss.Feed
		want1   *http.Response
		wantErr bool
	}{
		{
			name: "normal",
			fields: fields{
				lang: JPN,
			},
			args: args{
				ctx:   context.TODO(),
				topic: TOPIC_WORLD,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(tt.fields.lang)
			got, _, err := c.SearchByTopic(tt.args.ctx, tt.args.topic)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.SearchByTopic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got.Items) > 0 {
				fmt.Println(got.Items[0])
			}
		})
	}
}

func Test_client_SearchByQuery(t *testing.T) {
	type fields struct {
		base    string
		lang    Language
		hclient *http.Client
	}
	type args struct {
		ctx    context.Context
		params *QueryParameter
	}
	firstDay := time.Date(2022, 12, 10, 10, 0, 0, 0, time.UTC)
	secondDay := time.Date(2022, 12, 12, 10, 0, 0, 0, time.UTC)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *rss.Feed
		want1   *http.Response
		wantErr bool
	}{
		{
			name: "normal",
			fields: fields{
				lang: JPN,
			},
			args: args{
				ctx: context.TODO(),
				params: &QueryParameter{
					After:  &firstDay,
					Before: &secondDay,
					Words:  []string{"soccer"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(USEN)
			_, _, err := c.SearchByQuery(tt.args.ctx, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.SearchByQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_client_SearchByGeometry(t *testing.T) {
	type fields struct {
		base    string
		lang    Language
		hclient *http.Client
	}
	type args struct {
		ctx  context.Context
		word string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *rss.Feed
		want1   *http.Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				base:    tt.fields.base,
				lang:    tt.fields.lang,
				hclient: tt.fields.hclient,
			}
			got, got1, err := c.SearchByGeometry(tt.args.ctx, tt.args.word)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.SearchByGeometry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.SearchByGeometry() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Client.SearchByGeometry() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		lang Language
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.lang); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
