{{template "header" .}}
	{{template "menu" .}}
	{{if .Error}}
      <div class="alert alert-danger alert-dismissable">
        <button type="button" class="close" data-dismiss="alert" aria-hidden="true">&times;</button>
        <strong>Error!</strong> {{.Body}}
      </div>
	{{end}}

      <form action="/login" method="post" class="form-signin" role="form">
        <h2 class="form-signin-heading">Please sign in</h2>
        <input type="text" name="userid" class="form-control" placeholder="Account Name" required="" autofocus="">
        <input type="password" name="password" class="form-control" placeholder="Password" required="">
        <label class="checkbox">
          <input type="checkbox" value="remember-me"> Remember me
        </label>
        <button class="btn btn-lg btn-primary btn-block" type="submit">Sign in</button>
      </form>
{{template "footer" .}}
