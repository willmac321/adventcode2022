"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const day1_1 = __importDefault(require("./day1/day1"));
const day2_1 = __importDefault(require("./day2/day2"));
exports.default = [...day1_1.default, ...day2_1.default];
