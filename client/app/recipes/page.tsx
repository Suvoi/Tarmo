"use client"

import { useEffect, useState } from "react"
import { AppSidebar } from "@/components/app-sidebar"
import { Breadcrumb, BreadcrumbList, BreadcrumbItem, BreadcrumbLink, BreadcrumbSeparator, BreadcrumbPage } from "@/components/ui/breadcrumb"
import { Card, CardContent } from "@/components/ui/card"
import { Plus } from "lucide-react"
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar"
import Link from "next/link"

type Recipe = {
  id: string | number
  name: string
  instructions?: string
}

export default function Page() {
  const [recipes, setRecipes] = useState<Recipe[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    async function loadRecipes() {
      try {
        const base = process.env.NEXT_PUBLIC_APP_URL || ""
        const res = await fetch(`${base}/api/recipes`, { cache: "no-store" })

        if (!res.ok) throw new Error(`HTTP ${res.status}`)

        const data = await res.json()
        setRecipes(Array.isArray(data) ? data : data.recipes || [])
      } catch (err) {
        console.warn("API not available, using mockup recipes:", err)
        setRecipes([
          { id: "1", name: "Recipe 1", instructions: "Various instructions..." },
          { id: "2", name: "Recipe 2", instructions: "Various instructions..." },
          { id: "3", name: "Recipe 3", instructions: "Various instructions..." },
          { id: "4", name: "Recipe 4", instructions: "Various instructions..." },
          { id: "5", name: "Recipe 5", instructions: "Various instructions..." },
          { id: "6", name: "Recipe 6", instructions: "Various instructions..." },
        ])
      } finally {
        setLoading(false)
      }
    }

    loadRecipes()
  }, [])

  return (
    <SidebarProvider style={{ "--sidebar-width": "19rem" } as React.CSSProperties}>
      <AppSidebar />
      <SidebarInset>
        <header className="flex h-16 shrink-0 items-center gap-2 border-b px-4">
          <Breadcrumb>
            <BreadcrumbList>
              <BreadcrumbItem className="hidden md:block">
                <BreadcrumbLink href="/">Home</BreadcrumbLink>
              </BreadcrumbItem>
              <BreadcrumbSeparator className="hidden md:block" />
              <BreadcrumbItem>
                <BreadcrumbPage>Recipes</BreadcrumbPage>
              </BreadcrumbItem>
            </BreadcrumbList>
          </Breadcrumb>
        </header>

        <div className="h-full grid grid-cols-5 p-4 gap-4">
          {/* Card para crear una nueva receta */}
          <Link href="/recipes/new" className="w-48 h-48">
            <Card className="w-48 h-48 flex items-center justify-center cursor-pointer hover:shadow-lg transition">
              <CardContent className="h-full w-full flex items-center justify-center">
                <Plus size={36} />
              </CardContent>
            </Card>
          </Link>

          {/* Cards de recetas */}
          {loading ? (
            <div className="col-span-5 text-center text-gray-500">Loading...</div>
          ) : (
            recipes.map((recipe) => (
              <Link key={recipe.id} href={`/recipes/${recipe.id}`}>
                <Card className="w-48 h-48 flex flex-col justify-between cursor-pointer hover:shadow-lg transition">
                  <CardContent className="p-4">
                    <h3 className="text-lg font-semibold">{recipe.name}</h3>
                    <p className="text-sm text-muted-foreground line-clamp-3">
                      {recipe.instructions || "No instructions"}
                    </p>
                  </CardContent>
                </Card>
              </Link>
            ))
          )}
        </div>
      </SidebarInset>
    </SidebarProvider>
  )
}
