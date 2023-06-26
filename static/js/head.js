$(function () {

    start()

    // 退出登录按钮
    $("#sign_out_button").click(function () {
        $.removeCookie("miajio_token")
        window.location.href = "/index.html"
    });

    $("#sign_out_save_button").click(function () {
        let tk = $.cookie("miajio_token")
        if (null == tk || undefined == tk || "" == tk) {
            bootAlert($("#head_alert_div"), "登录信息无法获取,请重新登录", "danger")
        }

        let headPic = $("#sign_out_head_pic")[0].files[0]
        let username = $("#sign_out_username").val()

        let data = new FormData()
        data.append("headPic", headPic)
        data.append("username", username)

        $.ajax({
            url: "/user/update",
            type: "POST",
            headers: {
                "Authorization": tk
            },
            data: data,
            cache: false,
            processData: false,
            contentType: false,
            success: function(msg) {
                if (msg.code == 200) {
                    $.cookie("miajio_token", msg.data, { expires: 7 })
                    window.location.href = "/index.html"
                } else {
                    bootAlert($("#head_alert_div"), msg.error, "danger")
                }
            }
        })
    });

    $("#sign_out_user_head").click(function () {
        $("#sign_out_head_pic").click()
        $(document).on("change", "#sign_out_head_pic", function () {

            let inputObj = $("#sign_out_head_pic")[0]

            var reader = new FileReader()
            reader.onload = function(e) {
                if (reader.readyState === 2) {
                    $("#sign_out_user_head").attr("src", e.target.result)
                }
            }
            reader.readAsDataURL(inputObj.files[0])
        })

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
            success: function (msg) {
                if (msg.code == 200) {
                    let tk = msg.data
                    $.ajax({
                        url: "/user/auth",
                        type: "GET",
                        headers: {
                            "Authorization": tk
                        },
                        success: function (v) {
                            if (v.code == 200) {
                                let val = v.data
                                $.cookie("miajio_token", tk, { expires: 7 })
                                $("#login_user_name").html(val.username)
                                $("#sign_out_username").val(val.username)
                                $("#login_user_head").attr("src", val.headPic)
                                $("#sign_out_user_head").attr("src", val.headPic)

                                bootstrap.Modal.getInstance($("#SignIn")).hide()
                                $("#menu_button_group").hide()
                                $("#login_user_group").show()
                                bootAlert($("#head_alert_div"), msg.msg, "success")
                            } else {
                                bootAlert($("#sign_in_error"), msg.error, "danger")
                            }
                        }
                    })
                } else {
                    bootAlert($("#sign_in_error"), msg.error, "danger")
                }
            }
        })
    });

    // 注册按钮
    $("#sign_up_button").click(function () {
        let emailUid = $("#sign_up_email_uid").html()
        let username = $("#sign_up_user_name").val()
        let email = $("#sign_up_email").val()
        let checkCode = $("#sign_up_check_code").val()
        let password = $("#sign_up_password").val()
        let checkPassword = $("#sign_up_check_password").val()

        if (password != checkPassword) {
            bootAlert($("#sign_up_error"), "校验密码错误,请确认两次密码是否一致", "danger")
            return
        }

        let data = {
            "uid": emailUid,
            "username": username,
            "email": email,
            "checkCode": checkCode,
            "password": password
        }

        $.ajax({
            url: "/user/register",
            type: "POST",
            data: JSON.stringify(data),
            dataType: "json",
            cache: false,
            processData: false,
            success: function (msg) {
                if (msg.code == 200) {
                    let tk = msg.data

                    $.ajax({
                        url: "/user/auth",
                        type: "GET",
                        headers: {
                            "Authorization": tk
                        },
                        success: function (v) {
                            if (v.code == 200) {
                                let val = v.data
                                $.cookie("miajio_token", tk, { expires: 7 })
                                $("#login_user_name").html(val.username)
                                $("#sign_out_username").val(val.username)
                                $("#login_user_head").attr("src", val.headPic)
                                $("#sign_out_user_head").attr("src", val.headPic)

                                bootstrap.Modal.getInstance($("#SignUp")).hide()
                                $("#menu_button_group").hide()
                                $("#login_user_group").show()
                                bootAlert($("#head_alert_div"), msg.msg, "success")
                            } else {
                                bootAlert($("#sign_up_error"), msg.error, "danger")
                            }
                        }
                    })
                } else {
                    bootAlert($("#sign_up_error"), msg.error, "danger")
                }
            }
        })


    });

    // 发送验证码按钮
    $("#send_email_button").click(function () {
        let email = $("#sign_up_email").val()
        if (!IsEmail(email)) {
            bootAlert($("#sign_up_error"), "请填写正确的邮箱地址", "danger")
            return
        }

        let data = {
            "email": email,
            "emailType": "register"
        }

        $.ajax({
            url: "/email/sendCheckCode",
            type: "POST",
            data: JSON.stringify(data),
            dataType: "json",
            cache: false,
            processData: false,
            success: function (msg) {
                if (msg.code == 200) {
                    $("#sign_up_email_uid").html(msg.data)

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
                } else {
                    bootAlert($("#sign_up_error"), msg.error, "danger")
                }
            }
        })


    });
});

// 界面开始时加载操作
function start() {
    let tk = $.cookie("miajio_token")
    if (null != tk && undefined != tk && "" != tk) {
        $.ajax({
            url: "/user/auth",
            type: "GET",
            headers: {
                "Authorization": tk
            },
            success: function (v) {
                if (v.code == 200) {
                    let user = v.data
                    $("#login_user_name").html(user.username)
                    $("#login_user_head").attr("src", user.headPic)
                    $("#sign_out_user_head").attr("src", user.headPic)

                    $("#menu_button_group").hide()
                    $("#login_user_group").show()
                    $("#sign_out_username").val(user.username)
                }
            }
        })
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