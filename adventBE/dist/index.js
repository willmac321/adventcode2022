"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const application_1 = require("./application");
const router_1 = require("./routes/router");
const port = 4000;
const app = new application_1.App(port, [router_1.routes]);
app.listen();
