{{define "menu"}}
      <div class="masthead">
        <h3 class="muted">{{.ProjectName}}</h3>
        <div class="navbar">
          <div class="navbar-inner">
            <div class="container">
              <ul class="nav">
                <li class="active"><a href="/">Home</a></li>
                <li><a href="/stats">Statistics</a></li>
                <li><a href="/about">About</a></li>
                <li><a href="/login">Login</a></li>
              </ul>
            </div>
          </div>
        </div><!-- /.navbar -->
      </div>
{{end}}
