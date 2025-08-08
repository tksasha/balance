import { BOOTSTRAP_DATEPICKER_DEFAULTS, clearModalSize, hideModal, showModal } from "./_common";
import Datepicker from "../datepicker/js/Datepicker";

document.addEventListener("DOMContentLoaded", () => {
  Datepicker.locales.uk = {
    days: ["Неділя", "Понеділок", "Вівторок", "Середа", "Четвер", "П'ятниця", "Субота"],
    daysShort: ["Нед", "Пнд", "Втр", "Срд", "Чтв", "Птн", "Суб"],
    daysMin: ["Нд", "Пн", "Вт", "Ср", "Чт", "Пт", "Сб"],
    months: [
      "Січень",
      "Лютий",
      "Березень",
      "Квітень",
      "Травень",
      "Червень",
      "Липень",
      "Серпень",
      "Вересень",
      "Жовтень",
      "Листопад",
      "Грудень",
    ],
    monthsShort: ["Січ", "Лют", "Бер", "Кві", "Тра", "Чер", "Лип", "Сер", "Вер", "Жов", "Лис", "Гру"],
    today: "Сьогодні",
    clear: "Очистити",
    format: "dd.mm.yyyy",
    weekStart: 1,
  };
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

  if (Object.hasOwn(e.detail, "balancePath")) {
    await htmx.ajax("GET", e.detail.balancePath, { target: "#balance", swap: "outerHTML" });
  }

  if (Object.hasOwn(e.detail, "categoriesPath")) {
    await htmx.ajax("GET", e.detail.categoriesPath, { target: "#categories", swap: "outerHTML" });
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

  if (Object.hasOwn(e.detail, "balancePath")) {
    await htmx.ajax("GET", e.detail.balancePath, { target: "#balance", swap: "outerHTML" });
  }

  if (Object.hasOwn(e.detail, "categoriesPath")) {
    await htmx.ajax("GET", e.detail.categoriesPath, { target: "#categories", swap: "outerHTML" });
  }

  if (Object.hasOwn(e.detail, "itemsPath")) {
    await htmx.ajax("GET", e.detail.itemsPath, { target: "#items" });
  }
});

document.addEventListener("balance.item.deleted", async (e) => {
  hideModal();

  if (Object.hasOwn(e.detail, "balancePath")) {
    await htmx.ajax("GET", e.detail.balancePath, { target: "#balance", swap: "outerHTML" });
  }

  if (Object.hasOwn(e.detail, "categoriesPath")) {
    await htmx.ajax("GET", e.detail.categoriesPath, { target: "#categories", swap: "outerHTML" });
  }

  if (Object.hasOwn(e.detail, "itemsPath")) {
    await htmx.ajax("GET", e.detail.itemsPath, { target: "#items" });
  }
});
