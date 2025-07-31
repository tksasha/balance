export const BOOTSTRAP_DATEPICKER_DEFAULTS = {
  format: "dd.mm.yyyy", autohide: true, language: "uk", todayHighlight: true
};

export const hideModal = (event) => {
  const modal = bootstrap.Modal.getInstance("#modal");

  if (modal)
    modal.hide();
};

export const showModal = () => {
  bootstrap.Modal.getOrCreateInstance("#modal").show();
};

export const clearModalSize = () => {
  document
    .querySelector("#modal .modal-dialog")
    .classList.remove("modal-lg", "modal-sm", "modal-xl");
};

export const setModalSize = (size) => {
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
