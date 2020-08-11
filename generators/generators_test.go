package generators

import (
	"os"
	"testing"
)

func TestGenerateAPI(t *testing.T) {
	type args struct {
		name      string
		targetDir string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty args",
			args: args{
				name:      "",
				targetDir: "",
			},
			wantErr: true,
		},
		{
			name: "valid args",
			args: args{
				name:      "kitten",
				targetDir: "kittendir",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GenerateAPI(tt.args.name, tt.args.targetDir); (err != nil) != tt.wantErr {
				t.Errorf("GenerateAPI() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.args.targetDir != "" {
				t.Cleanup(func() {
					err := os.RemoveAll(tt.args.targetDir)
					if err != nil {
						t.Error(err.Error())
					}
				})
			}
		})
	}
}

func TestGenerateNetworkedService(t *testing.T) {
	type args struct {
		name      string
		targetDir string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty args",
			args: args{
				name:      "",
				targetDir: "",
			},
			wantErr: true,
		},
		{
			name: "valid args",
			args: args{
				name:      "kitten",
				targetDir: "kittensvc",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GenerateNetworkedService(tt.args.name, tt.args.targetDir); (err != nil) != tt.wantErr {
				t.Errorf("GenerateNetworkedService() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.args.targetDir != "" {
				t.Cleanup(func() {
					err := os.RemoveAll(tt.args.targetDir)
					if err != nil {
						t.Error(err.Error())
					}
				})
			}
		})
	}
}

func Test_makeDirs(t *testing.T) {
	type args struct {
		name      string
		targetDir string
		apiType   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty args",
			args: args{
				name:      "",
				targetDir: "",
				apiType:   "",
			},
			wantErr: true,
		},
		{
			name: "valid args for api",
			args: args{
				name:      "kitten",
				targetDir: "animals",
				apiType:   "api",
			},
			wantErr: false,
		},
		{
			name: "valid args for netsrv",
			args: args{
				name:      "kitten",
				targetDir: "netanimals",
				apiType:   "netsrv",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := makeDirs(tt.args.name, tt.args.targetDir, tt.args.apiType); (err != nil) != tt.wantErr {
				t.Errorf("makeAPIDirs() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.args.targetDir != "" {
				t.Cleanup(func() {
					err := os.RemoveAll(tt.args.targetDir)
					if err != nil {
						t.Error(err.Error())
					}
				})
			}
		})
	}
}

func Test_generateFilesFromTemplates(t *testing.T) {
	type args struct {
		apiType   string
		name      string
		targetDir string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty apiType",
			args: args{
				apiType:   "",
				name:      "kitten",
				targetDir: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := generateFilesFromTemplates(tt.args.apiType, tt.args.name, tt.args.targetDir); (err != nil) != tt.wantErr {
				t.Errorf("generateFilesFromTemplates() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
