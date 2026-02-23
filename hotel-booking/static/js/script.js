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
  const resetPinnedBtn = document.getElementById("resetPinnedBtn");
  const topPullZone = document.getElementById("topPullZone");
  if (!panel) return;

  const baseTop = panel.getBoundingClientRect().top + window.scrollY;

  const pinPanel = () => panel.classList.add("nav-pinned");
  const unpinPanel = () => panel.classList.remove("nav-pinned");

  const onScroll = () => {
    if (window.scrollY > baseTop - 80) pinPanel();
    else unpinPanel();
  };

  if (resetPinnedBtn) {
    resetPinnedBtn.addEventListener("click", () => {
      if (panel.classList.contains("nav-pinned")) unpinPanel();
      else pinPanel();
    });
  }

  let startY = null;
  const start = (y) => {
    if (window.scrollY <= 5) startY = y;
  };
  const move = (y) => {
    if (startY == null) return;
    if (y - startY > 45) {
      pinPanel();
      startY = null;
    }
  };
  const end = () => {
    startY = null;
  };

  if (topPullZone) {
    topPullZone.addEventListener("mousedown", (e) => start(e.clientY));
    window.addEventListener("mousemove", (e) => move(e.clientY));
    window.addEventListener("mouseup", end);

    topPullZone.addEventListener("touchstart", (e) => start(e.touches[0].clientY), { passive: true });
    window.addEventListener("touchmove", (e) => move(e.touches[0].clientY), { passive: true });
    window.addEventListener("touchend", end, { passive: true });
  }

  window.addEventListener("scroll", onScroll, { passive: true });
  onScroll();
})();
