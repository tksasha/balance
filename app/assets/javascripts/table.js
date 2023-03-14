class Table {
  constructor(element) {
    this.table = element;
  }

  summarize() {
    let sum = 0;

    (function(table) {
      const summary = table.querySelector('tr.summary');

      if(summary != null) {
        summary.remove();
      }
    }(this.table));

    if(this.table.rows.length <= 1) { return }

    for(let row of this.table.rows) {
      const value = (row.querySelector('td.sum a') || row.querySelector('td.sum')).innerHTML;

      sum += parseFloat(value.replace(/[^0-9\.]/, ''));
    }

    const tr = document.createElement('tr');

    tr.classList.add('summary');

    (function() {
      const td = document.createElement('td');

      td.innerHTML = 'ВСЬОГО';

      tr.appendChild(td);
    }());

    (function() {
      const td = document.createElement('td');

      td.innerHTML = new Money(sum);

      td.classList.add('sum');

      tr.appendChild(td);
    }());

    this.table.querySelector('tbody').appendChild(tr);
  }
}
