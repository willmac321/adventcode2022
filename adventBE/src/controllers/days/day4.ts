import { Request, Response } from "express";
import axios from "axios";

const day = async (req: Request, res: Response): Promise<Response> => {
  try {
    const result = await axios.get(
      "https://adventofcode.com/2022/day/4/input",
      {
        headers: { cookie: process.env.SESSION_COOKIE },
      }
    );

    if (result.status !== 200) {
      return res.send({
        input: ["uhoh"],
      });
    } else {
      const dayInput = result.data.split("\n").filter((s: string) => s);

      const output1 = dayInput.reduce((acc: number, curr: string): number => {
        const [set1, set2] = curr.split(",");
        const [set1x, set1y] = set1.split("-");
        const [set2x, set2y] = set2.split("-");

        // if set1 surrounds set2 and they are not equl
        if (
          parseInt(set1x) <= parseInt(set2x) &&
          parseInt(set1y) >= parseInt(set2y)
        )
          return ++acc;

        // if set2 surrounds set1 and they are not equl
        if (
          parseInt(set1x) >= parseInt(set2x) &&
          parseInt(set1y) <= parseInt(set2y)
        )
          return ++acc;

        return acc;
      }, 0);

      const output2 = dayInput.reduce((acc: number, curr: string): number => {
        const [set1, set2] = curr.split(",");
        const [set1x, set1y] = set1.split("-");
        const [set2x, set2y] = set2.split("-");

        // set set1min is in set2
        if (
          parseInt(set1x) >= parseInt(set2x) &&
          parseInt(set1x) <= parseInt(set2y)
        )
          return ++acc;

        // set set2min is in set1
        if (
          parseInt(set2x) >= parseInt(set1x) &&
          parseInt(set2x) <= parseInt(set1y)
        )
          return ++acc;

        // set set1max is in set1
        if (
          parseInt(set1y) >= parseInt(set2x) &&
          parseInt(set1y) <= parseInt(set2y)
        )
          return ++acc;

        // set set2max is in set1
        if (
          parseInt(set2y) >= parseInt(set1x) &&
          parseInt(set2y) <= parseInt(set1y)
        )
          return ++acc;

        return acc;
      }, 0);

      return res.send({
        input: dayInput,
        output1,
        output2,
      });
    }
  } catch (e) {
    console.error(e);

    return res.send({
      input: ["uhoh"],
    });
  }
};

export default [{ type: "get", func: day, path: "/day4" }];
