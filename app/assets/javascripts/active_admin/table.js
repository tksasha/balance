class Table {
  constructor(element) {
    this.table = element;
  }

  summarize() {
    let sum = 0;

    for(let row of this.table.rows) {
      const value = row.querySelector('td.sum');

      sum += parseFloat(value.innerHTML.replace(/[^0-9\.]/, ''));
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
