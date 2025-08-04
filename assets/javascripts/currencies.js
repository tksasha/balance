export class Currencies {
  constructor() {
    const el = document.querySelector("#currencies-widget");

    bootstrap.Popover.getOrCreateInstance(el, {
      placement: "bottom",
      html: true,
      content: function () {
        return document.querySelector("#currencies-widget-content").innerHTML;
      },
    });
  }
}
