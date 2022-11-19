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
  fetch("/createdir?name=", {
    method: 'POST'
  }).then((res) => res.json()).then((data) => {
    console.log(res)
  })
}