// src/components/Navbar.tsx
import React from "react";
import { AppBar, Toolbar, Typography, Button } from "@mui/material";

interface NavbarProps {
  onNavigate: (page: "ingredients" | "recipes") => void;
}

export const Navbar: React.FC<NavbarProps> = ({ onNavigate }) => {
  return (
    <AppBar position="static" style={{ height: "10vh"}}>
      <Toolbar style={{ display: "flex", justifyContent: "space-between" }}>
        <Typography variant="h6">Boca'recettes</Typography>
        <div>
          <Button onClick={() => onNavigate("ingredients")} color="inherit">
            Ingredients
          </Button>
          <Button onClick={() => onNavigate("recipes")} color="inherit">
            Recettes
          </Button>
        </div>
      </Toolbar>
    </AppBar>
  );
};

export default Navbar;
