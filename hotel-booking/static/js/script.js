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
  if (!panel) return;

  const baseTop = panel.getBoundingClientRect().top + window.scrollY;

  const onScroll = () => {
    if (window.scrollY > baseTop - 80) panel.classList.add("nav-pinned");
    else panel.classList.remove("nav-pinned");
  };

  window.addEventListener("scroll", onScroll, { passive: true });
  onScroll();
})();
