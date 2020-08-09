package templates

const (
	// MakePath is where to put the Makefile once it's generated.
	MakePath = "%s/Makefile"
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
