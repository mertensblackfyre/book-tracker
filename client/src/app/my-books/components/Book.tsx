import {
    Card,
    CardDescription,
    CardHeader,
    CardTitle,
} from "@/components/ui/card";
import Image from "next/image"

const Book = () => {
    return (
        <>
            <div className="flex px-3">
                <Card className="shadow-lg rounded px-3 pt-2 pb-3 px-4 ">
                    <div className="overflow-hidden rounded-md">
                        <Image
                            width={200}
                            height={200}
                            className={
                                "h-auto w-auto object-cover transition-all hover:scale-105"
                            }
                            src={
                                "https://images.pexels.com/photos/3225517/pexels-photo-3225517.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
                            }
                            alt={"Image of the book"}
                        />
                    </div>

                    <CardHeader className="space-y-1">
                        <CardTitle className="text-xl">Hello</CardTitle>
                        <CardDescription>By: Neetcode</CardDescription>
                    </CardHeader>
                </Card>
            </div>
        </>
    );
};

export default Book;
