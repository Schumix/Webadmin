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
          <li><a href="/">Home</a></li>
          <li><a href="/stats">Statistics</a></li>
          <li><a href="/status">Status</a></li>
          <li><a href="/about">About</a></li>
          <li><a href="/login">Login</a></li>
        </ul>
      </div>
{{end}}
