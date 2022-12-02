import { Paper, Stack, Typography } from "@mui/material";
import { Link as RouterLink } from "react-router-dom";

function App(): JSX.Element {
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
      <Stack
        sx={{
          verticalAlign: "center",
          alignItems: "center",
          justifyContent: "center",
        }}
      >
        <Typography variant="h2" sx={{ color: "primary.main" }}>
          🎄🎄🎄🎄 Advent of Code 2022 🎄🎄🎄🎄
        </Typography>
        <RouterLink to="/day1">
          <Typography
            variant="h4"
            component="span"
            sx={{ color: "secondary.main" }}
          >
            day 1
          </Typography>
        </RouterLink>
      </Stack>
    </Paper>
  );
}

export default App;
