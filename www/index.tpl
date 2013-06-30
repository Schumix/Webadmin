{{template "header"}}
      <div class="masthead">
        <h3 class="muted">Schumix Webadmin</h3>
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

      <!-- Jumbotron -->
      <div class="jumbotron">
        <h1>Teszt!</h1>
        <p class="lead">Test,Test,Test,Test,Test,Test,Test,Test,Test,Test,Test,Test,Test,Test,Test,Test,Test,Test</p>
        <a class="btn btn-large btn-success" href="#">Get started today</a>
      </div>

      <hr>

      <!-- Example row of columns -->
      <div class="row-fluid">
        <div class="span4">
          <h2>Admins</h2>
          <p>{{.Body}}</p>
          <p><a class="btn" href="#">View details &raquo;</a></p>
        </div>
        <div class="span4">
          <h2>Admins</h2>
          <p>{{.Body}}</p>
          <p><a class="btn" href="#">View details &raquo;</a></p>
       </div>
        <div class="span4">
          <h2>Admins</h2>
          <p>{{.Body}}</p>
          <p><a class="btn" href="#">View details &raquo;</a></p>
        </div>
      </div>
{{template "footer"}}