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
var Outcomes;
(function (Outcomes) {
    Outcomes["win"] = "Win";
    Outcomes["lose"] = "Lose";
    Outcomes["draw"] = "Draw";
})(Outcomes || (Outcomes = {}));
var OppType;
(function (OppType) {
    OppType["rock"] = "A";
    OppType["paper"] = "B";
    OppType["scissor"] = "C";
})(OppType || (OppType = {}));
var YouType;
(function (YouType) {
    YouType["rock"] = "X";
    YouType["paper"] = "Y";
    YouType["scissor"] = "Z";
})(YouType || (YouType = {}));
var YouType2;
(function (YouType2) {
    YouType2["lose"] = "X";
    YouType2["draw"] = "Y";
    YouType2["win"] = "Z";
})(YouType2 || (YouType2 = {}));
const isRock = (v) => v === OppType.rock || v === YouType.rock;
const isPaper = (v) => v === OppType.paper || v === YouType.paper;
const isScissor = (v) => v === OppType.scissor || v === YouType.scissor;
const getWinDrawLose = (opp, you) => {
    switch (opp) {
        case OppType.rock:
            if (isRock(you))
                return Outcomes.draw;
            if (isPaper(you))
                return Outcomes.win;
            if (isScissor(you))
                return Outcomes.lose;
            return null;
        case OppType.paper:
            if (isRock(you))
                return Outcomes.lose;
            if (isPaper(you))
                return Outcomes.draw;
            if (isScissor(you))
                return Outcomes.win;
            return null;
        case OppType.scissor:
            if (isRock(you))
                return Outcomes.win;
            if (isPaper(you))
                return Outcomes.lose;
            if (isScissor(you))
                return Outcomes.draw;
            return null;
        default:
            return null;
    }
};
const getShapeForCase = (opp, outcome) => {
    switch (outcome) {
        case YouType2.lose:
            if (isRock(opp))
                return YouType.scissor;
            if (isPaper(opp))
                return YouType.rock;
            if (isScissor(opp))
                return YouType.paper;
            return null;
        case YouType2.draw:
            if (isRock(opp))
                return YouType.rock;
            if (isPaper(opp))
                return YouType.paper;
            if (isScissor(opp))
                return YouType.scissor;
            return null;
        case YouType2.win:
            if (isRock(opp))
                return YouType.paper;
            if (isPaper(opp))
                return YouType.scissor;
            if (isScissor(opp))
                return YouType.rock;
            return null;
        default:
            return null;
    }
};
const getNumber = (numb, curr) => {
    let rv = numb;
    if (isRock(curr[1]))
        rv += 1;
    if (isPaper(curr[1]))
        rv += 2;
    if (isScissor(curr[1]))
        rv += 3;
    const evaluated = getWinDrawLose(curr[0], curr[1]);
    if (evaluated === Outcomes.lose)
        rv += 0;
    if (evaluated === Outcomes.draw)
        rv += 3;
    if (evaluated === Outcomes.win)
        rv += 6;
    return rv;
};
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
            // values
            const dayInput = result.data
                .split("\n")
                .map((item) => item.split(" "));
            const output1 = dayInput.reduce((acc, curr) => {
                return getNumber(acc, curr);
            }, 0);
            const output2 = dayInput.reduce((acc, curr) => {
                const youOut = getShapeForCase(curr[0], curr[1]);
                if (youOut === null)
                    return acc;
                const current = [curr[0], youOut];
                return getNumber(acc, current);
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
exports.default = [{ type: "get", func: day2, path: "/day2" }];
