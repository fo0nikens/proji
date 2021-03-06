package item

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClass(t *testing.T) {
	classExp := &Class{
		Name:      "test",
		Label:     "tst",
		IsDefault: false,
		Folders:   []*Folder{},
		Files:     []*File{},
		Scripts:   []*Script{},
	}

	classAct := NewClass("test", "tst", false)
	assert.Equal(t, classExp, classAct)
}

func TestClassImportFromConfig(t *testing.T) {
	tests := []struct {
		configName string
		class      *Class
		err        error
	}{
		{
			configName: "../../../../assets/examples/example-class-export.toml", class: &Class{
				Name:      "my-example",
				Label:     "mex",
				IsDefault: false,
				Folders: []*Folder{
					{Destination: "src/", Template: ""},
					{Destination: "docs/", Template: ""},
					{Destination: "tests/", Template: ""},
				},
				Files: []*File{
					{Destination: "src/main.py", Template: ""},
					{Destination: "README.md", Template: ""},
				},
				Scripts: []*Script{
					{
						Name:       "init_virtualenv.sh",
						Type:       "post",
						ExecNumber: 1,
						RunAsSudo:  false,
						Args:       []string{},
					},
					{
						Name:       "init_git.sh",
						Type:       "post",
						ExecNumber: 2,
						RunAsSudo:  false,
						Args:       []string{},
					},
				},
			},
			err: nil,
		},
		{
			configName: "example.yaml",
			class: &Class{
				Name:      "",
				Label:     "",
				IsDefault: false,
				Folders:   []*Folder{},
				Files:     []*File{},
				Scripts:   []*Script{},
			},
			err: errors.New(""),
		},
	}

	for _, test := range tests {
		c := NewClass("", "", false)
		err := c.ImportFromConfig(test.configName)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.class, c)
	}
}

func TestClassImportFromDirectory(t *testing.T) {
	tests := []struct {
		baseName string
		folders  []*Folder
		files    []*File
		class    *Class
		err      error
	}{
		{
			baseName: "new-project",
			folders: []*Folder{
				{Destination: "new-project", Template: ""},
				{Destination: "new-project/test", Template: ""},
				{Destination: "new-project/cmd", Template: ""},
				{Destination: "new-project/cmd/base", Template: ""},
				{Destination: "new-project/docs", Template: ""},
			},
			files: []*File{
				{Destination: "new-project/test.txt", Template: ""},
				{Destination: "new-project/README.md", Template: ""},
				{Destination: "new-project/cmd/main.go", Template: ""},
				{Destination: "new-project/test/main_test.go", Template: ""},
			},
			class: &Class{
				Name:      "new-project",
				Label:     "np",
				IsDefault: false,
				Folders: []*Folder{
					{Destination: "cmd", Template: ""},
					{Destination: "cmd/base", Template: ""},
					{Destination: "docs", Template: ""},
					{Destination: "test", Template: ""},
				},
				Files: []*File{
					{Destination: "README.md", Template: ""},
					{Destination: "cmd/main.go", Template: ""},
					{Destination: "test/main_test.go", Template: ""},
					{Destination: "test.txt", Template: ""},
				},
				Scripts: []*Script{},
			},
			err: nil,
		},
	}

	for _, test := range tests {
		for _, dir := range test.folders {
			assert.NoError(t, os.Mkdir(dir.Destination, 0755))
		}
		for _, file := range test.files {
			_, err := os.Create(file.Destination)
			assert.NoError(t, err)
		}

		c := NewClass("", "", false)
		assert.NoError(t, c.ImportFromDirectory(test.baseName, []string{}))
		conf, err := c.Export(".")
		assert.NoError(t, err)
		assert.NoError(t, c.ImportFromConfig(conf))
		assert.Equal(t, test.class, c)

		// Clean up
		_ = os.Remove(conf)
		_ = os.RemoveAll(test.baseName)
	}
}

func TestClassExport(t *testing.T) {
	tests := []struct {
		class      *Class
		configName string
		err        error
	}{
		{
			class: &Class{
				Name:      "example",
				Label:     "exp",
				IsDefault: false,
				Folders: []*Folder{
					{Destination: "exampleFolder/", Template: ""},
					{Destination: "foo/bar/", Template: ""},
				},
				Files: []*File{
					{Destination: "README.md", Template: "README.md"},
					{Destination: "exampleFolder/test.txt", Template: ""},
				},
				Scripts: []*Script{},
			},
			configName: "./proji-example.toml",
			err:        nil,
		},
	}

	for _, test := range tests {
		configName, err := test.class.Export(".")
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.configName, configName)
		assert.FileExists(t, test.configName, "Cannot find the exported config file.")
		_ = os.Remove(configName)
	}
}

func TestClassIsEmpty(t *testing.T) {
	tests := []struct {
		class   *Class
		isEmpty bool
	}{
		{
			class: &Class{
				Name:      "my-example",
				Label:     "mex",
				IsDefault: false,
				Folders: []*Folder{
					{Destination: "src/", Template: ""},
					{Destination: "docs/", Template: ""},
					{Destination: "tests/", Template: ""},
				},
				Files: []*File{
					{Destination: "src/main.py", Template: ""},
					{Destination: "README.md", Template: ""},
				},
				Scripts: []*Script{
					{
						Name:       "init_virtualenv.sh",
						Type:       "post",
						ExecNumber: 1,
						RunAsSudo:  false,
						Args:       []string{},
					},
					{
						Name:       "init_git.sh",
						Type:       "post",
						ExecNumber: 2,
						RunAsSudo:  false,
						Args:       []string{},
					},
				},
			},
			isEmpty: false,
		},
		{
			class: &Class{
				Name:      "blabla",
				Label:     "bl",
				IsDefault: false,
				Folders:   []*Folder{},
				Files:     []*File{},
				Scripts:   []*Script{},
			},
			isEmpty: true,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.class.isEmpty(), test.isEmpty)
	}
}
