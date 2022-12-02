import App from "./App";
import Day1 from "./days/Day1";
import { RouteObject } from "react-router-dom";

const routes: RouteObject[] = [
  {
    path: "/",
    element: <App />,
  },
  { path: "day1", element: <Day1 /> },
];

export default routes;
