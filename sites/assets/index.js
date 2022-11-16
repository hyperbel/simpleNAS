var enc = new TextEncoder();

function CreateAccountChanges() {
  var pass = enc.encode(document.createaccountform._password.value);
  var conf = enc.encode(document.createaccountform.confirm.value);
  if (pass == conf) crypto.subtle.digest('SHA-256', pass).then((pass_hash) => {
    console.log(pass);
    console.log(pass_hash);
    document.getElementById("password").value = pass_hash;
  })
}

function LogInChanges() {
  window.SubteCrypto.digest('SHA-256', enc.encode(document.loginform._passwd.value)).then((pw) => document.loginform.passwd = password);
  return true;
}
