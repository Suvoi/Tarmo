"use client"

import { useEffect, useState } from "react"

type Recipe = {
  id: string | number
  name: string
  instructions?: string
}

export default function Page({ params }: { params: { id: string } }) {
  const [recipe, setRecipe] = useState<Recipe | null>(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    async function loadRecipe() {
      try {
        const res = await fetch(`/api/recipes/${params.id}`, { cache: "no-store" })
        if (!res.ok) {
          setRecipe(null)
          return
        }
        const data = await res.json()
        setRecipe(data)
      } catch (err) {
        console.error(err)
        setRecipe(null)
      } finally {
        setLoading(false)
      }
    }

    loadRecipe()
  }, [params.id])

  if (loading) return <div className="p-4">Loading...</div>
  if (!recipe) return <div className="p-4">Recipe not found</div>

  return (
    <div className="p-6">
      <h1 className="text-3xl font-bold">{recipe.name}</h1>
      <p className="mt-4 whitespace-pre-wrap">
        {recipe.instructions || "No instructions available"}
      </p>
    </div>
  )
}
