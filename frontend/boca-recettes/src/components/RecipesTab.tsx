// src/components/RecipesTab.tsx
import React, { useEffect, useState } from "react";
import { fetchAllRecipes } from "../services/api";
import { Recipe } from "../types/domain";
import {
  Typography,
  Card,
  CardContent,
  List,
  ListItem,
  ListItemText,
  Divider,
} from "@mui/material";

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

  if (loading) {
    return <Typography>Loading recipes...</Typography>;
  }

  if (error) {
    return <Typography color="error">{error}</Typography>;
  }

  return (
    <div style={{ padding: "1rem" }}>
      <Typography variant="h4" gutterBottom>
        Toutes nos recettes
      </Typography>
      {recipes.map((recipe) => (
        <Card key={recipe.id} style={{ marginBottom: "2rem" }}>
          <CardContent>
            <Typography variant="h5">{recipe.name}</Typography>
            <Divider style={{ margin: "1rem 0" }} />
            <Typography variant="subtitle1" gutterBottom>
              Ingredients:
            </Typography>
            <List>
              {recipe.ingredients.map((ing) => (
                <ListItem key={ing.name} disablePadding>
                  <ListItemText
                    primary={`${ing.name}: ${ing.quantity} ${ing.unit}`}
                  />
                </ListItem>
              ))}
            </List>
            <Divider style={{ margin: "1rem 0" }} />
            <Typography variant="subtitle1" gutterBottom>
              Steps:
            </Typography>
            <ol>
              {recipe.steps.map((step) => (
                <li key={step.id} style={{ marginBottom: "1rem" }}>
                  <Typography>
                    <strong>{step.name}</strong>: {step.instructions}
                  </Typography>
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
                </li>
              ))}
            </ol>
          </CardContent>
        </Card>
      ))}
    </div>
  );
};

export default RecipesTab;
