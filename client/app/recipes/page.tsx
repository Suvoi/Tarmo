import { AppSidebar } from "@/components/app-sidebar"
import { Breadcrumb, BreadcrumbList, BreadcrumbItem, BreadcrumbLink, BreadcrumbSeparator, BreadcrumbPage } from "@/components/ui/breadcrumb"
import { Card, CardContent, } from "@/components/ui/card"
import { Plus } from "lucide-react"
import {
  SidebarInset,
  SidebarProvider,
} from "@/components/ui/sidebar"
import Link from "next/link"

type Recipe = {
  id: string | number
  name: string
  description?: string
}

async function getRecipes() {
  try {
    const res = await fetch("http://tarmo:9136/recipes", { cache: "no-store" });

    if (!res.ok) {
      console.warn("API responded with error:", res.status);
      return [];
    }

    const data = await res.json();
    return Array.isArray(data) ? data : data.recipes || [];
  } catch (err) {
    console.warn("API not available, returning empty recipe list");
    return [];
  }
}



export default async function Page() {
  const recipes = await getRecipes()

  return (
    <SidebarProvider
      style={
        {
          "--sidebar-width": "19rem",
        } as React.CSSProperties
      }
    >
      <AppSidebar />
      <SidebarInset>
       <header className="flex h-16 shrink-0 items-center gap-2 border-b px-4">
          <Breadcrumb>
            <BreadcrumbList>
              <BreadcrumbItem className="hidden md:block">
                <BreadcrumbLink href="/">
                  Home
                </BreadcrumbLink>
              </BreadcrumbItem>
              <BreadcrumbSeparator className="hidden md:block" />
              <BreadcrumbItem>
                <BreadcrumbPage>Recipes</BreadcrumbPage>
              </BreadcrumbItem>
            </BreadcrumbList>
          </Breadcrumb>
        </header>

        <div className="h-full grid grid-cols-5 p-4">
            <Link href="/recipes/new" className="w-48 h-48">
                <Card className="w-48 h-48 flex items-center justify-center cursor-pointer">
                    <CardContent className="h-full w-full flex items-center justify-center">
                        <Plus />
                    </CardContent>
                </Card>
            </Link>

            {recipes.map((recipe: Recipe) => (
            <Link key={recipe.id} href={`/recipes/${recipe.id}`}>
              <Card className="w-48 h-48 flex flex-col justify-between cursor-pointer hover:shadow-lg transition">
                <CardContent className="p-4">
                  <h3 className="text-lg font-semibold">{recipe.name}</h3>
                  <p className="text-sm text-muted-foreground line-clamp-3">
                    {recipe.description || "No description"}
                  </p>
                </CardContent>
              </Card>
            </Link>
          ))}
        </div>

      </SidebarInset>
    </SidebarProvider>
  )
}
