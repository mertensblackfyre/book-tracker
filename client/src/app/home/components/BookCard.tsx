import Image from "next/image";
import { PlusCircledIcon } from "@radix-ui/react-icons";

import { cn } from "@/lib/utils";
import {
   ContextMenu,
   ContextMenuContent,
   ContextMenuItem,
   ContextMenuSeparator,
   ContextMenuSub,
   ContextMenuSubContent,
   ContextMenuSubTrigger,
   ContextMenuTrigger,
} from "@/components/ui/context-menu";

export function BookCard() {
   return (
      <div className={cn("space-y-3 pl-12")}>
         <ContextMenu>
            <ContextMenuTrigger>
               <div className="overflow-hidden rounded-md">
                  <Image
                     width={200}
                     height={200}
                     className={cn(
                        "h-auto w-auto object-cover transition-all hover:scale-105"
                     )}
                     src={
                        "https://images.pexels.com/photos/3225517/pexels-photo-3225517.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
                     }
                     alt={""}
                  />
               </div>
            </ContextMenuTrigger>
            <ContextMenuContent className="w-40">
               <ContextMenuItem>Add to Library</ContextMenuItem>
               <ContextMenuSub>
                  <ContextMenuSubTrigger>Add to Playlist</ContextMenuSubTrigger>
                  <ContextMenuSubContent className="w-48">
                     <ContextMenuItem>
                        <PlusCircledIcon className="mr-2 h-4 w-4" />
                        New Playlist
                     </ContextMenuItem>
                     <ContextMenuSeparator />
                     {/* {playlists.map((playlist) => (
                        <ContextMenuItem key={playlist}>
                           <svg
                              xmlns="http://www.w3.org/2000/svg"
                              fill="none"
                              stroke="currentColor"
                              strokeLinecap="round"
                              strokeLinejoin="round"
                              strokeWidth="2"
                              className="mr-2 h-4 w-4"
                              viewBox="0 0 24 24"
                           >
                              <path d="M21 15V6M18.5 18a2.5 2.5 0 1 0 0-5 2.5 2.5 0 0 0 0 5ZM12 12H3M16 6H3M12 18H3" />
                           </svg>
                           {playlist}
                        </ContextMenuItem>
                     ))} */}
                  </ContextMenuSubContent>
               </ContextMenuSub>
               <ContextMenuSeparator />
               <ContextMenuItem>Play Next</ContextMenuItem>
               <ContextMenuItem>Play Later</ContextMenuItem>
               <ContextMenuItem>Create Station</ContextMenuItem>
               <ContextMenuSeparator />
               <ContextMenuItem>Like</ContextMenuItem>
               <ContextMenuItem>Share</ContextMenuItem>
            </ContextMenuContent>
         </ContextMenu>
         <div className="space-y-1 text-sm">
            <h3 className="font-medium leading-none">Amr</h3>
            <p className="text-xs text-muted-foreground">cool</p>
         </div>
      </div>
   );
}
