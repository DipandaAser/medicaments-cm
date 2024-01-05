package scrapper

import (
	"testing"
)

func TestSearchMedicaments(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name      string
		args      args
		wantEmpty bool
		wantErr   bool
	}{
		{
			name:      "Search medicaments with empty name",
			args:      args{name: ""},
			wantEmpty: true,
			wantErr:   false,
		},
		{
			name:      "Search medicaments with name",
			args:      args{name: "doliprane"},
			wantEmpty: false,
			wantErr:   false,
		},
		{
			name:      "Search medicaments with a letter",
			args:      args{name: "d"},
			wantEmpty: false,
			wantErr:   false,
		},
		{
			name:      "Search medicaments with wildcard name",
			args:      args{name: "dolipr*"},
			wantEmpty: true,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results, err := SearchMedicaments(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchMedicaments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (tt.wantEmpty && len(results) > 0) || (!tt.wantEmpty && len(results) == 0) {
				t.Errorf("SearchMedicaments() got = %v, want-empty = %v", results, tt.wantEmpty)
				return
			}
		})
	}
}
