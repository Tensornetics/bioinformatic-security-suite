$(document).ready(function () {
    $("#login-form").submit(function (event) {
        event.preventDefault();

        var username = $("#username").val();
        var password = $("#password").val();

        $.ajax({
            type: "POST",
            url: "/login",
            data: JSON.stringify({ username: username, password: password }),
            contentType: "application/json",
            success: function (data) {
                if (data.success) {
                    window.location.href = "/dashboard";
                } else {
                    $("#error-message").text(data.message);
                }
            }
        });
    });
});
