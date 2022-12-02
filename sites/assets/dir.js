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


document.getElementById("file_upload").onchange = function (ev) {
  let files = ev.target.files;
  Array.from(files).forEach((file) => {
    fetch('/uploadfile', {
      method: 'post',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        'file_upload': file,
        'url': window.location.pathname + window.location.search    
      })
    })
  })
}
