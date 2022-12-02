import { useEffect, useState } from "react";
import { CircularProgress, Paper } from "@mui/material";

type errorType = {
  message: string;
} | null;

type resultType = {
  hey: string;
} | null;

function Day1(): JSX.Element {
  const [result, setResult] = useState<resultType>(null);
  const [isLoaded, setIsLoaded] = useState(false);
  const [error, setError] = useState<errorType>(null);

  useEffect(() => {
    fetch("/api/day1")
      .then(async (res) => await res.json())
      .then(
        (result) => {
          setIsLoaded(true);
          setResult(result);
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

  return (
    <Paper
      sx={{
        py: 3,
        px: 3,
        border: 0,
        borderRadius: 0,
        display: "flex",
        minHeight: "100vh",
        flexDirection: "column",
      }}
    >
      {!isLoaded || result === null ?  <CircularProgress sx={{ height: 1, margin: "auto" }} /> : result.hey}
    </Paper>
  );
}

export default Day1;
