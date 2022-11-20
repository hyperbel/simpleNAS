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
      console.log(res)
    })
  }
}

function removefiles() {
  var checkBoxes = document.querySelectorAll('input[name=_checkboxes]:checked')
  for(var i = 0; i < checkBoxes.length-1; i++)
    console.log(checkBoxes[i])
}
