/*jshint esversion: 6 */

const navToggle = document.querySelector(".nav-toggle");
const closeBtn = document.querySelector(".close-btn");
const nav = document.querySelector(".list");
const header = document.querySelector(".header");

navToggle.addEventListener("click", () => {
  nav.classList.toggle("show-list");
  header.classList.toggle("header-black");
});
// closeBtn.addEventListener("click", () => {
//   if (nav.classList.contains("show-list")) {
//     nav.classList.remove("show-list");
//     header.classList.remove("header-black");
//   }
// });
