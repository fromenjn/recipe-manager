// src/components/RecipesTab.tsx
import React, { useEffect, useState } from "react";
import { fetchAllRecipes } from "../services/api";
import { Recipe } from "../types/domain";

const RecipesTab: React.FC = () => {
  const [recipes, setRecipes] = useState<Recipe[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const loadRecipes = async () => {
      try {
        setLoading(true);
        const data = await fetchAllRecipes();
        setRecipes(data);
      } catch (err: any) {
        setError(err.message || "Error fetching recipes");
      } finally {
        setLoading(false);
      }
    };

    loadRecipes();
  }, []);

  if (loading) return <div>Loading recipes...</div>;
  if (error) return <div style={{ color: "red" }}>{error}</div>;

  return (
    <div style={{ padding: "1rem" }}>
      <h2>All Recipes</h2>
      {recipes.map((recipe) => (
        <div key={recipe.id} style={{ marginBottom: "2rem" }}>
          <h3>{recipe.name}</h3>

          <p><strong>Ingredients:</strong></p>
          <ul>
            {recipe.ingredients.map((ing) => (
              <li key={ing.name}>
                {ing.name} - {ing.quantity} {ing.unit}
              </li>
            ))}
          </ul>

          <p><strong>Steps:</strong></p>
          <ol>
            {recipe.steps.map((step) => (
              <li key={step.id} style={{ marginBottom: "1rem" }}>
                <strong>{step.name}</strong>: {step.instructions}
                <div style={{ marginTop: "0.5rem" }}>
                  {step.illustration.map((ill) => (
                    <img
                      key={ill.id}
                      src={ill.filepath}
                      alt={ill.description}
                      style={{
                        maxWidth: "200px",
                        display: "block",
                        marginTop: "0.5rem",
                      }}
                    />
                  ))}
                </div>
              </li>
            ))}
          </ol>
        </div>
      ))}
    </div>
  );
};

export default RecipesTab;
