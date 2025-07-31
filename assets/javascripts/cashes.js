import { clearModalSize } from "./_common";

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
