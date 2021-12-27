items.tags.Load = function(url, tag_ids) {
  const render_tags_to_send = function() {
    let html = '';

    for(const tag of document.querySelectorAll('#tags-list .tag.active')) {
      html += '<input type="hidden" name="item[tag_ids][]" value="' + tag.dataset.id + '">'
    }

    if(0 == html.length) {
      html = '<input type="hidden" name="item[tag_ids][]">';
    }

    document.
      getElementById('tags-to-send').
      innerHTML = html;
  };

  const success = function(tags) {
    let html = '';

    for(const tag of tags) {
      let classes = 'tag';

      if(tag_ids.includes(tag.id)) {
        classes += ' active';
      }

      html += '<div data-id="' + tag.id + '" class="' + classes + '">' + tag.name + '</div>';
    }

    document.
      getElementById('tags-spinner').
      classList.
      add('d-none');

    document.
      getElementById('tags-list').
      innerHTML = html;

    render_tags_to_send();

    for(const tag of document.querySelectorAll('#tags-list .tag')) {
      tag.addEventListener('click', function() {
        this.classList.toggle('active');

        render_tags_to_send();
      });
    }
  };

  $.getJSON(url, success)
}
