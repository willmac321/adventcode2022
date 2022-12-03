import { Request, Response } from "express";
import axios from "axios";

const day = async (req: Request, res: Response): Promise<Response> => {
  try {
    const result = await axios.get(
      "https://adventofcode.com/2022/day/3/input",
      {
        headers: { cookie: process.env.SESSION_COOKIE },
      }
    );

    if (result.status !== 200) {
      return res.send({
        input: ["uhoh"],
      });
    } else {
      // values

      const dayInput = result.data.split("\n").filter((s: string) => s);

      const dayInput1 = dayInput.map((item: string) => {
        // seperate string in half
        return [item.slice(0, item.length / 2), item.slice(item.length / 2)];
      });

      const output1 = dayInput1.reduce((acc: number, curr: string[]) => {
        const [first, second] = curr;
        const hash: Record<string, number> | Record<string, never> = {};

        first.split("").forEach((item: string) => {
          if (second.split("").includes(item)) {
            let val = item.charCodeAt(0) - 96;
            val = val < 1 ? val + 6 + 52 : val;
            hash[item] = val;
          }
        });

        return (
          acc +
          Object.values(hash).reduce(
            (acc: number, curr: number) => acc + curr,
            0
          )
        );
      }, 0);

      let output2: number = 0;

      for (let i = 0; i < dayInput.length; i += 3) {
        const s1 = [...new Set(dayInput[i].split(""))];
        const s2 = [...new Set(dayInput[i + 1].split(""))];
        const s3 = [...new Set(dayInput[i + 2].split(""))];

        const out = s1.filter((item) => {
          return s2.includes(item) && s3.includes(item);
        })[0] as string;

        const val: number = out.charCodeAt(0) - 96;
        output2 += val < 1 ? val + 6 + 52 : val;
      }

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

export default [{ type: "get", func: day, path: "/day3" }];
