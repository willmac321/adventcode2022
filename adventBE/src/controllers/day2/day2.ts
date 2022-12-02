import { Request, Response } from "express";
import axios from "axios";
const day2 = async (req: Request, res: Response): Promise<Response> => {
  try {
    const result = await axios.get(
      "https://adventofcode.com/2022/day/2/input",
      {
        headers: { cookie: process.env.SESSION_COOKIE },
      }
    );

    if (result.status !== 200) {
      return res.send({
        input: ["uhoh"],
      });
    } else {
      const dayInput = result.data;

      return res.send({
        input: dayInput,
      });
    }
  } catch {
    return res.send({
      input: ["uhoh"],
    });
  }
};

export default [{ type: "get", func: day2, path: "/day2" }];
