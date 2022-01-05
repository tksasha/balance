items.tags.Render = function(form) {
  let html = '';

  for(const tag of form.querySelectorAll('#tags-list .tag.active')) {
    html += '<input type="hidden" name="item[tag_ids][]" value="' + tag.dataset.id + '">'
  }

  if(0 == html.length) {
    html = '<input type="hidden" name="item[tag_ids][]">';
  }

  form.
    querySelector('#tags-to-send').
    innerHTML = html;
};
