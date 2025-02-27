console.log(new Date);

const closeModal = (event) => {
  console.log(event);

  bootstrap.Modal.getInstance("#modal").hide();

  document.getElementById("modal-body").innerHTML = "";
}

document.addEventListener("balance.cash.updated", closeModal);
