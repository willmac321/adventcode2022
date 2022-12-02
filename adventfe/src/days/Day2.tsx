import { useEffect, useState } from "react";
import { CircularProgress, Paper, Container, Typography } from "@mui/material";

type errorType = {
  message: string;
} | null;

type resultType = {
  input: string[][];
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

  console.log(result);
  return (
    <Paper
      sx={{
        py: 3,
        px: 3,
        alignItems: "center",
        border: 0,
        borderRadius: 0,
        display: "flex",
        minHeight: "100vh",
        flexDirection: "column",
      }}
    >
      <Typography variant="h3" sx={{ color: "primary.main" }}>
        🎄🎄🎄🎄 Advent of Code 2022 Day 2 🎄🎄🎄🎄
      </Typography>
      {!isLoaded || result === null ? (
        <CircularProgress sx={{ height: 1, margin: "auto" }} />
      ) : (
        <Container>
          <Typography variant="h4" sx={{ p: 1, color: "primary.main" }}>
            Part 1 Answer
          </Typography>
          <Typography sx={{ m: 1, color: "primary.main" }}>
          </Typography>
          <Typography variant="h4" sx={{ p: 1, color: "primary.main" }}>
            Part 2 Answer
          </Typography>
          <Typography sx={{ m: 1, color: "primary.main" }}>
          </Typography>
          <Typography variant="h4" sx={{ p: 1, color: "secondary.main" }}>
            input
          </Typography>
          <>
            {result.input.map((input) => (
              <Typography
                sx={{ ml: 1, color: "secondary.main" }}
                key={JSON.stringify(input)}
              >
                {JSON.stringify(input)}
              </Typography>
            ))}
          </>
        </Container>
      )}
    </Paper>
  );
}

export default Day2;
