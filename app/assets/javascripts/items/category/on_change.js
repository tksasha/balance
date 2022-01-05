items.category.OnChange = function(form) {
  const category_id = form.querySelector('#item_category_id');

  const url = '/categories/' + category_id.value + '/tags';

  items.tags.Load(form, url, []);
};
