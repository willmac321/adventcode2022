import { Request, Response } from "express";
import axios from "axios";

const isDistinct = (
  charArr: string[],
  testLen: number
): boolean => {
  return [...(new Set(charArr))].length === testLen;
};

const day = async (req: Request, res: Response): Promise<Response> => {
  try {
    const result = await axios.get(
      "https://adventofcode.com/2022/day/6/input",
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

      const dayInput = result.data.split("\n")[0].split('')
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
  } catch (e) {
    console.error(e);

    return res.send({
      input: ["uhoh"],
    });
  }
};

export default [{ type: "get", func: day, path: "/day6" }];
