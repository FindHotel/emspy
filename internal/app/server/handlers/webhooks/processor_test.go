package webhooks

import (
	"context"
	"reflect"
	"testing"

	"github.com/FindHotel/emspy/internal/app/store"
	"github.com/FindHotel/emspy/internal/app/store/memory"
	"github.com/stretchr/testify/require"
)

func TestNewProcessor(t *testing.T) {
	memStrore, err := memory.New()
	require.NoError(t, err)

	type args struct {
		source string
		stores []store.Store
	}
	tests := []struct {
		name string
		args args
		want *Processor
	}{
		{
			name: "smth",
			args: args{
				source: "test source",
				stores: []store.Store{memStrore},
			},
			want: &Processor{
				source: "test source",
				stores: []store.Store{memStrore},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProcessor(tt.args.stores, tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProcessor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessor_Capture(t *testing.T) {
	memStrore, err := memory.New()
	require.NoError(t, err)

	type fields struct {
		stores []store.Store
		source string
	}
	type args struct {
		input []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test capture",
			fields: fields{
				stores: []store.Store{memStrore},
				source: "test",
			},
			args:    args{input: []byte("test input")},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Processor{
				stores: tt.fields.stores,
				source: tt.fields.source,
			}
			if err := p.Capture(context.Background(), tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("Processor.Capture() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
