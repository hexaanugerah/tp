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

  // Popup selections for city/date/guest
  const overlay = document.getElementById("searchPopup");
  if (!overlay) return;
  const title = document.getElementById("popupTitle");
  const optionsWrap = document.getElementById("popupOptions");
  const closeBtn = document.getElementById("closePopupBtn");

  const cityInput = document.getElementById("cityInput");
  const dateInput = document.getElementById("dateInput");
  const guestInput = document.getElementById("guestInput");
  const cityTrigger = document.getElementById("cityTrigger");
  const dateTrigger = document.getElementById("dateTrigger");
  const guestTrigger = document.getElementById("guestTrigger");

  const datasets = {
    city: {
      title: "Pilih Kota Hotel",
      values: ["Balikpapan", "Jakarta", "Bandung", "Yogyakarta", "Surabaya", "Makassar", "Bali", "Medan"],
      apply: (v) => {
        cityInput.value = v;
        cityTrigger.textContent = v;
      },
    },
    date: {
      title: "Pilih Tanggal",
      values: [
        "Sen, 23 Feb 2026 - Sel, 24 Feb 2026",
        "Rab, 25 Feb 2026 - Kam, 26 Feb 2026",
        "Jum, 27 Feb 2026 - Sab, 28 Feb 2026",
      ],
      apply: (v) => {
        dateInput.value = v;
        dateTrigger.textContent = v;
      },
    },
    guest: {
      title: "Pilih Tamu dan Kamar",
      values: ["1 Dewasa, 1 Kamar", "2 Dewasa, 1 Kamar", "2 Dewasa, 2 Kamar", "3 Dewasa, 2 Kamar"],
      apply: (v) => {
        guestInput.value = v;
        guestTrigger.textContent = v;
      },
    },
  };

  const openPopup = (key) => {
    const data = datasets[key];
    if (!data) return;
    title.textContent = data.title;
    optionsWrap.innerHTML = "";
    data.values.forEach((v) => {
      const btn = document.createElement("button");
      btn.type = "button";
      btn.textContent = v;
      btn.addEventListener("click", () => {
        data.apply(v);
        overlay.classList.remove("open");
      });
      optionsWrap.appendChild(btn);
    });
    overlay.classList.add("open");
  };

  document.querySelectorAll(".popup-trigger").forEach((el) => {
    el.addEventListener("click", () => openPopup(el.dataset.popup));
  });
  closeBtn.addEventListener("click", () => overlay.classList.remove("open"));
  overlay.addEventListener("click", (e) => {
    if (e.target === overlay) overlay.classList.remove("open");
  });
})();
