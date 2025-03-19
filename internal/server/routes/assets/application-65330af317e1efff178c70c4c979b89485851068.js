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
document.addEventListener("backoffice.index.shown", (e) => {
  setModalSize("modal-sm");
});

document.addEventListener("backoffice.cashes.shown", (e) => {
  clearModalSize();
});

document.addEventListener("backoffice.cash.updated", async (e) => {
  if (Object.hasOwn(e.detail, "backofficeCashesPath"))
    await htmx.ajax("GET", e.detail.backofficeCashesPath, { target: "#modal-body" });
});

document.addEventListener("backoffice.cash.deleted", async (e) => {
  if (Object.hasOwn(e.detail, "backofficeCashesPath"))
    await htmx.ajax("GET", e.detail.backofficeCashesPath, { target: "#modal-body" });
});

document.addEventListener("backoffice.cash.created", async (e) => {
  if (Object.hasOwn(e.detail, "backofficeCashesPath"))
    await htmx.ajax("GET", e.detail.backofficeCashesPath, { target: "#modal-body" });
});

document.addEventListener("backoffice.categories.shown", (e) => {
  setModalSize("modal-lg");
});

document.addEventListener("backoffice.category.updated", async (e) => {
  if (Object.hasOwn(e.detail, "backofficeCategoriesPath"))
    await htmx.ajax("GET", e.detail.backofficeCategoriesPath, { target: "#modal-body" });
});
document.addEventListener("balance.cash.updated", async (e) => {
  hideModal();

  if (Object.hasOwn(e.detail, "balancePath"))
    await htmx.ajax("GET", e.detail.balancePath, { target: "#balance", swap: "outerHTML" });
});

document.addEventListener("balance.cash.edit", clearModalSize);
htmx.on("balance.month.changed", (e) => {
  if (Object.hasOwn(e.detail, "balanceCategoriesPath"))
    htmx.ajax("GET", e.detail.balanceCategoriesPath, { target: "#categories", swap: "outerHTML" });
});

htmx.on("balance.year.changed", (e) => {
  if (Object.hasOwn(e.detail, "balanceCategoriesPath"))
    htmx.ajax("GET", e.detail.balanceCategoriesPath, { target: "#categories", swap: "outerHTML" });
});
document.addEventListener("balance.item.initialized", (e) => {
  $(".datepicker").datepicker(BOOTSTRAP_DATEPICKER_DEFAULTS);
});

document.addEventListener("balance.item.edit", clearModalSize);

document.addEventListener("balance.item.updated", async (e) => {
  hideModal();

  if (Object.hasOwn(e.detail, "balancePath"))
    await htmx.ajax("GET", e.detail.balancePath, { target: "#balance", swap: "outerHTML" });

  if (Object.hasOwn(e.detail, "categoriesPath"))
    await htmx.ajax("GET", e.detail.categoriesPath, { "target": "#categories", swap: "outerHTML" });
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

document.addEventListener("balance.item.create.error", (e) => {
  clearModalSize();

  const modal = bootstrap.Modal.getOrCreateInstance("#modal").show();

  $("#modal .datepicker").datepicker(BOOTSTRAP_DATEPICKER_DEFAULTS);
});

document.addEventListener("balance.item.created", async (e) => {
  hideModal();

  $(".datepicker").datepicker(BOOTSTRAP_DATEPICKER_DEFAULTS);

  if (Object.hasOwn(e.detail, "balancePath"))
    await htmx.ajax("GET", e.detail.balancePath, { target: "#balance", swap: "outerHTML" });

  if (Object.hasOwn(e.detail, "categoriesPath"))
    await htmx.ajax("GET", e.detail.categoriesPath, { "target": "#categories", swap: "outerHTML" });

  if (Object.hasOwn(e.detail, "itemsPath"))
    await htmx.ajax("GET", e.detail.itemsPath, { "target": "#items" });
});

document.addEventListener("balance.item.deleted", async (e) => {
  hideModal();

  if (Object.hasOwn(e.detail, "balancePath"))
    await htmx.ajax("GET", e.detail.balancePath, { target: "#balance", swap: "outerHTML" });

  if (Object.hasOwn(e.detail, "categoriesPath"))
    await htmx.ajax("GET", e.detail.categoriesPath, { "target": "#categories", swap: "outerHTML" });

  if (Object.hasOwn(e.detail, "itemsPath"))
    await htmx.ajax("GET", e.detail.itemsPath, { "target": "#items" });
});
