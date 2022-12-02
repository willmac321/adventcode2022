import { Request, Response } from "express";
import axios from "axios";
const day1 = async (req: Request, res: Response): Promise<Response> => {
  const result = await axios.get("https://adventofcode.com/2022/day/1/input", {
    headers: { cookie: process.env.SESSION_COOKIE },
  });

  const day1Input = result.data
    .split("\n\n")
    .map((elf: string) => elf.split("\n"));
  let day1Output = 0;
  const day1OutP2 = [0, 0, 0];

  day1Input.forEach((elf: string[]) => {
    const temp = elf.reduce((acc: number, cur: string) => {
      return acc + parseInt(cur);
    }, 0);
    day1Output = temp > day1Output ? temp : day1Output;
    if (temp > day1OutP2[0]) {
      day1OutP2[0] = temp;
    } else if (temp > day1OutP2[1]) {
      day1OutP2[1] = temp;
    } else if (temp > day1OutP2[2]) {
      day1OutP2[2] = temp;
    }
  });

  return res.send({
    input: day1Input,
    output1: day1Output,
    output2: day1OutP2.reduce((acc: number, curr: number) => {
      return acc + curr;
    }, 0),
  });
};

export default [{ type: "get", func: day1, path: "/day1" }];
