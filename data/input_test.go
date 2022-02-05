package data

import (
	"reflect"
	"testing"
)

func TestData_Templates(t *testing.T) {
	tests := []struct {
		name string
		file string
		want []string
	}{
		{
			"no templates",
			"../program/testdata/test0.yaml",
			[]string{},
		},
		{
			"no templates",
			"../program/testdata/test2.yaml",
			[]string{"tpl/template1.txt"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, e := New(tt.file)
			if e != nil {
				t.Fatal(e)
			}
			if got := d.Templates(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Templates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    Data
		wantErr bool
	}{
		{
			"basic",
			args{"../program/testdata/test0.yaml"},
			Data{Content: []interface{}{"one", "two", "three"}},
			false,
		},
		{
			"top map",
			args{"../program/testdata/test1.yaml"},
			Data{Content: map[string]interface{}{
				"type": []interface{}{"one", "two", "three"},
			},
			},
			false,
		},
		{
			"with templates",
			args{"../program/testdata/test2.yaml"},
			Data{Content: map[string]interface{}{
				"type":          []interface{}{"one", "two", "three"},
				"__templates__": []interface{}{"tpl/template1.txt"},
			},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}
