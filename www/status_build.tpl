{{template "header" .}}
      {{template "menu" .}}

      <!-- Jumbotron -->
      <div class="jumbotron">
        <h1>Realtime status</h1>
        <p class="lead">Server, Socket, Bot</p>
      </div>

      <div class="panel panel-default">
        <div class="panel-heading">Schumix</div>
        <div class="panel-body">
          <div class="row">
            <div class="col-lg-6">
              <table class="table table-hover table-bordered">
                <tr class="field-label active">
                  <td class="field-label active">
                    <b>Master</b>
                  </td>
                </tr>
                <tr>
                  <td>
                    <b>Platform</b></td>
                  <td>
                    <b>Configuration</b>
                  </td>
                  <td>
                    <b>Build status</b>
                  </td>
                  <td>
                    <b>Downloads</b>
                  </td>
                </tr>
                <tr>
                  <td>Ubuntu</td>
                  <td>Debug</td>
                  <td>
                    <a href="https://travis-ci.org/Schumix/Schumix2"><img src="https://api.travis-ci.org/Schumix/Schumix2.png?branch=master"></a>
                  </td>
                  <td>None</td>
                </tr>
                <tr>
                  <td>Windows</td>
                  <td>Debug</td>
                  <td>
                    <a href="https://ci.appveyor.com/project/Schumix2"><img src="https://ci.appveyor.com/api/projects/status?id=u8njmbb6ivqisgiu"></a>
                  </td>
                  <td>None</td>
                </tr>
              </table>
            </div>
            <div class="col-lg-6">
              <table class="table table-hover table-bordered">
                <tr>
                  <td class="field-label active">
                    <b>Stable</b>
                  </td>
                </tr>
                <tr>
                  <td>
                    <b>Platform</b></td>
                  <td>
                    <b>Configuration</b>
                  </td>
                  <td>
                    <b>Build status</b>
                  </td>
                  <td>
                    <b>Downloads</b>
                  </td>
                </tr>
                <tr>
                  <td>Ubuntu</td>
                  <td>Debug</td>
                  <td>
                    <a href="https://travis-ci.org/Schumix/Schumix2"><img src="https://api.travis-ci.org/Schumix/Schumix2.png?branch=stable"></a>
                  </td>
                  <td>None</td>
                </tr>
                <tr>
                  <td>Windows</td>
                  <td>Debug</td>
                  <td>
                    <a href="https://ci.appveyor.com/project/Schumix2-stable"><img src="https://ci.appveyor.com/api/projects/status?id=pjsrmeh2utc2100w"></a>
                  </td>
                  <td>None</td>
                </tr>
              </table>
            </div>
          </div>
        </div>
      </div>
{{template "footer" .}}
