// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: templates/metrics.gohtml

package templates

import (
	db "github.com/FlowingSPDG/get5-web-go/src/db"
	"github.com/FlowingSPDG/get5-web-go/templates/layout"
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strconv"
	"strings"
)

// Metrics generates templates/metrics.gohtml
func Metrics(u *db.MetricsDataPage) string {
	var _b strings.Builder
	RenderMetrics(&_b, u)
	return _b.String()
}

// RenderMetrics render templates/metrics.gohtml
func RenderMetrics(_buffer io.StringWriter, u *db.MetricsDataPage) {

	_body := func(_buffer io.StringWriter) {

	}

	_menu := func(_buffer io.StringWriter) {
		if u.LoggedIn == true {

			_buffer.WriteString("<li><a id=\"mymatches\" href=\"/mymatches\">My Matches</a></li>")

			_buffer.WriteString("<li><a id=\"match_create\" href=\"/match/create\">Create a Match</a></li>")

			_buffer.WriteString("<li><a id=\"myteams\" href=\"/myteams\">My Teams</a></li>")

			_buffer.WriteString("<li><a id=\"team_create\" href=\"/team/create\">Create a Team</a></li>")

			_buffer.WriteString("<li><a id=\"myservers\" href=\"/myservers\">My Servers</a></li>")

			_buffer.WriteString("<li><a id=\"server_create\" href=\"/server/create\">Add a Server</a></li>")

			_buffer.WriteString("<li><a href=\"/logout\">Logout</a></li>")

		} else {

			_buffer.WriteString("<li><a href=\"/login\"><img src=\"/static/img/login_small.png\" height=\"18\"/></a></li>")

		}
	}

	_content := func(_buffer io.StringWriter) {

		_buffer.WriteString("<div id=\"content\">\n  <h1>Get5 Metrics</h1>\n  <ul class=\"list-group\">\n    <li class=\"list-group-item\"> Registered users: ")
		_buffer.WriteString(gorazor.HTMLEscStr(strconv.Itoa(u.Data.RegisteredUsers)))
		_buffer.WriteString(" </li>\n    <li class=\"list-group-item\"> Saved teams: ")
		_buffer.WriteString(gorazor.HTMLEscStr(strconv.Itoa(u.Data.SavedTeams)))
		_buffer.WriteString("</li>\n    <li class=\"list-group-item\"> Matches created: ")
		_buffer.WriteString(gorazor.HTMLEscStr(strconv.Itoa(u.Data.MatchesCreated)))
		_buffer.WriteString("</li>\n    <li class=\"list-group-item\"> Completed matches: ")
		_buffer.WriteString(gorazor.HTMLEscStr(strconv.Itoa(u.Data.CompletedMatches)))
		_buffer.WriteString("</li>\n    <li class=\"list-group-item\"> Servers added: ")
		_buffer.WriteString(gorazor.HTMLEscStr(strconv.Itoa(u.Data.ServersAdded)))
		_buffer.WriteString("</li>\n    <li class=\"list-group-item\"> Maps with stats saved: ")
		_buffer.WriteString(gorazor.HTMLEscStr(strconv.Itoa(u.Data.MapsWithStatsSaved)))
		_buffer.WriteString("</li>\n    <li class=\"list-group-item\"> Unique players: ")
		_buffer.WriteString(gorazor.HTMLEscStr(strconv.Itoa(u.Data.UniquePlayers)))
		_buffer.WriteString("</li>\n  </ul>\n</div>")

	}

	layout.RenderBase(_buffer, _body, _menu, _content)
}
