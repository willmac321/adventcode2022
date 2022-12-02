import { Request, Response } from "express";
const day1 = (req: Request, res: Response): Response => {
  return res.send({hey:'hi'});
};

export default [{ type: "get", func: day1, path: "/day1" }];
