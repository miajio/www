$(function () {

    start()

    // 退出登录按钮
    $("#sign_out_button").click(function () {
        $.removeCookie("user_msg")
        window.location.href = "/index.html"
    });

    $("#sign_out_save_button").click(function () {
        let username = $("#sign_out_username").val()
    });

    // 登录按钮
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
                    $.cookie("user_msg", JSON.stringify(val), { expires: 7 })
                    $("#login_user_name").html(val.username)
                    $("#sign_out_username").val(val.username)

                    bootstrap.Modal.getInstance($("#SignIn")).hide()
                    $("#menu_button_group").hide()
                    $("#login_user_group").show()
                } else {
                    bootAlert($("#sign_in_error"), msg.error, "danger")
                }
            }
        })
    });

    // 注册按钮
    $("#sign_up_button").click(function () {
        let username = $("#sign_up_user_name").val()
        let email = $("#sign_up_email").val()
        let checkCode = $("#sign_up_check_code").val()
        let password = $("#sign_up_password").val()
        let checkPassword = $("#sign_up_check_password").val()

    });

    // 发送验证码按钮
    $("#send_email_button").click(function () {
        let email = $("#sign_up_email").val()
        if (!IsEmail(email)) {
            bootAlert($("#sign_up_error"), "请填写正确的邮箱地址", "danger")
            return
        }

        let outTime = 60
        $("#send_email_button").attr("disabled", "true")
        let interval = window.setInterval(function () {
            $("#send_email_button").html(outTime + "秒后重新获取")
            if (outTime <= 0) {
                $("#send_email_button").removeAttr("disabled")
                $("#send_email_button").html("重新获取")
                window.clearInterval(interval)
            }
            outTime--
        }, 1000)
    });
});

// 界面开始时加载操作
function start() {
    let user_msg = $.cookie("user_msg")
    if (null != user_msg && undefined != user_msg && "" != user_msg) {
        let user = JSON.parse(user_msg)
        $("#login_user_name").html(user.username)
        $("#menu_button_group").hide()
        $("#login_user_group").show()

        $("#sign_out_username").val(user.username)
    }
}

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