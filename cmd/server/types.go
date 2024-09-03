package server

type BasePage struct {
	PageTitle   string
	SiteTitle   string
	Description string
	Keywords    string
	Author      string
}

type ExampleType struct {
	Test string
	Base BasePage
}
