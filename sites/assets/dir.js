function back() {
  fetch("/back", data={
    method: 'POST'
  }).then((res) => res.json())
  .then((data) => console.log(data))
  window.location.href = 
}