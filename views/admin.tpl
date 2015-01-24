<!DOCTYPE html>
<html>

	<head>
		<title>Hackman</title>
                <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
                <meta name="viewport" content="width=device-width, initial-scale=1.0">

                <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.2/css/bootstrap.min.css">
	</head>

	<body>
		
		<div class="navbar navbar-default navbar-static-top">
			<div class="container">

				<div class="navbar-header">

					<a href="#" class="navbar-brand">
						<b style="color:#333; font-size:40px;">
							Hackman
						</b>
					</a>
				
					<button class="navbar-toggle" data-toggle="collapse" data-target=".navHeaderCollapse">
						<span class="icon-bar"></span>
						<span class="icon-bar"></span>
						<span class="icon-bar"></span>
					</button>
				</div>

				<div class="collapse navbar-collapse navHeaderCollapse" >

					<ul class="nav navbar-nav navbar-right">
						
						<li><a href="login.htm" style="color:2ac5ee;">Login</a></li>
						<li><a href="signup.htm">Organise Hackathon</a></li>
						
					</ul>
				</div>				
			</div>
		</div>
		<div class="container" style="background-color:#ffffff">
		<div class="jumbotron" style="background-color:#ffffff">
		<div class="row" >
			<div class="col-md-8">
                        <blockquote>Public Announcement</blockquote>
			<form class="form-inline">
				<input type="text" id="announcement" class="form-control" placeholder="Announcement Text" style="min-width:570px">
				<button type="submit" class="btn btn-default">Announce</button>
			</form>


                                <blockquote style="margin-top:20px;">
                                  Hackathon History
                                </blockquote>

				<table class="table table-hover" style="margin-top:20px">
					<thead>
 					<tr>
 						<th>Name</th>
 						<th>Time</th>
 					</tr>
 					</thead>
 					<tbody>
 					<tr>
 						<td>Gophergala</td>
 						<td>Jan. 23, 2015</td>
 					</tr>
 					</tbody>
 					<tbody>
 					<tr>
 						<td>Node knockout</td>
 						<td>Jan. 30, 2015</td>
 					</tr>
 					</tbody>
 					<tbody>
 					<tr>
 						<td>Django dash</td>
 						<td>Feb. 9, 2015</td>
 					</tr>
 					</tbody>
				</table>
			</div>
			<div class="col-md-4">
				<blockquote>Organise Hackathon</blockquote>
				<hr>
				<div class="form-group">
					<form>
				    <label for="hackathon">Hackathon</label>
				    <input type="text" class="form-control" id="hackathon" placeholder="Name of Hackathon">
				  </div>
				  <div class="form-group">
				    <label for="hackathon-description">Description</label>
				    <input type="text" class="form-control" id="hackathon-description" placeholder="Description for Hackathon">
				  </div>
				  <div class="form-group">
				    <label for="hackathon-organization">Github Organisation</label>
				    <input type="text" class="form-control" id="hackathon-organization" placeholder="GitHub Organization for Hackathon">
				  </div>
					</form>
				<button type="submit" class="btn btn-default">Create</button>
			</div>
		</div>
		</div>
		</div>



		<div class="navbar navbar-default navbar-fixed-bottom">

			<div class="container">
				<p class="navbar-text pull-left">Made With &lt3 and Golang</p>
				<a href="#" class="navbar-btn btn-danger btn pull-right">View Project on Github</a>
			</div>
			
		</div>

		<script type="text/javascript" src="http://code.jquery.com/jquery-2.1.3.min.js"></script>
                <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.2/js/bootstrap.min.js"></script>

	</body>

</html>
