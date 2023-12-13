// place files you want to import through the `$lib` alias in this folder.
import Cookies from "js-cookie"

export const foo = () => {
    let token = Cookies.get("Token");
    console.log(token);
}

