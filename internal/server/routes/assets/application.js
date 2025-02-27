console.log("done");

document.addEventListener("balance.ping.created", (event) => {
  bootstrap.Modal.getInstance("#modal").hide();

  document.getElementById("modal-body").innerHTML = "";
})
