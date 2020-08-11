package templates

const (
	// MakePath is where to put the Makefile once it's generated.
	MakePath = `{{ .TargetDir }}/`
	// Makefile is the filename.
	Makefile = "Makefile"
)

// TemplateFile holds the metadata for the files
// that are generated from templates.
type TemplateFile struct {
	Tmpl  string
	Path  string
	Title string
}

// APIFiles is a list of templated files
// for building a CRUD API.
var APIFiles = []TemplateFile{
	{
		APIMakeTemplate,
		MakePath,
		Makefile,
	},
	{
		APICRUDTemplate,
		APIPath,
		APITitle,
	},
	{
		APITestCRUDTemplate,
		APIPath,
		APITestFileName,
	},
	{
		APIServerTemplate,
		ServerAPIPath,
		ServerTitle,
	},
}

// NetSrvFiles are the templated files for generating
// a basic networked service.
var NetSrvFiles = []TemplateFile{
	{
		NetSrvMakeTemplate,
		MakePath,
		Makefile,
	},
	{
		NetSrvServerTemplate,
		ServerNetSrvPath,
		ServerTitle,
	},
	{
		HandlerTemplate,
		ServerNetSrvPath,
		HandlerTitle,
	},
}
