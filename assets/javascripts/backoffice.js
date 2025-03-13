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
