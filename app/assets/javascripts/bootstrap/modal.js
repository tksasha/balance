bootstrap.Modal.Show = function() {
  const el = document.querySelector('.modal');

  document.
    querySelector('.modal-header').
    classList.
    add('d-none');

  document.
    querySelector('.modal-footer').
    classList.
    add('d-none');

  const modal = bootstrap.Modal.getOrCreateInstance(el);

  modal.show();
};
