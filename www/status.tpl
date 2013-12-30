{{template "header" .}}
      {{template "menu" .}}

      <!-- Jumbotron -->
      <div class="jumbotron">
        <h1>Realtime status</h1>
        <p class="lead">Server, Socket, Bot</p>
        <a class="btn btn-large btn-success" href="https://github.com/Schumix/Webadmin">Get started today</a>
      </div>

      <div class="row">
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
        <div class="col-lg-4">
          <h2>Admins</h2>
          <p>{{.Body}}</p>
          <p><a class="btn btn-primary" href="#" role="button">View details &raquo;</a></p>
        </div>
      </div>
{{template "footer" .}}
