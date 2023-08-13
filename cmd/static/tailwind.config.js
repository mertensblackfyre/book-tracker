const withMT = require("@material-tailwind/html/utils/withMT");

module.exports = withMT({
   content: ["./index.html,", "./login.html"],
   theme: {
      extend: {},
   },
   plugins: [require("daisyui")],
});
