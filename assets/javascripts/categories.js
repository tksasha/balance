htmx.on("balance.month.changed", (e) => {
  if (Object.hasOwn(e.detail, "balanceCategoriesPath"))
    htmx.ajax("GET", e.detail.balanceCategoriesPath, { target: "#categories", swap: "outerHTML" });
});

htmx.on("balance.year.changed", (e) => {
  if (Object.hasOwn(e.detail, "balanceCategoriesPath"))
    htmx.ajax("GET", e.detail.balanceCategoriesPath, { target: "#categories", swap: "outerHTML" });
});
