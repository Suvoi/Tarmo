"use client"

import { AppSidebar } from "@/components/app-sidebar"
import { Breadcrumb, BreadcrumbList, BreadcrumbItem, BreadcrumbLink, BreadcrumbSeparator, BreadcrumbPage } from "@/components/ui/breadcrumb"
import { Button } from "@/components/ui/button"
import { toast } from "sonner"
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar"
import Link from "next/link"
import { useEffect, useState } from "react"
import { useRouter } from "next/navigation"

import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"

import { ChevronLeft, Pencil, Trash2 } from "lucide-react"
import { Spinner } from "@/components/ui/spinner"

type Recipe = {
  id: string | number
  name: string
  instructions?: string
}

export default function Page({ params }: { params: { id: string } }) {
  const [recipe, setRecipe] = useState<Recipe | null>(null)
  const [loading, setLoading] = useState(true)
  const router = useRouter()

  async function deleteRecipe() {
    try {
      await fetch(`/api/recipes/${params.id}`, { method: "DELETE" })
    } catch (err) {
      console.error(err)
    }
  }

  useEffect(() => {
    async function loadRecipe() {
      try {
        const res = await fetch(`/api/recipes/${params.id}`, { cache: "no-store" })
        if (!res.ok) {
          setRecipe({
            id: params.id,
            name: `Mock Recipe ${params.id}`,
            instructions: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris hendrerit neque purus, non facilisis velit aliquet eget. Maecenas bibendum maximus dolor, id rutrum tortor..."
          })
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

  return (
    <SidebarProvider style={{ "--sidebar-width": "14rem" } as React.CSSProperties}>
      <AppSidebar />
      <SidebarInset className="flex flex-col justify-between">

        {/* Breadcrumb */}
        <header className="flex h-16 shrink-0 items-center gap-2 border-b px-4">
          <Breadcrumb>
            <BreadcrumbList>
              <BreadcrumbItem className="hidden md:block">
                <BreadcrumbLink href="/">Home</BreadcrumbLink>
              </BreadcrumbItem>
              <BreadcrumbSeparator className="hidden md:block" />
              <BreadcrumbItem className="hidden md:block">
                <BreadcrumbLink href="/recipes">Recipes</BreadcrumbLink>
              </BreadcrumbItem>
              <BreadcrumbSeparator className="hidden md:block" />
              <BreadcrumbItem>
                <BreadcrumbPage>{recipe ? recipe.name : loading ? "Loading..." : "Not found"}</BreadcrumbPage>
              </BreadcrumbItem>
            </BreadcrumbList>
          </Breadcrumb>
        </header>

        {/* Main content */}
        <div className="p-4 h-full flex flex-col justify-center items-center">
          {loading ? (
            <Spinner />
          ) : !recipe ? (
            <h1 className="text-xl font-medium">Recipe not found</h1>
          ) : (
            <>
              <div className="flex-1 w-full">
                <h1 className="text-3xl font-bold">{recipe.name}</h1>
                <p className="mt-4 whitespace-pre-wrap">{recipe.instructions || "No instructions available"}</p>
              </div>

              {/* Footer actions */}
              <div className="h-16 flex items-center justify-between underline text-lg w-full">
                <Link href="/recipes">
                  <Button variant="outline" className="cursor-pointer">
                    <ChevronLeft />
                    Back
                  </Button>
                </Link>
                <div className="space-x-4">
                  <Button variant="secondary" className="cursor-pointer">
                    <Pencil />
                    Edit
                  </Button>
                  <Dialog>
                    <DialogTrigger>
                      <Button variant="destructive" className="cursor-pointer">
                        <Trash2 />
                        Delete
                      </Button>
                    </DialogTrigger>
                    <DialogContent>
                      <DialogHeader>
                        <DialogTitle>Delete Recipe?</DialogTitle>
                        <DialogDescription>
                          This action cannot be undone. This will permanently delete this recipe.
                        </DialogDescription>
                      </DialogHeader>
                      <DialogFooter className="flex justify-between!">
                        <DialogClose className="cursor-pointer">Back</DialogClose>
                        <Button
                          onClick={() => {
                            deleteRecipe()
                            router.push("/recipes")
                            toast.info(`Recipe Deleted`, {
                              description: `"${recipe.name}" has been permanently removed.`,
                            })
                          }}
                          variant="destructive"
                          className="cursor-pointer"
                        >
                          <Trash2 />
                          Delete
                        </Button>
                      </DialogFooter>
                    </DialogContent>
                  </Dialog>
                </div>
              </div>
            </>
          )}
        </div>
      </SidebarInset>
    </SidebarProvider>
  )
}
