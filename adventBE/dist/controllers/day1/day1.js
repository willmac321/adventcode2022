"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const day1 = (req, res) => {
    return res.send({ hey: 'hi' });
};
exports.default = [{ type: "get", func: day1, path: "/day1" }];
