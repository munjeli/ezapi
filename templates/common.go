package templates

const (
	// MakePath is where to put the Makefile once it's generated.
	MakePath = `{{ .TargetDir }}/`
	// Makefile is the filename.
	Makefile = "Makefile"
)

type TemplateFile struct {
	Tmpl  string
	Path  string
	Title string
}

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
		ServerTemplate,
		ServerAPIPath,
		ServerTitle,
	},
}

var NetSrvFiles = []TemplateFile{
	{
		NetSrvMakeTemplate,
		MakePath,
		Makefile,
	},
	{
		ServerTemplate,
		ServerNetSrvPath,
		ServerTitle,
	},
	{
		HandlerTemplate,
		ServerNetSrvPath,
		HandlerTitle,
	},
}
