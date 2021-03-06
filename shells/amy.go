package shells

import (
	"text/template"

	"github.com/pkg/errors"
)

var Amy *template.Template

func init() {
	text := `<!DOCTYPE html>
<!--[if lte IE 6]><html class="preIE7 preIE8 preIE9"><![endif]-->
<!--[if IE 7]><html class="preIE8 preIE9"><![endif]-->
<!--[if IE 8]><html class="preIE9"><![endif]-->
<!--[if gte IE 9]><!--><html><!--<![endif]-->
<head>
<meta charset="UTF-8" />
<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1" />
<meta name="viewport" content="width=device-width,initial-scale=1" />
{{if .PageTitle}}<title>{{.PageTitle}}</title>
{{end}}{{if .Author}}<meta name="author" content="{{.Author}}" />
{{end}}{{if .Description}}<meta name="description" content="{{.Description}}" />
{{end}}{{if .Keywords}}<meta name="keywords" content="{{.Keywords}}" />
{{end}}{{if .FavIcon}}<link rel="shortcut icon" href="{{.FavIcon}}" type="image/vnd.microsoft.icon" />
{{end}}{{if .GoogleFonts}}<link rel="stylesheet" href="//fonts.googleapis.com/css?family={{.GoogleFonts}}" type="text/css" />
{{end}}<link rel="stylesheet" href="//rawgit.com/suite911/cloud911/master/assets/css/cloud911.css" type="text/css" />
<link rel="stylesheet" href="/1.css" type="text/css" />
{{if .InlineCSS}}
<style type="text/css"><!-- /*<![CDATA[*/
{{.InlineCSS}}
/*]]>*/ --></style>
{{end}}
{{if .Head}}
{{.Head}}
{{end}}
<script type="text/javascript" src="//rawgit.com/suite911/cloud911/master/assets/js/early.js"></script>
{{if .ReCaptchaV3}}<script src='https://www.google.com/recaptcha/api.js?render={{.ReCaptchaV3}}'></script>
{{end}}</head>
<body>{{.BodyHead}}{{if .Body}}{{.Body}}{{else}}
<input type="checkbox" class="lights-off" id="lights-off" checked />
<div class="page-outer"><div class="page-inner">
	<header class="topnav">
		<div class="topnav">{{.TopNavHead}}
			<div class="topnavleft"{{range $k, $v := .TopNav}}
				><a href="{{$k}}"><span class="topnav">{{$v}}</span></a{{end}}
			></div>
			<div class="topnavright card"><label for="lights-off" class="lights-off"><span class="only-lights-on">Lights off &#x263d;</span><span class="only-lights-off">Lights on &#x263c;</span></label></div>
			<div class="topnavhack"></div>{{.TopNavTail}}
		</div>
	</header>
	<header class="header">
		<div class="header">{{.HeaderHead}}{{.Header}}{{.HeaderTail}}
		</div>
	</header>{{if .ContentTitle}}
	<div id="title">
		<h1>{{.ContentTitle}}</h1>{{if .ContentSubTitle}}
		<p>{{.ContentSubTitle}}</p>{{end}}
		<hr />
	</div>{{end}}
	<div id="content"{{if .NoScript}} class="display-none"{{end}}>
{{.ContentHead}}
{{if .Form}}
		<div class="form">
			<form id="form" action="{{.FormAction}}" method="POST">
{{if .ReCaptchaV3}}
				<input type="hidden" id="captcha-onload" name="captcha-onload" value="" />
				<input type="hidden" id="captcha-onchange" name="captcha-onchange" value="" />
				<input type="hidden" id="captcha-onsubmit" name="captcha-onsubmit" value="" />
{{end}}
				<input type="hidden" id="analytics-lights" />
{{end}}
{{.Content}}
{{if .Form}}
			<br />
			</form>
		</div>
{{end}}
{{.ContentTail}}
	</div>
{{if .NoScript}}<noscript>{{.NoScript}}</noscript>{{end}}
</div></div>
<footer class="footer">{{.FooterHead}}{{.Footer}}{{.FooterTail}}
	<div class="copyright">{{.Copyright}}</div>
</footer>
{{end}}
{{.BodyTail}}
{{if .InlineJavaScript}}<script type="text/javascript"><!-- //<![CDATA[
{{.InlineJavaScript}}
//]]> --></script>
{{end}}<script type="text/javascript" src='/0.js'></script>
<script type="text/javascript" src='/1.js'></script>
<script type="text/javascript" src="//rawgit.com/suite911/cloud911/master/assets/js/late.js"></script>
</body>
</html>
`
	var err error
	if Amy, err = template.New("Amy").Option("missingkey=zero").Parse(text); err != nil {
		panic(errors.Wrap(err, `template.New("Amy").Parse(text)`))
	}
}
