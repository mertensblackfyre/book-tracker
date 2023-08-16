'use client'

import { Menu } from "@/app/home/components/Menu";
import { Sidebar } from "@/app/home/components/SideBar";
import Image from "next/image";
import { FC, ReactNode } from "react";


const Container : FC<{children: ReactNode}> = ({children}) => {
    return (
      <>
         <div className="md:hidden">
            <Image
               src="/examples/music-light.png"
               width={1280}
               height={1114}
               alt="Music"
               className="block dark:hidden"
            />
            <Image
               src="/examples/music-dark.png"
               width={1280}
               height={1114}
               alt="Music"
               className="hidden dark:block"
            />
         </div>
         <div className="hidden md:block">
            <Menu />
            <div className="border-t">
               <div className="bg-background">
                  <div className="grid lg:grid-cols-5">
                     <Sidebar
                        playlists={[]}
                        className="hidden lg:block"
                     />
                     <div className="col-span-3 lg:col-span-4 lg:border-l">
                        <div className="h-full px-4 py-6 lg:px-8">
                            {children}
                        </div>
                     </div>
                  </div>
               </div>
            </div>
         </div>
      </>
   );
}

export default Container;