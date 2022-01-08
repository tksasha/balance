items.InlineForm.Reset = function() {
  const form = document.getElementById('item-inline-form');

  const event = new Event('reset');

  form.dispatchEvent(event);
}
