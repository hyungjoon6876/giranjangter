import { redirect } from "next/navigation";

// /listings is a canonical browse URL but the listings grid lives at /.
// Redirect server-side so direct visits (deep-links, SEO, share-back from /listings/[id])
// land on the actual browse page instead of a 404.
export default function ListingsIndexPage() {
  redirect("/");
}
