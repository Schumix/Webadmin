{{define "menu"}}
      <div class="masthead">
        <h3 class="muted">{{.ProjectName}}</h3>
        <div class="navbar">
          <div class="navbar-inner">
            <div class="container">
              <ul class="nav">
                <li{{ if eq .PageName "home" }} class="active"{{ end }}><a href="/">Home</a></li>
                <li{{ if eq .PageName "stats" }} class="active"{{ end }}><a href="/stats">Statistics</a></li>
                <li{{ if eq .PageName "about" }} class="active"{{ end }}><a href="/about">About</a></li>
                <li{{ if eq .PageName "login" }} class="active"{{ end }}><a href="/login">Login</a></li>
              </ul>
            </div>
          </div>
        </div><!-- /.navbar -->
      </div>
{{end}}
