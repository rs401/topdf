package main

import "testing"

func Test_convtopdf(t *testing.T) {
	type args struct {
		src string
		dst string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Test convtopdf", args{"downloaded", "out.pdf"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := convtopdf(tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("convtopdf() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
