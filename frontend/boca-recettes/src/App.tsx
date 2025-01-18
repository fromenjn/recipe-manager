// src/App.tsx
import { useState } from "react";
import IngredientsTab from "./components/IngredientsTab";
import RecipesTab from "./components/RecipesTab";

function App() {
  const [activeTab, setActiveTab] = useState<"ingredients" | "recipes">("ingredients");

  return (
    <div style={{ padding: "1rem" }}>
      <div style={{ marginBottom: "1rem" }}>
        <button
          onClick={() => setActiveTab("ingredients")}
          style={activeTab === "ingredients" ? { fontWeight: "bold" } : {}}
        >
          Ingredients
        </button>
        <button
          onClick={() => setActiveTab("recipes")}
          style={{ marginLeft: "1rem", ...(activeTab === "recipes" ? { fontWeight: "bold" } : {}) }
          }
        >
          Recipes
        </button>
      </div>

      {activeTab === "ingredients" && <IngredientsTab />}
      {activeTab === "recipes" && <RecipesTab />}
    </div>
  );
}

export default App;
