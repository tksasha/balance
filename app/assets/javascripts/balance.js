$(function() {
  $('#balance').on('balance.changed', function() {
    const value = parseFloat(this.innerHTML);

    if(value > 0) {
      this.classList = ['sum fw-bold text-success'];
    } else if(value < 0) {
      this.classList = ['sum fw-bold text-danger'];
    } else {
      this.classList = ['sum'];
    }
  });
});
