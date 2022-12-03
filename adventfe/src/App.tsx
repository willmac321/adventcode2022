import { Paper, Stack, Typography } from "@mui/material";
import { Link as RouterLink } from "react-router-dom";
import routes from "./routes";

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
        {routes.map((route, i) => {
          if (i === 0 || route.path === undefined) return <></>;
          return (
            <RouterLink key={route.path} to={route.path}>
              <Typography
                variant="h4"
                component="span"
                sx={{ color: "secondary.main" }}
              >
                {`${route.path.charAt(0).toUpperCase()}${route.path.slice(
                  1,
                  3
                )} ${route.path.slice(3)}`}
              </Typography>
            </RouterLink>
          );
        })}
      </Stack>
    </Paper>
  );
}

export default App;
