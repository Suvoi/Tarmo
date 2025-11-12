import { AppSidebar } from "@/components/app-sidebar"
import { Breadcrumb, BreadcrumbList, BreadcrumbItem, BreadcrumbLink, BreadcrumbSeparator, BreadcrumbPage } from "@/components/ui/breadcrumb"
import { Card, CardAction, CardContent, CardFooter, CardHeader, CardTitle, } from "@/components/ui/card"
import { Textarea } from "@/components/ui/textarea"
import {
    SidebarInset,
    SidebarProvider,
} from "@/components/ui/sidebar"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { FieldSet, FieldGroup, Field, FieldLabel } from "@/components/ui/field"
import Link from "next/link"

export default function Page() {
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
                        <CardTitle className="text-xl">
                            Create a New Recipe
                        </CardTitle>
                        <CardAction>
                            <Link href="/recipes">Back</Link>
                        </CardAction>
                    </CardHeader>
                    <CardContent className="h-full">
                        <FieldSet>
                            <FieldGroup>
                                <Field>
                                    <FieldLabel htmlFor="name">Name</FieldLabel>
                                    <Input id="name" autoComplete="off" />
                                </Field>
                                <Field>
                                    <FieldLabel>Instructions</FieldLabel>
                                    <Textarea className="max-h-64 h-64"/>
                                </Field>
                            </FieldGroup>
                        </FieldSet>
                    </CardContent>
                    <CardFooter>
                        <Button type="submit" className="w-32 cursor-pointer">
                            Create
                        </Button>
                    </CardFooter>

                </Card>

            </SidebarInset>
        </SidebarProvider>
    )
}
