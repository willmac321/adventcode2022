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
        const result = yield axios_1.default.get("https://adventofcode.com/2022/day/5/input", {
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
            // const dayInput = [
            //   "     [D]     ",
            //   " [N] [C]     ",
            //   " [Z] [M] [P] ",
            //   "1   2   3",
            //   "move 1 from 2 to 1",
            //   "move 3 from 1 to 3",
            //   "move 2 from 2 to 1",
            //   "move 1 from 1 to 2",
            // ];
            let stacks = [[""]];
            let rows = [];
            let directions = [];
            for (let i = 0; i < dayInput.length; i++) {
                if (dayInput[i].match(/[1-9]/g) != null) {
                    const columnCount = dayInput[i].split("   ").length;
                    stacks = Array(columnCount).fill([]);
                    rows = dayInput
                        .slice(0, i)
                        .map((row) => {
                        return row.replace(/\s{4}/g, " x").trim().split(" ");
                    })
                        .reverse();
                    directions = dayInput.slice(i + 1);
                    break;
                }
            }
            for (let y = 0; y < rows.length; y++) {
                const row = rows[y];
                for (let x = 0; x < row.length; x++) {
                    const item = row[x];
                    if (item !== "x")
                        stacks[x] = [...stacks[x], item];
                }
            }
            const stacks2 = stacks.map(stack => [...stack]);
            // part 1
            directions.forEach((raw) => {
                const direction = raw.split(" ");
                const from = parseInt(direction[3]) - 1;
                const to = parseInt(direction[5]) - 1;
                for (let i = 0; i < parseInt(direction[1]); i++) {
                    const pop = stacks[from].pop();
                    if (pop !== undefined)
                        stacks[to] = [...stacks[to], pop];
                }
            });
            directions.forEach((raw) => {
                const direction = raw.split(" ");
                const move = parseInt(direction[1]);
                const from = parseInt(direction[3]) - 1;
                const to = parseInt(direction[5]) - 1;
                const pop = stacks2[from].slice(-move);
                stacks2[from] = stacks2[from].slice(0, -move);
                if (pop !== undefined)
                    stacks2[to] = [...stacks2[to], ...pop];
            });
            const output1 = stacks
                .map((item) => item[item.length - 1])
                .join("")
                .replace(/\[/g, "")
                .replace(/\]/g, "");
            const output2 = stacks2
                .map((item) => item[item.length - 1])
                .join("")
                .replace(/\[/g, "")
                .replace(/\]/g, "");
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
exports.default = [{ type: "get", func: day, path: "/day5" }];
