items.tags.Load = function(form, url, tag_ids) {
  const success = function(tags) {
    let html = '';

    for(const tag of tags) {
      let classes = 'tag';

      if(tag_ids.includes(tag.id)) {
        classes += ' active';
      }

      html += '<div data-id="' + tag.id + '" class="' + classes + '">' + tag.name + '</div>';
    }

    form.
      querySelector('#tags-spinner').
      classList.
      add('d-none');

    form.
      querySelector('#tags-list').
      innerHTML = html;

    items.tags.Render(form);

    items.tags.OnClick(form);
  };

  $.getJSON(url, success)
}
