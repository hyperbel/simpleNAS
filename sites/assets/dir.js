function back() {
  console.log(window.location.href)
  fetch("/back", { 
    method: 'POST', 
    body: window.location.href,
  }).then((res) => res.json())
  .then((data) => {
    console.log(data)
    window.location.href = data["url"]
  })
}

function createdir() {
  let name = prompt("Directory name", "...")
  if (name != null || name != "") {
    fetch(`/createdir?name=${name}`, {
      method: 'POST',
      body: JSON.stringify({
        'search': window.location.search
      })
    }).then((res) => res.json()).then((data) => {
      console.log(data)
      location.reload()
    })
  }
}

function removefiles() {
  var checkBoxes = []; 
  document.querySelectorAll('input[name=_checkbox]:checked').forEach((el) => checkBoxes.push(el.id))
  fetch("/removefiles", {
    method: 'POST',
    body: JSON.stringify({
      'files': checkBoxes,
      'search': window.location.search
    })
  })
}

function file_upload_form_on_submit() {
  alert("submitting")
  document.getElementById("hidden_url").value = window.location.search
}
