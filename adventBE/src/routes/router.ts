import express from "express";

import controllers from "../controllers/index";
const routes = express.Router();

controllers.forEach((controller) => {
  switch (controller.type) {
    case "get":
      routes.get(controller.path, (req, res) => controller.func(req, res));
      break;
    default:
      break;
  }
});

export {routes};
