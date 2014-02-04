$(document).ajaxError (function(event, jqxhr, settings, exception) {
  alert (exception);
});

$("#login-form").submit (function (event) {
  event.preventDefault();
  $("#msg").hide();
  $.post ("/login",  $("#login-form").serialize(),
    function (data) {
      if (!data.Authenticated)
        $("#msg").show().text ("Login failed");
      else
        
    });
});

