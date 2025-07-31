import { BOOTSTRAP_DATEPICKER_DEFAULTS, clearModalSize, hideModal, showModal } from "./_common";
import Datepicker from "../datepicker/js/Datepicker";

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

  const element = document.querySelector("input[name=date]");

  const datepicker = new Datepicker(element, BOOTSTRAP_DATEPICKER_DEFAULTS);

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
