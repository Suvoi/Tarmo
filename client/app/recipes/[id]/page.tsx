"use client"

import { AppSidebar } from "@/components/app-sidebar"
import { Breadcrumb, BreadcrumbList, BreadcrumbItem, BreadcrumbLink, BreadcrumbSeparator, BreadcrumbPage } from "@/components/ui/breadcrumb"
import { Button } from "@/components/ui/button"
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar"
import Link from "next/link"
import { useEffect, useState } from "react"

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

import { ChevronLeft, Trash2, Undo2 } from "lucide-react"

type Recipe = {
  id: string | number
  name: string
  instructions?: string
}

export default function Page({ params }: { params: { id: string } }) {
  const [recipe, setRecipe] = useState<Recipe | null>(null)
  const [loading, setLoading] = useState(true)

  async function deleteRecipe() {
    try {
      const res = await fetch(`/api/recipes/${params.id}`, {
      method: "DELETE"})
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
            instructions: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris hendrerit neque purus, non facilisis velit aliquet eget. Maecenas bibendum maximus dolor, id rutrum tortor. Nulla condimentum, mi vel ornare facilisis, purus neque malesuada leo, eu facilisis orci tellus quis mauris. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Vivamus volutpat facilisis leo, ac tincidunt dui pharetra at. Integer orci diam, fermentum maximus consequat et, semper sit amet orci. Vestibulum vulputate justo lacus, tristique efficitur sapien dignissim euismod. Etiam tortor mi, tincidunt in scelerisque vitae, porta et tortor. Nam lorem dolor, tempus a ultrices et, facilisis. "
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

  if (loading) return <div className="p-4">Loading...</div>
  if (!recipe) return <div className="p-4">Recipe not found</div>

  return (
    <SidebarProvider style={{ "--sidebar-width": "14rem" } as React.CSSProperties}>
      <AppSidebar />
      <SidebarInset className="flex flex-col justify-between">

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
                <BreadcrumbPage>{recipe.name}</BreadcrumbPage>
              </BreadcrumbItem>
            </BreadcrumbList>
          </Breadcrumb>
        </header>

      <div className="p-6 h-full">
        <h1 className="text-3xl font-bold">{recipe.name}</h1>
        <p className="mt-4 whitespace-pre-wrap">
          {recipe.instructions || "No instructions available"}
        </p>
      </div>

      <div className="h-16 px-4 flex items-center justify-between underline text-lg">
        <Link href="/recipes">Back</Link>
        <div className="space-x-4">
          <Link href="#">Edit</Link>
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
                <DialogClose className="cursor-pointer">
                  Back
                </DialogClose>
                <Button onClick={deleteRecipe} variant="destructive" className="cursor-pointer">
                  <Trash2 />
                  Delete
                </Button>
              </DialogFooter>
            </DialogContent>
          </Dialog>
          
        </div>
      </div>
      </SidebarInset>
    </SidebarProvider>
  )
}
