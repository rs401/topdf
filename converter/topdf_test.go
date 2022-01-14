package converter

import (
	"log"
	"os"
	"testing"
)

func Test_convtopdf(t *testing.T) {
	input, err := os.CreateTemp("./", "input")
	if err != nil {
		log.Println("Error Creating the Temp Input File")
		log.Println(err)
		return
	}
	defer input.Close()
	defer os.Remove(input.Name())

	output, err := os.CreateTemp("./", "output-*.pdf")
	if err != nil {
		log.Println("Error Creating the Temp Output File")
		log.Println(err)
		return
	}
	defer output.Close()
	defer os.Remove(output.Name())
	type args struct {
		src string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Test convtopdf", args{input.Name()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if fname, err := Convtopdf(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("convtopdf() error = %v, wantErr %v, fname = %s", err.Error(), tt.wantErr, fname)
			}
		})
	}
}
