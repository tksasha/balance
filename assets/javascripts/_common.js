const BOOTSTRAP_DATEPICKER_DEFAULTS = {
  format: "dd.mm.yyyy", autoclose: true, language: "uk", todayHighlight: true
};

const hideModal = (event) => {
  bootstrap.Modal.getInstance("#modal").hide();
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

htmx.on("balance.month.changed", (e) => {
  if (Object.hasOwn(e.detail, "balanceCategoriesPath"))
    htmx.ajax("GET", e.detail.balanceCategoriesPath, { target: "#categories" });
});

htmx.on("balance.year.changed", (e) => {
  if (Object.hasOwn(e.detail, "balanceCategoriesPath"))
    htmx.ajax("GET", e.detail.balanceCategoriesPath, { target: "#categories" });
});
