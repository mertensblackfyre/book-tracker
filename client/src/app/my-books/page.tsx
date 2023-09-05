import Container from "@/components/layouts/Container";
import { FC } from "react";
import BookList from "./components/BookList";

type UserProfilePageType = {
   params: { username: string };
};

const UserProfilePage: FC<UserProfilePageType> = ({ params }) => {
   return (
      <>
         <Container>
            <header className="mb-[25px]">
               <h2 className="scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0">
                  books
               </h2>
            </header>
            <div>
               <BookList />
            </div>
         </Container>
      </>
   );
};

export default UserProfilePage;
