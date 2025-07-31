import "./backoffice";
import "./cashes";
import "./categories";
import "./items";

import { Months } from "./months";
import { Years } from "./years";
import { Header } from "./header";

document.addEventListener("DOMContentLoaded", () => {
  new Months();
  new Years();

  Header.refresh();
});
