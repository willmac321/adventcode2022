import { useEffect, useState } from "react";
import DayTemplate from "../DayTemplate";

type errorType = {
  message: string;
} | null;

type resultType = {
  input: string;
  output1: string;
  output2: string;
} | null;

function Day2(): JSX.Element {
  const [result, setResult] = useState<resultType>(null);
  const [isLoaded, setIsLoaded] = useState(false);
  const [error, setError] = useState<errorType>(null);

  useEffect(() => {
    fetch("/api/day2")
      .then(async (res) => await res.json())
      .then(
        (result) => {
          setIsLoaded(true);
          const inp = result.input
            .map((input: string[]) => {
              return JSON.stringify(input);
            })
            .join("\n\r");

          setResult({
            input: inp,
            output1: result.output1,
            output2: result.output2,
          });
        },
        // Note: it's important to handle errors here
        // instead of a catch() block so that we don't swallow
        // exceptions from actual bugs in components.
        (error) => {
          setIsLoaded(true);
          setError(error);
        }
      );
  }, []);

  if (error !== null) {
    return <div>Error: {error.message}</div>;
  }

  return <DayTemplate isLoaded={isLoaded} error={error} result={result} />;
}

export default Day2;
