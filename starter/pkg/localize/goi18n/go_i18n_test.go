// Internationalization implementation using nicksnyder/go-i18n
package goi18n

import (
	"io/fs"
	"testing"
	"testing/fstest"

	"github.com/aerogear/charmil/pkg/localize"
	"golang.org/x/text/language"
)

// nolint:funlen
func TestGoi18n_MustLocalize(t *testing.T) {
	type fields struct {
		fs       fs.FS
		language *language.Tag
		format   string
		path     string
	}
	type args struct {
		id          string
		tmplEntries []*localize.TemplateEntry
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      string
		wantErr   bool
		wantPanic bool
	}{
		{
			fields: fields{
				path:   "locales",
				format: "toml",
				fs: fstest.MapFS{
					"locales/en/active.en.toml": {
						Data: []byte(`
						[message-1]
						one = 'message 1'
						`),
					},
				},
			},
			args: args{
				id: "message-1",
			},
			want: "message 1",
		},
		{
			fields: fields{
				path:   "locales",
				format: "toml",
				fs: fstest.MapFS{
					"locales/en/active.en.toml": {
						Data: []byte(`
						[test-case-2]
						one = 'test case {{.Number}}'
						`),
					},
				},
			},
			args: args{
				id: "test-case-2",
				tmplEntries: []*localize.TemplateEntry{
					{
						Key:   "Number",
						Value: 2,
					},
				},
			},
			want: "test case 2",
		},
		{
			fields: fields{
				path:   "locales",
				format: "toml",
				fs: fstest.MapFS{
					"locales/en/active.en.toml": {
						Data: []byte(`
						[test-case-2]
						one = 'test case {{.Number}}'
						`),
					},
				},
			},
			args: args{
				id: "test-case-3",
			},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		// nolint:scopelint
		t.Run(tt.name, func(t *testing.T) {
			cfg := &Config{
				files:    tt.fields.fs,
				language: tt.fields.language,
				format:   tt.fields.format,
				path:     tt.fields.path,
			}
			l, err := New(cfg)
			if tt.wantErr != (err != nil) {
				t.Errorf("Goi18n.New(), wantErr = %v, got %v", tt.wantErr, err)
			}
			defer func() {
				if r := recover(); r == nil && tt.wantPanic {
					t.Errorf("Goi18n.MustLocalize(), expected panic, but code did not panic")
				}
			}()
			if got := l.MustLocalize(tt.args.id, tt.args.tmplEntries...); got != tt.want {
				t.Errorf("Goi18n.MustLocalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
