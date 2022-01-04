/*jshint esversion: 8 */
/* jshint expr: true */

const navToggle = document.querySelector(".nav-toggle");
const closeBtn = document.querySelector(".close-btn");
const nav = document.querySelector(".list");
const header = document.querySelector(".header");
const tabBtns = [...document.querySelectorAll(".btn")];
const titlePlace = document.querySelector(".title-place");
const descPlace = document.querySelector(".desc-place");
const distance = document.querySelector(".distance");
const time = document.querySelector(".time");
const place = document.querySelector(".place");

navToggle.addEventListener("click", () => {
  nav.classList.toggle("show-list");
  header.classList.toggle("header-black");
});

const getData = async () => {
  const url = fetch(`/data.json`);
  const resp = await url;
  const data = await resp.json();
  const dataDestination = data.destinations;

  const selectTabBtns = () => {
    tabBtns.map((button) => {
      button.addEventListener("click", (e) => {
        dataDestination.find((item) => {
          if (item.name === e.currentTarget.textContent) {
            // if (e.currentTarget !== item) {
            //   e.currentTarget.classList.remove("active-tab");
            // }
            place.src = item.images.png;
            titlePlace.textContent = item.name;
            descPlace.textContent = item.description;
            time.textContent = item.travel;
            distance.textContent = item.distance;
          }
        });
      });
    });
  };
  selectTabBtns();
};

getData();
