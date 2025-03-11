const BOOTSTRAP_DATEPICKER_DEFAULTS = {
  format: "dd.mm.yyyy", autoclose: true, language: "uk", todayHighlight: true
};

const hideModal = (event) => {
  bootstrap.Modal.getInstance("#modal").hide();
};

const setModalSize = (size) => {
  const dialog = document.querySelector("#modal .modal-dialog");

  dialog.classList.remove("modal-lg", "modal-sm", "modal-xl");

  dialog.classList.add(size);
};
