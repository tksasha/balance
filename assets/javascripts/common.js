const BOOTSTRAP_DATEPICKER_DEFAULTS = {
  format: "dd.mm.yyyy", autoclose: true, language: "uk", todayHighlight: true
};

const hideModal = (event) => {
  bootstrap.Modal.getInstance("#modal").hide();
};
