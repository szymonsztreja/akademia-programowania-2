package main

import "testing"

func TestGreeter_What(t *testing.T) {
	type fields struct {
		Doer Doer
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "foo",
			fields: fields{
				Doer: &DoerMock{Ret: "World"},
			},
			want: "World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Greeter{
				Doer: tt.fields.Doer,
			}
			if got := g.What(); got != tt.want {
				t.Errorf("Greeter.What() = %v, want %v", got, tt.want)
			}
		})
	}
}
