$(function () {
    $("#sign_in_button").click(function () {

        let email = $("#sign_in_email").val()
        let password = $("#sign_in_password").val()

        if (!IsEmail(email)) {
            bootAlert($("#sign_in_error"), "请填写正确的邮箱地址", "danger")
            return
        }

        if (password.length < 6) {
            bootAlert($("#sign_in_error"), "密码过短", "danger")
            return
        }

        let data = new FormData()
        data.append("email", email)
        data.append("password", password)

        $.ajax({
            url: "/user/login",
            type: "POST",
            data: data,
            cache: false,
            processData: false,
            contentType: false,
            async: false,
            success: function (msg) {
                if (msg.code == 200) {
                    let val = msg.data
                    $("#login_user_name").html(val.username)

                    bootstrap.Modal.getInstance($("#SignIn")).hide()
                    $("#menu_button_group").hide()
                    $("#login_user_group").show()
                } else {
                    bootAlert($("#sign_in_error"), msg.error, "danger")
                }
            }
        })


    });
});

function IsEmail(str) {
    let reg = /^([a-zA-Z]|[0-9])(\w|\-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/;
    return reg.test(str);
}

function bootAlert(alertPlaceholder, message, type) {
    alertPlaceholder.empty()
    const wrapper = document.createElement('div')
    wrapper.innerHTML = [
        `<div class="alert alert-${type} alert-dismissible fixed-top text-center" role="alert">`,
        `   <div>${message}</div>`,
        '   <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>',
        '</div>'
    ].join('')

    alertPlaceholder.append(wrapper)

    setTimeout(() => {
        alertPlaceholder.empty()
    }, 3000);
}