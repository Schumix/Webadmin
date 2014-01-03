{{template "header" .}}
      {{template "menu" .}}

      <!-- Jumbotron -->
      <div class="jumbotron">
        <h1>Realtime statistics</h1>
        <p class="lead">Get realtime statistics about the bot</p>
      </div>

      <div class="row">
        <div class="col-xs-6 col-sm-4 col-lg-4">
          <h2>Admins</h2>
          <p>{{.Body}}</p>
          <p><a class="btn btn-primary" href="#" role="button">View details &raquo;</a></p>
        </div>
        <div class="col-xs-6 col-sm-4 col-lg-4">
          <h2>Admins</h2>
          <p>{{.Body}}</p>
          <p><a class="btn btn-primary" href="#" role="button">View details &raquo;</a></p>
        </div>
        <div class="col-xs-6 col-sm-4 col-lg-4">
          <h2>Admins</h2>
          <p>{{.Body}}</p>
          <p><a class="btn btn-primary" href="#" role="button">View details &raquo;</a></p>
        </div>
      </div>
{{template "footer" .}}
