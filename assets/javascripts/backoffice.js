document.addEventListener("backoffice.index.shown", (e) => {
  setModalSize("modal-sm");
});

document.addEventListener("backoffice.cashes.shown", (e) => {
  clearModalSize();
});
