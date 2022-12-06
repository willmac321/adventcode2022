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
const isDistinct = (charArr, testLen) => {
    return [...(new Set(charArr))].length === testLen;
};
const day = (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        const result = yield axios_1.default.get("https://adventofcode.com/2022/day/6/input", {
            headers: { cookie: process.env.SESSION_COOKIE },
        });
        if (result.status !== 200) {
            return res.send({
                input: ["uhoh"],
            });
        }
        else {
            // values
            const dayInput = result.data.split("\n")[0].split('');
            // const dayInput = 'zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw'.split('');
            let output1 = 0;
            let output2 = 0;
            for (let i = 4; i < dayInput.length; i++) {
                if (isDistinct(dayInput.slice(i - 4, i), 4)) {
                    output1 = i;
                    break;
                }
            }
            for (let i = 14; i < dayInput.length; i++) {
                if (isDistinct(dayInput.slice(i - 14, i), 14)) {
                    output2 = i;
                    break;
                }
            }
            return res.send({
                input: dayInput,
                output1,
                output2,
            });
        }
    }
    catch (e) {
        console.error(e);
        return res.send({
            input: ["uhoh"],
        });
    }
});
exports.default = [{ type: "get", func: day, path: "/day6" }];
