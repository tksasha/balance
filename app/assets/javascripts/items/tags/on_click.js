items.tags.OnClick = function(form) {
  for(const tag of form.querySelectorAll('#tags-list .tag')) {
    tag.addEventListener('click', function() {
      this.classList.toggle('active');

      items.tags.Render(form);
    });
  }
};
