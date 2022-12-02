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
const day1 = (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    const result = yield axios_1.default.get("https://adventofcode.com/2022/day/1/input", {
        headers: { cookie: process.env.SESSION_COOKIE },
    });
    const day1Input = result.data
        .split("\n\n")
        .map((elf) => elf.split("\n"));
    let day1Output = 0;
    const day1OutP2 = [0, 0, 0];
    day1Input.forEach((elf) => {
        const temp = elf.reduce((acc, cur) => {
            return acc + parseInt(cur);
        }, 0);
        day1Output = temp > day1Output ? temp : day1Output;
        if (temp > day1OutP2[0]) {
            day1OutP2[0] = temp;
        }
        else if (temp > day1OutP2[1]) {
            day1OutP2[1] = temp;
        }
        else if (temp > day1OutP2[2]) {
            day1OutP2[2] = temp;
        }
    });
    return res.send({
        input: day1Input,
        output1: day1Output,
        output2: day1OutP2.reduce((acc, curr) => {
            return acc + curr;
        }, 0),
    });
});
exports.default = [{ type: "get", func: day1, path: "/day1" }];
