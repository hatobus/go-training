package templates

import "html/template"

var TemplateOfIssueList = template.Must(
	template.New("issues").Parse(`
<h1>{{.Issues | len}} issues</h1>
<table>
<tr style='text-align: left'>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th>
</tr>
{{range .Issues}}
<tr>
	<td><a href='{{.GenIssueDetailURL}}'>{{.Number}}</td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.GenIssueDetailURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

var TemplateOfIssue = template.Must(
	template.New("issue").Parse(`
<h1>{{.Title}}</h1>
<dl>
	<dt>user</dt>
	<dd><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></dd>
	<dt>state</dt>
	<dd>{{.State}}</dd>
</dl>
<br>
<h3>Body</h3>
<p>{{.Body}}</p>
`))
