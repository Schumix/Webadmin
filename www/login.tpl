{{template "header" .}}
	{{template "menu" .}}
{{if .SessionValue}}
{{else}}
      <form action="/login" method="post" class="form-signin">
        <h2 class="form-signin-heading">Please sign in</h2>
        <input type="text" name="userid" class="input-block-level" placeholder="Account Name">
        <input type="password" name="password" class="input-block-level" placeholder="Password">
        <label class="checkbox">
          <input type="checkbox" value="remember-me"> Remember me
        </label>
        <button class="btn btn-large btn-primary" type="submit">Sign in</button>
      </form>
{{end}}
{{template "footer" .}}
