const BOOTSTRAP_DATEPICKER_DEFAULTS = {
  format: "dd.mm.yyyy", autoclose: true, language: "uk", todayHighlight: true
};

const hideModal = (event) => {
  const modal = bootstrap.Modal.getInstance("#modal");

  if (modal)
    modal.hide();
};

const showModal = () => {
  bootstrap.Modal.getOrCreateInstance("#modal").show();
};

const clearModalSize = () => {
  document
    .querySelector("#modal .modal-dialog")
    .classList.remove("modal-lg", "modal-sm", "modal-xl");
};

const setModalSize = (size) => {
  clearModalSize();

  document
    .querySelector("#modal .modal-dialog")
    .classList.add(size);
};

document.getElementById("modal").addEventListener("shown.bs.modal", (e) => {
  const input = e.target.querySelector("[autofocus]")

  if (input)
    input.focus();
});
