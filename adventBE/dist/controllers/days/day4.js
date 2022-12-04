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
        const result = yield axios_1.default.get("https://adventofcode.com/2022/day/4/input", {
            headers: { cookie: process.env.SESSION_COOKIE },
        });
        if (result.status !== 200) {
            return res.send({
                input: ["uhoh"],
            });
        }
        else {
            const dayInput = result.data.split("\n").filter((s) => s);
            const output1 = dayInput.reduce((acc, curr) => {
                const [set1, set2] = curr.split(",");
                const [set1x, set1y] = set1.split("-");
                const [set2x, set2y] = set2.split("-");
                // if set1 surrounds set2 and they are not equl
                if (parseInt(set1x) <= parseInt(set2x) &&
                    parseInt(set1y) >= parseInt(set2y))
                    return ++acc;
                // if set2 surrounds set1 and they are not equl
                if (parseInt(set1x) >= parseInt(set2x) &&
                    parseInt(set1y) <= parseInt(set2y))
                    return ++acc;
                return acc;
            }, 0);
            const output2 = dayInput.reduce((acc, curr) => {
                const [set1, set2] = curr.split(",");
                const [set1x, set1y] = set1.split("-");
                const [set2x, set2y] = set2.split("-");
                // set set1min is in set2
                if (parseInt(set1x) >= parseInt(set2x) &&
                    parseInt(set1x) <= parseInt(set2y))
                    return ++acc;
                // set set2min is in set1
                if (parseInt(set2x) >= parseInt(set1x) &&
                    parseInt(set2x) <= parseInt(set1y))
                    return ++acc;
                // set set1max is in set1
                if (parseInt(set1y) >= parseInt(set2x) &&
                    parseInt(set1y) <= parseInt(set2y))
                    return ++acc;
                // set set2max is in set1
                if (parseInt(set2y) >= parseInt(set1x) &&
                    parseInt(set2y) <= parseInt(set1y))
                    return ++acc;
                return acc;
            }, 0);
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
exports.default = [{ type: "get", func: day, path: "/day4" }];
