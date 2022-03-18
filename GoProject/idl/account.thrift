include "base.thrift"
namespace go xintao.project.ecommerce

enum LoginResultCode {
    correct = 0 // 登录成功
    failed = 1 // 系统内部错误
    error = 2 // 密码错误
}

struct LoginResult {
    1: LoginResultCode LoginResultCode,
}

struct LoginRequest {
    1: required i64 PassportUid, // 账号
    2: required string PassportCode, // 密码
}
/*
登录之后 accesstoken塞到cookie里面
*/

struct LoginResponse {
    1: i32 code,
    2: string msg,
    3: LoginResult data,
}

struct LogoutRequest {
    1: required i64 PassportUid, // 账号
}

struct LogoutResponse {
    1: i32 code,
    2: string msg,
}

service EcommerceLogin {
    LoginResponse Login(1: LoginRequest req) (api.post="/xintao/project/ecommerce/login")
    LogoutResponse Logout(1: LogoutRequest req) (api.post="/xintao/project/ecommerce/logout")
}