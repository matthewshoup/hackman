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

						<li class="dropdown">
							
							<a href="#" class="dropdown-toggle" data-toggle="dropdown">
								<img style="max-width:25px;" src="{{.Avatar}}">
                                                                {{.Name}}
								<b class="caret"></b>
							</a>
							
							<ul class="dropdown-menu" role="menu" aria-labelledby="dropdownMenuDivider">
								
								<li><a href="#">My Account</a></li>
								<li><a href="#">Github</a></li>
								<li role="presentation" class="divider"></li>
								<li ><a href="/logout">Logout</a></li>
							
							</ul>
						
						</li>
						
					</ul>
				</div>
			</div>
		</div>


		<div class="container" style="background-color:#ffffff">
		<div class="jumbotron" style="background-color:#ffffff">
		<div class="row" >
			<div class="col-md-6">
                    <blockquote>Your Details </blockquote>
                    <form>
						<div class="form-group">
							<label for="teamName">Team Name</label>
						    <input type="text" class="form-control" id="teamName" placeholder="Name" required>
						</div>
				        <button type="submit" class="btn btn-default">Submit</button>
					</form>
			</div>
			<div class="col-md-6">
				<blockquote>Members</blockquote>
				<table class="table table-hover" style="margin-top:20px">
					<thead>
 					<tr>
 						<th>Name</th>
 						<th>Status</th>
 					</tr>
 					</thead>
 					<tbody>
 					<tr>
 						<td>Gophergala</td>
 						<td>Accepted</td>
 					</tr>
 					</tbody>
 					<tbody>
 					<tr>
 						<td>Node knockout</td>
 						<td>Pending</td>
 					</tr>
 					</tbody>
 					<tbody>
 					<tr>
 						<td>Django dash</td>
 						<td>Pending</td>
 					</tr>
 					</tbody>
 					<tbody>
 					<tr>
 						<td>Last Member</td>
 						<td>Pending</td>
 					</tr>
 					</tbody>
				</table>
				<form>
				  <div class="form-group">
				    <label for="exampleInputEmail1">Invite Member</label>
				    <input type="email" class="form-control" id="exampleInputEmail1" placeholder="Enter email" required>
				  </div>
				  <button type="submit" class="btn btn-default">Invite</button>
				</form>
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
