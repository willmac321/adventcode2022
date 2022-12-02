"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const axios_1 = __importDefault(require("axios"));
const day2 = (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        const result = yield axios_1.default.get("https://adventofcode.com/2022/day/2/input", {
            headers: { cookie: process.env.SESSION_COOKIE },
        });
        if (result.status !== 200) {
            return res.send({
                input: ["uhoh"],
            });
        }
        else {
            const dayInput = result.data;
            return res.send({
                input: dayInput,
            });
        }
    }
    catch (_a) {
        return res.send({
            input: ["uhoh"],
        });
    }
});
exports.default = [{ type: "get", func: day2, path: "/day2" }];
