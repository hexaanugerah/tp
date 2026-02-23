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
  if (!nav) return;

  const hideNav = () => {
    nav.classList.add("nav-hidden");
    nav.classList.remove("nav-reveal");
  };
  const showNav = () => {
    nav.classList.remove("nav-hidden");
    nav.classList.add("nav-reveal");
  };

  const onScrollNav = () => {
    if (window.scrollY <= 8) {
      showNav();
    } else {
      hideNav();
    }
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
})();
