// src/components/Navbar.tsx

import React from "react";
import { Link } from "react-router-dom";

const Navbar: React.FC = () => {
  return (
    <nav style={{ padding: "1rem", borderBottom: "1px solid #ccc" }}>
      <Link to="/ingredients" style={{ marginRight: "1rem" }}>
        Ingredients
      </Link>
      <Link to="/recipes">Recipes</Link>
    </nav>
  );
};

export default Navbar;
