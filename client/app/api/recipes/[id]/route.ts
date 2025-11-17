import { NextRequest, NextResponse } from "next/server";

export async function GET(
  req: NextRequest,
  context: { params: Promise<{ id: string }> }
) {
  const { id } = await context.params;

  const res = await fetch(`http://tarmo:9136/recipes/${encodeURIComponent(id)}`);

  if (!res.ok) {
    return NextResponse.json({ error: "Recipe not found" }, { status: res.status });
  }

  const data = await res.json();
  return NextResponse.json(data);
}
