"use client";

import { Button } from "@/components/ui/button";
import { Card, CardHeader, CardTitle, CardFooter } from "@/components/ui/card";
import { Icons } from "@/components/ui/icons";
import React from "react";
import axios from "axios";
import { API_URL } from "@/api/url";

const LoginPage = () => {
    const config = {
        headers: {
            "Access-Control-Allow-Origin": "*",
            "Access-Control-Allow-Methods": "GET,PUT,POST,DELETE,PATCH,OPTIONS",
        },
        origin:"http://localhost:3000",

        withCredentials: true,
    };

    const registerUser = async () => {
        console.log("Hello World 1:");
        try {
            const response = await axios.get(`${API_URL}/auth/google`, config);

            console.log("Hello World 2");
            if (response.status === 200) {
                localStorage.setItem("user", JSON.stringify(response.data));
                window.location.href = "/dashboard";
            }

            return response.data;
        } catch (error: any) {
            console.log(error);
        }
    };

    return (
        <>
            <div className="flex items-center justify-center h-screen">
                <Card className=" w-96 shadow-lg rounded px-8 pt-6 pb-8">
                    <CardHeader className="space-y-3">
                        <CardTitle className="text-2xl">Create an account</CardTitle>
                        <span>&nbsp;&nbsp;</span>
                    </CardHeader>

                    <CardFooter className="items-center justify-center flex flex-col">
                        <Button
                            className="my-2"
                            variant="outline"
                            onClick={registerUser}
                        >
                            <Icons.google className="mr-2 h-4 w-4" />
                            Google
                        </Button>
                        <Button
                            className="my-2"
                            variant="outline"
                            onClick={registerUser}
                        >
                            <Icons.gitHub className="mr-2 h-4 w-4" />
                            Github
                        </Button>
                    </CardFooter>
                </Card>
            </div>
        </>
    );
};

export default LoginPage;
