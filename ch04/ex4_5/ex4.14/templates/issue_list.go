package templates

import "html/template"

var TemplateOfIssueList = template.Must(
	template.New("issues").Parse(
		``,
	))

var TemplateOfIssue = template.Must(
	template.New("issue").Parse(`
<h1>{{.Title}}</h1>
<dl>
	<dt>user</dt>
	<dd><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></dd>
	<dt>state</dt>
	<dd>{{.State}}</dd>
</dl>
<p>{{.Body}}</p>
`))
