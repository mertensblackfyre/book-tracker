"use client";

import { BookCard } from "./components/BookCard";
import Container from "@/components/layouts/Container";

export default function HomePage() {
   return (
      <>
         <Container>
            <section className="my-[30px]">
               <header className="mb-[25px]">
                  <h2 className="scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0">
                     Top Rated
                  </h2>
               </header>
               <div className="flex  flex-row">
                  <BookCard />

                  <BookCard />
               </div>
            </section>

            <section className="my-[30px]">
               <header className="mb-[25px]">
                  <h2 className="scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0">
                     New
                  </h2>
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
