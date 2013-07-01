{{define "menu"}}
      <div class="masthead">
        <h3 class="muted">{{.Title}}</h3>
        <div class="navbar">
          <div class="navbar-inner">
            <div class="container">
              <ul class="nav">
                <li {{if equal .Name "index"}}class="active"{{end}}><a href="/">Home</a></li>
                <li {{if equal .Name "stats"}}{{else}}class="active"{{end}}><a href="/stats">Statistics</a></li>
                <li {{if equal .Name "about"}}class="active"{{end}}><a href="/about">About</a></li>
                <li {{if equal .Name "login"}}class="active"{{end}}><a href="/login">Login</a></li>
              </ul>
            </div>
          </div>
        </div><!-- /.navbar -->
      </div>
{{end}}