jQuery(document).ready(function($){

	// hide messages 
	$("#error").hide();
	$("#sent-form-msg").hide();

	// login submit...
	$("#loginForm #submit").click(function() {
		$("#error").hide();

		//required:

		//id
		var id = $("input#id").val();
		if(id == ""){
			$("#error").fadeIn().text("ID required.");
			$("input#id").focus();
			return false;
		}

		// password
		var password = $("input#password").val();
		if(password == ""){
			$("#error").fadeIn().text("Password required");
			$("input#password").focus();
			return false;
		}
		return true;
	});
	
	// on submit...
	$("#contactForm #submit").click(function() {
		$("#error").hide();
		
		//required:
		
		//name
		var name = $("input#name").val();
		if(name == ""){
			$("#error").fadeIn().text("Name required.");
			$("input#name").focus();
			return false;
		}
		
		// email
		var email = $("input#email").val();
		if(email == ""){
			$("#error").fadeIn().text("Email required");
			$("input#email").focus();
			return false;
		}
		
		// web
		var web = $("input#web").val();
		if(web == ""){
			$("#error").fadeIn().text("Web required");
			$("input#web").focus();
			return false;
		}
		
		// comments
		var comments = $("#comments").val();
		
		// send mail php
		var sendMailUrl = $("#sendMailUrl").val();
		
		//to, from & subject
		var to = $("#to").val();
		var from = $("#from").val();
		var subject = $("#subject").val();
		
		// data string
		var dataString = 'name='+ name
						+ '&email=' + email        
						+ '&web=' + web
						+ '&comments=' + comments
						+ '&to=' + to
						+ '&from=' + from
						+ '&subject=' + subject;						         
		// ajax
		$.ajax({
			type:"POST",
			url: sendMailUrl,
			data: dataString,
			success: function(){
                $("#sent-form-msg").fadeIn();
                $("#contactForm").fadeOut();
             }
		});
	});

	// project form submit...
	$("#projectform #submit").click(function() {
		$("#error").hide();

		//project name
		var proj_name = $("input#projectName").val();
		if(!proj_name){
			$("#error").fadeIn().text("ProjectName required.");
			$("input#projectName").focus();
			return false;
		}

		// project comments
		var proj_comment = $("textarea#projectComment").val();
		if(!proj_comment){
			$("#error").fadeIn().text("ProjectComment required.");
			$("textarea#projectComment").focus();
			return false;
		}

		// company name
		var comp_name = $("input#companyName").val();
		if(!comp_name){
			$("#error").fadeIn().text("CompanyName required.");
			$("input#companyName").focus();
			return false;
		}

		// position
		var position = $("input#position").val();
		if(!position){
			$("#error").fadeIn().text("Position required.");
			$("input#position").focus();
			return false;
		}

		// start date
		var start_dt = $("input#startDate").val();
		if(!start_dt){
			$("#error").fadeIn().text("StartDate required.");
			$("input#startDate").focus();
			return false;
		}

		// end date
		var end_dt = $("input#endDate").val();
		if(!end_dt){
			$("#error").fadeIn().text("EndDate required.");
			$("input#endDate").focus();
			return false;
		}

		//launch_url
		var launch_url = $("#launchUrl").val();

		// data string
		var dataString = 'project.ProjectName='+ proj_name
						+ '&project.ProjectComment=' + proj_comment
						+ '&project.CompanyName=' + comp_name
						+ '&project.Position=' + position
						+ '&project.StartDate=' + start_dt
						+ '&project.EndDate=' + end_dt
						+ '&project.LaunchUrl=' + launch_url;

        //add,update check
        var method = $("#method").val()
        var sendUrl = "/projects";
        if(method == "put") {
            sendUrl = sendUrl + "/"+$("input#projectId").val()
        }

		// ajax
		$.ajax({
			type:method,
			url: sendUrl,
			data: dataString,
			success: function(data) {
			    $("#sent-form-msg").fadeIn();
			    $("#projectform").fadeOut();
			},
			error: function(request,status,error) {
			    $("#error").fadeIn().text("status: "+request.status+", detail: "+request.responseText);
			}
		});
	});
	
    return false;
});