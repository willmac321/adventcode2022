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

      const dayInput = result.data.split("\n").filter((s:string)=>s);

      return res.send({
        input: dayInput,
        output1: "",
        output2: "",
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
