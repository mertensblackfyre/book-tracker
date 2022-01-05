/*jshint esversion: 8 */
/* jshint expr: true */

const navToggle = document.querySelector(".nav-toggle");
const closeBtn = document.querySelector(".close-btn");
const nav = document.querySelector(".list");
const header = document.querySelector(".header");
const tabBtns = [...document.querySelectorAll(".btns")];
const name = document.querySelector(".name");
const descVehicle = document.querySelector(".desc-vehicle");
const image = document.querySelector(".image");

navToggle.addEventListener("click", () => {
  nav.classList.toggle("show-list");
  header.classList.toggle("header-black");
});

const getData = async () => {
  const url = fetch(`/data.json`);
  const resp = await url;
  const data = await resp.json();
  const dataCrew = data.technology;
  const selectTabBtns = () => {
    tabBtns.map((button) => {
      button.addEventListener("click", (e) => {
        dataCrew.find((item) => {
          if (item.name === e.currentTarget.value) {
            name.textContent = item.name;
            descVehicle.textContent = item.description;
            image.src = item.images.portrait;
          }
        });
      });
    });
  };
  selectTabBtns();
};

getData();
