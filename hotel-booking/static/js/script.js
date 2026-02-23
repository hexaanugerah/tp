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
  const brand = document.getElementById("brandDragHandle");
  const resetPinnedBtn = document.getElementById("resetPinnedBtn");
  if (!panel || !brand) return;

  const pinPanel = () => panel.classList.add("nav-pinned");
  const unpinPanel = () => panel.classList.remove("nav-pinned");

  if (resetPinnedBtn) {
    resetPinnedBtn.addEventListener("click", () => {
      if (panel.classList.contains("nav-pinned")) unpinPanel();
      else pinPanel();
    });
  }

  let startY = null;
  const start = (y) => {
    startY = y;
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

  brand.addEventListener("mousedown", (e) => start(e.clientY));
  window.addEventListener("mousemove", (e) => move(e.clientY));
  window.addEventListener("mouseup", end);

  brand.addEventListener("touchstart", (e) => start(e.touches[0].clientY), { passive: true });
  window.addEventListener("touchmove", (e) => move(e.touches[0].clientY), { passive: true });
  window.addEventListener("touchend", end, { passive: true });
})();
