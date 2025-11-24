import { AppSidebar } from "@/components/app-sidebar"
import {
  SidebarInset,
  SidebarProvider
} from "@/components/ui/sidebar"
import { FlickeringGrid } from "@/components/ui/shadcn-io/flickering-grid"

export default function Page() {
  return (
    <SidebarProvider
      style={
        {
          "--sidebar-width": "14rem",
        } as React.CSSProperties
      }
    >
      <AppSidebar />
      <SidebarInset>
        <div className="relative h-screen">
          <FlickeringGrid
            className="absolute inset-0"
            squareSize={4}
            gridGap={6}
            flickerChance={0.3}
            color="rgb(100, 100, 100)"
            maxOpacity={0.2}
          />
          <div className="relative z-10 flex items-center justify-center h-full flex-col">
            <h1 className="text-4xl font-bold bg-linear-to-b from-white to-neutral-400 bg-clip-text text-transparent">Welcome Back</h1>
          </div>
        </div>
      </SidebarInset>
    </SidebarProvider>
  )
}
