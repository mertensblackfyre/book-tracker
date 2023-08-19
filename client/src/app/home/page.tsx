"use client";

import { BookCard } from "./components/BookCard";
import Container from "@/components/layouts/Container";

export default function HomePage() {
   return (
      <>
         <Container>
            <section className="my-[30px]">
               <header className="mb-[25px]">
                  <div className="space-y-1">
                     <h2 className="text-2xl font-semibold tracking-tight">
                        Listen Now
                     </h2>
                     <p className="text-sm text-muted-foreground">
                        Top picks for you. Updated daily.
                     </p>
                  </div>
               </header>
               <div className="flex  flex-row">
                  <BookCard />

                  <BookCard />
               </div>
            </section>

            <section className="my-[30px]">
               <header className="mb-[25px]">
                  <div className="space-y-1">
                     <h2 className="text-2xl font-semibold tracking-tight">
                        Listen Now
                     </h2>
                     <p className="text-sm text-muted-foreground">
                        Top picks for you. Updated daily.
                     </p>
                  </div>
               </header>
               <div className="flex  flex-row">
                  <BookCard />

                  <BookCard />
               </div>
            </section>
         </Container>
      </>
   );
}
