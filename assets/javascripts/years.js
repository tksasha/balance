import { Header } from "./header";

export class Years extends Header {
  constructor() {
    super().collection = document.querySelectorAll("#years div");

    this.init();
  }

  init() {
    super.init();

    const date = new Date();

    this.collection.forEach((y) => {
      const yearNumber = parseInt(y.dataset.number);

      if (Number.isNaN(yearNumber)) {
        return;
      }

      if (date.getFullYear() == yearNumber) {
        y.classList.add(this.className);
      }
    });
  }

  onClick(year) {
    year.classList.add(this.className);

    Header.refresh();
  }
}
