{{define "menu"}}
      <div class="masthead">
        <h3 class="muted">{{.ProjectName}}</h3>
        <div class="navbar">
          <div class="navbar-inner">
            <div class="container">
              <ul class="nav">
                <li{{ if eq .PageName "home" }} class="active"{{ end }}><a href="/">Home</a></li>
                <li{{ if eq .PageName "stats" }} class="active"{{ end }}><a href="/stats">Statistics</a></li>
                <li{{ if eq .PageName "status" }} class="active"{{ end }}><a href="/status">Status</a></li>
                <li{{ if eq .PageName "about" }} class="active"{{ end }}><a href="/about">About</a></li>
{{if .IsLoggedIn}}
                <li{{ if eq .PageName "logout" }} class="active"{{ end }}><a href="/logout">Logout</a></li>
{{else}}
                <li{{ if eq .PageName "login" }} class="active"{{ end }}><a href="/login">Login</a></li>
{{end}}
              </ul>
            </div>
          </div>
        </div><!-- /.navbar -->
      </div>
{{end}}
