import {
   Menubar,
   MenubarCheckboxItem,
   MenubarContent,
   MenubarItem,
   MenubarLabel,
   MenubarMenu,
   MenubarRadioGroup,
   MenubarRadioItem,
   MenubarSeparator,
   MenubarShortcut,
   MenubarSub,
   MenubarSubContent,
   MenubarSubTrigger,
   MenubarTrigger,
} from "@/components/ui/menubar";
import { UserNav } from "./UserNav";

export function Menu() {
   return (
      <Menubar className="rounded-none border-b border-none px-2 lg:px-4 flex items-center justify-between">
         <MenubarMenu>
            <MenubarTrigger className="font-bold">myBooks</MenubarTrigger>
         </MenubarMenu>
         <MenubarMenu>
            <UserNav />
         </MenubarMenu>
      </Menubar>
   );
}
