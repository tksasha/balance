const BOOTSTRAP_DATEPICKER_DEFAULTS = {
  format: 'dd.mm.yyyy', autoclose: true, language: 'uk', todayHighlight: true
};

const hideModal = (event) => {
  bootstrap.Modal.getInstance("#modal").hide();
};

document.addEventListener("balance.cash.updated", hideModal);

document.addEventListener("balance.item.initialized", (e) => {
  $(".datepicker").datepicker(BOOTSTRAP_DATEPICKER_DEFAULTS);
});

document.addEventListener("balance.items.shown", (e) => {
  const month = e.detail.month;
  const months = document.getElementById("months");

  for (const child of months.children) {
    child.classList.remove("active");

    if (child.dataset.number == month)
      child.classList.add("active");
  }

  const year = e.detail.year;
  const years = document.getElementById("years");

  for (const child of years.children) {
    child.classList.remove("active");

    if (child.dataset.number == year)
      child.classList.add("active");
  }
});
