"use client"

import { useEffect, useState } from "react"
import { AppSidebar } from "@/components/app-sidebar"
import { Breadcrumb, BreadcrumbList, BreadcrumbItem, BreadcrumbLink, BreadcrumbSeparator, BreadcrumbPage } from "@/components/ui/breadcrumb"
import { Card, CardContent } from "@/components/ui/card"
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar"
import Link from "next/link"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { ChevronRight, Inbox, LucideReceiptRussianRuble, Plus } from "lucide-react"
import { Spinner } from "@/components/ui/spinner"
import { Empty, EmptyContent, EmptyDescription, EmptyHeader, EmptyMedia, EmptyTitle } from "@/components/ui/empty"

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

      if (process.env.NODE_ENV === "development") {
        setRecipes(Array.from({ length: 12 }, (_, i) => ({
          id: i + 1,
          name: `Recipe ${i + 1}`,
        })))
        setLoading(false)
        return
      }

      try {
        const res = await fetch(`/api/recipes`, { cache: "no-store" })

        if (!res.ok) throw new Error(`HTTP ${res.status}`)

        const data = await res.json()
        setRecipes(Array.isArray(data) ? data : data.recipes || [])
      } catch (err) {
        console.warn("API not available, ", err)
        setRecipes([])
      } finally {
        setLoading(false)
      }
    }

    loadRecipes()
  }, [])

  return (
    <SidebarProvider style={{ "--sidebar-width": "14rem" } as React.CSSProperties}>
      <AppSidebar />
      <SidebarInset className="flex flex-col h-screen">
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

        <div className="h-16 flex p-4 items-center space-x-2">
          <Link href="/recipes/new">
            <Button className="cursor-pointer">
              <Plus/>
            </Button>
          </Link>
          <Input className="w-1/3"placeholder="Search..."/>
        </div>

        <div className="flex-1 overflow-y-auto p-4">
          {loading ? (
            <div className="h-full flex items-center justify-center">
              <Spinner/>
            </div>
          ) : recipes?.length ? (
            <div className="grid grid-cols-3 gap-2">
              {recipes.map((recipe) => (
                <Link key={recipe.id} href={`/recipes/${recipe.id}`}>
                  <Card className="flex flex-col justify-between cursor-pointer hover:shadow-lg transition">
                    <CardContent>
                      <h3 className="text-lg font-semibold">{recipe.name}</h3>
                      <p className="text-sm text-muted-foreground line-clamp-3">
                        {recipe.instructions || "No instructions"}
                      </p>
                    </CardContent>
                  </Card>
                </Link>
              ))}
            </div>
          ) : (
            <div className="h-full flex items-center justify-center">
              <Empty>
                <EmptyHeader>
                  <EmptyMedia variant="icon">
                    <Inbox />
                  </EmptyMedia>
                  <EmptyTitle>
                    <h1>Looks a bit empty…</h1>
                  </EmptyTitle>
                  <EmptyDescription>
                    <p>Add a recipe to start building your collection.</p>
                  </EmptyDescription>
                </EmptyHeader>
                <EmptyContent>
                  <Link href="/recipes/new">
                    <Button className="cursor-pointer items-center">
                      Add recipe
                      <ChevronRight />
                    </Button>
                  </Link>
                </EmptyContent>
              </Empty>
            </div>
          )}
        </div>
      </SidebarInset>
    </SidebarProvider>
  )
}
