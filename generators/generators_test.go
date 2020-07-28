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
				targetDir: "kittenapi",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GenerateAPI(tt.args.name, tt.args.targetDir); (err != nil) != tt.wantErr {
				t.Errorf("GenerateAPI() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		if tt.args.targetDir != "" {
			t.Cleanup( func(){
				err := os.RemoveAll(tt.args.targetDir)
				if err != nil {
					t.Error(err.Error())
				}
			})
		}
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
				targetDir: "kittenapi",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GenerateNetworkedService(tt.args.name, tt.args.targetDir); (err != nil) != tt.wantErr {
				t.Errorf("GenerateNetworkedService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		if tt.args.targetDir != "" {
			t.Cleanup( func(){
				err := os.RemoveAll(tt.args.targetDir)
				if err != nil {
					t.Error(err.Error())
				}
			})
		}
	}
}

func Test_makeAPIDirs(t *testing.T) {
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
				targetDir: "kittennetsrv",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := makeAPIDirs(tt.args.name, tt.args.targetDir); (err != nil) != tt.wantErr {
				t.Errorf("makeAPIDirs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		if tt.args.targetDir != "" {
			t.Cleanup( func(){
				err := os.RemoveAll(tt.args.targetDir)
				if err != nil {
					t.Error(err.Error())
				}
			})
		}
	}
}

func Test_makeNetSrvDirs(t *testing.T) {
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
				targetDir: "kittennetsrv",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := makeNetSrvDirs(tt.args.name, tt.args.targetDir); (err != nil) != tt.wantErr {
				t.Errorf("makeNetSrvDirs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		if tt.args.targetDir != "" {
			t.Cleanup( func(){
				err := os.RemoveAll(tt.args.targetDir)
				if err != nil {
					t.Error(err.Error())
				}
			})
		}
	}
}
