import Axios from "axios";
const instance = Axios.create({
  baseURL: "http://localhost:3000/api/",
  headers: { "Content-Type": "application/json" },
});
const baseURL = "http://localhost:3000/api/";

export default {
  createFairy: (_, payload) => {
    Axios.post(
      baseURL + "fairy",
      {
        name: payload.name,
        element: payload.element,
        percentage: parseInt(payload.percentage),
      },
      instance
    )
      .then((response) => {
        console.log(response);
      })
      .catch(function(error) {
        console.log(error);
      });
  },
  getFairy: () => {
    Axios.get("fairy", instance)
      .then((response) => {
        console.log(response);
      })
      .catch(function(error) {
        console.log(error);
      });
  },
};
