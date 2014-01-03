{{define "menu"}}
      <div class="masthead">
        <div class="row">
            <h3 class="text-muted col-lg-3">{{.ProjectName}}</h3>

            <div class="col-lg-3" style="float:right;">
              <form action="{{.PageName}}" id="serverlist">
                <div class="input-group">
                  <select name="server_change" form="serverlist" class="form-control">
                    <option value="rizon">Rizon</option>
                    <option value="teszt">Teszt</option>
                    <option value="teszt2" selected>Teszt2</option>
                    <option value="teszt3">Teszt3</option>
                  </select>

                  <div class="input-group-btn">
                    <input type="submit" class="btn btn-primary" value="Submit">
                  </div>
                </div>
              </form>
            </div>
        </div>

        <ul class="nav nav-justified">
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
{{end}}
