package googlenews

import (
	"testing"
	"time"
)

func TestQueryProperty_buildQueryString(t *testing.T) {
	type fields struct {
		After  *time.Time
		Before *time.Time
		Words  []string
		Media  string

		Loc *time.Location
	}
	jpn := time.FixedZone("Asia/Tokyo", 9*60*60)
	firstDay := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	firstDayJPN := time.Date(2023, 1, 1, 0, 0, 0, 0, jpn)
	secondDay := time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "one word",
			fields: fields{
				Words: []string{"soccer"},
			},
			want: "(\"soccer\")",
		},
		{
			name: "two word",
			fields: fields{
				Words: []string{"soccer", "basketball"},
			},
			want: "(\"soccer\" OR \"basketball\")",
		}, {
			name: "after timezone UTC",
			fields: fields{
				After: &firstDay,
			},
			want: "after:2023-01-01",
		}, {
			name: "after timezone Asia/Tokyo",
			fields: fields{
				After: &firstDayJPN,
			},
			want: "after:2023-01-01",
		}, {
			name: "all parameters",
			fields: fields{
				After:  &firstDay,
				Before: &secondDay,
				Words: []string{
					"soccer",
					"basketball",
				},
				Media: "bbc.co",
			},
			want: "(\"soccer\" OR \"basketball\") + inurl:bbc.co + after:2023-01-01 + before:2023-01-02",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qp := &QueryParameter{
				After:  tt.fields.After,
				Before: tt.fields.Before,
				Words:  tt.fields.Words,
				Media:  tt.fields.Media,
			}
			if got := qp.buildQueryString(); got != tt.want {
				t.Errorf("QueryProperty.BuildQueryString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryProperty_validate(t *testing.T) {
	type fields struct {
		After  *time.Time
		Before *time.Time
		Words  []string
		Media  string
		Loc    *time.Location
	}
	firstDay := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	secondDay := time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "normal",
			fields: fields{
				After:  &firstDay,
				Before: &secondDay,
				Words: []string{
					"soccer",
					"basketball",
				},
				Media: "bbc.co",
			},
			wantErr: false,
		}, {
			name: "error(after after before)",
			fields: fields{
				After:  &secondDay,
				Before: &firstDay,
				Words: []string{
					"soccer",
					"basketball",
				},
				Media: "bbc.co",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qp := &QueryParameter{
				After:  tt.fields.After,
				Before: tt.fields.Before,
				Words:  tt.fields.Words,
				Media:  tt.fields.Media,
			}
			if err := qp.validate(); (err != nil) != tt.wantErr {
				t.Errorf("QueryProperty.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
