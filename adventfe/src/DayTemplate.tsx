import { CircularProgress, Paper, Container, Typography } from "@mui/material";

function DayTemplate({
  isLoaded,
  error,
  result,
}: {
  isLoaded: boolean;
  error: { message: string } | null;
  result: { input: string; output1: string; output2: string } | null;
}): JSX.Element {
  if (error !== null) {
    return <div>Error: {error.message}</div>;
  }

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
        🎄🎄🎄🎄 Advent of Code 2022 Day 1🎄🎄🎄🎄
      </Typography>
      {!isLoaded || result === null ? (
        <CircularProgress sx={{ height: 1, margin: "auto" }} />
      ) : (
        <Container>
          <Typography variant="h4" sx={{ p: 1, color: "primary.main" }}>
            Part 1 Answer
          </Typography>
          <Typography sx={{ m: 1, color: "primary.main" }}>
            {result.output1}
          </Typography>
          <Typography variant="h4" sx={{ p: 1, color: "primary.main" }}>
            Part 2 Answer
          </Typography>
          <Typography sx={{ m: 1, color: "primary.main" }}>
            {result.output2}
          </Typography>
          <Typography variant="h4" sx={{ p: 1, color: "secondary.main" }}>
            input
          </Typography>
          <Typography variant="body1" component="p" whiteSpace="pre-wrap"  sx={{ ml: 1, color: "secondary.main" }}>
            {result.input}
          </Typography>
        </Container>
      )}
    </Paper>
  );
}

export default DayTemplate;
