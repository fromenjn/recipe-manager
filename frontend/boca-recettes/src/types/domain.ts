// src/types/domain.ts

export interface Recipe {
    id: string;
    name: string;
    ingredients: {
      name: string;
      quantity: number;
      unit: string;
    }[];
    steps: {
      id: string;
      name: string;
      instructions: string;
      illustration: {
        id: string;
        description: string;
        filepath: string;
      }[];
    }[];
  }
  