// src/App.tsx
import { useState } from "react";
import Navbar from "./components/Navbar";
import IngredientsTab from "./components/IngredientsTab";
import RecipesTab from "./components/RecipesTab";
import { Container } from "@mui/material";

function App() {
  const [activePage, setActivePage] = useState<"ingredients" | "recipes">("ingredients");

  return (
    <div style={{ height: "100vh", display: "flex", flexDirection: "column", width: "100vw" }}>
      <Navbar onNavigate={setActivePage} />
      <Container style={{ flex: 1, overflow: "auto", padding: "1rem" }}>
        {activePage === "ingredients" && <IngredientsTab />}
        {activePage === "recipes" && <RecipesTab />}
      </Container>
    </div>
  );
}

export default App;
