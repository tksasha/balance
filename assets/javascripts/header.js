export class Header {
  collection;
  className = "active";

  reset() {
    this.collection.forEach((m) => m.classList.remove(this.className));
  }

  init() {
    this.collection.forEach((m) =>
      m.addEventListener("click", (e) => {
        this.reset();

        this.onClick(m);
      }),
    );
  }

  onClick() {}

  static refresh() {
    const params = new URLSearchParams(window.location.search);

    params.set("month", this.#month());

    params.set("year", this.#year());

    htmx.ajax("GET", "/items?" + params.toString(), "#items");

    htmx.ajax("GET", "/categories?" + params.toString(), { target: "#categories", swap: "outerHTML" });
  }

  static #month() {
    return document.querySelector("#months div.active").dataset.number;
  }

  static #year() {
    return document.querySelector("#years div.active").dataset.number;
  }
}
