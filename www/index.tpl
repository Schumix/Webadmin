{{template "header" .}}
      {{template "menu" .}}

	{{if .Success}}
      <div class="alert alert-success alert-dismissable">  
        <button type="button" class="close" data-dismiss="alert" aria-hidden="true">&times;</button>
        <strong>Success!</strong> {{.Body}}  
      </div>  
	{{end}}

      <!-- Jumbotron -->
      <div class="jumbotron">
        <h1>Schumix</h1>
        <p class="lead">Schumix Irc Bot and Framework</p>
        <a class="btn btn-large btn-success" href="https://github.com/Schumix/Schumix2" role="button">Get started today</a>
      </div>

      <div class="row">
        <div class="col-lg-4">
          <h2>Multifunctional bot</h2>
          <p>More irc servers at a time</p>
        </div>
        <div class="col-lg-4">
          <h2>Built-in addon system</h2>
          <p>
            C# compiler, Calendar and many more addons<br>
            Create your own addon easily with the API
          </p>
       </div>
        <div class="col-lg-4">
          <h2>Settings</h2>
          <p>Edit the bot's settings easily in the console</p>
        </div>
      </div>
{{template "footer" .}}
