import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { FC } from "react"

const Book : FC<{data: BookDetails}> = ({data}) => {
    return (
        <div>
            <Card className="shadow-lg rounded px-8 pt-6 pb-8">
                <img src={data.coverImgUrl} className="card-imp-top" alt="..." />
                <CardHeader className="space-y-1">
                    <CardTitle className="text-2xl">{data.title}</CardTitle>
                    <CardDescription>
                        By: {data.author}
                    </CardDescription>
                </CardHeader>
                <CardContent>
                    <CardDescription>
                        Sample book description
                    </CardDescription>
                </CardContent>
                <CardFooter className="items-center justify-center">
                    {/* <Button variant="outline">
                                <Icons.google className="mr-2 h-4 w-4" />
                                Google Authentication
                            </Button> */}
                </CardFooter>
            </Card>
        </div>
    )
}

export default Book;
