'use client'

import { SAMPLE_BOOK } from "../data/books";
import Book from "./Book";

const BookList = () => {
    return (
        <ul className="flex flex-wrap gap-2">
            {SAMPLE_BOOK.map(v => {
                return (
                    <li key={v.title}>
                        <Book data={v} />
                    </li>
                )
            })}
        </ul>
    )
}

export default BookList;