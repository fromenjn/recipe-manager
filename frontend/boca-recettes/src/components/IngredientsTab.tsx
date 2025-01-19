import React, { useEffect, useState } from "react";
import {
  Typography,
  List,
  ListItem,
  Select,
  MenuItem,
  TextField,
  Button,
  Box,
  CircularProgress,
  Alert,
  Card,
  CardContent,
} from "@mui/material";
import { fetchAllIngredients, fetchAllRecipes, fetchRecipeById } from "../services/api";
import { Recipe } from "../types/domain";

const IngredientsTab: React.FC = () => {
  const [allIngredients, setAllIngredients] = useState<string[]>([]);
  const [allRecipes, setAllRecipes] = useState<Recipe[]>([]);
  const [ingredientList, setIngredientList] = useState<string[]>([]);
  const [recipes, setRecipes] = useState<Recipe[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const [selectedRecipeId, setSelectedRecipeId] = useState("");
  const [selectedIngredient, setSelectedIngredient] = useState("");
  const [quantity, setQuantity] = useState<number>(0);
  const [scaledRecipe, setScaledRecipe] = useState<Recipe | null>(null);

  // Fetch all data on mount
  useEffect(() => {
    const loadData = async () => {
      try {
        setLoading(true);
        const [ingredientsData, recipesData] = await Promise.all([
          fetchAllIngredients(),
          fetchAllRecipes(),
        ]);
        setAllIngredients(ingredientsData);
        setAllRecipes(recipesData);
        setIngredientList(ingredientsData);
        setRecipes(recipesData);
      } catch (err: any) {
        setError(err.message || "Echec du chargement des données.");
      } finally {
        setLoading(false);
      }
    };
    loadData();
  }, []);

  // Refresh ingredients when a recipe is selected
  useEffect(() => {
    if (selectedRecipeId) {
      const selectedRecipe = allRecipes.find((recipe) => recipe.id === selectedRecipeId);
      if (selectedRecipe) {
        setIngredientList(selectedRecipe.ingredients.map((ing) => ing.name));
      }
    } else {
      setIngredientList(allIngredients);
    }
  }, [selectedRecipeId, allRecipes, allIngredients]);

  // Refresh recipes when an ingredient is selected
  useEffect(() => {
    if (selectedIngredient) {
      const filteredRecipes = allRecipes.filter((recipe) =>
        recipe.ingredients.some((ing) => ing.name === selectedIngredient)
      );
      setRecipes(filteredRecipes);
    } else {
      setRecipes(allRecipes);
    }
  }, [selectedIngredient, allRecipes]);

  // Scale a recipe based on the selected ingredient and quantity
  const handleScaleRecipe = async () => {
    if (!selectedRecipeId || !selectedIngredient || quantity <= 0) {
      setError("Tous les champs ne sont pas remplis correctement !");
      return;
    }
    try {
      setError(null);
      setScaledRecipe(null);
      const scaled = await fetchRecipeById(selectedRecipeId, selectedIngredient, quantity);
      setScaledRecipe(scaled);
    } catch (err: any) {
      setError(err.message || "Failed to scale the recipe.");
    }
  };

  if (loading) {
    return (
      <Box display="flex" justifyContent="center" alignItems="center" height="50vh">
        <CircularProgress />
      </Box>
    );
  }

  return (
    <Box padding="1rem">
      {/*<Typography variant="h5" gutterBottom>
      </Typography>*/}
      <Box marginBottom={2}>
        <Select
          value={selectedRecipeId}
          onChange={(e) => setSelectedRecipeId(e.target.value)}
          displayEmpty
          fullWidth
          sx={{
            marginBottom: 2, 
            backgroundColor: "white", // White background for the dropdown
            color: "black",           // Black text color
            "& .MuiSelect-icon": { color: "black" }, // Black dropdown icon
          }}
        >
          <MenuItem value="">-- Choisir une recette --</MenuItem>
          {recipes.map((recipe) => (
            <MenuItem key={recipe.id} value={recipe.id}>
              {recipe.name}
            </MenuItem>
          ))}
        </Select>
        <Select
          value={selectedIngredient}
          onChange={(e) => setSelectedIngredient(e.target.value)}
          displayEmpty
          fullWidth
          sx={{
            marginBottom: 2, 
            backgroundColor: "white", // White background for the dropdown
            color: "black",           // Black text color
            "& .MuiSelect-icon": { color: "black" }, // Black dropdown icon
          }}
        >
          <MenuItem value="">-- Choisir un ingrédient --</MenuItem>
          {ingredientList.map((ing) => (
            <MenuItem key={ing} value={ing}>
              {ing}
            </MenuItem>
          ))}
        </Select>
        <TextField
            type="number"
            label="Quantity"
            value={quantity}
            onChange={(e) => setQuantity(Number(e.target.value))}
            fullWidth
            sx={{
                backgroundColor: "white", // Ensures the input box has a white background
                borderRadius: "4px",      // Rounded corners for consistency
                "& .MuiInputLabel-root": {
                backgroundColor: "white", // Prevent label from overlapping with the background
                padding: "0 4px",         // Add padding around the label
                },
                "& .MuiOutlinedInput-root": {
                "& fieldset": {
                    borderColor: "#ccc",    // Light border for consistency
                },
                "&:hover fieldset": {
                    borderColor: "#1976d2", // Highlight on hover
                },
                },
            }}
        />
        <Button variant="contained" color="primary" onClick={handleScaleRecipe} fullWidth>
          Ajuster la recette
        </Button>
      </Box>
      {error && (
        <Alert severity="error" sx={{ marginBottom: 2 }}>
          {error}
        </Alert>
      )}
      {scaledRecipe && (
        <Card>
          <CardContent>
            <Typography variant="h6" gutterBottom>
              Scaled Recipe: {scaledRecipe.name}
            </Typography>
            <Typography variant="subtitle1">Ingredients:</Typography>
            <List>
              {scaledRecipe.ingredients.map((ing) => (
                <ListItem key={ing.name}>
                  {ing.name}: {ing.quantity} {ing.unit}
                </ListItem>
              ))}
            </List>
            <Typography variant="subtitle1">Steps:</Typography>
            <ol>
              {scaledRecipe.steps.map((step) => (
                <li key={step.id}>
                  {step.name}: {step.instructions}
                </li>
              ))}
            </ol>
          </CardContent>
        </Card>
      )}
    </Box>
  );
};

export default IngredientsTab;
