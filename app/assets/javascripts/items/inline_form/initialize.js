items.InlineForm.Initialize = function() {
  const form = document.querySelector('#item-inline-form form');

  const category_id = form.querySelector('#item_category_id');

  const url = '/categories/' + category_id.value + '/tags';

  category_id.addEventListener('change', function() { items.category.OnChange(form) });

  items.tags.Load(form, url, []);
}
