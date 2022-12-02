"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.routes = void 0;
const express_1 = __importDefault(require("express"));
const index_1 = __importDefault(require("../controllers/index"));
const routes = express_1.default.Router();
exports.routes = routes;
index_1.default.forEach((controller) => {
    switch (controller.type) {
        case "get":
            routes.get(controller.path, (req, res) => controller.func(req, res));
            break;
        default:
            break;
    }
});
