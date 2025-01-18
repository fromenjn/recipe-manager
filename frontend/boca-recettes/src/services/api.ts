// src/services/api.ts
import { Recipe } from "../types/domain";

const BASE_URL = import.meta.env.VITE_BASE_URL || "";
// Adjust if your backend runs on a different host or port

/**
 * Fetch all ingredients (array of strings) from /ingredients
 */
export async function fetchAllIngredients(): Promise<string[]> {
  const response = await fetch(`${BASE_URL}/ingredients`);
  if (!response.ok) {
    throw new Error(`Failed to fetch ingredients: ${response.statusText}`);
  }
  return response.json() as Promise<string[]>;
}

/**
 * Fetch all recipes from /recipes
 */
export async function fetchAllRecipes(): Promise<Recipe[]> {
  const response = await fetch(`${BASE_URL}/recipes`);
  if (!response.ok) {
    throw new Error(`Failed to fetch recipes: ${response.statusText}`);
  }
  return response.json() as Promise<Recipe[]>;
}

/**
 * Fetch a specific recipe by ID, optionally scaled by providing
 * ingredient and quantity
 */
export async function fetchRecipeById(
  recipeId: string,
  ingredient?: string,
  quantity?: number
): Promise<Recipe> {
  const params = new URLSearchParams();
  if (ingredient) params.append("ingredient", ingredient);
  if (quantity !== undefined) params.append("quantity", String(quantity));

  const response = await fetch(`${BASE_URL}/recipe/${recipeId}?${params}`);
  if (!response.ok) {
    throw new Error(`Failed to fetch recipe ${recipeId}: ${response.statusText}`);
  }
  return response.json() as Promise<Recipe>;
}
