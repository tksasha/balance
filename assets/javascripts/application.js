import "./backoffice";
import "./cashes";
import "./categories";
import "./items";

import { Months } from "./months";
import { Years } from "./years";
import { Header } from "./header";
import { Currencies } from "./currencies";

document.addEventListener("DOMContentLoaded", () => {
  new Months();
  new Years();

  new Currencies();

  Header.refresh();
});
