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
const day = (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        const result = yield axios_1.default.get("https://adventofcode.com/2022/day/3/input", {
            headers: { cookie: process.env.SESSION_COOKIE },
        });
        if (result.status !== 200) {
            return res.send({
                input: ["uhoh"],
            });
        }
        else {
            // values
            const dayInput = result.data.split("\n").filter((s) => s);
            const dayInput1 = dayInput.map((item) => {
                // seperate string in half
                return [item.slice(0, item.length / 2), item.slice(item.length / 2)];
            });
            const output1 = dayInput1.reduce((acc, curr) => {
                const [first, second] = curr;
                const hash = {};
                first.split("").forEach((item) => {
                    if (second.split("").includes(item)) {
                        let val = item.charCodeAt(0) - 96;
                        val = val < 1 ? val + 6 + 52 : val;
                        hash[item] = val;
                    }
                });
                return (acc +
                    Object.values(hash).reduce((acc, curr) => acc + curr, 0));
            }, 0);
            let output2 = 0;
            for (let i = 0; i < dayInput.length; i += 3) {
                const s1 = [...new Set(dayInput[i].split(""))];
                const s2 = [...new Set(dayInput[i + 1].split(""))];
                const s3 = [...new Set(dayInput[i + 2].split(""))];
                const out = s1.filter((item) => {
                    return s2.includes(item) && s3.includes(item);
                })[0];
                const val = out.charCodeAt(0) - 96;
                output2 += val < 1 ? val + 6 + 52 : val;
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
exports.default = [{ type: "get", func: day, path: "/day3" }];
