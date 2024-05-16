$(function() {
  if(location.pathname != '/admin/items') {
    return false;
  }

  const currency = function() {
    const params = new URLSearchParams(location.search);
    const currency = params.get('scope') || 'uah';

    return {
      uah: 0,
      usd: 1,
      eur: 3,
    }[currency];
  }();

  $.getJSON(
    '/admin/categories',
    { q: { currency_eq: currency } },
    function(categories) {
      $.each(categories, function(idx, category) {
        $('#q_category_id').append('<option value="' + category.id+ '">' + category.name + '</option>');
      });

      (function() {
        const params = new URLSearchParams(location.search);

        $('#q_category_id').val(params.get('q[category_id_eq]'));
      }());
    }
  );
});
