import React, { useState } from "react";
import PropTypes from "prop-types";

import Box from "@mui/material/Box";
import Drawer from "@mui/material/Drawer";
import TextField from "@mui/material/TextField";
import LoadingButton from "@mui/lab/LoadingButton";
import SendIcon from "@mui/icons-material/Send";
import ExitToAppIcon from "@mui/icons-material/ExitToApp";
import Typography from "@mui/material/Typography";
import { useStore } from "../store";
import { toast } from "react-toastify";

export const LoginDrawer = ({ open, toggleDrawer }) => {
  const [loading, setLoading] = useState(false);
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const token = useStore((state) => state.token);
  const setToken = useStore((state) => state.setToken);
  const setUserId = useStore((state) => state.setUserId);

  const sendLogin = async (data) => {
    try {
      const res = await fetch(
        `${import.meta.env.VITE_CORE_BASE}/v1/user/login`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(data),
        }
      );
      const response = await res.json();
      if (response.token) {
        setToken(response.token);
        setUserId(response.user.userId);
        setLoading(false);
        toast.success("Login successful");
        toggleDrawer(false)();
      } else {
        setLoading(false);
        toast.error("Invalid credentials");
      }
    } catch (e) {
      setLoading(false);
      toast.error("Invalid credentials");
    }
  };

  const handleLogin = (event) => {
    event.preventDefault();
    const data = { username, password };
    setLoading(true);
    sendLogin(data);
  };

  const DrawerList = (
    <Box sx={{ width: 350, p: 3 }}>
      {!token && (
        <>
          <Typography variant="h6" component="h2">
            Login
          </Typography>
          <form onSubmit={handleLogin}>
            <TextField
              label="Username"
              variant="outlined"
              fullWidth
              margin="normal"
              onChange={(e) => setUsername(e.target.value)}
              required
            />
            <TextField
              label="Password"
              type="password"
              variant="outlined"
              fullWidth
              margin="normal"
              onChange={(e) => setPassword(e.target.value)}
              required
            />
            <LoadingButton
              onClick={handleLogin}
              endIcon={<SendIcon />}
              loading={loading}
              loadingPosition="end"
              variant="contained"
            >
              <span>Send</span>
            </LoadingButton>
          </form>
        </>
      )}
      {token && (
        <LoadingButton
          onClick={() => setToken("")}
          endIcon={<ExitToAppIcon />}
          loading={loading}
          loadingPosition="end"
          variant="contained"
        >
          <span>Logout</span>
        </LoadingButton>
      )}
    </Box>
  );

  return (
    <div>
      <Drawer open={open} onClose={toggleDrawer(false)} anchor={"right"}>
        {DrawerList}
      </Drawer>
    </div>
  );
};

LoginDrawer.propTypes = {
  open: PropTypes.bool,
  toggleDrawer: PropTypes.func,
};
