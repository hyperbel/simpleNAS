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
  let text;
  let name = prompt("Directory name", "...")
  if (name != null || name != "") {
    fetch(`/createdir?name=${name}`, {
      method: 'POST'
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

document.getElementById("file_upload").onchange = () => {
  document.getElementById("hidden_url").innerHTML = window.location.pathname + window.location.search;
  document.getElementById("file_form").submit();
}
