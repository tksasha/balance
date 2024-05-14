class Money {
  constructor(value) {
    this.value = value;
  }

  Value() {
    return new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD'}).format(this.value);
  }

  toString() {
    return this.Value().replace('\$', '').replace(/,/g, ' ');
  }
}
