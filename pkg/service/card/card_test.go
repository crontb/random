package card

import (
	"testing"
)

func TestGenerateNumbers(t *testing.T) {
	type args struct {
		totalNumbers int
		maxNumber    int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantLen int
		wantMin int
		wantMax int
		wantErr bool
	}{
		{
			name: "Case 1",
			args: args{
				totalNumbers: 20,
				maxNumber:    99,
			},
			wantLen: 20,
			wantMin: 0,
			wantMax: 99,
		},
		{
			name: "Case 2",
			args: args{
				totalNumbers: 200,
				maxNumber:    10000,
			},
			wantLen: 200,
			wantMin: 0,
			wantMax: 10000,
		},
		{
			name: "Case 3",
			args: args{
				totalNumbers: 500,
				maxNumber:    100,
			},
			wantLen: 500,
			wantMin: 0,
			wantMax: 100,
		},
		{
			name: "Case 4",
			args: args{
				totalNumbers: 0,
				maxNumber:    1000,
			},
			wantErr: true,
		},
		{
			name: "Case 5",
			args: args{
				totalNumbers: -9,
				maxNumber:    10,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateNumbers(tt.args.totalNumbers, tt.args.maxNumber)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GenerateNumbers() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantLen > 0 && len(got) != tt.wantLen {
				t.Errorf("GenerateNumbers() = %v, wantLen %v", got, tt.wantLen)
			}
			for _, item := range got {
				if item < tt.wantMin || item > tt.wantMax {
					t.Errorf("GenerateNumbers() = %v, item %v, wantRange %v-%v", got, item, tt.wantMin, tt.wantMax)
				}
			}
		})
	}
}

func BenchmarkGenerateNumbers(b *testing.B) {
	type args struct {
		totalNumbers int
		maxNumber    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Benchmark 1",
			args: args{
				totalNumbers: 20,
				maxNumber:    99,
			},
		},
		{
			name: "Benchmark 2",
			args: args{
				totalNumbers: 100,
				maxNumber:    999,
			},
		},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				GenerateNumbers(tt.args.totalNumbers, tt.args.maxNumber)
			}
		})
	}
}
