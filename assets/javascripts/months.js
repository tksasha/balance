import { Header } from "./header";
import { Years } from "./years";

export class Months extends Header {
  constructor() {
    super().collection = document.querySelectorAll("#months div");

    this.init();
  }

  init() {
    super.init();

    const date = new Date();

    this.collection.forEach((m) => {
      const monthNumber = parseInt(m.dataset.number);

      if (Number.isNaN(monthNumber)) {
        return;
      }

      if (date.getMonth() + 1 == monthNumber) {
        m.classList.add(this.className);
      }
    });
  }

  onClick(month) {
    month.classList.add(this.className);

    Header.refresh();
  }
}
