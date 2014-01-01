{{template "header" .}}
      {{template "menu" .}}

      <!-- Jumbotron -->
      <div class="jumbotron">
        <h1>Realtime status</h1>
        <p class="lead">Server, Socket, Bot</p>
      </div>

      <div class="row">
        <div class="col-lg-4">
          <h2>Build status</h2>
          <p>Láthatóvá válik a bot és a webadmin fordításának állapota.</p>
          <p><a class="btn btn-primary" href="/status-build" role="button">View details &raquo;</a></p>
        </div>
        <div class="col-lg-4">
          <h2>Admins</h2>
          <p>{{.Body}}</p>
          <p><a class="btn btn-primary" href="#" role="button">View details &raquo;</a></p>
        </div>
        <div class="col-lg-4">
          <h2>Admins</h2>
          <p>{{.Body}}</p>
          <p><a class="btn btn-primary" href="#" role="button">View details &raquo;</a></p>
        </div>
      </div>
{{template "footer" .}}
