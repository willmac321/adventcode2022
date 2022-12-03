import { Request, Response } from "express";
import axios from "axios";

enum Outcomes {
  win = "Win",
  lose = "Lose",
  draw = "Draw",
}
enum OppType {
  rock = "A",
  paper = "B",
  scissor = "C",
}
enum YouType {
  rock = "X",
  paper = "Y",
  scissor = "Z",
}
enum YouType2 {
  lose = "X",
  draw = "Y",
  win = "Z",
}
const isRock = (v: string): boolean => v === OppType.rock || v === YouType.rock;
const isPaper = (v: string): boolean =>
  v === OppType.paper || v === YouType.paper;
const isScissor = (v: string): boolean =>
  v === OppType.scissor || v === YouType.scissor;

const getWinDrawLose = (opp: string, you: string): Outcomes | null => {
  switch (opp) {
    case OppType.rock:
      if (isRock(you)) return Outcomes.draw;
      if (isPaper(you)) return Outcomes.win;
      if (isScissor(you)) return Outcomes.lose;
      return null;
    case OppType.paper:
      if (isRock(you)) return Outcomes.lose;
      if (isPaper(you)) return Outcomes.draw;
      if (isScissor(you)) return Outcomes.win;
      return null;
    case OppType.scissor:
      if (isRock(you)) return Outcomes.win;
      if (isPaper(you)) return Outcomes.lose;
      if (isScissor(you)) return Outcomes.draw;
      return null;
    default:
      return null;
  }
};

const getShapeForCase = (opp: OppType, outcome:YouType2): YouType | null => {
  switch (outcome) {
    case YouType2.lose:
      if (isRock(opp)) return YouType.scissor;
      if (isPaper(opp)) return YouType.rock;
      if (isScissor(opp)) return YouType.paper;
      return null;
    case YouType2.draw:
      if (isRock(opp)) return YouType.rock;
      if (isPaper(opp)) return YouType.paper;
      if (isScissor(opp)) return YouType.scissor;
      return null;
    case YouType2.win:
      if (isRock(opp)) return YouType.paper;
      if (isPaper(opp)) return YouType.scissor;
      if (isScissor(opp)) return YouType.rock;
      return null;
    default:
      return null;
  }
};

const getNumber = (numb: number, curr:Array<OppType|YouType>):number => {
        let rv: number = numb;
        if (isRock(curr[1])) rv += 1;
        if (isPaper(curr[1])) rv += 2;
        if (isScissor(curr[1])) rv += 3;
        const evaluated = getWinDrawLose(curr[0], curr[1]);
        if (evaluated === Outcomes.lose) rv += 0;
        if (evaluated === Outcomes.draw) rv += 3;
        if (evaluated === Outcomes.win) rv += 6;
        return rv;
};

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
      // values

      const dayInput = result.data
        .split("\n")
        .map((item: string) => item.split(" "));

      const output1 = dayInput.reduce((acc:number, curr:Array<OppType & YouType>) => {
        return getNumber(acc, curr);
      }, 0);

      const output2 = dayInput.reduce((acc:number, curr:Array<OppType & YouType2>) => {
        const youOut = getShapeForCase(curr[0],curr[1]);
        if (youOut === null) return acc;
        const current: Array<OppType|YouType> = [curr[0], youOut];
        return getNumber(acc, current);
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

export default [{ type: "get", func: day2, path: "/day2" }];
