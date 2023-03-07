package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseJSON(t *testing.T) {
	type args struct {
		json string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr error
	}{
		//	{
		//		name: "parse null value",
		//		args: args{json: "null"},
		//	},
		{
			name: "parse true value",
			args: args{json: "true"},
			want: true,
		},
		{
			name: "parse true value",
			args: args{json: "false"},
			want: false,
		},
		{
			name: "parse numeric value",
			args: args{json: "11.5"},
			want: 11.5,
		},
		{
			name: "parse string value",
			args: args{json: `"This is a string"`},
			want: "This is a string",
		},
		{
			name: "parse array value",
			args: args{json: `[ "This is a string" , "and another" ]`},
			want: []interface{}{"This is a string", "and another"},
		},
		{
			name: "parse obj value",
			args: args{json: `{ "some field" : true }`},
			want: map[string]interface{}{"some field": true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJSON(tt.args.json)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
