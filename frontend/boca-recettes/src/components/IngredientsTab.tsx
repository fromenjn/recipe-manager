// src/components/IngredientsTab.tsx
import React, { useEffect, useState } from "react";
import { fetchAllIngredients, fetchAllRecipes, fetchRecipeById } from "../services/api";
import { Recipe } from "../types/domain";

const IngredientsTab: React.FC = () => {
  const [ingredientList, setIngredientList] = useState<string[]>([]);
  const [recipes, setRecipes] = useState<Recipe[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  // User selections
  const [selectedRecipeId, setSelectedRecipeId] = useState("");
  const [selectedIngredient, setSelectedIngredient] = useState("");
  const [quantity, setQuantity] = useState<number>(0);

  // Scaled recipe result
  const [scaledRecipe, setScaledRecipe] = useState<Recipe | null>(null);

  // On mount: load all ingredients and all recipes
  useEffect(() => {
    const loadData = async () => {
      try {
        setLoading(true);
        // Parallel fetch
        const [ingredientsData, recipesData] = await Promise.all([
          fetchAllIngredients(), // string[]
          fetchAllRecipes(),     // Recipe[]
        ]);
        setIngredientList(ingredientsData);
        setRecipes(recipesData);
      } catch (err: any) {
        setError(err.message || "Error fetching data");
      } finally {
        setLoading(false);
      }
    };

    loadData();
  }, []);

  // Handle scaling a recipe
  const handleScaleRecipe = async () => {
    if (!selectedRecipeId || !selectedIngredient || quantity <= 0) {
      return;
    }
    try {
      setError(null);
      setScaledRecipe(null);
      const scaled = await fetchRecipeById(selectedRecipeId, selectedIngredient, quantity);
      setScaledRecipe(scaled);
    } catch (err: any) {
      setError(err.message || "Error scaling recipe");
    }
  };

  if (loading) return <div>Loading...</div>;
  if (error) return <div style={{ color: "red" }}>{error}</div>;

  return (
    <div style={{ padding: "1rem" }}>
      <h2>Ingredients</h2>
      <p>Here is the list of all available ingredients:</p>
      <ul>
        {ingredientList.map((ing) => (
          <li key={ing}>{ing}</li>
        ))}
      </ul>

      <hr />

      <h3>Scale a Recipe by an Ingredient</h3>
      <div style={{ marginBottom: "1rem" }}>
        <label style={{ marginRight: "0.5rem" }}>Recipe:</label>
        <select
          value={selectedRecipeId}
          onChange={(e) => setSelectedRecipeId(e.target.value)}
        >
          <option value="">-- Select a Recipe --</option>
          {recipes.map((r) => (
            <option key={r.id} value={r.id}>
              {r.name}
            </option>
          ))}
        </select>
      </div>

      <div style={{ marginBottom: "1rem" }}>
        <label style={{ marginRight: "0.5rem" }}>Ingredient to Scale:</label>
        <select
          value={selectedIngredient}
          onChange={(e) => setSelectedIngredient(e.target.value)}
        >
          <option value="">-- Select an Ingredient --</option>
          {ingredientList.map((ing) => (
            <option key={ing} value={ing}>
              {ing}
            </option>
          ))}
        </select>
      </div>

      <div style={{ marginBottom: "1rem" }}>
        <label style={{ marginRight: "0.5rem" }}>Desired Quantity:</label>
        <input
          type="number"
          value={quantity}
          onChange={(e) => setQuantity(Number(e.target.value))}
        />
      </div>

      <button onClick={handleScaleRecipe}>Scale Recipe</button>

      {scaledRecipe && (
        <div style={{ marginTop: "1.5rem" }}>
          <h4>Scaled Recipe: {scaledRecipe.name}</h4>
          <ul>
            {scaledRecipe.ingredients.map((ing) => (
              <li key={ing.name}>
                {ing.name}: {ing.quantity} {ing.unit}
              </li>
            ))}
          </ul>
          <ol>
            {scaledRecipe.steps.map((step) => (
              <li key={step.id} style={{ marginBottom: "1rem" }}>
                <strong>{step.name}</strong>: {step.instructions}
                {step.illustration.map((ill) => (
                  <div key={ill.id} style={{ marginTop: "0.5rem" }}>
                    <img
                      src={ill.filepath}
                      alt={ill.description}
                      style={{ maxWidth: "200px" }}
                    />
                  </div>
                ))}
              </li>
            ))}
          </ol>
        </div>
      )}
    </div>
  );
};

export default IngredientsTab;
