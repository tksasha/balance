const BOOTSTRAP_DATEPICKER_DEFAULTS = {
  format: "dd.mm.yyyy", autohide: true, language: "uk", todayHighlight: true
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
document.addEventListener("backoffice.index.shown", (e) => {
  setModalSize("modal-sm");
});

document.addEventListener("backoffice.cashes.shown", (e) => {
  clearModalSize();
});

document.addEventListener("backoffice.cash.deleted", async (e) => {
  if (Object.hasOwn(e.detail, "backofficeCashesPath"))
    await htmx.ajax("GET", e.detail.backofficeCashesPath, { target: "#modal-body" });
});

document.addEventListener("backoffice.cash.created", async (e) => {
  if (Object.hasOwn(e.detail, "backofficeCashesPath"))
    await htmx.ajax("GET", e.detail.backofficeCashesPath, { target: "#modal-body" });
});
htmx.on("balance.cash.updated", async (e) => {
  hideModal();

  if (Object.hasOwn(e.detail, "balancePath"))
    await htmx.ajax("GET", e.detail.balancePath, { target: "#balance", swap: "outerHTML" });
});

htmx.on("balance.cash.edit", clearModalSize);

htmx.on("backoffice.cash.updated", async (e) => {
  if (Object.hasOwn(e.detail, "backofficeCashesPath"))
    await htmx.ajax("GET", e.detail.backofficeCashesPath, { target: "#modal-body" });

  if (Object.hasOwn(e.detail, "balancePath"))
    await htmx.ajax("GET", e.detail.balancePath, { target: "#balance-and-cashes-row" });

  if (Object.hasOwn(e.detail, "cashesPath"))
    await htmx.ajax("GET", e.detail.cashesPath, { target: "#balance-and-cashes-row", swap: "beforeend" });
});
htmx.on("balance.month.changed", (e) => {
  if (Object.hasOwn(e.detail, "balanceCategoriesPath"))
    htmx.ajax("GET", e.detail.balanceCategoriesPath, { target: "#categories", swap: "outerHTML" });
});

htmx.on("balance.year.changed", (e) => {
  if (Object.hasOwn(e.detail, "balanceCategoriesPath"))
    htmx.ajax("GET", e.detail.balanceCategoriesPath, { target: "#categories", swap: "outerHTML" });
});

htmx.on("backoffice.category.created", (e) => {
  if (Object.hasOwn(e.detail, "backofficeCategoriesPath"))
    htmx.ajax("GET", e.detail.backofficeCategoriesPath, { target: "#modal-body" });
});

htmx.on("backoffice.categories.shown", (e) => {
  setModalSize("modal-lg");
});

htmx.on("backoffice.category.updated", (e) => {
  if (Object.hasOwn(e.detail, "backofficeCategoriesPath"))
    htmx.ajax("GET", e.detail.backofficeCategoriesPath, { target: "#modal-body" });
});
document.addEventListener("balance.item.initialized", (e) => {
  const element = document.querySelector("input[name=date]");

  const datepicker = new Datepicker(element, BOOTSTRAP_DATEPICKER_DEFAULTS);
});

document.addEventListener("balance.item.edit", (e) => {
  clearModalSize();

  const element = document.querySelector(".modal input[name=date]");

  const datepicker = new Datepicker(element, BOOTSTRAP_DATEPICKER_DEFAULTS);
});

document.addEventListener("balance.item.updated", async (e) => {
  hideModal();

  if (Object.hasOwn(e.detail, "balancePath"))
    await htmx.ajax("GET", e.detail.balancePath, { target: "#balance", swap: "outerHTML" });

  if (Object.hasOwn(e.detail, "categoriesPath"))
    await htmx.ajax("GET", e.detail.categoriesPath, { "target": "#categories", swap: "outerHTML" });
});

document.addEventListener("balance.items.shown", (e) => {
  const month = e.detail.month;
  const year = e.detail.year;

  const months = document.getElementById("months");

  for (const child of months.children) {
    child.classList.remove("active");

    if (child.dataset.number == month)
      child.classList.add("active");

    let url = new URL(child.getAttribute("hx-get"), window.location.origin);

    url.searchParams.set("year", year);

    child.setAttribute("hx-get", url.toString());

    htmx.process(child);
  }

  const years = document.getElementById("years");

  for (const child of years.children) {
    child.classList.remove("active");

    if (child.dataset.number == year)
      child.classList.add("active");

    let url = new URL(child.getAttribute("hx-get"), window.location.origin);

    url.searchParams.set("month", month);

    child.setAttribute("hx-get", url.toString());

    htmx.process(child);
  }
});

document.addEventListener("balance.item.create.error", (e) => {
  showModal();

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
