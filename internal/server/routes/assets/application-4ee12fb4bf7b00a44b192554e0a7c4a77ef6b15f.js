(() => {
  // javascripts/_common.js
  var BOOTSTRAP_DATEPICKER_DEFAULTS = {
    format: "dd.mm.yyyy",
    autohide: true,
    language: "uk",
    todayHighlight: true
  };
  var clearModalSize = () => {
    document.querySelector("#modal .modal-dialog").classList.remove("modal-lg", "modal-sm", "modal-xl");
  };
  document.getElementById("modal").addEventListener("shown.bs.modal", (e) => {
    const input = e.target.querySelector("[autofocus]");
    if (input)
      input.focus();
  });

  // javascripts/backoffice.js
  document.addEventListener("backoffice.index.shown", (e) => {
    setModalSize("modal-sm");
  });
  document.addEventListener("backoffice.cashes.shown", (e) => {
    clearModalSize();
  });
  document.addEventListener("backoffice.cash.deleted", async (e) => {
    if (Object.hasOwn(e.detail, "backofficeCashesPath"))
      await htmx.ajax("GET", e.detail.backofficeCashesPath, { target: "#modal-body" });
  });
  document.addEventListener("backoffice.cash.created", async (e) => {
    if (Object.hasOwn(e.detail, "backofficeCashesPath"))
      await htmx.ajax("GET", e.detail.backofficeCashesPath, { target: "#modal-body" });
  });

  // javascripts/cashes.js
  htmx.on("balance.cash.updated", async (e) => {
    hideModal();
    if (Object.hasOwn(e.detail, "balancePath"))
      await htmx.ajax("GET", e.detail.balancePath, { target: "#balance", swap: "outerHTML" });
  });
  htmx.on("balance.cash.edit", clearModalSize);
  htmx.on("backoffice.cash.updated", async (e) => {
    if (Object.hasOwn(e.detail, "backofficeCashesPath"))
      await htmx.ajax("GET", e.detail.backofficeCashesPath, { target: "#modal-body" });
    if (Object.hasOwn(e.detail, "balancePath"))
      await htmx.ajax("GET", e.detail.balancePath, { target: "#balance-and-cashes-row" });
    if (Object.hasOwn(e.detail, "cashesPath"))
      await htmx.ajax("GET", e.detail.cashesPath, { target: "#balance-and-cashes-row", swap: "beforeend" });
  });

  // javascripts/categories.js
  htmx.on("balance.month.changed", (e) => {
    if (Object.hasOwn(e.detail, "balanceCategoriesPath"))
      htmx.ajax("GET", e.detail.balanceCategoriesPath, { target: "#categories", swap: "outerHTML" });
  });
  htmx.on("balance.year.changed", (e) => {
    if (Object.hasOwn(e.detail, "balanceCategoriesPath"))
      htmx.ajax("GET", e.detail.balanceCategoriesPath, { target: "#categories", swap: "outerHTML" });
  });
  htmx.on("backoffice.category.created", (e) => {
    if (Object.hasOwn(e.detail, "backofficeCategoriesPath"))
      htmx.ajax("GET", e.detail.backofficeCategoriesPath, { target: "#modal-body" });
  });
  htmx.on("backoffice.categories.shown", (e) => {
    setModalSize("modal-lg");
  });
  htmx.on("backoffice.category.updated", (e) => {
    if (Object.hasOwn(e.detail, "backofficeCategoriesPath"))
      htmx.ajax("GET", e.detail.backofficeCategoriesPath, { target: "#modal-body" });
  });

  // datepicker/js/lib/utils.js
  function lastItemOf(arr) {
    return arr[arr.length - 1];
  }
  function pushUnique(arr, ...items) {
    items.forEach((item) => {
      if (arr.includes(item)) {
        return;
      }
      arr.push(item);
    });
    return arr;
  }
  function stringToArray(str, separator) {
    return str ? str.split(separator) : [];
  }
  function isInRange(testVal, min, max) {
    const minOK = min === void 0 || testVal >= min;
    const maxOK = max === void 0 || testVal <= max;
    return minOK && maxOK;
  }
  function limitToRange(val, min, max) {
    if (val < min) {
      return min;
    }
    if (val > max) {
      return max;
    }
    return val;
  }
  function createTagRepeat(tagName, repeat, attributes = {}, index = 0, html = "") {
    const openTagSrc = Object.keys(attributes).reduce((src, attr) => {
      let val = attributes[attr];
      if (typeof val === "function") {
        val = val(index);
      }
      return `${src} ${attr}="${val}"`;
    }, tagName);
    html += `<${openTagSrc}></${tagName}>`;
    const next = index + 1;
    return next < repeat ? createTagRepeat(tagName, repeat, attributes, next, html) : html;
  }
  function optimizeTemplateHTML(html) {
    return html.replace(/>\s+/g, ">").replace(/\s+</, "<");
  }

  // datepicker/js/lib/date.js
  function stripTime(timeValue) {
    return new Date(timeValue).setHours(0, 0, 0, 0);
  }
  function today() {
    return (/* @__PURE__ */ new Date()).setHours(0, 0, 0, 0);
  }
  function dateValue(...args) {
    switch (args.length) {
      case 0:
        return today();
      case 1:
        return stripTime(args[0]);
    }
    const newDate = /* @__PURE__ */ new Date(0);
    newDate.setFullYear(...args);
    return newDate.setHours(0, 0, 0, 0);
  }
  function addDays(date, amount) {
    const newDate = new Date(date);
    return newDate.setDate(newDate.getDate() + amount);
  }
  function addWeeks(date, amount) {
    return addDays(date, amount * 7);
  }
  function addMonths(date, amount) {
    const newDate = new Date(date);
    const monthsToSet = newDate.getMonth() + amount;
    let expectedMonth = monthsToSet % 12;
    if (expectedMonth < 0) {
      expectedMonth += 12;
    }
    const time = newDate.setMonth(monthsToSet);
    return newDate.getMonth() !== expectedMonth ? newDate.setDate(0) : time;
  }
  function addYears(date, amount) {
    const newDate = new Date(date);
    const expectedMonth = newDate.getMonth();
    const time = newDate.setFullYear(newDate.getFullYear() + amount);
    return expectedMonth === 1 && newDate.getMonth() === 2 ? newDate.setDate(0) : time;
  }
  function dayDiff(day, from) {
    return (day - from + 7) % 7;
  }
  function dayOfTheWeekOf(baseDate, dayOfWeek, weekStart = 0) {
    const baseDay = new Date(baseDate).getDay();
    return addDays(baseDate, dayDiff(dayOfWeek, weekStart) - dayDiff(baseDay, weekStart));
  }
  function calcWeekNum(dayOfTheWeek, sameDayOfFirstWeek) {
    return Math.round((dayOfTheWeek - sameDayOfFirstWeek) / 6048e5) + 1;
  }
  function getIsoWeek(date) {
    const thuOfTheWeek = dayOfTheWeekOf(date, 4, 1);
    const firstThu = dayOfTheWeekOf(new Date(thuOfTheWeek).setMonth(0, 4), 4, 1);
    return calcWeekNum(thuOfTheWeek, firstThu);
  }
  function calcTraditionalWeekNumber(date, weekStart) {
    const startOfFirstWeek = dayOfTheWeekOf(new Date(date).setMonth(0, 1), weekStart, weekStart);
    const startOfTheWeek = dayOfTheWeekOf(date, weekStart, weekStart);
    const weekNum = calcWeekNum(startOfTheWeek, startOfFirstWeek);
    if (weekNum < 53) {
      return weekNum;
    }
    const weekOneOfNextYear = dayOfTheWeekOf(new Date(date).setDate(32), weekStart, weekStart);
    return startOfTheWeek === weekOneOfNextYear ? 1 : weekNum;
  }
  function getWesternTradWeek(date) {
    return calcTraditionalWeekNumber(date, 0);
  }
  function getMidEasternWeek(date) {
    return calcTraditionalWeekNumber(date, 6);
  }
  function startOfYearPeriod(date, years) {
    const year = new Date(date).getFullYear();
    return Math.floor(year / years) * years;
  }
  function regularizeDate(date, timeSpan, useLastDate) {
    if (timeSpan !== 1 && timeSpan !== 2) {
      return date;
    }
    const newDate = new Date(date);
    if (timeSpan === 1) {
      useLastDate ? newDate.setMonth(newDate.getMonth() + 1, 0) : newDate.setDate(1);
    } else {
      useLastDate ? newDate.setFullYear(newDate.getFullYear() + 1, 0, 0) : newDate.setMonth(0, 1);
    }
    return newDate.setHours(0, 0, 0, 0);
  }

  // datepicker/js/lib/date-format.js
  var reFormatTokens = /dd?|DD?|mm?|MM?|yy?(?:yy)?/;
  var reNonDateParts = /[\s!-/:-@[-`{-~年月日]+/;
  var knownFormats = {};
  var parseFns = {
    y(date, year) {
      return new Date(date).setFullYear(parseInt(year, 10));
    },
    m(date, month, locale) {
      const newDate = new Date(date);
      let monthIndex = parseInt(month, 10) - 1;
      if (isNaN(monthIndex)) {
        if (!month) {
          return NaN;
        }
        const monthName = month.toLowerCase();
        const compareNames = (name) => name.toLowerCase().startsWith(monthName);
        monthIndex = locale.monthsShort.findIndex(compareNames);
        if (monthIndex < 0) {
          monthIndex = locale.months.findIndex(compareNames);
        }
        if (monthIndex < 0) {
          return NaN;
        }
      }
      newDate.setMonth(monthIndex);
      return newDate.getMonth() !== normalizeMonth(monthIndex) ? newDate.setDate(0) : newDate.getTime();
    },
    d(date, day) {
      return new Date(date).setDate(parseInt(day, 10));
    }
  };
  var formatFns = {
    d(date) {
      return date.getDate();
    },
    dd(date) {
      return padZero(date.getDate(), 2);
    },
    D(date, locale) {
      return locale.daysShort[date.getDay()];
    },
    DD(date, locale) {
      return locale.days[date.getDay()];
    },
    m(date) {
      return date.getMonth() + 1;
    },
    mm(date) {
      return padZero(date.getMonth() + 1, 2);
    },
    M(date, locale) {
      return locale.monthsShort[date.getMonth()];
    },
    MM(date, locale) {
      return locale.months[date.getMonth()];
    },
    y(date) {
      return date.getFullYear();
    },
    yy(date) {
      return padZero(date.getFullYear(), 2).slice(-2);
    },
    yyyy(date) {
      return padZero(date.getFullYear(), 4);
    }
  };
  function normalizeMonth(monthIndex) {
    return monthIndex > -1 ? monthIndex % 12 : normalizeMonth(monthIndex + 12);
  }
  function padZero(num, length) {
    return num.toString().padStart(length, "0");
  }
  function parseFormatString(format) {
    if (typeof format !== "string") {
      throw new Error("Invalid date format.");
    }
    if (format in knownFormats) {
      return knownFormats[format];
    }
    const separators = format.split(reFormatTokens);
    const parts = format.match(new RegExp(reFormatTokens, "g"));
    if (separators.length === 0 || !parts) {
      throw new Error("Invalid date format.");
    }
    const partFormatters = parts.map((token) => formatFns[token]);
    const partParserKeys = Object.keys(parseFns).reduce((keys, key) => {
      const token = parts.find((part) => part[0] !== "D" && part[0].toLowerCase() === key);
      if (token) {
        keys.push(key);
      }
      return keys;
    }, []);
    return knownFormats[format] = {
      parser(dateStr, locale) {
        const dateParts = dateStr.split(reNonDateParts).reduce((dtParts, part, index) => {
          if (part.length > 0 && parts[index]) {
            const token = parts[index][0];
            if (token === "M") {
              dtParts.m = part;
            } else if (token !== "D") {
              dtParts[token] = part;
            }
          }
          return dtParts;
        }, {});
        return partParserKeys.reduce((origDate, key) => {
          const newDate = parseFns[key](origDate, dateParts[key], locale);
          return isNaN(newDate) ? origDate : newDate;
        }, today());
      },
      formatter(date, locale) {
        let dateStr = partFormatters.reduce((str, fn, index) => {
          return str += `${separators[index]}${fn(date, locale)}`;
        }, "");
        return dateStr += lastItemOf(separators);
      }
    };
  }
  function parseDate(dateStr, format, locale) {
    if (dateStr instanceof Date || typeof dateStr === "number") {
      const date = stripTime(dateStr);
      return isNaN(date) ? void 0 : date;
    }
    if (!dateStr) {
      return void 0;
    }
    if (dateStr === "today") {
      return today();
    }
    if (format && format.toValue) {
      const date = format.toValue(dateStr, format, locale);
      return isNaN(date) ? void 0 : stripTime(date);
    }
    return parseFormatString(format).parser(dateStr, locale);
  }
  function formatDate(date, format, locale) {
    if (isNaN(date) || !date && date !== 0) {
      return "";
    }
    const dateObj = typeof date === "number" ? new Date(date) : date;
    if (format.toDisplay) {
      return format.toDisplay(dateObj, format, locale);
    }
    return parseFormatString(format).formatter(dateObj, locale);
  }

  // datepicker/js/lib/dom.js
  var range = document.createRange();
  function parseHTML(html) {
    return range.createContextualFragment(html);
  }
  function getParent(el) {
    return el.parentElement || (el.parentNode instanceof ShadowRoot ? el.parentNode.host : void 0);
  }
  function isActiveElement(el) {
    return el.getRootNode().activeElement === el;
  }
  function hideElement(el) {
    if (el.style.display === "none") {
      return;
    }
    if (el.style.display) {
      el.dataset.styleDisplay = el.style.display;
    }
    el.style.display = "none";
  }
  function showElement(el) {
    if (el.style.display !== "none") {
      return;
    }
    if (el.dataset.styleDisplay) {
      el.style.display = el.dataset.styleDisplay;
      delete el.dataset.styleDisplay;
    } else {
      el.style.display = "";
    }
  }
  function emptyChildNodes(el) {
    if (el.firstChild) {
      el.removeChild(el.firstChild);
      emptyChildNodes(el);
    }
  }
  function replaceChildNodes(el, newChildNodes) {
    emptyChildNodes(el);
    if (newChildNodes instanceof DocumentFragment) {
      el.appendChild(newChildNodes);
    } else if (typeof newChildNodes === "string") {
      el.appendChild(parseHTML(newChildNodes));
    } else if (typeof newChildNodes.forEach === "function") {
      newChildNodes.forEach((node) => {
        el.appendChild(node);
      });
    }
  }

  // datepicker/js/lib/event.js
  var listenerRegistry = /* @__PURE__ */ new WeakMap();
  var { addEventListener, removeEventListener } = EventTarget.prototype;
  function registerListeners(keyObj, listeners) {
    let registered = listenerRegistry.get(keyObj);
    if (!registered) {
      registered = [];
      listenerRegistry.set(keyObj, registered);
    }
    listeners.forEach((listener) => {
      addEventListener.call(...listener);
      registered.push(listener);
    });
  }
  function unregisterListeners(keyObj) {
    let listeners = listenerRegistry.get(keyObj);
    if (!listeners) {
      return;
    }
    listeners.forEach((listener) => {
      removeEventListener.call(...listener);
    });
    listenerRegistry.delete(keyObj);
  }
  if (!Event.prototype.composedPath) {
    const getComposedPath = (node, path = []) => {
      path.push(node);
      let parent;
      if (node.parentNode) {
        parent = node.parentNode;
      } else if (node.host) {
        parent = node.host;
      } else if (node.defaultView) {
        parent = node.defaultView;
      }
      return parent ? getComposedPath(parent, path) : path;
    };
    Event.prototype.composedPath = function() {
      return getComposedPath(this.target);
    };
  }
  function findFromPath(path, criteria, currentTarget) {
    const [node, ...rest] = path;
    if (criteria(node)) {
      return node;
    }
    if (node === currentTarget || node.tagName === "HTML" || rest.length === 0) {
      return;
    }
    return findFromPath(rest, criteria, currentTarget);
  }
  function findElementInEventPath(ev, selector) {
    const criteria = typeof selector === "function" ? selector : (el) => el instanceof Element && el.matches(selector);
    return findFromPath(ev.composedPath(), criteria, ev.currentTarget);
  }

  // datepicker/js/i18n/base-locales.js
  var base_locales_default = {
    en: {
      days: ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"],
      daysShort: ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"],
      daysMin: ["Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"],
      months: ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"],
      monthsShort: ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"],
      today: "Today",
      clear: "Clear",
      titleFormat: "MM y"
    }
  };

  // datepicker/js/options/defaultOptions.js
  var defaultOptions_default = {
    autohide: false,
    beforeShowDay: null,
    beforeShowDecade: null,
    beforeShowMonth: null,
    beforeShowYear: null,
    clearButton: false,
    dateDelimiter: ",",
    datesDisabled: [],
    daysOfWeekDisabled: [],
    daysOfWeekHighlighted: [],
    defaultViewDate: void 0,
    // placeholder, defaults to today() by the program
    disableTouchKeyboard: false,
    enableOnReadonly: true,
    format: "mm/dd/yyyy",
    language: "en",
    maxDate: null,
    maxNumberOfDates: 1,
    maxView: 3,
    minDate: null,
    nextArrow: "\xBB",
    orientation: "auto",
    pickLevel: 0,
    prevArrow: "\xAB",
    showDaysOfWeek: true,
    showOnClick: true,
    showOnFocus: true,
    startView: 0,
    title: "",
    todayButton: false,
    todayButtonMode: 0,
    todayHighlight: false,
    updateOnBlur: true,
    weekNumbers: 0,
    weekStart: 0
  };

  // datepicker/js/options/processOptions.js
  var {
    language: defaultLang,
    format: defaultFormat,
    weekStart: defaultWeekStart
  } = defaultOptions_default;
  function sanitizeDOW(dow, day) {
    return dow.length < 6 && day >= 0 && day < 7 ? pushUnique(dow, day) : dow;
  }
  function determineGetWeekMethod(numberingMode, weekStart) {
    const methodId = numberingMode === 4 ? weekStart === 6 ? 3 : !weekStart + 1 : numberingMode;
    switch (methodId) {
      case 1:
        return getIsoWeek;
      case 2:
        return getWesternTradWeek;
      case 3:
        return getMidEasternWeek;
    }
  }
  function updateWeekStart(newValue, config, weekNumbers) {
    config.weekStart = newValue;
    config.weekEnd = (newValue + 6) % 7;
    if (weekNumbers === 4) {
      config.getWeekNumber = determineGetWeekMethod(4, newValue);
    }
    return newValue;
  }
  function validateDate(value, format, locale, origValue) {
    const date = parseDate(value, format, locale);
    return date !== void 0 ? date : origValue;
  }
  function validateViewId(value, origValue, max = 3) {
    const viewId = parseInt(value, 10);
    return viewId >= 0 && viewId <= max ? viewId : origValue;
  }
  function replaceOptions(options, from, to, convert = void 0) {
    if (from in options) {
      if (!(to in options)) {
        options[to] = convert ? convert(options[from]) : options[from];
      }
      delete options[from];
    }
  }
  function processOptions(options, datepicker) {
    const inOpts = Object.assign({}, options);
    const config = {};
    const locales = datepicker.constructor.locales;
    const rangeEnd = !!datepicker.rangeSideIndex;
    let {
      datesDisabled,
      format,
      language,
      locale,
      maxDate,
      maxView,
      minDate,
      pickLevel,
      startView,
      weekNumbers,
      weekStart
    } = datepicker.config || {};
    replaceOptions(inOpts, "calendarWeeks", "weekNumbers", (val) => val ? 1 : 0);
    replaceOptions(inOpts, "clearBtn", "clearButton");
    replaceOptions(inOpts, "todayBtn", "todayButton");
    replaceOptions(inOpts, "todayBtnMode", "todayButtonMode");
    if (inOpts.language) {
      let lang;
      if (inOpts.language !== language) {
        if (locales[inOpts.language]) {
          lang = inOpts.language;
        } else {
          lang = inOpts.language.split("-")[0];
          if (!locales[lang]) {
            lang = false;
          }
        }
      }
      delete inOpts.language;
      if (lang) {
        language = config.language = lang;
        const origLocale = locale || locales[defaultLang];
        locale = Object.assign({
          format: defaultFormat,
          weekStart: defaultWeekStart
        }, locales[defaultLang]);
        if (language !== defaultLang) {
          Object.assign(locale, locales[language]);
        }
        config.locale = locale;
        if (format === origLocale.format) {
          format = config.format = locale.format;
        }
        if (weekStart === origLocale.weekStart) {
          weekStart = updateWeekStart(locale.weekStart, config, weekNumbers);
        }
      }
    }
    if (inOpts.format) {
      const hasToDisplay = typeof inOpts.format.toDisplay === "function";
      const hasToValue = typeof inOpts.format.toValue === "function";
      const validFormatString = reFormatTokens.test(inOpts.format);
      if (hasToDisplay && hasToValue || validFormatString) {
        format = config.format = inOpts.format;
      }
      delete inOpts.format;
    }
    let newPickLevel = pickLevel;
    if ("pickLevel" in inOpts) {
      newPickLevel = validateViewId(inOpts.pickLevel, pickLevel, 2);
      delete inOpts.pickLevel;
    }
    if (newPickLevel !== pickLevel) {
      if (newPickLevel > pickLevel) {
        if (!("minDate" in inOpts)) {
          inOpts.minDate = minDate;
        }
        if (!("maxDate" in inOpts)) {
          inOpts.maxDate = maxDate;
        }
      }
      if (datesDisabled && !inOpts.datesDisabled) {
        inOpts.datesDisabled = [];
      }
      pickLevel = config.pickLevel = newPickLevel;
    }
    let minDt = minDate;
    let maxDt = maxDate;
    if ("minDate" in inOpts) {
      const defaultMinDt = dateValue(0, 0, 1);
      minDt = inOpts.minDate === null ? defaultMinDt : validateDate(inOpts.minDate, format, locale, minDt);
      if (minDt !== defaultMinDt) {
        minDt = regularizeDate(minDt, pickLevel, false);
      }
      delete inOpts.minDate;
    }
    if ("maxDate" in inOpts) {
      maxDt = inOpts.maxDate === null ? void 0 : validateDate(inOpts.maxDate, format, locale, maxDt);
      if (maxDt !== void 0) {
        maxDt = regularizeDate(maxDt, pickLevel, true);
      }
      delete inOpts.maxDate;
    }
    if (maxDt < minDt) {
      minDate = config.minDate = maxDt;
      maxDate = config.maxDate = minDt;
    } else {
      if (minDate !== minDt) {
        minDate = config.minDate = minDt;
      }
      if (maxDate !== maxDt) {
        maxDate = config.maxDate = maxDt;
      }
    }
    if (inOpts.datesDisabled) {
      const dtsDisabled = inOpts.datesDisabled;
      if (typeof dtsDisabled === "function") {
        config.datesDisabled = null;
        config.checkDisabled = (timeValue, viewId) => dtsDisabled(
          new Date(timeValue),
          viewId,
          rangeEnd
        );
      } else {
        const disabled = config.datesDisabled = dtsDisabled.reduce((dates, dt) => {
          const date = parseDate(dt, format, locale);
          return date !== void 0 ? pushUnique(dates, regularizeDate(date, pickLevel, rangeEnd)) : dates;
        }, []);
        config.checkDisabled = (timeValue) => disabled.includes(timeValue);
      }
      delete inOpts.datesDisabled;
    }
    if ("defaultViewDate" in inOpts) {
      const viewDate = parseDate(inOpts.defaultViewDate, format, locale);
      if (viewDate !== void 0) {
        config.defaultViewDate = viewDate;
      }
      delete inOpts.defaultViewDate;
    }
    if ("weekStart" in inOpts) {
      const wkStart = Number(inOpts.weekStart) % 7;
      if (!isNaN(wkStart)) {
        weekStart = updateWeekStart(wkStart, config, weekNumbers);
      }
      delete inOpts.weekStart;
    }
    if (inOpts.daysOfWeekDisabled) {
      config.daysOfWeekDisabled = inOpts.daysOfWeekDisabled.reduce(sanitizeDOW, []);
      delete inOpts.daysOfWeekDisabled;
    }
    if (inOpts.daysOfWeekHighlighted) {
      config.daysOfWeekHighlighted = inOpts.daysOfWeekHighlighted.reduce(sanitizeDOW, []);
      delete inOpts.daysOfWeekHighlighted;
    }
    if ("weekNumbers" in inOpts) {
      let method = inOpts.weekNumbers;
      if (method) {
        const getWeekNumber = typeof method === "function" ? (timeValue, startOfWeek) => method(new Date(timeValue), startOfWeek) : determineGetWeekMethod(method = parseInt(method, 10), weekStart);
        if (getWeekNumber) {
          weekNumbers = config.weekNumbers = method;
          config.getWeekNumber = getWeekNumber;
        }
      } else {
        weekNumbers = config.weekNumbers = 0;
        config.getWeekNumber = null;
      }
      delete inOpts.weekNumbers;
    }
    if ("maxNumberOfDates" in inOpts) {
      const maxNumberOfDates = parseInt(inOpts.maxNumberOfDates, 10);
      if (maxNumberOfDates >= 0) {
        config.maxNumberOfDates = maxNumberOfDates;
        config.multidate = maxNumberOfDates !== 1;
      }
      delete inOpts.maxNumberOfDates;
    }
    if (inOpts.dateDelimiter) {
      config.dateDelimiter = String(inOpts.dateDelimiter);
      delete inOpts.dateDelimiter;
    }
    let newMaxView = maxView;
    if ("maxView" in inOpts) {
      newMaxView = validateViewId(inOpts.maxView, maxView);
      delete inOpts.maxView;
    }
    newMaxView = pickLevel > newMaxView ? pickLevel : newMaxView;
    if (newMaxView !== maxView) {
      maxView = config.maxView = newMaxView;
    }
    let newStartView = startView;
    if ("startView" in inOpts) {
      newStartView = validateViewId(inOpts.startView, newStartView);
      delete inOpts.startView;
    }
    if (newStartView < pickLevel) {
      newStartView = pickLevel;
    } else if (newStartView > maxView) {
      newStartView = maxView;
    }
    if (newStartView !== startView) {
      config.startView = newStartView;
    }
    if (inOpts.prevArrow) {
      const prevArrow = parseHTML(inOpts.prevArrow);
      if (prevArrow.childNodes.length > 0) {
        config.prevArrow = prevArrow.childNodes;
      }
      delete inOpts.prevArrow;
    }
    if (inOpts.nextArrow) {
      const nextArrow = parseHTML(inOpts.nextArrow);
      if (nextArrow.childNodes.length > 0) {
        config.nextArrow = nextArrow.childNodes;
      }
      delete inOpts.nextArrow;
    }
    if ("disableTouchKeyboard" in inOpts) {
      config.disableTouchKeyboard = "ontouchstart" in document && !!inOpts.disableTouchKeyboard;
      delete inOpts.disableTouchKeyboard;
    }
    if (inOpts.orientation) {
      const orientation = inOpts.orientation.toLowerCase().split(/\s+/g);
      config.orientation = {
        x: orientation.find((x) => x === "left" || x === "right") || "auto",
        y: orientation.find((y) => y === "top" || y === "bottom") || "auto"
      };
      delete inOpts.orientation;
    }
    if ("todayButtonMode" in inOpts) {
      switch (inOpts.todayButtonMode) {
        case 0:
        case 1:
          config.todayButtonMode = inOpts.todayButtonMode;
      }
      delete inOpts.todayButtonMode;
    }
    Object.entries(inOpts).forEach(([key, value]) => {
      if (value !== void 0 && key in defaultOptions_default) {
        config[key] = value;
      }
    });
    return config;
  }

  // datepicker/js/options/shortcutKeys.js
  var defaultShortcutKeys = {
    show: { key: "ArrowDown" },
    hide: null,
    toggle: { key: "Escape" },
    prevButton: { key: "ArrowLeft", ctrlOrMetaKey: true },
    nextButton: { key: "ArrowRight", ctrlOrMetaKey: true },
    viewSwitch: { key: "ArrowUp", ctrlOrMetaKey: true },
    clearButton: { key: "Backspace", ctrlOrMetaKey: true },
    todayButton: { key: ".", ctrlOrMetaKey: true },
    exitEditMode: { key: "ArrowDown", ctrlOrMetaKey: true }
  };
  function createShortcutKeyConfig(options) {
    return Object.keys(defaultShortcutKeys).reduce((keyDefs, shortcut) => {
      const keyDef = options[shortcut] === void 0 ? defaultShortcutKeys[shortcut] : options[shortcut];
      const key = keyDef && keyDef.key;
      if (!key || typeof key !== "string") {
        return keyDefs;
      }
      const normalizedDef = {
        key,
        ctrlOrMetaKey: !!(keyDef.ctrlOrMetaKey || keyDef.ctrlKey || keyDef.metaKey)
      };
      if (key.length > 1) {
        normalizedDef.altKey = !!keyDef.altKey;
        normalizedDef.shiftKey = !!keyDef.shiftKey;
      }
      keyDefs[shortcut] = normalizedDef;
      return keyDefs;
    }, {});
  }

  // datepicker/js/picker/templates/pickerTemplate.js
  var getButtons = (buttonList) => buttonList.map((classes) => `<button type="button" class="%buttonClass% ${classes}" tabindex="-1"></button>`).join("");
  var pickerTemplate_default = optimizeTemplateHTML(`<div class="datepicker">
  <div class="datepicker-picker">
    <div class="datepicker-header">
      <div class="datepicker-title"></div>
      <div class="datepicker-controls">
        ${getButtons([
    "prev-button prev-btn",
    "view-switch",
    "next-button next-btn"
  ])}
      </div>
    </div>
    <div class="datepicker-main"></div>
    <div class="datepicker-footer">
      <div class="datepicker-controls">
        ${getButtons([
    "today-button today-btn",
    "clear-button clear-btn"
  ])}
      </div>
    </div>
  </div>
</div>`);

  // datepicker/js/picker/templates/daysTemplate.js
  var daysTemplate_default = optimizeTemplateHTML(`<div class="days">
  <div class="days-of-week">${createTagRepeat("span", 7, { class: "dow" })}</div>
  <div class="datepicker-grid">${createTagRepeat("span", 42)}</div>
</div>`);

  // datepicker/js/picker/templates/weekNumbersTemplate.js
  var weekNumbersTemplate_default = optimizeTemplateHTML(`<div class="week-numbers calendar-weeks">
  <div class="days-of-week"><span class="dow"></span></div>
  <div class="weeks">${createTagRepeat("span", 6, { class: "week" })}</div>
</div>`);

  // datepicker/js/picker/views/View.js
  var View = class {
    constructor(picker, config) {
      Object.assign(this, config, {
        picker,
        element: parseHTML(`<div class="datepicker-view"></div>`).firstChild,
        selected: [],
        isRangeEnd: !!picker.datepicker.rangeSideIndex
      });
      this.init(this.picker.datepicker.config);
    }
    init(options) {
      if ("pickLevel" in options) {
        this.isMinView = this.id === options.pickLevel;
      }
      this.setOptions(options);
      this.updateFocus();
      this.updateSelection();
    }
    prepareForRender(switchLabel, prevButtonDisabled, nextButtonDisabled) {
      this.disabled = [];
      const picker = this.picker;
      picker.setViewSwitchLabel(switchLabel);
      picker.setPrevButtonDisabled(prevButtonDisabled);
      picker.setNextButtonDisabled(nextButtonDisabled);
    }
    setDisabled(date, classList) {
      classList.add("disabled");
      pushUnique(this.disabled, date);
    }
    // Execute beforeShow() callback and apply the result to the element
    // args:
    performBeforeHook(el, timeValue) {
      let result = this.beforeShow(new Date(timeValue));
      switch (typeof result) {
        case "boolean":
          result = { enabled: result };
          break;
        case "string":
          result = { classes: result };
      }
      if (result) {
        const classList = el.classList;
        if (result.enabled === false) {
          this.setDisabled(timeValue, classList);
        }
        if (result.classes) {
          const extraClasses = result.classes.split(/\s+/);
          classList.add(...extraClasses);
          if (extraClasses.includes("disabled")) {
            this.setDisabled(timeValue, classList);
          }
        }
        if (result.content) {
          replaceChildNodes(el, result.content);
        }
      }
    }
    renderCell(el, content, cellVal, date, { selected, range: range2 }, outOfScope, extraClasses = []) {
      el.textContent = content;
      if (this.isMinView) {
        el.dataset.date = date;
      }
      const classList = el.classList;
      el.className = `datepicker-cell ${this.cellClass}`;
      if (cellVal < this.first) {
        classList.add("prev");
      } else if (cellVal > this.last) {
        classList.add("next");
      }
      classList.add(...extraClasses);
      if (outOfScope || this.checkDisabled(date, this.id)) {
        this.setDisabled(date, classList);
      }
      if (range2) {
        const [rangeStart, rangeEnd] = range2;
        if (cellVal > rangeStart && cellVal < rangeEnd) {
          classList.add("range");
        }
        if (cellVal === rangeStart) {
          classList.add("range-start");
        }
        if (cellVal === rangeEnd) {
          classList.add("range-end");
        }
      }
      if (selected.includes(cellVal)) {
        classList.add("selected");
      }
      if (cellVal === this.focused) {
        classList.add("focused");
      }
      if (this.beforeShow) {
        this.performBeforeHook(el, date);
      }
    }
    refreshCell(el, cellVal, selected, [rangeStart, rangeEnd]) {
      const classList = el.classList;
      classList.remove("range", "range-start", "range-end", "selected", "focused");
      if (cellVal > rangeStart && cellVal < rangeEnd) {
        classList.add("range");
      }
      if (cellVal === rangeStart) {
        classList.add("range-start");
      }
      if (cellVal === rangeEnd) {
        classList.add("range-end");
      }
      if (selected.includes(cellVal)) {
        classList.add("selected");
      }
      if (cellVal === this.focused) {
        classList.add("focused");
      }
    }
    changeFocusedCell(cellIndex) {
      this.grid.querySelectorAll(".focused").forEach((el) => {
        el.classList.remove("focused");
      });
      this.grid.children[cellIndex].classList.add("focused");
    }
  };

  // datepicker/js/picker/views/DaysView.js
  var DaysView = class extends View {
    constructor(picker) {
      super(picker, {
        id: 0,
        name: "days",
        cellClass: "day"
      });
    }
    init(options, onConstruction = true) {
      if (onConstruction) {
        const inner = parseHTML(daysTemplate_default).firstChild;
        this.dow = inner.firstChild;
        this.grid = inner.lastChild;
        this.element.appendChild(inner);
      }
      super.init(options);
    }
    setOptions(options) {
      let updateDOW;
      if ("minDate" in options) {
        this.minDate = options.minDate;
      }
      if ("maxDate" in options) {
        this.maxDate = options.maxDate;
      }
      if (options.checkDisabled) {
        this.checkDisabled = options.checkDisabled;
      }
      if (options.daysOfWeekDisabled) {
        this.daysOfWeekDisabled = options.daysOfWeekDisabled;
        updateDOW = true;
      }
      if (options.daysOfWeekHighlighted) {
        this.daysOfWeekHighlighted = options.daysOfWeekHighlighted;
      }
      if ("todayHighlight" in options) {
        this.todayHighlight = options.todayHighlight;
      }
      if ("weekStart" in options) {
        this.weekStart = options.weekStart;
        this.weekEnd = options.weekEnd;
        updateDOW = true;
      }
      if (options.locale) {
        const locale = this.locale = options.locale;
        this.dayNames = locale.daysMin;
        this.switchLabelFormat = locale.titleFormat;
        updateDOW = true;
      }
      if ("beforeShowDay" in options) {
        this.beforeShow = typeof options.beforeShowDay === "function" ? options.beforeShowDay : void 0;
      }
      if ("weekNumbers" in options) {
        if (options.weekNumbers && !this.weekNumbers) {
          const weeksElem = parseHTML(weekNumbersTemplate_default).firstChild;
          this.weekNumbers = {
            element: weeksElem,
            dow: weeksElem.firstChild,
            weeks: weeksElem.lastChild
          };
          this.element.insertBefore(weeksElem, this.element.firstChild);
        } else if (this.weekNumbers && !options.weekNumbers) {
          this.element.removeChild(this.weekNumbers.element);
          this.weekNumbers = null;
        }
      }
      if ("getWeekNumber" in options) {
        this.getWeekNumber = options.getWeekNumber;
      }
      if ("showDaysOfWeek" in options) {
        if (options.showDaysOfWeek) {
          showElement(this.dow);
          if (this.weekNumbers) {
            showElement(this.weekNumbers.dow);
          }
        } else {
          hideElement(this.dow);
          if (this.weekNumbers) {
            hideElement(this.weekNumbers.dow);
          }
        }
      }
      if (updateDOW) {
        Array.from(this.dow.children).forEach((el, index) => {
          const dow = (this.weekStart + index) % 7;
          el.textContent = this.dayNames[dow];
          el.className = this.daysOfWeekDisabled.includes(dow) ? "dow disabled" : "dow";
        });
      }
    }
    // Apply update on the focused date to view's settings
    updateFocus() {
      const viewDate = new Date(this.picker.viewDate);
      const viewYear = viewDate.getFullYear();
      const viewMonth = viewDate.getMonth();
      const firstOfMonth = dateValue(viewYear, viewMonth, 1);
      const start = dayOfTheWeekOf(firstOfMonth, this.weekStart, this.weekStart);
      this.first = firstOfMonth;
      this.last = dateValue(viewYear, viewMonth + 1, 0);
      this.start = start;
      this.focused = this.picker.viewDate;
    }
    // Apply update on the selected dates to view's settings
    updateSelection() {
      const { dates, rangepicker } = this.picker.datepicker;
      this.selected = dates;
      if (rangepicker) {
        this.range = rangepicker.dates;
      }
    }
    // Update the entire view UI
    render() {
      this.today = this.todayHighlight ? today() : void 0;
      this.prepareForRender(
        formatDate(this.focused, this.switchLabelFormat, this.locale),
        this.first <= this.minDate,
        this.last >= this.maxDate
      );
      if (this.weekNumbers) {
        const weekStart = this.weekStart;
        const startOfWeek = dayOfTheWeekOf(this.first, weekStart, weekStart);
        Array.from(this.weekNumbers.weeks.children).forEach((el, index) => {
          const dateOfWeekStart = addWeeks(startOfWeek, index);
          el.textContent = this.getWeekNumber(dateOfWeekStart, weekStart);
          if (index > 3) {
            el.classList[dateOfWeekStart > this.last ? "add" : "remove"]("next");
          }
        });
      }
      Array.from(this.grid.children).forEach((el, index) => {
        const current = addDays(this.start, index);
        const dateObj = new Date(current);
        const day = dateObj.getDay();
        const extraClasses = [];
        if (this.today === current) {
          extraClasses.push("today");
        }
        if (this.daysOfWeekHighlighted.includes(day)) {
          extraClasses.push("highlighted");
        }
        this.renderCell(
          el,
          dateObj.getDate(),
          current,
          current,
          this,
          current < this.minDate || current > this.maxDate || this.daysOfWeekDisabled.includes(day),
          extraClasses
        );
      });
    }
    // Update the view UI by applying the changes of selected and focused items
    refresh() {
      const range2 = this.range || [];
      Array.from(this.grid.children).forEach((el) => {
        this.refreshCell(el, Number(el.dataset.date), this.selected, range2);
      });
    }
    // Update the view UI by applying the change of focused item
    refreshFocus() {
      this.changeFocusedCell(Math.round((this.focused - this.start) / 864e5));
    }
  };

  // datepicker/js/picker/views/MonthsView.js
  function computeMonthRange(range2, thisYear) {
    if (!range2 || !range2[0] || !range2[1]) {
      return;
    }
    const [[startY, startM], [endY, endM]] = range2;
    if (startY > thisYear || endY < thisYear) {
      return;
    }
    return [
      startY === thisYear ? startM : -1,
      endY === thisYear ? endM : 12
    ];
  }
  var MonthsView = class extends View {
    constructor(picker) {
      super(picker, {
        id: 1,
        name: "months",
        cellClass: "month"
      });
    }
    init(options, onConstruction = true) {
      if (onConstruction) {
        this.grid = this.element;
        this.element.classList.add("months", "datepicker-grid");
        this.grid.appendChild(parseHTML(createTagRepeat("span", 12, { "data-month": (ix) => ix })));
        this.first = 0;
        this.last = 11;
      }
      super.init(options);
    }
    setOptions(options) {
      if (options.locale) {
        this.monthNames = options.locale.monthsShort;
      }
      if ("minDate" in options) {
        if (options.minDate === void 0) {
          this.minYear = this.minMonth = this.minDate = void 0;
        } else {
          const minDateObj = new Date(options.minDate);
          this.minYear = minDateObj.getFullYear();
          this.minMonth = minDateObj.getMonth();
          this.minDate = minDateObj.setDate(1);
        }
      }
      if ("maxDate" in options) {
        if (options.maxDate === void 0) {
          this.maxYear = this.maxMonth = this.maxDate = void 0;
        } else {
          const maxDateObj = new Date(options.maxDate);
          this.maxYear = maxDateObj.getFullYear();
          this.maxMonth = maxDateObj.getMonth();
          this.maxDate = dateValue(this.maxYear, this.maxMonth + 1, 0);
        }
      }
      if (options.checkDisabled) {
        this.checkDisabled = this.isMinView || options.datesDisabled === null ? options.checkDisabled : () => false;
      }
      if ("beforeShowMonth" in options) {
        this.beforeShow = typeof options.beforeShowMonth === "function" ? options.beforeShowMonth : void 0;
      }
    }
    // Update view's settings to reflect the viewDate set on the picker
    updateFocus() {
      const viewDate = new Date(this.picker.viewDate);
      this.year = viewDate.getFullYear();
      this.focused = viewDate.getMonth();
    }
    // Update view's settings to reflect the selected dates
    updateSelection() {
      const { dates, rangepicker } = this.picker.datepicker;
      this.selected = dates.reduce((selected, timeValue) => {
        const date = new Date(timeValue);
        const year = date.getFullYear();
        const month = date.getMonth();
        if (selected[year] === void 0) {
          selected[year] = [month];
        } else {
          pushUnique(selected[year], month);
        }
        return selected;
      }, {});
      if (rangepicker && rangepicker.dates) {
        this.range = rangepicker.dates.map((timeValue) => {
          const date = new Date(timeValue);
          return isNaN(date) ? void 0 : [date.getFullYear(), date.getMonth()];
        });
      }
    }
    // Update the entire view UI
    render() {
      this.prepareForRender(
        this.year,
        this.year <= this.minYear,
        this.year >= this.maxYear
      );
      const selected = this.selected[this.year] || [];
      const yrOutOfRange = this.year < this.minYear || this.year > this.maxYear;
      const isMinYear = this.year === this.minYear;
      const isMaxYear = this.year === this.maxYear;
      const range2 = computeMonthRange(this.range, this.year);
      Array.from(this.grid.children).forEach((el, index) => {
        const date = regularizeDate(new Date(this.year, index, 1), 1, this.isRangeEnd);
        this.renderCell(
          el,
          this.monthNames[index],
          index,
          date,
          { selected, range: range2 },
          yrOutOfRange || isMinYear && index < this.minMonth || isMaxYear && index > this.maxMonth
        );
      });
    }
    // Update the view UI by applying the changes of selected and focused items
    refresh() {
      const selected = this.selected[this.year] || [];
      const range2 = computeMonthRange(this.range, this.year) || [];
      Array.from(this.grid.children).forEach((el, index) => {
        this.refreshCell(el, index, selected, range2);
      });
    }
    // Update the view UI by applying the change of focused item
    refreshFocus() {
      this.changeFocusedCell(this.focused);
    }
  };

  // datepicker/js/picker/views/YearsView.js
  function toTitleCase(word) {
    return [...word].reduce((str, ch, ix) => str += ix ? ch : ch.toUpperCase(), "");
  }
  var YearsView = class extends View {
    constructor(picker, config) {
      super(picker, config);
    }
    init(options, onConstruction = true) {
      if (onConstruction) {
        this.navStep = this.step * 10;
        this.beforeShowOption = `beforeShow${toTitleCase(this.cellClass)}`;
        this.grid = this.element;
        this.element.classList.add(this.name, "datepicker-grid");
        this.grid.appendChild(parseHTML(createTagRepeat("span", 12)));
      }
      super.init(options);
    }
    setOptions(options) {
      if ("minDate" in options) {
        if (options.minDate === void 0) {
          this.minYear = this.minDate = void 0;
        } else {
          this.minYear = startOfYearPeriod(options.minDate, this.step);
          this.minDate = dateValue(this.minYear, 0, 1);
        }
      }
      if ("maxDate" in options) {
        if (options.maxDate === void 0) {
          this.maxYear = this.maxDate = void 0;
        } else {
          this.maxYear = startOfYearPeriod(options.maxDate, this.step);
          this.maxDate = dateValue(this.maxYear, 11, 31);
        }
      }
      if (options.checkDisabled) {
        this.checkDisabled = this.isMinView || options.datesDisabled === null ? options.checkDisabled : () => false;
      }
      if (this.beforeShowOption in options) {
        const beforeShow = options[this.beforeShowOption];
        this.beforeShow = typeof beforeShow === "function" ? beforeShow : void 0;
      }
    }
    // Update view's settings to reflect the viewDate set on the picker
    updateFocus() {
      const viewDate = new Date(this.picker.viewDate);
      const first = startOfYearPeriod(viewDate, this.navStep);
      const last = first + 9 * this.step;
      this.first = first;
      this.last = last;
      this.start = first - this.step;
      this.focused = startOfYearPeriod(viewDate, this.step);
    }
    // Update view's settings to reflect the selected dates
    updateSelection() {
      const { dates, rangepicker } = this.picker.datepicker;
      this.selected = dates.reduce((years, timeValue) => {
        return pushUnique(years, startOfYearPeriod(timeValue, this.step));
      }, []);
      if (rangepicker && rangepicker.dates) {
        this.range = rangepicker.dates.map((timeValue) => {
          if (timeValue !== void 0) {
            return startOfYearPeriod(timeValue, this.step);
          }
        });
      }
    }
    // Update the entire view UI
    render() {
      this.prepareForRender(
        `${this.first}-${this.last}`,
        this.first <= this.minYear,
        this.last >= this.maxYear
      );
      Array.from(this.grid.children).forEach((el, index) => {
        const current = this.start + index * this.step;
        const date = regularizeDate(new Date(current, 0, 1), 2, this.isRangeEnd);
        el.dataset.year = current;
        this.renderCell(
          el,
          current,
          current,
          date,
          this,
          current < this.minYear || current > this.maxYear
        );
      });
    }
    // Update the view UI by applying the changes of selected and focused items
    refresh() {
      const range2 = this.range || [];
      Array.from(this.grid.children).forEach((el) => {
        this.refreshCell(el, Number(el.textContent), this.selected, range2);
      });
    }
    // Update the view UI by applying the change of focused item
    refreshFocus() {
      this.changeFocusedCell(Math.round((this.focused - this.start) / this.step));
    }
  };

  // datepicker/js/events/functions.js
  function triggerDatepickerEvent(datepicker, type) {
    const options = {
      bubbles: true,
      cancelable: true,
      detail: {
        date: datepicker.getDate(),
        viewDate: new Date(datepicker.picker.viewDate),
        viewId: datepicker.picker.currentView.id,
        datepicker
      }
    };
    datepicker.element.dispatchEvent(new CustomEvent(type, options));
  }
  function goToPrevOrNext(datepicker, direction) {
    const { config, picker } = datepicker;
    const { currentView, viewDate } = picker;
    let newViewDate;
    switch (currentView.id) {
      case 0:
        newViewDate = addMonths(viewDate, direction);
        break;
      case 1:
        newViewDate = addYears(viewDate, direction);
        break;
      default:
        newViewDate = addYears(viewDate, direction * currentView.navStep);
    }
    newViewDate = limitToRange(newViewDate, config.minDate, config.maxDate);
    picker.changeFocus(newViewDate).render();
  }
  function switchView(datepicker) {
    const viewId = datepicker.picker.currentView.id;
    if (viewId === datepicker.config.maxView) {
      return;
    }
    datepicker.picker.changeView(viewId + 1).render();
  }
  function clearSelection(datepicker) {
    datepicker.setDate({ clear: true });
  }
  function goToOrSelectToday(datepicker) {
    const currentDate = today();
    if (datepicker.config.todayButtonMode === 1) {
      datepicker.setDate(currentDate, { forceRefresh: true, viewDate: currentDate });
    } else {
      datepicker.setFocusedDate(currentDate, true);
    }
  }
  function unfocus(datepicker) {
    const onBlur = () => {
      if (datepicker.config.updateOnBlur) {
        datepicker.update({ revert: true });
      } else {
        datepicker.refresh("input");
      }
      datepicker.hide();
    };
    const element = datepicker.element;
    if (isActiveElement(element)) {
      element.addEventListener("blur", onBlur, { once: true });
    } else {
      onBlur();
    }
  }

  // datepicker/js/events/pickerListeners.js
  function goToSelectedMonthOrYear(datepicker, selection) {
    const picker = datepicker.picker;
    const viewDate = new Date(picker.viewDate);
    const viewId = picker.currentView.id;
    const newDate = viewId === 1 ? addMonths(viewDate, selection - viewDate.getMonth()) : addYears(viewDate, selection - viewDate.getFullYear());
    picker.changeFocus(newDate).changeView(viewId - 1).render();
  }
  function onClickViewSwitch(datepicker) {
    switchView(datepicker);
  }
  function onClickPrevButton(datepicker) {
    goToPrevOrNext(datepicker, -1);
  }
  function onClickNextButton(datepicker) {
    goToPrevOrNext(datepicker, 1);
  }
  function onClickView(datepicker, ev) {
    const target = findElementInEventPath(ev, ".datepicker-cell");
    if (!target || target.classList.contains("disabled")) {
      return;
    }
    const { id, isMinView } = datepicker.picker.currentView;
    const data = target.dataset;
    if (isMinView) {
      datepicker.setDate(Number(data.date));
    } else if (id === 1) {
      goToSelectedMonthOrYear(datepicker, Number(data.month));
    } else {
      goToSelectedMonthOrYear(datepicker, Number(data.year));
    }
  }
  function onMousedownPicker(ev) {
    ev.preventDefault();
  }

  // datepicker/js/picker/Picker.js
  var orientClasses = ["left", "top", "right", "bottom"].reduce((obj, key) => {
    obj[key] = `datepicker-orient-${key}`;
    return obj;
  }, {});
  var toPx = (num) => num ? `${num}px` : num;
  function processPickerOptions(picker, options) {
    if ("title" in options) {
      if (options.title) {
        picker.controls.title.textContent = options.title;
        showElement(picker.controls.title);
      } else {
        picker.controls.title.textContent = "";
        hideElement(picker.controls.title);
      }
    }
    if (options.prevArrow) {
      const prevButton = picker.controls.prevButton;
      emptyChildNodes(prevButton);
      options.prevArrow.forEach((node) => {
        prevButton.appendChild(node.cloneNode(true));
      });
    }
    if (options.nextArrow) {
      const nextButton = picker.controls.nextButton;
      emptyChildNodes(nextButton);
      options.nextArrow.forEach((node) => {
        nextButton.appendChild(node.cloneNode(true));
      });
    }
    if (options.locale) {
      picker.controls.todayButton.textContent = options.locale.today;
      picker.controls.clearButton.textContent = options.locale.clear;
    }
    if ("todayButton" in options) {
      if (options.todayButton) {
        showElement(picker.controls.todayButton);
      } else {
        hideElement(picker.controls.todayButton);
      }
    }
    if ("minDate" in options || "maxDate" in options) {
      const { minDate, maxDate } = picker.datepicker.config;
      picker.controls.todayButton.disabled = !isInRange(today(), minDate, maxDate);
    }
    if ("clearButton" in options) {
      if (options.clearButton) {
        showElement(picker.controls.clearButton);
      } else {
        hideElement(picker.controls.clearButton);
      }
    }
  }
  function computeResetViewDate(datepicker) {
    const { dates, config, rangeSideIndex } = datepicker;
    const viewDate = dates.length > 0 ? lastItemOf(dates) : regularizeDate(config.defaultViewDate, config.pickLevel, rangeSideIndex);
    return limitToRange(viewDate, config.minDate, config.maxDate);
  }
  function setViewDate(picker, newDate) {
    if (!("_oldViewDate" in picker) && newDate !== picker.viewDate) {
      picker._oldViewDate = picker.viewDate;
    }
    picker.viewDate = newDate;
    const { id, year, first, last } = picker.currentView;
    const viewYear = new Date(newDate).getFullYear();
    switch (id) {
      case 0:
        return newDate < first || newDate > last;
      case 1:
        return viewYear !== year;
      default:
        return viewYear < first || viewYear > last;
    }
  }
  function getTextDirection(el) {
    return window.getComputedStyle(el).direction;
  }
  function findScrollParents(el) {
    const parent = getParent(el);
    if (parent === document.body || !parent) {
      return;
    }
    return window.getComputedStyle(parent).overflow !== "visible" ? parent : findScrollParents(parent);
  }
  var Picker = class {
    constructor(datepicker) {
      const { config, inputField } = this.datepicker = datepicker;
      const template = pickerTemplate_default.replace(/%buttonClass%/g, config.buttonClass);
      const element = this.element = parseHTML(template).firstChild;
      const [header, main, footer] = element.firstChild.children;
      const title = header.firstElementChild;
      const [prevButton, viewSwitch, nextButton] = header.lastElementChild.children;
      const [todayButton, clearButton] = footer.firstChild.children;
      const controls = {
        title,
        prevButton,
        viewSwitch,
        nextButton,
        todayButton,
        clearButton
      };
      this.main = main;
      this.controls = controls;
      const elementClass = inputField ? "dropdown" : "inline";
      element.classList.add(`datepicker-${elementClass}`);
      processPickerOptions(this, config);
      this.viewDate = computeResetViewDate(datepicker);
      registerListeners(datepicker, [
        [element, "mousedown", onMousedownPicker],
        [main, "click", onClickView.bind(null, datepicker)],
        [controls.viewSwitch, "click", onClickViewSwitch.bind(null, datepicker)],
        [controls.prevButton, "click", onClickPrevButton.bind(null, datepicker)],
        [controls.nextButton, "click", onClickNextButton.bind(null, datepicker)],
        [controls.todayButton, "click", goToOrSelectToday.bind(null, datepicker)],
        [controls.clearButton, "click", clearSelection.bind(null, datepicker)]
      ]);
      this.views = [
        new DaysView(this),
        new MonthsView(this),
        new YearsView(this, { id: 2, name: "years", cellClass: "year", step: 1 }),
        new YearsView(this, { id: 3, name: "decades", cellClass: "decade", step: 10 })
      ];
      this.currentView = this.views[config.startView];
      this.currentView.render();
      this.main.appendChild(this.currentView.element);
      if (config.container) {
        config.container.appendChild(this.element);
      } else {
        inputField.after(this.element);
      }
    }
    setOptions(options) {
      processPickerOptions(this, options);
      this.views.forEach((view) => {
        view.init(options, false);
      });
      this.currentView.render();
    }
    detach() {
      this.element.remove();
    }
    show() {
      if (this.active) {
        return;
      }
      const { datepicker, element } = this;
      const inputField = datepicker.inputField;
      if (inputField) {
        const inputDirection = getTextDirection(inputField);
        if (inputDirection !== getTextDirection(getParent(element))) {
          element.dir = inputDirection;
        } else if (element.dir) {
          element.removeAttribute("dir");
        }
        this.place();
        element.classList.add("active");
        if (datepicker.config.disableTouchKeyboard) {
          inputField.blur();
        }
      } else {
        element.classList.add("active");
      }
      this.active = true;
      triggerDatepickerEvent(datepicker, "show");
    }
    hide() {
      if (!this.active) {
        return;
      }
      this.datepicker.exitEditMode();
      this.element.classList.remove("active");
      this.active = false;
      triggerDatepickerEvent(this.datepicker, "hide");
    }
    place() {
      const { classList, style } = this.element;
      style.display = "block";
      const {
        width: calendarWidth,
        height: calendarHeight
      } = this.element.getBoundingClientRect();
      const offsetParent = this.element.offsetParent;
      style.display = "";
      const { config, inputField } = this.datepicker;
      const {
        left: inputLeft,
        top: inputTop,
        right: inputRight,
        bottom: inputBottom,
        width: inputWidth,
        height: inputHeight
      } = inputField.getBoundingClientRect();
      let { x: orientX, y: orientY } = config.orientation;
      let left = inputLeft;
      let top = inputTop;
      if (offsetParent === document.body || !offsetParent) {
        left += window.scrollX;
        top += window.scrollY;
      } else {
        const offsetParentRect = offsetParent.getBoundingClientRect();
        left -= offsetParentRect.left - offsetParent.scrollLeft;
        top -= offsetParentRect.top - offsetParent.scrollTop;
      }
      const scrollParent = findScrollParents(inputField);
      let scrollAreaLeft = 0;
      let scrollAreaTop = 0;
      let {
        clientWidth: scrollAreaRight,
        clientHeight: scrollAreaBottom
      } = document.documentElement;
      if (scrollParent) {
        const scrollParentRect = scrollParent.getBoundingClientRect();
        if (scrollParentRect.top > 0) {
          scrollAreaTop = scrollParentRect.top;
        }
        if (scrollParentRect.left > 0) {
          scrollAreaLeft = scrollParentRect.left;
        }
        if (scrollParentRect.right < scrollAreaRight) {
          scrollAreaRight = scrollParentRect.right;
        }
        if (scrollParentRect.bottom < scrollAreaBottom) {
          scrollAreaBottom = scrollParentRect.bottom;
        }
      }
      let adjustment = 0;
      if (orientX === "auto") {
        if (inputLeft < scrollAreaLeft) {
          orientX = "left";
          adjustment = scrollAreaLeft - inputLeft;
        } else if (inputLeft + calendarWidth > scrollAreaRight) {
          orientX = "right";
          if (scrollAreaRight < inputRight) {
            adjustment = scrollAreaRight - inputRight;
          }
        } else if (getTextDirection(inputField) === "rtl") {
          orientX = inputRight - calendarWidth < scrollAreaLeft ? "left" : "right";
        } else {
          orientX = "left";
        }
      }
      if (orientX === "right") {
        left += inputWidth - calendarWidth;
      }
      left += adjustment;
      if (orientY === "auto") {
        if (inputTop - calendarHeight > scrollAreaTop) {
          orientY = inputBottom + calendarHeight > scrollAreaBottom ? "top" : "bottom";
        } else {
          orientY = "bottom";
        }
      }
      if (orientY === "top") {
        top -= calendarHeight;
      } else {
        top += inputHeight;
      }
      classList.remove(...Object.values(orientClasses));
      classList.add(orientClasses[orientX], orientClasses[orientY]);
      style.left = toPx(left);
      style.top = toPx(top);
    }
    setViewSwitchLabel(labelText) {
      this.controls.viewSwitch.textContent = labelText;
    }
    setPrevButtonDisabled(disabled) {
      this.controls.prevButton.disabled = disabled;
    }
    setNextButtonDisabled(disabled) {
      this.controls.nextButton.disabled = disabled;
    }
    changeView(viewId) {
      const currentView = this.currentView;
      if (viewId !== currentView.id) {
        if (!this._oldView) {
          this._oldView = currentView;
        }
        this.currentView = this.views[viewId];
        this._renderMethod = "render";
      }
      return this;
    }
    // Change the focused date (view date)
    changeFocus(newViewDate) {
      this._renderMethod = setViewDate(this, newViewDate) ? "render" : "refreshFocus";
      this.views.forEach((view) => {
        view.updateFocus();
      });
      return this;
    }
    // Apply the change of the selected dates
    update(viewDate = void 0) {
      const newViewDate = viewDate === void 0 ? computeResetViewDate(this.datepicker) : viewDate;
      this._renderMethod = setViewDate(this, newViewDate) ? "render" : "refresh";
      this.views.forEach((view) => {
        view.updateFocus();
        view.updateSelection();
      });
      return this;
    }
    // Refresh the picker UI
    render(quickRender = true) {
      const { currentView, datepicker, _oldView: oldView } = this;
      const oldViewDate = new Date(this._oldViewDate);
      const renderMethod = quickRender && this._renderMethod || "render";
      delete this._oldView;
      delete this._oldViewDate;
      delete this._renderMethod;
      currentView[renderMethod]();
      if (oldView) {
        this.main.replaceChild(currentView.element, oldView.element);
        triggerDatepickerEvent(datepicker, "changeView");
      }
      if (!isNaN(oldViewDate)) {
        const newViewDate = new Date(this.viewDate);
        if (newViewDate.getFullYear() !== oldViewDate.getFullYear()) {
          triggerDatepickerEvent(datepicker, "changeYear");
        }
        if (newViewDate.getMonth() !== oldViewDate.getMonth()) {
          triggerDatepickerEvent(datepicker, "changeMonth");
        }
      }
    }
  };

  // datepicker/js/events/elementListeners.js
  function findNextAvailableOne(date, addFn, increase, testFn, min, max) {
    if (!isInRange(date, min, max)) {
      return;
    }
    if (testFn(date)) {
      const newDate = addFn(date, increase);
      return findNextAvailableOne(newDate, addFn, increase, testFn, min, max);
    }
    return date;
  }
  function moveByArrowKey(datepicker, direction, vertical) {
    const picker = datepicker.picker;
    const currentView = picker.currentView;
    const step = currentView.step || 1;
    let viewDate = picker.viewDate;
    let addFn;
    switch (currentView.id) {
      case 0:
        viewDate = addDays(viewDate, vertical ? direction * 7 : direction);
        addFn = addDays;
        break;
      case 1:
        viewDate = addMonths(viewDate, vertical ? direction * 4 : direction);
        addFn = addMonths;
        break;
      default:
        viewDate = addYears(viewDate, direction * (vertical ? 4 : 1) * step);
        addFn = addYears;
    }
    viewDate = findNextAvailableOne(
      viewDate,
      addFn,
      direction < 0 ? -step : step,
      (date) => currentView.disabled.includes(date),
      currentView.minDate,
      currentView.maxDate
    );
    if (viewDate !== void 0) {
      picker.changeFocus(viewDate).render();
    }
  }
  function onKeydown(datepicker, ev) {
    const { config, picker, editMode } = datepicker;
    const active = picker.active;
    const { key, altKey, shiftKey } = ev;
    const ctrlOrMetaKey = ev.ctrlKey || ev.metaKey;
    const cancelEvent = () => {
      ev.preventDefault();
      ev.stopPropagation();
    };
    if (key === "Tab") {
      unfocus(datepicker);
      return;
    }
    if (key === "Enter") {
      if (!active) {
        datepicker.update();
      } else if (editMode) {
        datepicker.exitEditMode({ update: true, autohide: config.autohide });
      } else {
        const currentView = picker.currentView;
        if (currentView.isMinView) {
          datepicker.setDate(picker.viewDate);
        } else {
          picker.changeView(currentView.id - 1).render();
          cancelEvent();
        }
      }
      return;
    }
    const shortcutKeys = config.shortcutKeys;
    const keyInfo = { key, ctrlOrMetaKey, altKey, shiftKey };
    const shortcut = Object.keys(shortcutKeys).find((item) => {
      const keyDef = shortcutKeys[item];
      return !Object.keys(keyDef).find((prop) => keyDef[prop] !== keyInfo[prop]);
    });
    if (shortcut) {
      let action;
      if (shortcut === "toggle") {
        action = shortcut;
      } else if (editMode) {
        if (shortcut === "exitEditMode") {
          action = shortcut;
        }
      } else if (active) {
        if (shortcut === "hide") {
          action = shortcut;
        } else if (shortcut === "prevButton") {
          action = [goToPrevOrNext, [datepicker, -1]];
        } else if (shortcut === "nextButton") {
          action = [goToPrevOrNext, [datepicker, 1]];
        } else if (shortcut === "viewSwitch") {
          action = [switchView, [datepicker]];
        } else if (config.clearButton && shortcut === "clearButton") {
          action = [clearSelection, [datepicker]];
        } else if (config.todayButton && shortcut === "todayButton") {
          action = [goToOrSelectToday, [datepicker]];
        }
      } else if (shortcut === "show") {
        action = shortcut;
      }
      if (action) {
        if (Array.isArray(action)) {
          action[0].apply(null, action[1]);
        } else {
          datepicker[action]();
        }
        cancelEvent();
        return;
      }
    }
    if (!active || editMode) {
      return;
    }
    const handleArrowKeyPress = (direction, vertical) => {
      if (shiftKey || ctrlOrMetaKey || altKey) {
        datepicker.enterEditMode();
      } else {
        moveByArrowKey(datepicker, direction, vertical);
        ev.preventDefault();
      }
    };
    if (key === "ArrowLeft") {
      handleArrowKeyPress(-1, false);
    } else if (key === "ArrowRight") {
      handleArrowKeyPress(1, false);
    } else if (key === "ArrowUp") {
      handleArrowKeyPress(-1, true);
    } else if (key === "ArrowDown") {
      handleArrowKeyPress(1, true);
    } else if (key === "Backspace" || key === "Delete" || key && key.length === 1 && !ctrlOrMetaKey) {
      datepicker.enterEditMode();
    }
  }
  function onFocus(datepicker) {
    if (datepicker.config.showOnFocus && !datepicker._showing) {
      datepicker.show();
    }
  }
  function onMousedown(datepicker, ev) {
    const el = ev.target;
    if (datepicker.picker.active || datepicker.config.showOnClick) {
      el._active = isActiveElement(el);
      el._clicking = setTimeout(() => {
        delete el._active;
        delete el._clicking;
      }, 2e3);
    }
  }
  function onClickInput(datepicker, ev) {
    const el = ev.target;
    if (!el._clicking) {
      return;
    }
    clearTimeout(el._clicking);
    delete el._clicking;
    if (el._active) {
      datepicker.enterEditMode();
    }
    delete el._active;
    if (datepicker.config.showOnClick) {
      datepicker.show();
    }
  }
  function onPaste(datepicker, ev) {
    if (ev.clipboardData.types.includes("text/plain")) {
      datepicker.enterEditMode();
    }
  }

  // datepicker/js/events/otherListeners.js
  function onClickOutside(datepicker, ev) {
    const { element, picker } = datepicker;
    if (!picker.active && !isActiveElement(element)) {
      return;
    }
    const pickerElem = picker.element;
    if (findElementInEventPath(ev, (el) => el === element || el === pickerElem)) {
      return;
    }
    unfocus(datepicker);
  }

  // datepicker/js/Datepicker.js
  function stringifyDates(dates, config) {
    return dates.map((dt) => formatDate(dt, config.format, config.locale)).join(config.dateDelimiter);
  }
  function processInputDates(datepicker, inputDates, clear = false) {
    if (inputDates.length === 0) {
      return clear ? [] : void 0;
    }
    const { config, dates: origDates, rangeSideIndex } = datepicker;
    const { pickLevel, maxNumberOfDates } = config;
    let newDates = inputDates.reduce((dates, dt) => {
      let date = parseDate(dt, config.format, config.locale);
      if (date === void 0) {
        return dates;
      }
      date = regularizeDate(date, pickLevel, rangeSideIndex);
      if (isInRange(date, config.minDate, config.maxDate) && !dates.includes(date) && !config.checkDisabled(date, pickLevel) && (pickLevel > 0 || !config.daysOfWeekDisabled.includes(new Date(date).getDay()))) {
        dates.push(date);
      }
      return dates;
    }, []);
    if (newDates.length === 0) {
      return;
    }
    if (config.multidate && !clear) {
      newDates = newDates.reduce((dates, date) => {
        if (!origDates.includes(date)) {
          dates.push(date);
        }
        return dates;
      }, origDates.filter((date) => !newDates.includes(date)));
    }
    return maxNumberOfDates && newDates.length > maxNumberOfDates ? newDates.slice(maxNumberOfDates * -1) : newDates;
  }
  function refreshUI(datepicker, mode = 3, quickRender = true, viewDate = void 0) {
    const { config, picker, inputField } = datepicker;
    if (mode & 2) {
      const newView = picker.active ? config.pickLevel : config.startView;
      picker.update(viewDate).changeView(newView).render(quickRender);
    }
    if (mode & 1 && inputField) {
      inputField.value = stringifyDates(datepicker.dates, config);
    }
  }
  function setDate(datepicker, inputDates, options) {
    const config = datepicker.config;
    let { clear, render, autohide, revert, forceRefresh, viewDate } = options;
    if (render === void 0) {
      render = true;
    }
    if (!render) {
      autohide = forceRefresh = false;
    } else if (autohide === void 0) {
      autohide = config.autohide;
    }
    viewDate = parseDate(viewDate, config.format, config.locale);
    const newDates = processInputDates(datepicker, inputDates, clear);
    if (!newDates && !revert) {
      return;
    }
    if (newDates && newDates.toString() !== datepicker.dates.toString()) {
      datepicker.dates = newDates;
      refreshUI(datepicker, render ? 3 : 1, true, viewDate);
      triggerDatepickerEvent(datepicker, "changeDate");
    } else {
      refreshUI(datepicker, forceRefresh ? 3 : 1, true, viewDate);
    }
    if (autohide) {
      datepicker.hide();
    }
  }
  function getOutputConverter(datepicker, format) {
    return format ? (date) => formatDate(date, format, datepicker.config.locale) : (date) => new Date(date);
  }
  var Datepicker = class {
    /**
     * Create a date picker
     * @param  {Element} element - element to bind a date picker
     * @param  {Object} [options] - config options
     * @param  {DateRangePicker} [rangepicker] - DateRangePicker instance the
     * date picker belongs to. Use this only when creating date picker as a part
     * of date range picker
     */
    constructor(element, options = {}, rangepicker = void 0) {
      element.datepicker = this;
      this.element = element;
      this.dates = [];
      const config = this.config = Object.assign({
        buttonClass: options.buttonClass && String(options.buttonClass) || "button",
        container: null,
        defaultViewDate: today(),
        maxDate: void 0,
        minDate: void 0
      }, processOptions(defaultOptions_default, this));
      let inputField;
      if (element.tagName === "INPUT") {
        inputField = this.inputField = element;
        inputField.classList.add("datepicker-input");
        if (options.container) {
          config.container = options.container instanceof HTMLElement ? options.container : document.querySelector(options.container);
        }
      } else {
        config.container = element;
      }
      if (rangepicker) {
        const index = rangepicker.inputs.indexOf(inputField);
        const datepickers = rangepicker.datepickers;
        if (index < 0 || index > 1 || !Array.isArray(datepickers)) {
          throw Error("Invalid rangepicker object.");
        }
        datepickers[index] = this;
        this.rangepicker = rangepicker;
        this.rangeSideIndex = index;
      }
      this._options = options;
      Object.assign(config, processOptions(options, this));
      config.shortcutKeys = createShortcutKeyConfig(options.shortcutKeys || {});
      const initialDates = stringToArray(
        element.value || element.dataset.date,
        config.dateDelimiter
      );
      delete element.dataset.date;
      const inputDateValues = processInputDates(this, initialDates);
      if (inputDateValues && inputDateValues.length > 0) {
        this.dates = inputDateValues;
      }
      if (inputField) {
        inputField.value = stringifyDates(this.dates, config);
      }
      const picker = this.picker = new Picker(this);
      const keydownListener = [element, "keydown", onKeydown.bind(null, this)];
      if (inputField) {
        registerListeners(this, [
          keydownListener,
          [inputField, "focus", onFocus.bind(null, this)],
          [inputField, "mousedown", onMousedown.bind(null, this)],
          [inputField, "click", onClickInput.bind(null, this)],
          [inputField, "paste", onPaste.bind(null, this)],
          // To detect a click on outside, just listening to mousedown is enough,
          // no need to listen to touchstart.
          // Actually, listening to touchstart can be a problem because, while
          // mousedown is fired only on tapping but not on swiping/pinching,
          // touchstart is fired on swiping/pinching as well.
          // (issue #95)
          [document, "mousedown", onClickOutside.bind(null, this)],
          [window, "resize", picker.place.bind(picker)]
        ]);
      } else {
        registerListeners(this, [keydownListener]);
        this.show();
      }
    }
    /**
     * Format Date object or time value in given format and language
     * @param  {Date|Number} date - date or time value to format
     * @param  {String|Object} format - format string or object that contains
     * toDisplay() custom formatter, whose signature is
     * - args:
     *   - date: {Date} - Date instance of the date passed to the method
     *   - format: {Object} - the format object passed to the method
     *   - locale: {Object} - locale for the language specified by `lang`
     * - return:
     *     {String} formatted date
     * @param  {String} [lang=en] - language code for the locale to use
     * @return {String} formatted date
     */
    static formatDate(date, format, lang) {
      return formatDate(date, format, lang && base_locales_default[lang] || base_locales_default.en);
    }
    /**
     * Parse date string
     * @param  {String|Date|Number} dateStr - date string, Date object or time
     * value to parse
     * @param  {String|Object} format - format string or object that contains
     * toValue() custom parser, whose signature is
     * - args:
     *   - dateStr: {String|Date|Number} - the dateStr passed to the method
     *   - format: {Object} - the format object passed to the method
     *   - locale: {Object} - locale for the language specified by `lang`
     * - return:
     *     {Date|Number} parsed date or its time value
     * @param  {String} [lang=en] - language code for the locale to use
     * @return {Number} time value of parsed date
     */
    static parseDate(dateStr, format, lang) {
      return parseDate(dateStr, format, lang && base_locales_default[lang] || base_locales_default.en);
    }
    /**
     * @type {Object} - Installed locales in `[languageCode]: localeObject` format
     * en`:_English (US)_ is pre-installed.
     */
    static get locales() {
      return base_locales_default;
    }
    /**
     * @type {Boolean} - Whether the picker element is shown. `true` when shown
     */
    get active() {
      return !!(this.picker && this.picker.active);
    }
    /**
     * @type {HTMLDivElement} - DOM object of picker element
     */
    get pickerElement() {
      return this.picker ? this.picker.element : void 0;
    }
    /**
     * Set new values to the config options
     * @param {Object} options - config options to update
     */
    setOptions(options) {
      const newOptions = processOptions(options, this);
      Object.assign(this._options, options);
      Object.assign(this.config, newOptions);
      this.picker.setOptions(newOptions);
      refreshUI(this, 3);
    }
    /**
     * Show the picker element
     */
    show() {
      if (this.inputField) {
        const { config, inputField } = this;
        if (inputField.disabled || inputField.readOnly && !config.enableOnReadonly) {
          return;
        }
        if (!isActiveElement(inputField) && !config.disableTouchKeyboard) {
          this._showing = true;
          inputField.focus();
          delete this._showing;
        }
      }
      this.picker.show();
    }
    /**
     * Hide the picker element
     * Not available on inline picker
     */
    hide() {
      if (!this.inputField) {
        return;
      }
      this.picker.hide();
      this.picker.update().changeView(this.config.startView).render();
    }
    /**
     * Toggle the display of the picker element
     * Not available on inline picker
     *
     * Unlike hide(), the picker does not return to the start view when hiding.
     */
    toggle() {
      if (!this.picker.active) {
        this.show();
      } else if (this.inputField) {
        this.picker.hide();
      }
    }
    /**
     * Destroy the Datepicker instance
     * @return {Detepicker} - the instance destroyed
     */
    destroy() {
      this.hide();
      unregisterListeners(this);
      this.picker.detach();
      const element = this.element;
      element.classList.remove("datepicker-input");
      delete element.datepicker;
      return this;
    }
    /**
     * Get the selected date(s)
     *
     * The method returns a Date object of selected date by default, and returns
     * an array of selected dates in multidate mode. If format string is passed,
     * it returns date string(s) formatted in given format.
     *
     * @param  {String} [format] - format string to stringify the date(s)
     * @return {Date|String|Date[]|String[]} - selected date(s), or if none is
     * selected, empty array in multidate mode and undefined in sigledate mode
     */
    getDate(format = void 0) {
      const callback = getOutputConverter(this, format);
      if (this.config.multidate) {
        return this.dates.map(callback);
      }
      if (this.dates.length > 0) {
        return callback(this.dates[0]);
      }
    }
    /**
     * Set selected date(s)
     *
     * In multidate mode, you can pass multiple dates as a series of arguments
     * or an array. (Since each date is parsed individually, the type of the
     * dates doesn't have to be the same.)
     * The given dates are used to toggle the select status of each date. The
     * number of selected dates is kept from exceeding the length set to
     * maxNumberOfDates.
     *
     * With clear: true option, the method can be used to clear the selection
     * and to replace the selection instead of toggling in multidate mode.
     * If the option is passed with no date arguments or an empty dates array,
     * it works as "clear" (clear the selection then set nothing), and if the
     * option is passed with new dates to select, it works as "replace" (clear
     * the selection then set the given dates)
     *
     * When render: false option is used, the method omits re-rendering the
     * picker element. In this case, you need to call refresh() method later in
     * order for the picker element to reflect the changes. The input field is
     * refreshed always regardless of this option.
     *
     * When invalid (unparsable, repeated, disabled or out-of-range) dates are
     * passed, the method ignores them and applies only valid ones. In the case
     * that all the given dates are invalid, which is distinguished from passing
     * no dates, the method considers it as an error and leaves the selection
     * untouched. (The input field also remains untouched unless revert: true
     * option is used.)
     * Replacing the selection with the same date(s) also causes a similar
     * situation. In both cases, the method does not refresh the picker element
     * unless forceRefresh: true option is used.
     *
     * If viewDate option is used, the method changes the focused date to the
     * specified date instead of the last item of the selection.
     *
     * @param {...(Date|Number|String)|Array} [dates] - Date strings, Date
     * objects, time values or mix of those for new selection
     * @param {Object} [options] - function options
     * - clear: {boolean} - Whether to clear the existing selection
     *     default: false
     * - render: {boolean} - Whether to re-render the picker element
     *     default: true
     * - autohide: {boolean} - Whether to hide the picker element after re-render
     *     Ignored when used with render: false
     *     default: config.autohide
     * - revert: {boolean} - Whether to refresh the input field when all the
     *     passed dates are invalid
     *     default: false
     * - forceRefresh: {boolean} - Whether to refresh the picker element when
     *     passed dates don't change the existing selection
     *     default: false
     * - viewDate: {Date|Number|String} - Date to be focused after setiing date(s)
     *     default: The last item of the resulting selection, or defaultViewDate
     *     config option if none is selected
     */
    setDate(...args) {
      const dates = [...args];
      const opts = {};
      const lastArg = lastItemOf(args);
      if (lastArg && typeof lastArg === "object" && !Array.isArray(lastArg) && !(lastArg instanceof Date)) {
        Object.assign(opts, dates.pop());
      }
      const inputDates = Array.isArray(dates[0]) ? dates[0] : dates;
      setDate(this, inputDates, opts);
    }
    /**
     * Update the selected date(s) with input field's value
     * Not available on inline picker
     *
     * The input field will be refreshed with properly formatted date string.
     *
     * In the case that all the entered dates are invalid (unparsable, repeated,
     * disabled or out-of-range), which is distinguished from empty input field,
     * the method leaves the input field untouched as well as the selection by
     * default. If revert: true option is used in this case, the input field is
     * refreshed with the existing selection.
     * The method also doesn't refresh the picker element in this case and when
     * the entered dates are the same as the existing selection. If
     * forceRefresh: true option is used, the picker element is refreshed in
     * these cases too.
     *
     * @param  {Object} [options] - function options
     * - autohide: {boolean} - whether to hide the picker element after refresh
     *     default: false
     * - revert: {boolean} - Whether to refresh the input field when all the
     *     passed dates are invalid
     *     default: false
     * - forceRefresh: {boolean} - Whether to refresh the picer element when
     *     input field's value doesn't change the existing selection
     *     default: false
     */
    update(options = void 0) {
      if (!this.inputField) {
        return;
      }
      const opts = Object.assign(options || {}, { clear: true, render: true, viewDate: void 0 });
      const inputDates = stringToArray(this.inputField.value, this.config.dateDelimiter);
      setDate(this, inputDates, opts);
    }
    /**
     * Get the focused date
     *
     * The method returns a Date object of focused date by default. If format
     * string is passed, it returns date string formatted in given format.
     *
     * @param  {String} [format] - format string to stringify the date
     * @return {Date|String} - focused date (viewDate)
     */
    getFocusedDate(format = void 0) {
      return getOutputConverter(this, format)(this.picker.viewDate);
    }
    /**
     * Set focused date
     *
     * By default, the method updates the focus on the view shown at the time,
     * or the one set to the startView config option if the picker is hidden.
     * When resetView: true is passed, the view displayed is changed to the
     * pickLevel config option's if the picker is shown.
     *
     * @param {Date|Number|String} viewDate - date string, Date object, time
     * values of the date to focus
     * @param {Boolean} [resetView] - whether to change the view to pickLevel
     * config option's when the picker is shown. Ignored when the picker is
     * hidden
     */
    setFocusedDate(viewDate, resetView = false) {
      const { config, picker, active, rangeSideIndex } = this;
      const pickLevel = config.pickLevel;
      const newViewDate = parseDate(viewDate, config.format, config.locale);
      if (newViewDate === void 0) {
        return;
      }
      picker.changeFocus(regularizeDate(newViewDate, pickLevel, rangeSideIndex));
      if (active && resetView) {
        picker.changeView(pickLevel);
      }
      picker.render();
    }
    /**
     * Refresh the picker element and the associated input field
     * @param {String} [target] - target item when refreshing one item only
     * 'picker' or 'input'
     * @param {Boolean} [forceRender] - whether to re-render the picker element
     * regardless of its state instead of optimized refresh
     */
    refresh(target = void 0, forceRender = false) {
      if (target && typeof target !== "string") {
        forceRender = target;
        target = void 0;
      }
      let mode;
      if (target === "picker") {
        mode = 2;
      } else if (target === "input") {
        mode = 1;
      } else {
        mode = 3;
      }
      refreshUI(this, mode, !forceRender);
    }
    /**
     * Enter edit mode
     * Not available on inline picker or when the picker element is hidden
     */
    enterEditMode() {
      const inputField = this.inputField;
      if (!inputField || inputField.readOnly || !this.picker.active || this.editMode) {
        return;
      }
      this.editMode = true;
      inputField.classList.add("in-edit");
    }
    /**
     * Exit from edit mode
     * Not available on inline picker
     * @param  {Object} [options] - function options
     * - update: {boolean} - whether to call update() after exiting
     *     If false, input field is revert to the existing selection
     *     default: false
     */
    exitEditMode(options = void 0) {
      if (!this.inputField || !this.editMode) {
        return;
      }
      const opts = Object.assign({ update: false }, options);
      delete this.editMode;
      this.inputField.classList.remove("in-edit");
      if (opts.update) {
        this.update(opts);
      }
    }
  };

  // javascripts/items.js
  document.addEventListener("balance.item.initialized", (e) => {
    const element = document.querySelector("input[name=date]");
    const datepicker = new Datepicker(element, BOOTSTRAP_DATEPICKER_DEFAULTS);
  });
  document.addEventListener("balance.item.edit", (e) => {
    clearModalSize();
    const element = document.querySelector(".modal input[name=date]");
    const datepicker = new Datepicker(element, BOOTSTRAP_DATEPICKER_DEFAULTS);
  });
  document.addEventListener("balance.item.updated", async (e) => {
    hideModal();
    if (Object.hasOwn(e.detail, "balancePath"))
      await htmx.ajax("GET", e.detail.balancePath, { target: "#balance", swap: "outerHTML" });
    if (Object.hasOwn(e.detail, "categoriesPath"))
      await htmx.ajax("GET", e.detail.categoriesPath, { "target": "#categories", swap: "outerHTML" });
  });
  document.addEventListener("balance.items.shown", (e) => {
    const month = e.detail.month;
    const year = e.detail.year;
    const months = document.getElementById("months");
    for (const child of months.children) {
      child.classList.remove("active");
      if (child.dataset.number == month)
        child.classList.add("active");
      let url = new URL(child.getAttribute("hx-get"), window.location.origin);
      url.searchParams.set("year", year);
      child.setAttribute("hx-get", url.toString());
      htmx.process(child);
    }
    const years = document.getElementById("years");
    for (const child of years.children) {
      child.classList.remove("active");
      if (child.dataset.number == year)
        child.classList.add("active");
      let url = new URL(child.getAttribute("hx-get"), window.location.origin);
      url.searchParams.set("month", month);
      child.setAttribute("hx-get", url.toString());
      htmx.process(child);
    }
  });
  document.addEventListener("balance.item.create.error", (e) => {
    showModal();
    clearModalSize();
    const modal = bootstrap.Modal.getOrCreateInstance("#modal").show();
    $("#modal .datepicker").datepicker(BOOTSTRAP_DATEPICKER_DEFAULTS);
  });
  document.addEventListener("balance.item.created", async (e) => {
    hideModal();
    const element = document.querySelector("input[name=date]");
    const datepicker = new Datepicker(element, BOOTSTRAP_DATEPICKER_DEFAULTS);
    if (Object.hasOwn(e.detail, "balancePath"))
      await htmx.ajax("GET", e.detail.balancePath, { target: "#balance", swap: "outerHTML" });
    if (Object.hasOwn(e.detail, "categoriesPath"))
      await htmx.ajax("GET", e.detail.categoriesPath, { "target": "#categories", swap: "outerHTML" });
    if (Object.hasOwn(e.detail, "itemsPath"))
      await htmx.ajax("GET", e.detail.itemsPath, { "target": "#items" });
  });
  document.addEventListener("balance.item.deleted", async (e) => {
    hideModal();
    if (Object.hasOwn(e.detail, "balancePath"))
      await htmx.ajax("GET", e.detail.balancePath, { target: "#balance", swap: "outerHTML" });
    if (Object.hasOwn(e.detail, "categoriesPath"))
      await htmx.ajax("GET", e.detail.categoriesPath, { "target": "#categories", swap: "outerHTML" });
    if (Object.hasOwn(e.detail, "itemsPath"))
      await htmx.ajax("GET", e.detail.itemsPath, { "target": "#items" });
  });

  // javascripts/header.js
  var Header = class {
    collection;
    className = "active";
    reset() {
      this.collection.forEach((m) => m.classList.remove(this.className));
    }
    init() {
      this.collection.forEach(
        (m) => m.addEventListener("click", (e) => {
          this.reset();
          this.onClick(m);
        })
      );
    }
    onClick() {
    }
    static refresh() {
      const params = new URLSearchParams(window.location.search);
      params.set("month", this.#month());
      params.set("year", this.#year());
      console.log(params.toString());
      htmx.ajax("GET", "/items?" + params.toString(), "#items");
    }
    static #month() {
      return document.querySelector("#months div.active").dataset.number;
    }
    static #year() {
      return document.querySelector("#years div.active").dataset.number;
    }
  };

  // javascripts/years.js
  var Years = class extends Header {
    constructor() {
      super().collection = document.querySelectorAll("#years div");
      this.init();
    }
    init() {
      super.init();
      const date = /* @__PURE__ */ new Date();
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
  };

  // javascripts/months.js
  var Months = class extends Header {
    constructor() {
      super().collection = document.querySelectorAll("#months div");
      this.init();
    }
    init() {
      super.init();
      const date = /* @__PURE__ */ new Date();
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
  };

  // javascripts/application.js
  document.addEventListener("DOMContentLoaded", () => {
    new Months();
    new Years();
    Header.refresh();
  });
})();
