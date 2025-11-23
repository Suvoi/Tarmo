"use client"

import { useState } from "react"
import { AppSidebar } from "@/components/app-sidebar"
import { Breadcrumb, BreadcrumbList, BreadcrumbItem, BreadcrumbLink, BreadcrumbSeparator, BreadcrumbPage } from "@/components/ui/breadcrumb"
import { Card, CardAction, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Textarea } from "@/components/ui/textarea"
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { FieldSet, FieldGroup, Field, FieldLabel } from "@/components/ui/field"
import Link from "next/link"

export default function Page() {
  const [name, setName] = useState("")
  const [instructions, setInstructions] = useState("")
  const [loading, setLoading] = useState(false)

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault()
    setLoading(true)

    try {
      const res = await fetch("/api/recipes", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name, description: instructions }),
      })

      if (!res.ok) throw new Error(`HTTP ${res.status}`)
      alert("Recipe created successfully!")
      setName("")
      setInstructions("")
    } catch (err) {
      console.error("Failed to create recipe:", err)
      alert("Could not connect to the API.")
    } finally {
      setLoading(false)
    }
  }

  return (
    <SidebarProvider style={{ "--sidebar-width": "14rem" } as React.CSSProperties}>
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
                <BreadcrumbLink href="/recipes">Recipes</BreadcrumbLink>
              </BreadcrumbItem>
              <BreadcrumbSeparator className="hidden md:block" />
              <BreadcrumbItem>
                <BreadcrumbPage>New</BreadcrumbPage>
              </BreadcrumbItem>
            </BreadcrumbList>
          </Breadcrumb>
        </header>

        <Card className="h-full m-4">
          <CardHeader>
            <CardTitle className="text-xl">Create a New Recipe</CardTitle>
            <CardAction>
              <Link href="/recipes">Back</Link>
            </CardAction>
          </CardHeader>

          <form onSubmit={handleSubmit}>
            <CardContent className="h-full">
              <FieldSet>
                <FieldGroup>
                  <Field>
                    <FieldLabel htmlFor="name">Name</FieldLabel>
                    <Input
                      id="name"
                      value={name}
                      onChange={(e) => setName(e.target.value)}
                      autoComplete="off"
                      required
                    />
                  </Field>
                  <Field>
                    <FieldLabel>Instructions</FieldLabel>
                    <Textarea
                      value={instructions}
                      onChange={(e) => setInstructions(e.target.value)}
                      className="max-h-64 h-64"
                    />
                  </Field>
                </FieldGroup>
              </FieldSet>
            </CardContent>

            <CardFooter>
              <Button type="submit" disabled={loading} className="w-32 cursor-pointer">
                {loading ? "Creating..." : "Create"}
              </Button>
            </CardFooter>
          </form>
        </Card>
      </SidebarInset>
    </SidebarProvider>
  )
}
