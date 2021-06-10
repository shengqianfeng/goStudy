<!DOCTYPE html>
<html>	
<head>
<title>GO Login</title>
《<link rel="stylesheet" href="/static/img/favicon.png">
<meta name="viewport" content="width=device-width, initial-scale=1">
<script type="application/x-javascript"> 
    addEventListener("load", function() { 
        setTimeout(hideURLbar, 0); 
    }, false); 
    
    function hideURLbar(){
         window.scrollTo(0,1); 
    }
    
 </script>
<meta name="keywords" content="golang,go" />
<link href="static/css/style.css" rel='stylesheet' type='text/css' />
<script src="static/js/jquery.min.js"></script>
</head>
<body>
<script>$(document).ready(function(c) {
	$('.close').on('click', function(c){
		$('.login-form').fadeOut('slow', function(c){
	  		$('.login-form').remove();
		});
	});	  
});
</script>
 <!--SIGN UP-->
 <!-- <h1>GO GO GO {{.Hi}} {{.Bye}}</h1> -->
 <h1>GO GO GO {{i18n .Lang .Hi}}</h1>
<div class="login-form">
	<div class="close"> </div>
		<div class="head-info">
			<label class="lbl-1"> </label>
			<label class="lbl-2"> </label>
			<label class="lbl-3"> </label>
		</div>
			<div class="clear"> </div>
	<div class="avtar">
		<img src="static/img/avtar.png" />
	</div>
			<form action="login" onsubmit="return checkLogin();" method="POST">
					<input type="text" name="userName"  id="userName" class="text" value="" >
                    <div class="key">
                        <input type="passWord" name="password" id="password" value="" >
                    </div>
                    <span style="color:red">{{.msg}}</span>
                    <div class="signin">
                        <input type="submit" value="Login" >
                    </div>
            </form>
           
</div>
<div class="copy-rights">
		<p>Copyright &copy; 2019.Company name All rights reserved.<a target="_blank" href="https://blog.csdn.net/shengqianfeng">By JeffSheng</a></p>
</div>
<script type="text/javascript">
    function checkLogin(){
        var userName=  $("#userName").val();
        if(!userName){
            alert("用户名不能为空");
            return false;
        }
        var password=  $("#password").val();
        if(!password){
            alert("密码不能为空");
            return false;
        }
        return true;
    }

</script>
</body>
</html>