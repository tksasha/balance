htmx.on("balance.cash.updated", (e) => {
  hideModal();

  if (Object.hasOwn(e.detail, "balancePath"))
    htmx.ajax("GET", e.detail.balancePath, { target: "#balance", swap: "outerHTML" });
});

htmx.on("balance.cash.edit", clearModalSize);

htmx.on("backoffice.cash.updated", (e) => {
  if (Object.hasOwn(e.detail, "balancePath"))
    htmx.ajax("GET", e.detail.balancePath, { target: "#balance" });

  if (Object.hasOwn(e.detail, "balanceCashesPath"))
    htmx.ajax("GET", e.detail.balanceCashesPath, { target: "#cashes" });
});
