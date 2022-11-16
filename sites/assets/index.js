var enc = new TextEncoder();

function CreateAccountChanges() {
  var password = document.createaccountform._password.value;
  var pw_confirm = document.createaccountform.confirm.value;
  var res_pw = enc.encode(password)
  var res_conf = enc.encode(pw_confirm)
  if (res_pw == res_conf) crypto.subtle.digest('SHA-256', res_pw).then((pass_hash) => {
    console.log(res_pw)
    console.log(pass_hash)
  })
}

function LogInChanges() {
  window.SubteCrypto.digest('SHA-256', enc.encode(document.loginform._passwd.value)).then((pw) => document.loginform.passwd = password);
  return true;
}
