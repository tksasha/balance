const hideModal = (event) => {
  bootstrap.Modal.getInstance("#modal").hide();
};

document.addEventListener("balance.cash.updated", hideModal);
