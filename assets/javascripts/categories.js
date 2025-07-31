import { setModalSize } from "./_common";

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
