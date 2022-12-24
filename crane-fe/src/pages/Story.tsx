import * as React from "react";
import { IconButton, Toolbar, Typography } from "@mui/material";
import SearchIcon from "@mui/icons-material/Search";
import SaveIcon from "@mui/icons-material/Save";
import HomeIcon from "@mui/icons-material/Home";
import StoryBook from "../components/StoryBook";

function Story() {
  return (
    <React.Fragment>
      <Toolbar sx={{ borderBottom: 1, borderColor: "divider" }}>
        <IconButton>
          <HomeIcon />
        </IconButton>
        <IconButton>
          <SaveIcon />
        </IconButton>
        <Typography
          color="inherit"
          align="center"
          noWrap
          sx={{ flex: 1 }}
        ></Typography>
        <IconButton>
          <SearchIcon />
        </IconButton>
      </Toolbar>
      <StoryBook placeholder="记录这一刻" focused fullWidth multiline />
    </React.Fragment>
  );
}

export default Story;
