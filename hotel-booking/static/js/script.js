function openHotelModal(name, desc, city) {
  const m = document.getElementById("hotelModal");
  if (!m) return;
  document.getElementById("mTitle").innerText = name;
  document.getElementById("mDesc").innerText = desc;
  document.getElementById("mCity").innerText = city;
  m.style.display = "block";
}
function closeModal() {
  const m = document.getElementById("hotelModal");
  if (m) m.style.display = "none";
}
window.onclick = (e) => {
  const m = document.getElementById("hotelModal");
  if (m && e.target === m) m.style.display = "none";
};

(function () {
  const panel = document.getElementById("stickySearch");
  if (panel) {
    const baseTop = panel.getBoundingClientRect().top + window.scrollY;
    const onScrollPanel = () => {
      if (window.scrollY > baseTop - 80) panel.classList.add("nav-pinned");
      else panel.classList.remove("nav-pinned");
    };
    window.addEventListener("scroll", onScrollPanel, { passive: true });
    onScrollPanel();
  }

  const nav = document.getElementById("mainNavbar");
  const pullZone = document.getElementById("navPullZone");
  if (nav) {
    const hideNav = () => {
      nav.classList.add("nav-hidden");
      nav.classList.remove("nav-reveal");
      document.body.classList.remove("navbar-visible");
    };
    const showNav = () => {
      nav.classList.remove("nav-hidden");
      nav.classList.add("nav-reveal");
      document.body.classList.add("navbar-visible");
    };
    const onScrollNav = () => {
      if (window.scrollY <= 8) showNav();
      else hideNav();
    };

    let startY = null;
    const start = (y) => {
      if (window.scrollY > 8) startY = y;
    };
    const move = (y) => {
      if (startY == null) return;
      if (y - startY > 42) {
        showNav();
        startY = null;
      }
    };
    const end = () => {
      startY = null;
    };

    if (pullZone) {
      pullZone.addEventListener("mousedown", (e) => start(e.clientY));
      window.addEventListener("mousemove", (e) => move(e.clientY));
      window.addEventListener("mouseup", end);
      pullZone.addEventListener("touchstart", (e) => start(e.touches[0].clientY), { passive: true });
      window.addEventListener("touchmove", (e) => move(e.touches[0].clientY), { passive: true });
      window.addEventListener("touchend", end, { passive: true });
    }

    window.addEventListener("scroll", onScrollNav, { passive: true });
    onScrollNav();
  }


  // Rooms page section nav (sticky + scroll spy + smooth underline)
  const roomsNav = document.getElementById("roomsSectionNav");
  if (roomsNav) {
    const navLinks = Array.from(roomsNav.querySelectorAll(".rooms-section-link[href^='#']"));
    const underline = document.getElementById("roomsNavUnderline");
    const sections = navLinks
      .map((l) => document.querySelector(l.getAttribute("href")))
      .filter(Boolean);

    const setActive = (link) => {
      navLinks.forEach((l) => l.classList.remove("active"));
      link.classList.add("active");
      if (!underline) return;
      const rect = link.getBoundingClientRect();
      const navRect = roomsNav.getBoundingClientRect();
      underline.style.left = `${rect.left - navRect.left}px`;
      underline.style.width = `${rect.width}px`;
      underline.style.opacity = "1";
    };

    navLinks.forEach((link) => {
      link.addEventListener("click", (e) => {
        const id = link.getAttribute("href");
        const target = document.querySelector(id);
        if (!target) return;
        e.preventDefault();
        target.scrollIntoView({ behavior: "smooth", block: "start" });
        setActive(link);
        history.replaceState(null, "", id);
      });
    });

    const onScrollSpy = () => {
      let current = sections[0];
      sections.forEach((sec) => {
        if (window.scrollY >= sec.offsetTop - 170) current = sec;
      });
      const activeLink = navLinks.find((l) => l.getAttribute("href") === `#${current.id}`);
      if (activeLink) setActive(activeLink);
    };

    window.addEventListener("scroll", onScrollSpy, { passive: true });
    window.addEventListener("resize", onScrollSpy);
    onScrollSpy();
  }

  // Popup selections for city/date/guest
  const overlay = document.getElementById("searchPopup");
  if (!overlay) return;
  const title = document.getElementById("popupTitle");
  const optionsWrap = document.getElementById("popupOptions");
  const closeBtn = document.getElementById("closePopupBtn");

  const cityInput = document.getElementById("cityInput");
  const checkinInput = document.getElementById("checkinInput");
  const checkoutInput = document.getElementById("checkoutInput");
  const dateInput = document.getElementById("dateInput");
  const guestInput = document.getElementById("guestInput");
  const cityTrigger = document.getElementById("cityTrigger");
  const checkinTrigger = document.getElementById("checkinTrigger");
  const checkoutTrigger = document.getElementById("checkoutTrigger");
  const guestTrigger = document.getElementById("guestTrigger");

  const syncDateRange = () => {
    if (dateInput && checkinInput && checkoutInput) {
      dateInput.value = `${checkinInput.value} - ${checkoutInput.value}`;
    }
    if (checkinTrigger && checkinInput) {
      checkinTrigger.innerHTML = `<span class="date-caption">Check-in</span><strong>${checkinInput.value}</strong>`;
    }
    if (checkoutTrigger && checkoutInput) {
      checkoutTrigger.innerHTML = `<span class="date-caption">Check-out</span><strong>${checkoutInput.value}</strong>`;
    }
  };

  const datasets = {
    city: {
      title: "Pilih Kota Hotel",
      values: ["📍 Near me", "Balikpapan", "Jakarta", "Bandung", "Yogyakarta", "Surabaya", "Makassar", "Bali", "Medan"],
      apply: (v) => {
        if (cityInput) cityInput.value = v;
        if (cityTrigger) cityTrigger.textContent = v;
      },
    },
    guest: {
      title: "Pilih Tamu dan Kamar",
      values: ["1 Dewasa, 1 Kamar", "2 Dewasa, 1 Kamar", "2 Dewasa, 2 Kamar", "3 Dewasa, 2 Kamar"],
      apply: (v) => {
        if (guestInput) guestInput.value = v;
        if (guestTrigger) guestTrigger.textContent = v;
      },
    },
  };

  const parseDay = (value) => parseInt((value || "01/03/2026").split("/")[0], 10);
  const buildDate = (day) => `${String(day).padStart(2, "0")}/03/2026`;

  const renderCalendar = (mode) => {
    title.textContent = mode === "checkin" ? "Pilih Check-in" : "Pilih Check-out";
    const currentDay = mode === "checkin" ? parseDay(checkinInput?.value) : parseDay(checkoutInput?.value);
    const prevMonthTail = [23, 24, 25, 26, 27, 28];
    const days = Array.from({ length: 31 }, (_, i) => i + 1);

    optionsWrap.innerHTML = `
      <div class="calendar-shell">
        <div class="calendar-header">
          <button type="button" class="cal-nav">◀</button>
          <div class="cal-month">March 2026</div>
          <button type="button" class="cal-nav">▶</button>
        </div>
        <div class="calendar-weekdays">
          <span>Mon</span><span>Tue</span><span>Wed</span><span>Thu</span><span>Fri</span><span>Sat</span><span>Sun</span>
        </div>
        <div class="calendar-grid" id="calendarGrid"></div>
        <div class="popup-foot"><button type="button" id="closeCalBtn">Close</button></div>
      </div>`;

    const grid = optionsWrap.querySelector("#calendarGrid");
    prevMonthTail.forEach((d) => {
      const b = document.createElement("button");
      b.type = "button";
      b.className = "cal-day muted";
      b.textContent = String(d);
      grid.appendChild(b);
    });
    days.forEach((d) => {
      const b = document.createElement("button");
      b.type = "button";
      b.className = `cal-day${d === currentDay ? " selected" : ""}`;
      b.textContent = String(d);
      b.addEventListener("click", () => {
        if (mode === "checkin" && checkinInput) {
          checkinInput.value = buildDate(d);
          if (checkoutInput && parseDay(checkoutInput.value) < d) checkoutInput.value = buildDate(d + 1 <= 31 ? d + 1 : d);
        }
        if (mode === "checkout" && checkoutInput) {
          const inDay = parseDay(checkinInput?.value);
          checkoutInput.value = buildDate(d < inDay ? inDay + 1 : d);
        }
        syncDateRange();
        overlay.classList.remove("open");
      });
      grid.appendChild(b);
    });

    const closeCalBtn = optionsWrap.querySelector("#closeCalBtn");
    if (closeCalBtn) closeCalBtn.addEventListener("click", () => overlay.classList.remove("open"));
  };

  const openPopup = (key, triggerEl) => {
    if (key === "checkin" || key === "checkout") {
      renderCalendar(key);
    } else {
      const data = datasets[key];
      if (!data) return;
      title.textContent = data.title;
      optionsWrap.innerHTML = "";
      data.values.forEach((v) => {
        const btn = document.createElement("button");
        btn.type = "button";
        btn.className = "popup-option";
        btn.textContent = v;
        btn.addEventListener("click", () => {
          data.apply(v);
          overlay.classList.remove("open");
        });
        optionsWrap.appendChild(btn);
      });
    }

    if (triggerEl) {
      const rect = triggerEl.getBoundingClientRect();
      overlay.style.top = `${rect.bottom + 8}px`;
      overlay.style.left = `${Math.max(16, rect.left)}px`;
      overlay.style.width = `${Math.max(rect.width, 360)}px`;
    }
    overlay.classList.add("open");
  };
  syncDateRange();

  document.querySelectorAll(".popup-trigger").forEach((el) => {
    el.addEventListener("click", () => openPopup(el.dataset.popup, el));
  });
  closeBtn.addEventListener("click", () => overlay.classList.remove("open"));
  document.addEventListener("click", (e) => {
    const insidePopup = overlay.contains(e.target);
    const onTrigger = e.target.closest && e.target.closest(".popup-trigger");
    if (!insidePopup && !onTrigger) overlay.classList.remove("open");
  });
})();
