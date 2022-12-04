import App from "./App";
import Day1 from "./days/Day1";
import Day2 from "./days/Day2";
import Day3 from "./days/Day3";
import Day4 from "./days/Day4";
import { RouteObject } from "react-router-dom";

const routes: RouteObject[] = [
  {
    path: "/",
    element: <App />,
  },
  { path: "day1", element: <Day1 /> },
  { path: "day2", element: <Day2 /> },
  { path: "day3", element: <Day3 /> },
  { path: "day4", element: <Day4 /> },
];

export default routes;
